package main

import (
	"ChiRouter/application"
	"context"
	"fmt"
)

func main() {
	app := application.New()

	err := app.Start(context.TODO())
	if err != nil{
		fmt.Println("app not strted", err)
	}
}
