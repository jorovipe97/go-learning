package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials/stscreds"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/service/ec2/ec2iface"
)

// Cleaner is a wrapper of aws-sdk client
type Cleaner struct {
	ec2conn         ec2iface.EC2API
	used            map[string]*Used
	resolvedAliases map[string]string
}

// Used is metadata about the details of the image being used
type Used struct {
	Type string
	ID   string
}

func (u *Used) String() string {
	return fmt.Sprintf("(type: '%s': launchTemplate: '%s')", u.Type, u.ID)
}

// NewCleaner returns a new cleaner
func NewCleaner(sess *session.Session) (*Cleaner, error) {
	// Create the credentials from AssumeRoleProvider to assume the role
	// referenced by the "myRoleARN" ARN.
	creds := stscreds.NewCredentials(sess, "arn:aws:iam::211125335344:role/testing-permissions-ami-build-repo-renderly-ci")

	cleaner := &Cleaner{
		ec2conn:         ec2.New(sess, &aws.Config{Credentials: creds}),
		used:            map[string]*Used{},
		resolvedAliases: map[string]string{},
	}

	// err := cleaner.setInstanceUsed()
	// if err != nil {
	// 	return nil, err
	// }

	// err = cleaner.setLaunchConfigurationUsed()
	// if err != nil {
	// 	return nil, err
	// }

	err := cleaner.setLaunchTemplateUsed()
	if err != nil {
		return nil, err
	}

	return cleaner, nil
}

// https://docs.aws.amazon.com/autoscaling/ec2/userguide/using-systems-manager-parameters.html
const ResolveAliases = true

func (c *Cleaner) setLaunchTemplateUsed() error {
	ret, err := c.ec2conn.DescribeLaunchTemplates(&ec2.DescribeLaunchTemplatesInput{})
	if err != nil {
		return err
	}

	for _, lt := range ret.LaunchTemplates {
		versions, err := c.ec2conn.DescribeLaunchTemplateVersions(&ec2.DescribeLaunchTemplateVersionsInput{
			LaunchTemplateId: lt.LaunchTemplateId,
		})
		if err != nil {
			return err
		}

		for _, ltv := range versions.LaunchTemplateVersions {
			if ltv.LaunchTemplateData == nil || ltv.LaunchTemplateData.ImageId == nil {
				continue
			}

			fmt.Println(*ltv.VersionNumber)

			imageId := *ltv.LaunchTemplateData.ImageId

			if ResolveAliases {
				resolvedImageId, err := c.resolveImageAlias(imageId, lt.LaunchTemplateId, ltv.VersionNumber)

				if err != nil {
					return err
				}

				imageId = resolvedImageId
			}

			c.used[imageId] = &Used{
				ID:   fmt.Sprintf("%s (%d)", *ltv.LaunchTemplateName, *ltv.VersionNumber),
				Type: "launch template",
			}
		}
	}
	return nil
}

// If the passed imageAlias is an alias, It will return the resolved imageId.
// Otherwise, it will return the original imageId
func (c *Cleaner) resolveImageAlias(imageAlias string, launchTemplateId *string, launchTemplateVersion *int64) (string, error) {
	// aws ec2 describe-launch-template-versions --launch-template-name poc-blender-render-ssm-param --versions 1 --resolve-alias --region us-east-1
	// If we need to resolve aliases. We need to perform an additional DescribeLaunchTemplateVersions
	// for each version we want to resolve. Because when calling the DescribeLaunchTemplateVersions operation: Resource aliasing (resolveAlias)
	// is only supported when doing single version describe
	imageIdIsAliased := strings.HasPrefix(imageAlias, "resolve:ssm:")
	resolvedImageId, isAliasResolved := c.resolvedAliases[imageAlias]
	imageId := imageAlias

	if imageIdIsAliased && isAliasResolved {
		fmt.Println("Alias already resolved getting from cache: ", imageAlias)
		// If already resolved, Just use the resolved value.
		imageId = resolvedImageId
	} else if imageIdIsAliased && !isAliasResolved {
		fmt.Println("Resolving alias: ", imageAlias)
		// If have not been already resolved, resolve the alias.
		aliasedVersions, err := c.ec2conn.DescribeLaunchTemplateVersions(&ec2.DescribeLaunchTemplateVersionsInput{
			LaunchTemplateId: launchTemplateId,
			Versions: []*string{
				aws.String(strconv.FormatInt(*launchTemplateVersion, 10)),
			},
			ResolveAlias: aws.Bool(true),
		})

		if err != nil {
			return "", err
		}

		if len(aliasedVersions.LaunchTemplateVersions) == 0 {
			return imageAlias, nil
		}

		// Save the resolved value.
		ltv := aliasedVersions.LaunchTemplateVersions[0]
		imageId = *ltv.LaunchTemplateData.ImageId

		// We track if a given alias was already resolved to avoid unnecesary additional DescribeLaunchTemplateVersions requests.
		c.resolvedAliases[imageAlias] = imageId
	}

	return imageId, nil
}

func main() {
	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	}))

	cleanner, error := NewCleaner(sess)

	if error != nil {
		log.Fatal(error)
	}

	for k, v := range cleanner.used {
		fmt.Printf("%s: '%s'\n", k, v)
	}
}
