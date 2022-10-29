package engine

// making field lowercase makes it private property
type OrderBook struct {
	Bids []Order `json:"bids"`
	Asks []Order `json:"asks"`
}

// APIs
// addBuyOrder(order)
// addSellOrder(order)
// removeBuyOrder(orderId)

// Add the new Order to end of orderbook in bids
func (book *OrderBook) AddBuyOrder(order Order) {
	n := len(book.Bids)
	//fmt.Println("------------------------------------------------")
	//fmt.Println("Iteration ", n)
	//fmt.Println(book.Bids, n) //
	if n == 0 {
		book.Bids = append(book.Bids, order)
	} else {
		//var i int = -1
		var k int
		var c bool = false
		//var mccounter
		for i := n - 1; i >= 0; i-- {
			buyOrder := book.Bids[i]

			// check the price of existing order
			// convert decimal to Signed int
			k = i
			if buyOrder.Price.LessThan(order.Price) {
				c = true

				//fmt.Println(i)
				//fmt.Println("Entering the less than loop")
				k = i
				break
			}
			k = i
			//fmt.Println(i)
		}

		if k == 0 && c == false {
			k = -1
		}

		//fmt.Println(c)
		//fmt.Println(k)
		// if new order price is less than the last order price
		if k == n-1 {
			// append the new order at end
			//fmt.Println("n-1", book.Bids, k, n)
			book.Bids = append(book.Bids, order)
			//fmt.Println("n-1", book.Bids, k, n)
		} else if k == -1 {
			//fmt.Println("i == -1", book.Bids, k, n)
			var j int = 0
			book.Bids = append(book.Bids, order)
			copy(book.Bids[j+1:], book.Bids[j:])
			book.Bids[j] = order
			//fmt.Println("i == -1", book.Bids, k, n)
		} else {
			// add order to the index before the order which
			//fmt.Println("else", book.Bids, k, n)
			book.Bids = append(book.Bids, order)
			copy(book.Bids[k+1:], book.Bids[k:])
			book.Bids[k+1] = order
			//fmt.Println("else", book.Bids, k, n)
		}
	}

	//fmt.Println(book.Bids, n) //
	//fmt.Println("------------------------------------------------")
}

func (book *OrderBook) AddSellOrder(order Order) {
	n := len(book.Asks)

	if n == 0 {
		book.Asks = append(book.Asks, order)
	} else {
		//var i int = -1
		var k int
		var c bool = false
		//var mccounter
		for i := n - 1; i >= 0; i-- {
			buyOrder := book.Asks[i]

			// check the price of existing order
			// convert decimal to Signed int
			k = i
			if buyOrder.Price.LessThan(order.Price) {
				c = true

				//fmt.Println(i)
				//fmt.Println("Entering the less than loop")
				k = i
				break
			}
			k = i
			//fmt.Println(i)
		}

		if k == 0 && c == false {
			k = -1
		}

		//fmt.Println(c)
		//fmt.Println(k)
		// if new order price is less than the last order price
		if k == n-1 {
			// append the new order at end
			//fmt.Println("n-1", book.Asks, k, n)
			book.Asks = append(book.Asks, order)
			//fmt.Println("n-1", book.Asks, k, n)
		} else if k == -1 {
			//fmt.Println("i == -1", book.Asks, k, n)
			var j int = 0
			book.Asks = append(book.Asks, order)
			copy(book.Asks[j+1:], book.Asks[j:])
			book.Asks[j] = order
			//fmt.Println("i == -1", book.Asks, k, n)
		} else {
			// add order to the index before the order which
			//fmt.Println("else", book.Asks, k, n)
			book.Asks = append(book.Asks, order)
			copy(book.Asks[k+1:], book.Asks[k:])
			book.Asks[k+1] = order
			//fmt.Println("else", book.Asks, k, n)
		}
	}

	//fmt.Println(book.Asks, n) //
	//fmt.Println("------------------------------------------------")
}

func (book *OrderBook) RemoveBuyOrder(index int) {
	book.Bids = append(book.Bids[:index], book.Bids[index+1:]...)
}

func (book *OrderBook) RemoveSellOrder(index int) {
	book.Asks = append(book.Asks[:index], book.Asks[index+1:]...)
}
