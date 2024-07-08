package utils

import "github.com/gofiber/fiber/v2"

func GetLanguage(c *fiber.Ctx) string {
	lang := c.Get("Accept-Language")
	if lang == "" {
		lang = "en"
	}
	return lang
}
