package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"encoding/json"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

type DateSQL struct {
	ID   int
	NAME string
}

func main() {
	connStr := "root:password@tcp(localhost:3306)/dbGo"
	// Connect to database
	db, err := sql.Open("mysql", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", indexHandler)

	app.Get("/scripts/*", func(c *fiber.Ctx) error {
		script := c.Params("*")
		response := "./scripts/" + script
		return c.SendFile(response, true)
	})

	app.Get("/SQLIndex", func(c *fiber.Ctx) error {
		data := getQuery(db)
		out, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}
		return c.JSON(string(out))
	})

	app.Post("/", postHandler) // Add this

	app.Put("/update", putHandler) // Add this

	app.Delete("/delete", deleteHandler) // Add this

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatalln(app.Listen(fmt.Sprintf(":%v", port)))
}

func indexHandler(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title": "Hello, World!",
	}, "layouts/main")
}

func postHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
func putHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}
func deleteHandler(c *fiber.Ctx) error {
	return c.SendString("Hello")
}

func getQuery(db *sql.DB) []DateSQL {
	result := []DateSQL{}
	results, err := db.Query("SELECT * FROM data")
	if err != nil {
		log.Fatal(err)
	}
	for results.Next() {
		var data DateSQL
		err = results.Scan(&data.ID, &data.NAME)
		if err != nil {
			panic(err.Error())
		}
		result = append(result, data)
	}
	return result
}
