package main

import (
	"log"

	"github.com/Nexters/myply/application"
)

func main() {
	app, _ := application.New()
	log.Fatal(app.Listen(":8080"))
}
