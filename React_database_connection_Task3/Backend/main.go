package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/lib/pq"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var db *sql.DB

func main() {
	var err error

	db, err = sql.Open("postgres", "host=localhost port=5432 user=app password=app123 dbname=myappdb sslmode=disable")
	if err != nil {
		log.Fatal("Database connection error:", err)
	}
	defer db.Close()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5174", // Adjust if your React dev server uses another port
		AllowMethods: "GET,POST,OPTIONS",
		AllowHeaders: "Content-Type",
	}))
	// POST /users endpoint
	app.Post("/users", func(c *fiber.Ctx) error {
		user := new(User)

		if err := c.BodyParser(user); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Cannot parse JSON",
			})
		}

		result, err := db.Exec("INSERT INTO Users (name, email, password) VALUES ($1, $2, $3)",
			user.Name, user.Email, user.Password)

		if err != nil {
			log.Println("Insert error:", err)
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Failed to insert user",
			})
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Println("RowsAffected error:", err)
		} else {
			log.Printf("Rows affected: %d\n", rowsAffected)
		}

		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "User created",
		})
	})

	log.Println("Server running at http://localhost:8080")
	log.Fatal(app.Listen(":8080"))
}
