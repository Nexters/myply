package main

import (
	"fmt"

	"github.com/Nexters/myply/infrastructure/configs" // TODO: wire
	"github.com/Nexters/myply/infrastructure/logger"  //TODO: wire
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "github.com/Nexters/myply/docs"
)

// @title MYPLY SERVER
// @version 1.0
// @description This is a sample swagger for Fiber
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email minkj1992@gmail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	config := configs.NewConfig()
	logger := logger.NewLogger(config)

	logger.Infof("Configuration settings\n%+v", config)

	app := fiber.New()

	app.Get("/swagger/*", swagger.HandlerDefault)     // swagger
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("[%s] Hello, myply ✈️", config.Phase))
	})
	app.Listen(":8080")
}
