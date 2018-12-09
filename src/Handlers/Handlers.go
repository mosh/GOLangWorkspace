package Handlers

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/aws/aws-lambda-go/events"
)

var (
	// ErrNameNotProvided is thrown when a name is not provided
	ErrNameNotProvided = errors.New("no name was provided in the HTTP body")
)

type platformResponse struct {
	Body string
}

func asJSON(obj interface{}) string {

	blob, _ := json.MarshalIndent(obj, "", " ")
	s := string(blob[:])
	return s
}

func asJSONResponse(obj interface{}) (events.APIGatewayProxyResponse, error) {

	return events.APIGatewayProxyResponse{
		Body:       asJSON(obj),
		StatusCode: 200,
	}, nil

}

// Handler is your Lambda function handler
// It uses Amazon API Gateway request/responses provided by the aws-lambda-go/events package,
// However you could use other event sources (S3, Kinesis etc), or JSON-decoded primitive types such as 'string'.
func Handler(request events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Processing Lambda request %s\n", request.RequestContext.RequestID)

	// If no name is provided in the HTTP request body, throw an error
	if len(request.Body) < 1 {
		return events.APIGatewayProxyResponse{}, ErrNameNotProvided
	}

	obj := platformResponse{Body: request.Body}

	return asJSONResponse(obj)
}
