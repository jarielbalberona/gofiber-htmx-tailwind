package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sujit-baniya/flash"
)

/********** Handlers for rendering views **********/

// Render Register Page with success/error messages
func HandleLandingPage(ctx *fiber.Ctx) error {
	return ctx.Render("home", fiber.Map{
		"Page":         "Home",
		"Error":        flash.Get(ctx)["Error"],
		"Success":      flash.Get(ctx)["Success"],
		"FlashTitle":   flash.Get(ctx)["FlashTitle"],
		"FlashMessage": flash.Get(ctx)["FlashMessage"],
	})
}

func HandleIdeasPage(ctx *fiber.Ctx) error {
	return ctx.Render("ideas", fiber.Map{
		"Page":         "Ideas",
		"Error":        flash.Get(ctx)["Error"],
		"Success":      flash.Get(ctx)["Success"],
		"FlashTitle":   flash.Get(ctx)["FlashTitle"],
		"FlashMessage": flash.Get(ctx)["FlashMessage"],
	})
}
func HandleLoginPage(ctx *fiber.Ctx) error {
	return ctx.Render("login", fiber.Map{
		"Page":         "Login",
		"Error":        flash.Get(ctx)["Error"],
		"Success":      flash.Get(ctx)["Success"],
		"FlashTitle":   flash.Get(ctx)["FlashTitle"],
		"FlashMessage": flash.Get(ctx)["FlashMessage"],
	})
}

func HandleFlashLib(ctx *fiber.Ctx) error {
	fm := fiber.Map{
		"Page":         "Login",
		"Success":      true,
		"FlashTitle":   "Welcome",
		"FlashMessage": "You have logged in successfully!!",
	}

	return flash.WithSuccess(ctx, fm).Redirect("/")
}
