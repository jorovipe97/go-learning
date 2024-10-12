
# https://docs.aws.amazon.com/lambda/latest/dg/golang-package.html
build:
	@echo "Building..."
	# For the x86_64 architecture use GOARCH=amd64:
	GOOS=linux GOARCH=amd64 go build -tags lambda.norpc -o bootstrap main.go
	
	@echo "Zipping executable..."
	zip hello-lambda-go.zip bootstrap

# https://docs.aws.amazon.com/lambda/latest/dg/lambda-intro-execution-role.html
permissions:
	@echo "Setting permissions..."
	aws iam create-role \
		--role-name hello-lambda-go-role \
		--assume-role-policy-document file://trust-policy.json \
		--region us-east-1
	
	aws iam attach-role-policy --role-name hello-lambda-go-role \
		--policy-arn arn:aws:iam::aws:policy/service-role/AWSLambdaBasicExecutionRole \
		--region us-east-1

deploy: build
	@echo "Deploying..."
	aws lambda create-function --function-name hello-lambda-go \
		--runtime provided.al2023 --handler bootstrap \
		--architectures x86_64 \
		--role arn:aws:iam::490924455933:role/hello-lambda-go-role \
		--zip-file fileb://hello-lambda-go.zip \
		--region us-east-1

update: build
	@echo "Updating..."
	aws lambda update-function-code --function-name hello-lambda-go \
		--zip-file fileb://hello-lambda-go.zip \
		--region us-east-1

# https://stackoverflow.com/a/67526542
invoke:
	@echo "Invoking..."
	aws lambda invoke --function-name hello-lambda-go \
		--region us-east-1 \
		--cli-binary-format raw-in-base64-out \
		--payload file://payload.json \
		response.json