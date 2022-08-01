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
		{NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0"))},
		{NewOrder("b2", Buy, DecimalBig("10.0"), DecimalBig("7000.0"))},
		{NewOrder("b3", Buy, DecimalBig("11.0"), DecimalBig("7000.0"))},
		{NewOrder("b4", Buy, DecimalBig("1.0"), DecimalBig("7000.0"))},
	}
	on := NewOrderNode()
	volume := DecimalBig("0.0")
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
		{NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0"))},
		{NewOrder("b2", Buy, DecimalBig("10.0"), DecimalBig("7000.0"))},
		{NewOrder("b3", Buy, DecimalBig("11.0"), DecimalBig("7000.0"))},
		{NewOrder("b4", Buy, DecimalBig("1.0"), DecimalBig("7000.0"))},
	}
	on := NewOrderNode()
	volume := DecimalBig("0.0")
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
		{NewOrder("b1", Buy, DecimalBig("5.0"), DecimalBig("7000.0"))},
		{NewOrder("b2", Buy, DecimalBig("10.0"), DecimalBig("7000.0"))},
		{NewOrder("b3", Buy, DecimalBig("11.0"), DecimalBig("7000.0"))},
		{NewOrder("b4", Buy, DecimalBig("1.0"), DecimalBig("7000.0"))},
	}
	on := NewOrderNode()
	volume := DecimalBig("0.0")
	for _, tt := range tests {
		on.updateVolume(tt.input.Amount)
		volume = volume.Add(tt.input.Amount)
	}

	if on.Volume.Cmp(volume) != 0 {
		t.Fatalf("Invalid order volume (have: %s, want: %s", on.Volume.String(), volume.String())
	}
}
