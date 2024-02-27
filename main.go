package main

import (
	"context"
	"fmt"

	"github.com/amanraghuvanshi/microservice-golang-redis/app"
)

func main() {
	application := app.New()

	err := application.Start(context.TODO())

	if err != nil {
		fmt.Println("Failed to start the application")
	}
}
