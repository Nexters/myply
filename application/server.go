//go:build wireinject
// +build wireinject

package application

import (
	"fmt"

	"github.com/Nexters/myply/application/controller"
	"github.com/Nexters/myply/application/router"
	"github.com/Nexters/myply/domain/service"
	"github.com/Nexters/myply/infrastructure/clients"

	"github.com/Nexters/myply/infrastructure/configs"
	"github.com/Nexters/myply/infrastructure/logger"
	"github.com/Nexters/myply/infrastructure/persistence"
	"github.com/Nexters/myply/infrastructure/persistence/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/google/wire"
	"go.uber.org/zap"

	"github.com/Nexters/myply/docs"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
)

func New() (*fiber.App, error) {
	panic(wire.Build(wire.NewSet(
		NewServer,
		logger.Set,
		configs.Set,
		db.Set,
		clients.Set,
		router.Set,
		controller.Set,
		service.Set,
		persistence.Set)))
}

func NewServer(
	config *configs.Config,
	logger *zap.SugaredLogger,
	mongo *db.MongoInstance,
	musicsRouter router.MusicsRouter,
) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(etag.New())

	setSwagger(config.BaseURI)
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("[%s] Hello, myply ✈️", config.Phase))
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	musicsRouter.Init(&v1)

	return app
}

func setSwagger(baseURI string) {
	docs.SwaggerInfo.Title = "Myply Server ✈️"
	docs.SwaggerInfo.Description = "This is a My Playlist server."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = baseURI
	docs.SwaggerInfo.BasePath = "/api/v1"
}
