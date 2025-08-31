package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "app"
	password = "app123"
	dbname   = "usersdb"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Cannot ping DB: %v", err)
	}
	fmt.Println("Connected to database!")

	createTables(db)

	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		// Password: "",
		// DB=0
	})
	err1 := rdb.Ping(ctx).Err()
	if err != nil {
		log.Fatalf("Could not connect to REDIS: %v", err1)
	}
	fmt.Println("Connected to Redis!")

	userID := seedUser(db)
	err = createOrder(db, rdb, userID, 99.99)
	if err != nil {
		log.Printf("Error creating order: %v", err)
	}

	err = createOrder(db, rdb, 1, 99.99)
	if err != nil {
		log.Printf("Error creating order: %v", err)
	}
}

func createTables(db *sql.DB) {
	userTable := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL
	);`

	orderTable := `
CREATE TABLE IF NOT EXISTS orders (
	id SERIAL PRIMARY KEY,
	user_id INTEGER REFERENCES users(id),
	amount DECIMAL(10, 2) NOT NULL,
	status TEXT DEFAULT 'pending'
);`

	_, err := db.Exec(userTable)
	if err != nil {
		log.Fatalf("Failed to create users table: %v", err)
	}

	_, err = db.Exec(orderTable)
	if err != nil {
		log.Fatalf("Failed to create orders table: %v", err)
	}

	fmt.Println("Tables created successfully.")
}
func createOrder(db *sql.DB, rdb *redis.Client, userID int, amount float64) error {
	// Insert into orders table
	var orderID int
	err := db.QueryRow(
		`INSERT INTO orders (user_id, amount) VALUES ($1, $2) RETURNING id`,
		userID, amount).Scan(&orderID)
	if err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}

	// Prepare message to publish (you can use JSON or any format)
	message := fmt.Sprintf(`{"order_id": %d, "user_id": %d, "amount": %.2f}`, orderID, userID, amount)

	// Publish to Redis channel "orders"
	err = rdb.Publish(ctx, "orders", message).Err()
	if err != nil {
		return fmt.Errorf("failed to publish to Redis: %w", err)
	}

	fmt.Println("Order created and message published to Redis channel.")
	return nil
}
func seedUser(db *sql.DB) int {
	var userID int
	err := db.QueryRow(`
		INSERT INTO users (name, email)
		VALUES ($1, $2)
		ON CONFLICT (email) DO UPDATE SET name = EXCLUDED.name
		RETURNING id
	`, "Alice", "alice@example.com").Scan(&userID)

	if err != nil {
		log.Fatalf("Failed to insert or get user: %v", err)
	}
	fmt.Println("User seeded with ID:", userID)
	return userID
}
