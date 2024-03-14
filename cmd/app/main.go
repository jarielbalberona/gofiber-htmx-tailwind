package main

import (
	"fmt"
	"log"

	"github.com/experience-digital/idea-generation/internal/db"
	"github.com/experience-digital/idea-generation/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

func Init() *fiber.App {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Can't load env file. Err: %s", err)
	}

	db.Init()
	engine := html.New("web/templates", ".html")
	return fiber.New(fiber.Config{
		Views:         engine,
		ViewsLayout:   "base.layout",
		CaseSensitive: true,
		StrictRouting: true,
		ServerHeader:  "Fiber",
		AppName:       "Idea Generation App v1.0",
	})
}

func main() {
	app := Init()

	app.Static("/", "web/static")
	routes.SetupRoutes(app)

	fmt.Println("Server is running on port 9000")
	log.Fatal(app.Listen(":9000"))
}
