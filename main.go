package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)



func welcome(c *fiber.Ctx) error{
	return c.SendString("Welcome to kitobz")
}

func main() {
	app := fiber.New()
	connectDatabase()
	app.Get("/api",welcome)

	log.Fatal()
	//app.Listen(":3000")
}