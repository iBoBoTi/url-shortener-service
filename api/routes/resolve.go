package routes

import (
	"github.com/go-redis/redis/v8"
	"github.com/gofiber/fiber/v2"
	"github.com/iBoBoTi/url-shortener/api/repository"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	r := repository.CreateClient(0)
	defer r.Close()

	val, err := r.Get(repository.Ctx, url).Result()
	if err != redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "short not found in database",
		})
	} else if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "cannot connect to DB",
		})
	}

	rInr := repository.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(repository.Ctx, "counter")
	return c.Redirect(val, 301)
}
