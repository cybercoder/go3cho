package v1

import (
	user "3cho/router/api/v1/user"

	"github.com/gofiber/fiber"
)

func Home(c *fiber.Ctx) {
	c.Send("Hello, World!")
}

func InitAPIv1(restApp *fiber.App) {
	user.InitRoutes(restApp.Group("/api/v1/user"))
}
