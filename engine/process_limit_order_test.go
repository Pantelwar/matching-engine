package engine

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
)

func TestProcessLimitOrder(t *testing.T) {
	var tests = []struct {
		bookGen        []*Order
		input          *Order
		processedOrder []*Order
		partialOrder   *Order
	}{
		{
			[]*Order{
				NewOrder("b1", Buy, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0)),
			},
			NewOrder("s2", Sell, decimal.NewFromFloat(5.0), decimal.NewFromFloat(8000.0)),
			[]*Order{},
			nil,
		},
		{
			[]*Order{
				NewOrder("s2", Sell, decimal.NewFromFloat(5.0), decimal.NewFromFloat(8000.0)),
			},
			NewOrder("b1", Buy, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0)),
			[]*Order{},
			nil,
		},
		////////////////////////////////////////////////////////////////////////
		{
			[]*Order{
				NewOrder("b1", Buy, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0)),
			},
			NewOrder("s2", Sell, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0)),
			[]*Order{
				NewOrder("b1", Buy, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0)),
				NewOrder("s2", Sell, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0)),
			},
			nil,
		},
		{
			[]*Order{
				NewOrder("s1", Sell, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0)),
			},
			NewOrder("b2", Buy, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0)),
			[]*Order{
				NewOrder("s1", Sell, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0)),
				NewOrder("b2", Buy, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0)),
			},
			nil,
		},
		////////////////////////////////////////////////////////////////////////
		{
			[]*Order{
				NewOrder("b1", Buy, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0)),
			},
			NewOrder("s2", Sell, decimal.NewFromFloat(1.0), decimal.NewFromFloat(7000.0)),
			[]*Order{
				NewOrder("s2", Sell, decimal.NewFromFloat(1.0), decimal.NewFromFloat(7000.0)),
			},
			NewOrder("b1", Buy, decimal.NewFromFloat(4.0), decimal.NewFromFloat(7000.0)),
		},
		{
			[]*Order{
				NewOrder("s1", Sell, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0)),
			},
			NewOrder("b2", Buy, decimal.NewFromFloat(1.0), decimal.NewFromFloat(7000.0)),
			[]*Order{
				NewOrder("b2", Buy, decimal.NewFromFloat(1.0), decimal.NewFromFloat(7000.0)),
			},
			NewOrder("s1", Sell, decimal.NewFromFloat(4.0), decimal.NewFromFloat(7000.0)),
		},
		////////////////////////////////////////////////////////////////////////
		{
			[]*Order{
				NewOrder("b1", Buy, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0)),
			},
			NewOrder("s2", Sell, decimal.NewFromFloat(1.0), decimal.NewFromFloat(6000.0)),
			[]*Order{
				NewOrder("s2", Sell, decimal.NewFromFloat(1.0), decimal.NewFromFloat(7000.0)),
			},
			NewOrder("b1", Buy, decimal.NewFromFloat(4.0), decimal.NewFromFloat(7000.0)),
		},
	}

	for i, tt := range tests {
		ob := NewOrderBook()

		// Order book generation.
		for _, o := range tt.bookGen {
			ob.Process(*o)
		}

		processedOrder, partialOrder := ob.Process(*tt.input)
		fmt.Println("result ", i, processedOrder, partialOrder)
		for i, po := range processedOrder {
			if *po != *tt.processedOrder[i] {
				t.Fatalf("Incorrect processedOrder: (have: \n%s\n, want: \n%s\n)", processedOrder, tt.processedOrder)
			}
		}

		if tt.partialOrder == nil {
			if partialOrder != tt.partialOrder {
				// fmt.Println(len(partialOrder.String()), len((tt.partialOrder.String())))
				t.Fatalf("Incorrect partialOrder: (have: \n%s\n, want: \n%s)", partialOrder, tt.partialOrder)
			}
		} else {
			if *partialOrder != *tt.partialOrder {
				// fmt.Println(len(partialOrder.String()), len((tt.partialOrder.String())))
				t.Fatalf("Incorrect partialOrder: (have: \n%s\n, want: \n%s)", partialOrder, tt.partialOrder)
			}
		}
	}
}
