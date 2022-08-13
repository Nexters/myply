package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Nexters/myply/application"
)

func main() {
	app, err := application.New()
	if err != nil {
		panic(err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
