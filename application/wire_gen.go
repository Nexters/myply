// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package application

import (
	"fmt"

	"github.com/Nexters/myply/application/controller"
	"github.com/Nexters/myply/application/router"
	"github.com/Nexters/myply/docs"
	"github.com/Nexters/myply/domain/member"
	"github.com/Nexters/myply/domain/memos"
	"github.com/Nexters/myply/domain/service"
	"github.com/Nexters/myply/infrastructure/clients"
	"github.com/Nexters/myply/infrastructure/configs"
	"github.com/Nexters/myply/infrastructure/logger"
	"github.com/Nexters/myply/infrastructure/persistence"
	"github.com/Nexters/myply/infrastructure/persistence/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"go.uber.org/zap"
)

// Injectors from server.go:

func New() (*fiber.App, error) {
	config, err := configs.NewConfig()
	if err != nil {
		return nil, err
	}
	sugaredLogger, err := logger.NewLogger(config)
	if err != nil {
		return nil, err
	}
	mongoInstance, err := db.NewMongoDB(config)
	if err != nil {
		return nil, err
	}
	memberRepository := persistence.NewMemberRepository(mongoInstance, config)
	memberService := member.NewMemberService(memberRepository)
	memberController := controller.NewMemberController(memberService)
	memberRouter := router.NewMemberRouter(memberController)
	repository := persistence.NewMemoRepository(mongoInstance)
	memosService := memos.NewMemoService(repository)
	memoController := controller.NewMemoController(memosService)
	memoRouter := router.NewMemoRouter(memoController)
	youtubeClient, err := clients.NewYoutubeClient(config)
	if err != nil {
		return nil, err
	}
	musicRepository := persistence.NewMusicRepository()
	musicsService := service.NewMusicService(sugaredLogger, youtubeClient, musicRepository)
	musicController := controller.NewMusicController(sugaredLogger, musicsService)
	musicsRouter := router.NewMusicsRouter(musicController)
	app := NewServer(config, sugaredLogger, mongoInstance, memberRouter, memoRouter, musicsRouter)
	return app, nil
}

// server.go:

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
func NewServer(
	config *configs.Config, logger2 *zap.SugaredLogger,
	mongo *db.MongoInstance,
	memberRouter router.MemberRouter,
	memoRouter router.MemoRouter,
	musicsRouter router.MusicsRouter,
) *fiber.App {
	app := fiber.New(fiber.Config{

		ErrorHandler: CustomErrorHandler,
	})

	app.Use(cors.New())
	app.Use(etag.New())
	app.Use(recover2.New())

	setSwagger(config.BaseURI)
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("[%s] Hello, myply ✈️", config.Phase))
	})

	api := app.Group("/api")
	v1 := api.Group("/v1")

	memberRouter.Init(&v1)
	memoRouter.Init(&v1)
	musicsRouter.Init(&v1)

	return app
}

func setSwagger(baseURI string) {
	docs.SwaggerInfo.
		Title = "Myply Server ✈️"
	docs.SwaggerInfo.
		Description = "This is a My Playlist server."
	docs.SwaggerInfo.
		Version = "1.0"
	docs.SwaggerInfo.
		Host = baseURI
	docs.SwaggerInfo.
		BasePath = "/api/v1"
}

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {

	httpCode := fiber.StatusInternalServerError
	code := controller.Unknown
	message := "Internal Server Error"

	if e, ok := err.(*fiber.Error); ok {
		httpCode = e.Code
		message = e.Message

	}

	ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
	return ctx.Status(httpCode).JSON(controller.BaseResponse{
		Code:    code,
		Message: message,
	})
}
