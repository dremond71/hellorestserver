package main

import (
	"fmt"

	"github.com/dremond71/hellorestbook"
	"github.com/dremond71/hellorestdatabase"
	"github.com/gofiber/fiber"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/book", hellorestbook.GetBooks)
	app.Get("/api/v1/book/:id", hellorestbook.GetBook)
	app.Post("/api/v1/book", hellorestbook.NewBook)
	app.Delete("/api/v1/book/:id", hellorestbook.DeleteBook)
}

func initDatabase() {

	var err error

	dsn := "root:my-secret-pw@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"
	hellorestdatabase.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println("Connection Opened to Database")

	hellorestdatabase.DB.AutoMigrate(&hellorestbook.Book{})
	fmt.Println("Database Migrated")

}

func main() {
	app := fiber.New()
	initDatabase()

	setupRoutes(app)
	app.Listen(3000)

	// no need to close db connections anymore in 2021
	//defer hellorestdatabase.DB.Close()

}
