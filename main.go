package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/websocket/v2"
)

func main() {
	app := fiber.New(fiber.Config{
		DisableStartupMessage: Config.Env != "local" && Config.Env != "default",
	})

	app.Use(cors.New(cors.Config{
		AllowMethods: Config.CorsMethods,
		AllowOrigins: Config.CorsOrigins,
	}))

	app.Get("/", websocket.New(websocketHandler))

	log.Println("Running '" + Config.Env + "' environment on port: " + Config.Port)
	app.Listen(":" + Config.Port)
}

func websocketHandler(c *websocket.Conn) {
	for {
		msgType, msg, err := c.ReadMessage()
		if err != nil {
			break
		}
		c.WriteMessage(msgType, msg)
	}
}
