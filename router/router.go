package router

import (
	v1 "3cho/router/api/v1"

	"github.com/gofiber/fiber"
)

func Init(restApp *fiber.App) {
	v1.InitAPIv1(restApp)
}
