// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package application

import (
	"fmt"
	"github.com/Nexters/myply/application/controller"
	"github.com/Nexters/myply/application/middleware"
	"github.com/Nexters/myply/application/router"
	"github.com/Nexters/myply/docs"
	"github.com/Nexters/myply/domain/member"
	"github.com/Nexters/myply/domain/memos"
	"github.com/Nexters/myply/domain/musics"
	"github.com/Nexters/myply/domain/tag"
	"github.com/Nexters/myply/infrastructure/clients"
	"github.com/Nexters/myply/infrastructure/configs"
	"github.com/Nexters/myply/infrastructure/logger"
	"github.com/Nexters/myply/infrastructure/persistence"
	"github.com/Nexters/myply/infrastructure/persistence/cache"
	"github.com/Nexters/myply/infrastructure/persistence/db"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	recover2 "github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
)

// Injectors from server.go:

func New() (*fiber.App, error) {
	config, err := configs.NewConfig()
	if err != nil {
		return nil, err
	}
	mongoInstance, err := db.NewMongoDB(config)
	if err != nil {
		return nil, err
	}
	memberRepository := persistence.NewMemberRepository(mongoInstance, config)
	authMiddleware := middleware.NewAuthMiddleware(memberRepository)
	memberService := member.NewMemberService(memberRepository)
	memberController := controller.NewMemberController(memberService)
	memberRouter := router.NewMemberRouter(memberController)
	repository := persistence.NewMemoRepository(mongoInstance)
	sugaredLogger, err := logger.NewLogger(config)
	if err != nil {
		return nil, err
	}
	cacheCache, err := cache.NewMongoCacheDB(mongoInstance, config)
	if err != nil {
		return nil, err
	}
	youtubeClient, err := clients.NewYoutubeClient(config)
	if err != nil {
		return nil, err
	}
	musicRepository := persistence.NewMusicRepository(config, cacheCache, youtubeClient)
	service := musics.NewMusicService(sugaredLogger, musicRepository)
	memosService := memos.NewMemoService(repository, service)
	memoController := controller.NewMemoController(memosService, service)
	memoRouter := router.NewMemoRouter(memoController)
	tagRepository := persistence.NewTagRepository()
	tagService := tag.NewTagService(tagRepository)
	musicController := controller.NewMusicController(sugaredLogger, service, memosService, tagService)
	musicsRouter := router.NewMusicsRouter(musicController)
	tagController := controller.NewTagController(tagService)
	tagRouter := router.NewTagRouter(tagController)
	app := NewServer(config, authMiddleware, memberRouter, memoRouter, musicsRouter, tagRouter)
	return app, nil
}

// server.go:

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
	app.Use(recover2.New())

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
