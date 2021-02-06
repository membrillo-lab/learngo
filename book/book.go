package book

import (
	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
)

type Book struct {
	gorm.Model
	Title  string `json:"title"`
	Author string `json:"author"`
	Rating int    `json:"rating"`
}

func GetBooks(c *fiber.Ctx) {
	c.Send("All books")
}

func GetBook(c *fiber.Ctx) {
	c.Send("A single book")
}

func NewBook(c *fiber.Ctx) {
	c.Send("Add a new book")
}

func DeleteBook(c *fiber.Ctx) {
	c.Send("Delete a single book")
}
