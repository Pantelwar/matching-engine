package engine

// CancelOrder remove the order from book and returns
func (ob *OrderBook) CancelOrder(id string) *Order {
	ob.mutex.Lock()
	orderNode := ob.orders[id]
	ob.mutex.Unlock()

	if orderNode == nil {
		return nil
	}

	for i, order := range orderNode.Orders {
		if order.ID == id {
			// ob.orders[id].addOrder(*order)
			// ob.removeOrder(order, i)
			orderNode.removeOrder(i)
			// orderNode.updateVolume(-order.Amount)
			// orderNode.Orders = append(orderNode.Orders[:i], orderNode.Orders[i+1:]...)
			if len(orderNode.Orders) == 0 {
				ob.removeOrder(order)
			}
			ob.mutex.Lock()
			delete(ob.orders, id)
			ob.mutex.Unlock()
			return order
		}
	}
	return nil
}
