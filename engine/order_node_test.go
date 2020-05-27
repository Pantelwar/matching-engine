package engine

import (
	"math/big"
	"testing"
)

func TestNewOrderNode(t *testing.T) {
	t.Log(NewOrderNode())
}

func TestAddOrderInNode(t *testing.T) {
	var tests = []struct {
		input *Order
	}{
		{NewOrder("b1", Buy, "5.0", "7000.0")},
		{NewOrder("b2", Buy, "10.0", "7000.0")},
		{NewOrder("b3", Buy, "11.0", "7000.0")},
		{NewOrder("b4", Buy, "1.0", "7000.0")},
	}
	on := NewOrderNode()
	volume := new(big.Float).SetFloat64(0)
	for _, tt := range tests {
		on.addOrder(*tt.input)
		volume = new(big.Float).Add(volume, tt.input.Amount)
	}

	if len(on.Orders) != len(tests) {
		t.Fatalf("Invalid order length (have: %d, want: %d)", len(on.Orders), len(tests))
	}

	if on.Volume.Cmp(volume) != 0 {
		t.Fatalf("Invalid order volume (have: %f, want: %f)", on.Volume, volume)
	}
}

func TestRemoveOrderFromNode(t *testing.T) {
	var tests = []struct {
		input *Order
	}{
		{NewOrder("b1", Buy, "5.0", "7000.0")},
		{NewOrder("b2", Buy, "10.0", "7000.0")},
		{NewOrder("b3", Buy, "11.0", "7000.0")},
		{NewOrder("b4", Buy, "1.0", "7000.0")},
	}
	on := NewOrderNode()
	volume := new(big.Float).SetFloat64(0)
	for _, tt := range tests {
		on.addOrder(*tt.input)
		volume = new(big.Float).Add(volume, tt.input.Amount)
	}

	on.removeOrder(0)
	volume = new(big.Float).Sub(volume, tests[0].input.Amount)

	if len(on.Orders) != len(tests)-1 {
		t.Fatalf("Invalid order length (have: %d, want: %d)", len(on.Orders), len(tests))
	}

	if on.Volume.Cmp(volume) != 0 {
		t.Fatalf("Invalid order volume (have: %f, want: %f)", on.Volume, volume)
	}
}

func TestUpdateVolume(t *testing.T) {
	var tests = []struct {
		input *Order
	}{
		{NewOrder("b1", Buy, "5.0", "7000.0")},
		{NewOrder("b2", Buy, "10.0", "7000.0")},
		{NewOrder("b3", Buy, "11.0", "7000.0")},
		{NewOrder("b4", Buy, "1.0", "7000.0")},
	}
	on := NewOrderNode()
	volume := new(big.Float).SetFloat64(0)
	for _, tt := range tests {
		on.updateVolume(tt.input.Amount)
		volume = new(big.Float).Add(volume, tt.input.Amount)
	}

	if on.Volume.Cmp(volume) != 0 {
		t.Fatalf("Invalid order volume (have: %f, want: %f)", on.Volume, volume)
	}
}
