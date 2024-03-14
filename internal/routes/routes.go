package routes

import (
	"github.com/experience-digital/idea-generation/internal/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", handler.HandleLandingPage)
	app.Get("/ideas", handler.HandleIdeasPage)
	app.Get("/login", handler.HandleLoginPage)

	dashboard := app.Group("dashboard/")
	dashboard.Get("index", func(ctx *fiber.Ctx) error {
		return ctx.Render("index", fiber.Map{
			"Title": "Dashboard!",
		})
	})

	// 404 Handler
	app.Use(func(ctx *fiber.Ctx) error {
		return ctx.Render("404", fiber.Map{})
	})
}
