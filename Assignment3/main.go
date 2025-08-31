package Assignment3

import (
	"fmt"
)

type Order interface {
	Execute() error
}

type MarketOrder struct {
	Symbol   string
	Quantity int
}

type LimitOrder struct {
	Symbol      string
	Quantity    int
	LimitPrice  float64
	MarketPrice float64
}

func (m MarketOrder) Execute() error {
	fmt.Printf("Processing Market Order: Buying %d %s at Market Price\n", m.Quantity, m.Symbol)
	return nil
}

func (l LimitOrder) Execute() error {
	if l.LimitPrice < l.MarketPrice {
		return fmt.Errorf(
			"Limit Order Error: Limit price ₹%.2f is below market price ₹%.2f for %s",
			l.LimitPrice, l.MarketPrice, l.Symbol,
		)
	}
	fmt.Printf("Processing Limit Order: Buying %d %s at ₹%.2f\n", l.Quantity, l.Symbol, l.LimitPrice)
	return nil
}

func ProcessOrder(o Order) {
	err := o.Execute()
	if err != nil {
		fmt.Println("Order Execution Failed:", err)
	}
}

func main() {
	marketOrder := MarketOrder{
		Symbol:   "INFY",
		Quantity: 50,
	}

	limitOrderGood := LimitOrder{
		Symbol:      "AAPL",
		Quantity:    100,
		LimitPrice:  174.25,
		MarketPrice: 170.00,
	}

	limitOrderBad := LimitOrder{
		Symbol:      "GOOG",
		Quantity:    10,
		LimitPrice:  2500.00,
		MarketPrice: 2600.00,
	}

	ProcessOrder(marketOrder)
	ProcessOrder(limitOrderGood)
	ProcessOrder(limitOrderBad)
}
