package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/template/pug/v2"

	"github.com/gofiber/fiber/v2"
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

	engine := pug.New("./views", ".pug")
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Get("/", indexHandler)

	app.Get("/SQLIndex", func(c *fiber.Ctx) error {
		data := getQuery(db)
		return c.Render("sql", fiber.Map{
			"Title":  "Hello",
			"pple":   data,
			"length": len(data),
		}, "layouts/main")
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
		"Title": "PUG!",
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
