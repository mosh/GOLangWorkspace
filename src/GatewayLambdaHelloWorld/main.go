package main

import (
	"Handlers"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {

	lambda.Start(Handlers.Handler)
}
