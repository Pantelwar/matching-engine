package engine

import (
	"testing"
)

func TestNewOrder(t *testing.T) {
	t.Log(NewOrder("b1", Sell, 5.0, 7000.0))
}

func TestOrder(t *testing.T) {
	data := NewOrder("b1", Buy, 5.0, 7000.0)

	result, _ := data.ToJSON()
	if string(result) != "{\"amount\":5,\"price\":7000,\"id\":\"b1\",\"type\":\"buy\"}" {
		t.Fatal("Result should be: {\"amount\":5,\"price\":7000,\"id\":\"b1\",\"type\":\"buy\"}, got: " + string(result))
	}

	dataO := &Order{}
	err := dataO.FromJSON(result)
	if err == nil {
		t.Log("Successful unmarshalling")
	} else {
		t.Fatal("Incorrect json")
	}

	data1 := &Order{}
	err = data1.FromJSON([]byte("{\"amount\":5,\"price\":7000,\"id\":\"b1\",\"tye\":\"buy\"}"))
	if err != nil {
		t.Log("Successful Detection")
	} else {
		t.Fatal("Approving invalid json")
	}
}
