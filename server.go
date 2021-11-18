package main

import (
	"fmt"
	"os"
	"time"

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

	var MYSQL_USER string
	var MYSQL_PASSWORD string
	var MYSQL_HOST string
	var MYSQL_PORT string
	var env_variable_present bool

	MYSQL_USER, env_variable_present = os.LookupEnv("MYSQL_USER")
	if !env_variable_present {
		MYSQL_USER = "user-not-provided"
	}

	MYSQL_PASSWORD, env_variable_present = os.LookupEnv("MYSQL_PASSWORD")
	if !env_variable_present {
		MYSQL_PASSWORD = "password-not-provided"
	}

	MYSQL_HOST, env_variable_present = os.LookupEnv("MYSQL_HOST")
	if !env_variable_present {
		MYSQL_HOST = "host-not-provided"
	}

	MYSQL_PORT, env_variable_present = os.LookupEnv("MYSQL_PORT")
	if !env_variable_present {
		MYSQL_PORT = "port-not-provided"
	}

	//fmt.Printf("Host: %s, Port: %s, User: %s, Password: %s\n", MYSQL_HOST, MYSQL_PORT, MYSQL_USER, MYSQL_PASSWORD)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/testdb?charset=utf8mb4&parseTime=True&loc=Local", MYSQL_USER, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT)

	var connection_attempts int = 1
	hellorestdatabase.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Printf("Attempting to get database connection %d\n", connection_attempts)

	// attempt getting a connection until it succeeds
	// or fails 10 times
	for err != nil {

		if connection_attempts > 10 {
			// could not connect after 10 attempts
			panic("failed to connect database")
		}

		time.Sleep(5 * time.Second)
		hellorestdatabase.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

		connection_attempts++

		fmt.Printf("Attempting to get database connection %d\n", connection_attempts)

	}

	// at this point, we definitely have a database connection

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
