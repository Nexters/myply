//go:build wireinject
// +build wireinject

package application

import (
	"fmt"
	"github.com/Nexters/myply/domain"

	"github.com/Nexters/myply/application/controller"
	"github.com/Nexters/myply/application/middleware"
	"github.com/Nexters/myply/application/router"

	"github.com/Nexters/myply/infrastructure/clients"
	"github.com/Nexters/myply/infrastructure/configs"
	"github.com/Nexters/myply/infrastructure/logger"
	"github.com/Nexters/myply/infrastructure/persistence"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/google/wire"

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
		clients.Set,
		router.Set,
		middleware.Set,
		controller.Set,
		persistence.Set,
		domain.Set)))
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

// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Device-Token
func NewServer(
	config *configs.Config,
	authMiddleware middleware.AuthMiddleware,
	memberRouter router.MemberRouter,
	memoRouter router.MemoRouter,
	musicsRouter router.MusicsRouter,
	tagRouter router.TagRouter,
) *fiber.App {
	app := fiber.New(fiber.Config{
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

	v1.Use(authMiddleware.New())
	memberRouter.Init(&v1)
	memoRouter.Init(&v1)
	musicsRouter.Init(&v1)
	tagRouter.Init(&v1)

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
