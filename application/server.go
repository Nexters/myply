//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package application

import (
	"context"
	"fmt"
	"github.com/Nexters/myply/application/controller"
	"github.com/Nexters/myply/application/router"
	"github.com/Nexters/myply/domain/service"
	"github.com/Nexters/myply/infrastructure/configs"
	"github.com/Nexters/myply/infrastructure/logger"
	"github.com/Nexters/myply/infrastructure/persistence"
	"github.com/Nexters/myply/infrastructure/persistence/db"
	"github.com/Nexters/myply/infrastructure/persistence/thirdparty"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	"github.com/google/uuid"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"

	_ "github.com/Nexters/myply/docs"
)

func New() (*fiber.App, error) {
	panic(wire.Build(wire.NewSet(NewServer, logger.Set, configs.Set, db.Set)))
}

func NewMusicsController() (controller.MusicsController, error) {
	panic(wire.Build(wire.NewSet(controller.NewMusicsController, service.NewMusicsService, persistence.NewMusicRepository, thirdparty.NewYoutubeApiV3, configs.Set)))
}

func SetRoutes(root *fiber.Router) {
	musicsController, err := NewMusicsController()
	if err == nil {
		router.SetMusicsRouter(root, musicsController)
	}
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
func NewServer(config *configs.Config, logger *zap.SugaredLogger, mongo *db.MongoInstance) *fiber.App {
	// TODO: move to repository
	collection := mongo.Db.Collection("members")
	member := persistence.Member{
		ID:      uuid.NewString(),
		Name:    "leoo",
		MemoIDs: []primitive.ObjectID{},
	}
	// TODO: move to api
	insertionResult, _ := collection.InsertOne(context.Background(), member)
	logger.Infof("Instance\n%+v", insertionResult)

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

	api := app.Group("/api")
	v1 := api.Group("/v1")

	SetRoutes(&v1)

	return app
}
