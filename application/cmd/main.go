package main

import (
	"log"

	"github.com/Nexters/myply/application"
)

func main() {
	app, err := application.New()
	if err != nil {
		panic(err)
	}
	log.Fatal(app.Listen(":8080"))
}
