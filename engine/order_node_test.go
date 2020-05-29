package engine

import (
	"testing"

	"github.com/shopspring/decimal"
)

func TestNewOrderNode(t *testing.T) {
	t.Log(NewOrderNode())
}

func TestAddOrderInNode(t *testing.T) {
	var tests = []struct {
		input *Order
	}{
		{NewOrder("b1", Buy, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0))},
		{NewOrder("b2", Buy, decimal.NewFromFloat(10.0), decimal.NewFromFloat(7000.0))},
		{NewOrder("b3", Buy, decimal.NewFromFloat(11.0), decimal.NewFromFloat(7000.0))},
		{NewOrder("b4", Buy, decimal.NewFromFloat(1.0), decimal.NewFromFloat(7000.0))},
	}
	on := NewOrderNode()
	volume := decimal.NewFromFloat(0.0)
	for _, tt := range tests {
		on.addOrder(*tt.input)
		volume = volume.Add(tt.input.Amount)
	}

	if len(on.Orders) != len(tests) {
		t.Fatalf("Invalid order length (have: %d, want: %d", len(on.Orders), len(tests))
	}

	if on.Volume.Cmp(volume) != 0 {
		t.Fatalf("Invalid order volume (have: %s, want: %s", on.Volume.String(), volume.String())
	}
}

func TestRemoveOrderFromNode(t *testing.T) {
	var tests = []struct {
		input *Order
	}{
		{NewOrder("b1", Buy, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0))},
		{NewOrder("b2", Buy, decimal.NewFromFloat(10.0), decimal.NewFromFloat(7000.0))},
		{NewOrder("b3", Buy, decimal.NewFromFloat(11.0), decimal.NewFromFloat(7000.0))},
		{NewOrder("b4", Buy, decimal.NewFromFloat(1.0), decimal.NewFromFloat(7000.0))},
	}
	on := NewOrderNode()
	volume := decimal.NewFromFloat(0.0)
	for _, tt := range tests {
		on.addOrder(*tt.input)
		volume = volume.Add(tt.input.Amount)
	}

	on.removeOrder(0)
	volume = volume.Sub(tests[0].input.Amount)

	if len(on.Orders) != len(tests)-1 {
		t.Fatalf("Invalid order length (have: %d, want: %d", len(on.Orders), len(tests))
	}

	if on.Volume.Cmp(volume) != 0 {
		t.Fatalf("Invalid order volume (have: %s, want: %s", on.Volume.String(), volume.String())
	}
}

func TestUpdateVolume(t *testing.T) {
	var tests = []struct {
		input *Order
	}{
		{NewOrder("b1", Buy, decimal.NewFromFloat(5.0), decimal.NewFromFloat(7000.0))},
		{NewOrder("b2", Buy, decimal.NewFromFloat(10.0), decimal.NewFromFloat(7000.0))},
		{NewOrder("b3", Buy, decimal.NewFromFloat(11.0), decimal.NewFromFloat(7000.0))},
		{NewOrder("b4", Buy, decimal.NewFromFloat(1.0), decimal.NewFromFloat(7000.0))},
	}
	on := NewOrderNode()
	volume := decimal.NewFromFloat(0.0)
	for _, tt := range tests {
		on.updateVolume(tt.input.Amount)
		volume = volume.Add(tt.input.Amount)
	}

	if on.Volume.Cmp(volume) != 0 {
		t.Fatalf("Invalid order volume (have: %s, want: %s", on.Volume.String(), volume.String())
	}
}
