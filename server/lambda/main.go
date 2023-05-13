package lambda

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

func handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	message := "Hello from Lambda!"
	fmt.Println(message)

	return events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(message),
	}, nil
}

func main() {
	lambda.Start(handler)
}
