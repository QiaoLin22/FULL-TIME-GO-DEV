package main

import (
	"flag"

	"github.com/QiaoLin22/FULL-TIME-GO-DEV/hotel-reservation/api"
	"github.com/gofiber/fiber/v2"
)

func main() {
	listenAddr := flag.String("listenaddr", ":5000", "The listen address of the API server")
	flag.Parse()
	app := fiber.New()
	apiv1 := app.Group("/api/v1")

	app.Get("/foo", handleFoo)

	apiv1.Get("/user", api.HandleGetUsers)
	apiv1.Get("/user:id", api.HandleGetUser)
	app.Listen(*listenAddr)
}

func handleFoo(c *fiber.Ctx) error {
	return c.JSON(map[string]string{"msg": "working just fine"})
}
