package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"runtime"

	"github.com/aws/aws-lambda-go/lambda"
)

// Following this tutorial
// https://docs.aws.amazon.com/lambda/latest/dg/golang-handler.html

type Order struct {
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
	Item    string  `json:"item"`
}

type CountRequest struct {
	Init          int `json:"init"`
	Main          int `json:"main"`
	HandleRequest int `json:"handleRequest"`
	ThreadCount   int `json:"threadCount"`
}

func NewCountRequest() *CountRequest {
	return &CountRequest{
		Init:          0,
		Main:          0,
		HandleRequest: 0,
		ThreadCount:   runtime.GOMAXPROCS(0),
	}
}

var (
	countRequest = NewCountRequest()
)

func init() {
	log.Println("Init")
	countRequest.Init++
}

func handleRequest(ctx context.Context, event json.RawMessage) (CountRequest, error) {
	log.Println("HandleRequest")
	countRequest.HandleRequest++

	// Parse the input event
	var order Order
	if err := json.Unmarshal(event, &order); err != nil {
		log.Printf("Failed to unmarshal event: %v", err)
		return *countRequest, err
	}

	fmt.Printf("Order: %+v\n", order)
	fmt.Printf("CountRequest: %+v\n", countRequest)

	return *countRequest, nil
}

func main() {
	log.Println("Main")
	countRequest.Main++
	lambda.Start(handleRequest)
}
