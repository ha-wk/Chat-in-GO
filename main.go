package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/gofiber/websocket/v2"
	"github.com/steelthedev/go-chat/handlers"
)

func main() {


    // Create views engine
    viewsEngine := html.New("./views", ".html")


    // Start new fiber instance
    app := fiber.New(fiber.Config{      //UPDATED to handle html rendering
        Views: viewsEngine,
    })

    // Static route and directory
    app.Static("/static/", "./static")

    // Create a "ping" handler to test the server
    app.Get("/ping", func(ctx *fiber.Ctx) error{
        return ctx.SendString("Welcome to fiber")
    })

	// create new App Handler
    appHandler := handlers.NewAppHandler()

    // Add appHandler routes
    app.Get("/", appHandler.HandleGetIndex)

	server := NewWebSocket()
    app.Get("/ws", websocket.New(func(ctx *websocket.Conn) {
        server.HandleWebSocket(ctx)
    }))

    go server.HandleMessages()


    // Start the http server
    app.Listen(":3000")

}