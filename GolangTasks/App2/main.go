package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()

type OrderMessage struct {
	OrderID int     `json:"order_id"`
	UserID  int     `json:"user_id"`
	Amount  float64 `json:"amount"`
}

func main() {
	// Connect to Redis
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	defer rdb.Close()

	// Connect to Postgres
	connStr := "host=127.0.0.1 port=5432 user=app password=app123 dbname=usersdb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Failed to connect to DB: %v", err)
	}
	defer db.Close()

	// Subscribe to the Redis channel
	subscriber := rdb.Subscribe(ctx, "orders")
	ch := subscriber.Channel()

	fmt.Println("Listening for new orders...")

	for msg := range ch {
		var order OrderMessage
		err := json.Unmarshal([]byte(msg.Payload), &order)
		if err != nil {
			log.Printf("Invalid message format: %v", err)
			continue
		}

		fmt.Printf("Received order: %+v\n", order)

		// Process payment
		status := processPayment(order)

		// Update the order status in DB
		err = updateOrderStatus(db, order.OrderID, status)
		if err != nil {
			log.Printf("Failed to update order status: %v", err)
		} else {
			fmt.Printf("Order %d marked as '%s'\n", order.OrderID, status)
		}
	}
}

func processPayment(order OrderMessage) string {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100)
	if randomNumber%2 == 0 {
		return "success"
	}
	return "failure"
}

func updateOrderStatus(db *sql.DB, orderID int, status string) error {
	_, err := db.Exec(`UPDATE orders SET status = $1 WHERE id = $2`, status, orderID)
	return err
}
