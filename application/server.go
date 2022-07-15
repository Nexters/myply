//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package application

import (
	"fmt"

	"github.com/Nexters/myply/infrastructure/configs"
	"github.com/Nexters/myply/infrastructure/logger"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/google/wire"
	"go.uber.org/zap"

	_ "github.com/Nexters/myply/docs"
)

func New() (*fiber.App, error) {
	wire.Build(wire.NewSet(NewServer, logger.Set, configs.Set))
	return nil, nil
}

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
func NewServer(config *configs.Config, logger *zap.SugaredLogger) *fiber.App {
	logger.Infof("Configuration settings\n%+v", config)

	app := fiber.New()

	// TODO: wire routes
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
	// TODO wire routes

	return app
}
