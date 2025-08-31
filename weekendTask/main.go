package weekendTask

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID          uint
	Name, Email string `gorm:"unique"`
}
type Order struct {
	ID                   uint
	UserID               uint
	Status               string
	TotalCents           int
	CreatedAt, UpdatedAt time.Time
}

type Job struct {
	JobID   string
	OrderID uint
}

func connectDB() *gorm.DB {
	dsn := os.Getenv("DATABASE_URL")
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	if err := db.AutoMigrate(&User{}, &Order{}); err != nil {
		log.Fatal(err)
	}
	return db
}

func main() {
	db := connectDB()
	app := fiber.New()

	app.Post("/users", createUserHandler(db))
	app.Get("/users/:id", getUserHandler(db))
	app.Post("/orders", createOrderHandler(db))
	app.Get("/orders/:id", getOrderHandler(db))

	jobsBuf := mustGetIntEnv("JOBS_BUFFER", 100)
	workerCount := mustGetIntEnv("WORKERS", 3)
	jobs := make(chan Job, jobsBuf)
	ctx, cancel := context.WithCancel(context.Background())

	if !fiber.IsChild() {
		startWorkers(ctx, workerCount, jobs, db)
	}

	app.Post("/orders/:id/confirm", func(c *fiber.Ctx) error {
		id := c.Params("id")

	})

	go func() { app.Listen(":3000") }()
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	log.Info().Msg("Shutting down server...")

	cancel()
	time.Sleep(time.Second * 5)

}
