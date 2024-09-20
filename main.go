package main

import (
	"context"
	"log"
	"net/http"
	"path/filepath"
	routing "stori/src/controllers"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/awslabs/aws-lambda-go-api-proxy/httpadapter"
	"github.com/joho/godotenv"
)

var httpLambda *httpadapter.HandlerAdapter

func init() {
	http.HandleFunc("/api/user/summary", func(w http.ResponseWriter, r *http.Request) {
		routing.UserController(w, r)
	})

	httpLambda = httpadapter.New(http.DefaultServeMux)
}

func Handler(ctx context.Context, req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	return httpLambda.ProxyWithContext(ctx, req)
}

func main() {
	errEnv := godotenv.Load(filepath.Join(".env"))
	if errEnv != nil {
		log.Fatal(errEnv)
	}

	// setup.Connection()
	lambda.Start(Handler)
}
