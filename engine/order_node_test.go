package engine

import (
	"testing"
)

func TestNewOrderNode(t *testing.T) {
	t.Log(NewOrderNode())
}

func TestAddOrderInNode(t *testing.T) {
	var tests = []struct {
		input *Order
	}{
		{NewOrder("b1", Buy, 5.0, 7000.0)},
		{NewOrder("b2", Buy, 10.0, 7000.0)},
		{NewOrder("b3", Buy, 11.0, 7000.0)},
		{NewOrder("b4", Buy, 1.0, 7000.0)},
	}
	on := NewOrderNode()
	volume := 0.0
	for _, tt := range tests {
		on.addOrder(*tt.input)
		volume += tt.input.Amount
	}

	if len(on.Orders) != len(tests) {
		t.Fatalf("Invalid order length (have: %d, want: %d", len(on.Orders), len(tests))
	}

	if on.Volume != volume {
		t.Fatalf("Invalid order volume (have: %f, want: %f", on.Volume, volume)
	}
}

func TestRemoveOrderFromNode(t *testing.T) {
	var tests = []struct {
		input *Order
	}{
		{NewOrder("b1", Buy, 5.0, 7000.0)},
		{NewOrder("b2", Buy, 10.0, 7000.0)},
		{NewOrder("b3", Buy, 11.0, 7000.0)},
		{NewOrder("b4", Buy, 1.0, 7000.0)},
	}
	on := NewOrderNode()
	volume := 0.0
	for _, tt := range tests {
		on.addOrder(*tt.input)
		volume += tt.input.Amount
	}

	on.removeOrder(0)
	volume -= tests[0].input.Amount

	if len(on.Orders) != len(tests)-1 {
		t.Fatalf("Invalid order length (have: %d, want: %d", len(on.Orders), len(tests))
	}

	if on.Volume != volume {
		t.Fatalf("Invalid order volume (have: %f, want: %f", on.Volume, volume)
	}
}

func TestUpdateVolume(t *testing.T) {
	var tests = []struct {
		input *Order
	}{
		{NewOrder("b1", Buy, 5.0, 7000.0)},
		{NewOrder("b2", Buy, 10.0, 7000.0)},
		{NewOrder("b3", Buy, 11.0, 7000.0)},
		{NewOrder("b4", Buy, 1.0, 7000.0)},
	}
	on := NewOrderNode()
	volume := 0.0
	for _, tt := range tests {
		on.updateVolume(tt.input.Amount)
		volume += tt.input.Amount
	}

	if on.Volume != volume {
		t.Fatalf("Invalid order volume (have: %f, want: %f", on.Volume, volume)
	}
}
