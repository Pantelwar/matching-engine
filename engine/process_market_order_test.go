package engine

import (
	"fmt"
	"testing"
)

// var decimalZero, _ = util.NewDecimalFromString("0.0")

func TestProcessMarketOrder(t *testing.T) {
	var tests = []struct {
		bookGen        []*Order
		input          *Order
		processedOrder []*Order
		partialOrder   *Order
		book           string
	}{
		// 		////////////////////////////////////////////////////////////////////////
		// 		{
		// 			[]*Order{},
		// 			NewOrder("b1", Buy, DecimalBig("2.0"), DecimalBig("7000.0")),
		// 			[]*Order{},
		// 			nil,
		// 			`------------------------------------------
		// `},
		// 		////////////////////////////////////////////////////////////////////////
		// 		{
		// 			[]*Order{},
		// 			NewOrder("s1", Sell, DecimalBig("2.0"), DecimalBig("7000.0")),
		// 			[]*Order{},
		// 			nil,
		// 			`------------------------------------------
		// `},
		{
			[]*Order{
				NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0")),
			},
			NewOrder("s2", Sell, DecimalBig("5.0"), DecimalBig("8000.0")),
			[]*Order{
				NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0")),
				NewOrder("s2", Sell, DecimalBig("5.0"), decimalZero),
			},
			nil,
			`------------------------------------------
`},
		// 		{
		// 			[]*Order{
		// 				NewOrder("s2", Sell, DecimalBig("5.0"), DecimalBig("8000.0")),
		// 			},
		// 			NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0")),
		// 			[]*Order{
		// 				NewOrder("s2", Sell, DecimalBig("5.0"), DecimalBig("8000.0")),
		// 				NewOrder("b1", Buy, DecimalBig("5.0"), decimalZero),
		// 			},
		// 			nil,
		// 			`------------------------------------------
		// `},
		// 		//////////////////////////////////////////////////////////////////////
		// 		{
		// 			[]*Order{
		// 				NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0")),
		// 			},
		// 			NewOrder("s2", Sell, DecimalBig("5.0"), DecimalBig("7000.0")),
		// 			[]*Order{
		// 				NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0")),
		// 				NewOrder("s2", Sell, DecimalBig("5.0"), decimalZero),
		// 			},
		// 			nil,
		// 			`------------------------------------------
		// `},
		// 		{
		// 			[]*Order{
		// 				NewOrder("s1", Sell, DecimalBig("5.0"), DecimalBig("7000.0")),
		// 			},
		// 			NewOrder("b2", Buy, DecimalBig("5.0"), DecimalBig("7000.0")),
		// 			[]*Order{
		// 				NewOrder("s1", Sell, DecimalBig("5.0"), DecimalBig("7000.0")),
		// 				NewOrder("b2", Buy, DecimalBig("5.0"), decimalZero),
		// 			},
		// 			nil,
		// 			`------------------------------------------
		// `},
		// 		//////////////////////////////////////////////////////////////////////
		// 		{
		// 			[]*Order{
		// 				NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0")),
		// 			},
		// 			NewOrder("s2", Sell, DecimalBig("1.0"), DecimalBig("7000.0")),
		// 			[]*Order{
		// 				NewOrder("s2", Sell, DecimalBig("1.0"), decimalZero),
		// 			},
		// 			NewOrder("b1", Buy, DecimalBig("4.0"), DecimalBig("7000.0")),
		// 			`------------------------------------------
		// 7000 -> 4
		// `},
		// 		{
		// 			[]*Order{
		// 				NewOrder("s1", Sell, DecimalBig("5.0"), DecimalBig("7000.0")),
		// 			},
		// 			NewOrder("b2", Buy, DecimalBig("1.0"), DecimalBig("7000.0")),
		// 			[]*Order{
		// 				NewOrder("b2", Buy, DecimalBig("1.0"), decimalZero),
		// 			},
		// 			NewOrder("s1", Sell, DecimalBig("4.0"), DecimalBig("7000.0")),
		// 			`7000 -> 4
		// ------------------------------------------
		// `},
		// 		//////////////////////////////////////////////////////////////////////
		// 		{
		// 			[]*Order{
		// 				NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0")),
		// 			},
		// 			NewOrder("s2", Sell, DecimalBig("1.0"), DecimalBig("6000.0")),
		// 			[]*Order{
		// 				NewOrder("s2", Sell, DecimalBig("1.0"), decimalZero),
		// 			},
		// 			NewOrder("b1", Buy, DecimalBig("4.0"), DecimalBig("7000.0")),
		// 			`------------------------------------------
		// 7000 -> 4
		// `},
		// 		{
		// 			[]*Order{
		// 				NewOrder("s1", Sell, DecimalBig("5.0"), DecimalBig("7000.0")),
		// 			},
		// 			NewOrder("b2", Buy, DecimalBig("1.0"), DecimalBig("8000.0")),
		// 			[]*Order{
		// 				NewOrder("b2", Buy, DecimalBig("1.0"), decimalZero),
		// 			},
		// 			NewOrder("s1", Sell, DecimalBig("4.0"), DecimalBig("7000.0")),
		// 			`7000 -> 4
		// ------------------------------------------
		// `},

		// 		////////////////////////////////////////////////////////////////////////
		// 		{
		// 			[]*Order{
		// 				NewOrder("b1", Buy, DecimalBig("1.0"), DecimalBig("7000.0")),
		// 				NewOrder("b2", Buy, DecimalBig("1.0"), DecimalBig("6000.0")),
		// 			},
		// 			NewOrder("s3", Sell, DecimalBig("2.0"), DecimalBig("6000.0")),
		// 			[]*Order{
		// 				NewOrder("b1", Buy, DecimalBig("1.0"), DecimalBig("7000.0")),
		// 				NewOrder("b2", Buy, DecimalBig("1.0"), DecimalBig("6000.0")),
		// 				NewOrder("s3", Sell, DecimalBig("2.0"), decimalZero),
		// 			},
		// 			nil,
		// 			`------------------------------------------
		// `},
		// 		////////////////////////////////////////////////////////////////////////
		// 		{
		// 			[]*Order{
		// 				NewOrder("b1", Buy, DecimalBig("1.0"), DecimalBig("7000.0")),
		// 				NewOrder("b2", Buy, DecimalBig("2.0"), DecimalBig("6000.0")),
		// 			},
		// 			NewOrder("s3", Sell, DecimalBig("2.0"), DecimalBig("6000.0")),
		// 			[]*Order{
		// 				NewOrder("b1", Buy, DecimalBig("1.0"), DecimalBig("7000.0")),
		// 				NewOrder("s3", Sell, DecimalBig("2.0"), decimalZero),
		// 			},
		// 			NewOrder("b2", Buy, DecimalBig("1.0"), DecimalBig("6000.0")),
		// 			`------------------------------------------
		// 6000 -> 1
		// `},
		//////////////////////////////////////////////////////////////////////
		{
			[]*Order{
				NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0")),
			},
			NewOrder("s2", Sell, DecimalBig("6.0"), DecimalBig("6000.0")),
			[]*Order{
				NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0")),
				NewOrder("s2", Sell, DecimalBig("6.0"), decimalZero),
			},
			nil,
			`------------------------------------------
`},
		{
			[]*Order{
				NewOrder("s1", Sell, DecimalBig("5.0"), DecimalBig("7000.0")),
			},
			NewOrder("b2", Buy, DecimalBig("6.0"), DecimalBig("8000.0")),
			[]*Order{
				NewOrder("s1", Sell, DecimalBig("5.0"), DecimalBig("7000.0")),
				NewOrder("b2", Buy, DecimalBig("6.0"), decimalZero),
			},
			nil,
			`------------------------------------------
`},

		{
			[]*Order{
				// NewOrder("b1", Buy, DecimalBig("10.0"), DecimalBig("74.0")),
				// NewOrder("b2", Buy, DecimalBig("10.0"), DecimalBig("75.0")),
				NewOrder("b1", Buy, DecimalBig("0.001"), DecimalBig("4000000.00")),
				NewOrder("b2", Buy, DecimalBig("0.001"), DecimalBig("3990000.00")),
			},
			NewOrder("s1", Sell, DecimalBig("0.2"), DecimalBig("4000000.00")),
			[]*Order{
				NewOrder("b1", Buy, DecimalBig("0.001"), DecimalBig("4000000.00")),
				NewOrder("b2", Buy, DecimalBig("0.001"), DecimalBig("3990000.00")),
				NewOrder("s1", Sell, DecimalBig("0.2"), decimalZero),
			},
			nil,
			"",
		},

		////////////////////////////////////////////////////////////////////////

		////////////////////////////////////////////////////////////////////////

		{
			[]*Order{
				// NewOrder("b1", Buy, DecimalBig("10.0"), DecimalBig("74.0")),
				// NewOrder("b2", Buy, DecimalBig("10.0"), DecimalBig("75.0")),
				NewOrder("b1", Buy, DecimalBig("0.2"), DecimalBig("4200000.00")),
				NewOrder("b2", Buy, DecimalBig("0.001"), DecimalBig("4100000.00")),
			},
			NewOrder("s1", Sell, DecimalBig("0.001"), DecimalBig("4200000.00")),
			[]*Order{
				NewOrder("s1", Sell, DecimalBig("0.001"), decimalZero),
			},
			// nil,
			NewOrder("b1", Buy, DecimalBig("0.199"), DecimalBig("4200000.00")),
			"",
		},

		////////////////////////////////////////////////////////////////////////

	}

	for i, tt := range tests {
		ob := NewOrderBook()

		// Order book generation.
		for _, o := range tt.bookGen {
			ob.Process(*o)
		}

		fmt.Println("before:", ob)
		processedOrder, partialOrder := ob.ProcessMarket(*tt.input)
		fmt.Println("result ", i, processedOrder, partialOrder)
		fmt.Println("after:", ob)
		if len(processedOrder) != len(tt.processedOrder) {
			t.Fatalf("Incorrect processedOrder: (have: \n%s\n, want: \n%s\n)", processedOrder, tt.processedOrder)
		}
		for i, po := range processedOrder {
			if po.String() != tt.processedOrder[i].String() {
				fmt.Println(*po, *tt.processedOrder[i], *po == *tt.processedOrder[i])
				fmt.Println(len(po.String()), len((tt.processedOrder[i].String())))
				t.Fatalf("Incorrect processedOrder: (have: \n%s\n, want: \n%s\n)", processedOrder, tt.processedOrder)
			}
		}

		// if ob.String() != tt.book {
		// 	// fmt.Println(len(partialOrder.String()), len((tt.partialOrder.String())))
		// 	t.Fatalf("Incorrect book: (have: \n%s\n, want: \n%s)", ob.String(), tt.book)
		// }

		if tt.partialOrder == nil {
			if partialOrder != tt.partialOrder {
				// fmt.Println(len(partialOrder.String()), len((tt.partialOrder.String())))
				t.Fatalf("Incorrect partialOrder: (have: \n%s\n, want: \n%s)", partialOrder, tt.partialOrder)
			}
		} else {
			if partialOrder.String() != tt.partialOrder.String() {
				// fmt.Println(len(partialOrder.String()), len((tt.partialOrder.String())))
				t.Fatalf("Incorrect partialOrder: (have: \n%s\n, want: \n%s)", partialOrder, tt.partialOrder)
			}
		}
	}
}
