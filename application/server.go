//go:build wireinject
// +build wireinject

package application

import (
	"fmt"
	"github.com/Nexters/myply/domain/memos"
	"github.com/Nexters/myply/domain/service"

	"github.com/Nexters/myply/application/controller"
	"github.com/Nexters/myply/application/router"
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
	"github.com/gofiber/fiber/v2/middleware/recover"
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
		memos.Set,
		persistence.Set)))
}

func NewServer(
	config *configs.Config,
	logger *zap.SugaredLogger,
	mongo *db.MongoInstance,
	musicsRouter router.MusicsRouter,
	mc *controller.MemoController,
) *fiber.App {
	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: CustomErrorHandler,
	})

	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(recover.New())

	setSwagger(config.BaseURI)
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("[%s] Hello, myply ✈️", config.Phase))
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	app.Get("/v1/memos/:id", (*mc).GetMemo)
	app.Post("/v1/memos", (*mc).AddMemo)

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

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {

	httpCode := fiber.StatusInternalServerError
	code := controller.Unknown
	message := "Internal Server Error"

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		httpCode = e.Code
		message = e.Message
		// TODO: handle code by httpcode
	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return ctx.Status(httpCode).JSON(
		controller.BaseResponse{
			Code:    code,
			Message: message,
		})
}
