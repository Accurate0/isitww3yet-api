package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	runtime "github.com/aws/aws-lambda-go/lambda"
)

type StatusResponse struct {
	Emoji     string `json:"emoji"`
	Status    string `json:"status"`
	Reference string `json:"reference"`
}

func handleRequest(ctx context.Context, event events.SQSEvent) (events.APIGatewayProxyResponse, error) {
	resp := &StatusResponse{
		Emoji:     "😳",
		Status:    "Maybe",
		Reference: "https://en.wikipedia.org/wiki/2022_Russian_invasion_of_Ukraine",
	}
	// there will literally never be an error :D
	json, _ := json.Marshal(resp)
	return events.APIGatewayProxyResponse{Body: string(json), StatusCode: 200}, nil

}

func main() {
	runtime.Start(handleRequest)
}
