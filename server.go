package main

import (
	"3cho/router"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {
	restApp := fiber.New()
	restApp.Use(middleware.Recover())
	restApp.Use(middleware.Logger())

	router.Init(restApp)
	restApp.Listen(3000)
}
