package main

import (
	"flag"
	"foodapp/user-service/database"
	"foodapp/user-service/handlers"
	"foodapp/user-service/models"
	"os"

	fiberprometheus "github.com/ansrivas/fiberprometheus/v2"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

var (
	DSN   string
	PORT  string
	debug bool
)

type OrderRequest struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

type OrderResponse struct {
	OrderID string `json:"order_id"`
}

func main() {
	service := "users-service"
	flag.BoolVar(&debug, "debug", false, "sets log level to debug")
	flag.Parse()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	DSN = os.Getenv("DSN")
	if DSN == "" {
		DSN = `host=localhost user=app password=app123 dbname=usersdb port=5432 sslmode=disable`
		log.Info().Msg(DSN)
	}

	PORT = os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	db, err := database.GetConnection(DSN)
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", service).
			Msgf("unable to connect to the database %s", service)
	}

	log.Info().Str("service", service).Msg("database connection is established")
	Init(db)

	// Migrate DB
	db.AutoMigrate(&models.Order{})

	app := fiber.New()

	prom := fiberprometheus.New(service)
	prom.RegisterAt(app, "/metrics") // exposes Prometheus metrics here
	app.Use(prom.Middleware)         // automatic request metrics

	app.Get("/", handlers.Root)
	app.Get("ping", handlers.Ping)
	app.Get("/health", handlers.Health)

	userHandler := handlers.NewUserHandler(database.NewOrderRepository(db))

	order_group := app.Group("/api/v1/users/orders")
	order_group.Post("/", userHandler.CreateOrder)
}
func Init(db *gorm.DB) {
	db.AutoMigrate(&models.Order{})
}
