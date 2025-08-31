package Assignment2

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

type update struct {
	Ticker string
	Price  float64
	Time   time.Time
}

func SimulateTicker(ctx context.Context, ticker string, ch chan<- update) {
	price := rand.Float64()*10 + 9
	t := time.NewTicker(1 * time.Second)
	defer t.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case now := <-t.C:

			change := (rand.Float64() - 0.5) * 2
			price += change
			ch <- update{
				Ticker: ticker,
				Price:  float64(int(price*100)) / 100,
				Time:   now,
			}
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	tickers := []string{"AAPL", "GOOG", "INFY"}
	ch := make(chan update)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, ticker := range tickers {
		go SimulateTicker(ctx, ticker, ch)
	}

loop:
	for {
		select {
		case update := <-ch:
			fmt.Printf("[%s] %s: %.2f\n", update.Time.Format("15:04:05"), update.Ticker, update.Price)
		case <-ctx.Done():
			break loop
		}
	}
}
