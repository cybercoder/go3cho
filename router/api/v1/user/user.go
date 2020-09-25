package user

import "github.com/gofiber/fiber"

func Home(c *fiber.Ctx) {
	c.Send("Hello User, World!")
}

func InitRoutes(v1group fiber.Router) {
	v1group.Get("/", Home)
}
