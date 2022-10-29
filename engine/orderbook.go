package engine

import (
	"fmt"
)

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
	fmt.Println("------------------------------------------------")
	fmt.Println("Iteration ",n)
	fmt.Println(book.Bids,n)  //
	if n == 0 {
		book.Bids = append(book.Bids, order)
	} else {
		var i int
		var k int
		for i := n - 1; i >= 0; i-- {
			buyOrder := book.Bids[i]

			// check the price of existing order
			// convert decimal to Signed int
			if buyOrder.Price.LessThan(order.Price) {
				fmt.Println(i)
				k = i
				break
			}
			k = i
			//fmt.Println(i)
		}
		
		fmt.Println(k)
		// if new order price is less than the last order price
		if i == n-1 {
			// append the new order at end
			fmt.Println("n-1",book.Bids,i,n)
			book.Bids = append(book.Bids, order)
			fmt.Println("n-1",book.Bids, i,n)
		}else if i == -1{
			fmt.Println("i == -1",book.Bids,i, n)
			var j int = 0
			book.Bids = append(book.Bids, order)
			copy(book.Bids[j+1:], book.Bids[j:]) 
			book.Bids[j] = order
			fmt.Println("i == -1",book.Bids, i,n)
		}else {
			// add order to the index before the order which
			fmt.Println("else",book.Bids, i, n)
			book.Bids = append(book.Bids, order)
			copy(book.Bids[k+1:], book.Bids[k:]) 
			book.Bids[k] = order 
			fmt.Println("else",book.Bids, i,n)
		}
	}

	fmt.Println(book.Bids, n) //
	fmt.Println("------------------------------------------------")
}

func (book *OrderBook) AddSellOrder(order Order) {
	n := len(book.Asks)

	if n == 0 {
		book.Asks = append(book.Asks, order)
	} else {
		var i int
		for i := n - 1; i >= 0; i-- {
			sellOrder := book.Asks[i]

			if sellOrder.Price.LessThan(order.Price) {
				break
			}
		}
		if i == n-1 {
			// append the new order at end
			book.Asks = append(book.Asks, order)
		} else {
			// add order to the index before the order which
			copy(book.Asks[i+1:], book.Asks[i:])
			book.Asks[i] = order
		}
	}
}

func (book *OrderBook) RemoveBuyOrder(index int) {
	book.Bids = append(book.Bids[:index], book.Bids[index+1:]...)
}

func (book *OrderBook) RemoveSellOrder(index int) {
	book.Asks = append(book.Asks[:index], book.Asks[index+1:]...)
}