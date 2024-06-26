package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/jfelipearaujo-org/lambda-authorizer/internal/handler"
)

func main() {
	lambda.Start(handler.HandleRequest)
}
