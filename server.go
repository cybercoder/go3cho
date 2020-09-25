package main

import (
	"3cho/migrations"
	"3cho/router"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
)

func main() {
	migrations.InitMigrate()
	restApp := fiber.New()
	restApp.Use(middleware.Recover())
	restApp.Use(middleware.Logger())

	router.Init(restApp)
	restApp.Listen(3000)
}
