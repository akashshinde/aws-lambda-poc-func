package main

import (
	"context"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequest(ctx context.Context, event MyEvent) (MyEvent, error) {
        time.Sleep(3* time.Minute)
	return MyEvent{
		Name: event.Name,
	}, nil
}

func main() {
	lambda.Start(HandleRequest)
}
