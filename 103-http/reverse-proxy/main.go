package main

import (
	"reverseproxy/proxy"

	"github.com/gofiber/fiber/v2"
)

func main() {

	handler := proxy.NewHandler()

	app := fiber.New()

	app.Get("/proxy", handler.ProxyHandler)

	app.Listen(":3000")

}
