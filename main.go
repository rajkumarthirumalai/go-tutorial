package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int    `json:"id"`
	Completed bool   `json:"completed"`
	Body      string `json:"body"`
}

func main() {
	fmt.Println("Hello World")
	app := fiber.New()
	app.Get("/", func(c *fiber.Ctx) error { return c.Status(200).JSON(fiber.Map{"msg": "hello world"}) })
	todos := []Todo{}
	app.Post("/api/todos", func(c *fiber.Ctx) error {
		todo := &Todo{} //it will empty object with ID zero,completed false and body with empty string
		
		if err := c.BodyParser(todo);err != nil {
			return err
		}

		if todo.Body == "" {
			return c.Status(400).JSON(fiber.Map{"error":"Body must not be empty"})
		}

		todo.ID = len(todos) + 1
		todos =append(todos,*todo )
		return c.Status(201).JSON(todos)
	})
	log.Fatal(app.Listen(":4000"))
}
