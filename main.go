package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/jinzhu/gorm"
	"github.com/membrillo-lab/learngo.git/book"
	"github.com/membrillo-lab/learngo.git/database"

	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func helloWorld(c *fiber.Ctx) {
	c.Send("Hello World!!!")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", book.GetBooks)
	app.Get("/api/v1/book/:id", book.GetBook)
	app.Post("/api/v1/book", book.NewBook)
	app.Delete("/api/v1/book/:id", book.DeleteBook)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "books.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database successfully opened")

	database.DBConn.AutoMigrate(&book.Book{})
	fmt.Println("Database automigrated")
}

func main() {
	app := fiber.New()

	initDatabase()
	defer database.DBConn.Close()

	// app.Get("/", helloWorld)
	setupRoutes(app)

	app.Listen(3000)
}
