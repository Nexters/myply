package application

//
//import (
//	"fmt"
//	"github.com/Nexters/myply/application/controller"
//	"github.com/gofiber/fiber/v2"
//	"github.com/gofiber/swagger"
//	"github.com/google/wire"
//)
//
//var Set = wire.NewSet(NewRouter)
//
//type Router struct {
//	app            *fiber.App
//	memoController *controller.MemoController
//}
//
//func NewRouter(app *fiber.App, memoController *controller.MemoController) *Router {
//	return &Router{
//		app:            app,
//		memoController: memoController,
//	}
//}
//
//func (r *Router) SetupRoutes() {
//	r.app.Get("/swagger/*", swagger.HandlerDefault)     // swagger
//	r.app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
//		URL:         "http://example.com/doc.json",
//		DeepLinking: false,
//		// Expand ("list") or Collapse ("none") tag groups by default
//		DocExpansion: "none",
//		// Prefill OAuth ClientId on Authorize popup
//	}))
//	r.app.Get("/", func(c *fiber.Ctx) error {
//		return c.SendString(fmt.Sprintf("Hello, myply ✈️"))
//	})
//	r.app.Get("/v1/memos/:id", (*r.memoController).GetMemos)
//}
