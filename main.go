package main

import (
	"exchange/engine"
	"fmt"

	"github.com/shopspring/decimal"
	//"fmt"
	//"encoding/json"
	//"time"
)

func CreateOrder(id string, buy_or_sell engine.Side, quantity decimal.Decimal, price decimal.Decimal, time int64) *engine.Order {
	var created_order engine.Order
	created_order.ID = id
	created_order.Side = buy_or_sell
	created_order.Quantity = quantity
	created_order.Price = price
	created_order.Timestamp = time

	return &created_order
}

func main() {

	data1 := `{
        "id" : "1",
        "side": "buy",
        "quantity": 10,
		"price" : 300,
		"timestamp" : 123
    }`

	/////////////////////////////////////////////////////// BUY ORDERS ARE HERE ////////////////////////////////////////////////////////////////////////////////

	var order1 *engine.Order = &engine.Order{}
	order1.FromJSON([]byte(data1))

	data2 := `{
        "id" : "2",
        "side": "buy",
        "quantity": 10,
		"price" : 200,
		"timestamp" : 123
    }`

	var order2 *engine.Order = &engine.Order{}
	order2.FromJSON([]byte(data2))

	data3 := `{
        "id" : "3",
        "side": "buy",
        "quantity": 10,
		"price" : 100,
		"timestamp" : 123
    }`

	var order3 *engine.Order = &engine.Order{}
	order3.FromJSON([]byte(data3))

	data4 := `{
        "id" : "4",
        "side": "buy",
        "quantity": 10,
		"price" : 150,
		"timestamp" : 123
    }`

	var order4 *engine.Order = &engine.Order{}
	order4.FromJSON([]byte(data4))

	data5 := `{
        "id" : "5",
        "side": "buy",
        "quantity": 10,
		"price" : 125,
		"timestamp" : 123
    }`

	var order5 *engine.Order = &engine.Order{}
	order5.FromJSON([]byte(data5))

	data6 := `{
        "id" : "4",
        "side": "buy",
        "quantity": 10,
		"price" : 25,
		"timestamp" : 123
    }`

	var order6 *engine.Order = &engine.Order{}
	order6.FromJSON([]byte(data6))

	data7 := `{
        "id" : "7",
        "side": "buy",
        "quantity": 10,
		"price" : 780,
		"timestamp" : 123
    }`

	var order7 *engine.Order = &engine.Order{}
	order7.FromJSON([]byte(data7))

	/*
		log.Println(order1)
		log.Println(order2)
		log.Println(order3)
		log.Println(order4)
	*/

	var book *engine.OrderBook = &engine.OrderBook{}
	book.AddBuyOrder(*order4)
	book.AddBuyOrder(*order3)
	book.AddBuyOrder(*order1)
	book.AddBuyOrder(*order2)
	book.AddBuyOrder(*order7)
	book.AddBuyOrder(*order5)
	book.AddBuyOrder(*order6)
	fmt.Println(book)
	// BUY ORDERS ARE ADDED AND THE BOOK IS UPDATED AND CHECKED
	// ADDING THE SELL ORDERS NOW
	/////////////////////////////////////////////////////////////////////////////SELL ORDERS ARE ADDED HERE ///////////////////////////////////////////////////
	sell_data1 := `{
        "id" : "18",
        "side": "sell",
        "quantity": 10,
		"price" : 780,
		"timestamp" : 123
    }`

	var sell_order1 *engine.Order = &engine.Order{}
	sell_order1.FromJSON([]byte(sell_data1))
	sell_data2 := `{
        "id" : "19",
        "side": "sell",
        "quantity": 10,
		"price" : 25,
		"timestamp" : 123
    }`

	var sell_order2 *engine.Order = &engine.Order{}
	sell_order2.FromJSON([]byte(sell_data2))
	sell_data3 := `{
        "id" : "3",
        "side": "sell",
        "quantity": 10,
		"price" : 100,
		"timestamp" : 123
    }`

	var sell_order3 *engine.Order = &engine.Order{}
	sell_order3.FromJSON([]byte(sell_data3))
	sell_data4 := `{
        "id" : "4",
        "side": "sell",
        "quantity": 10,
		"price" : 734,
		"timestamp" : 123
    }`

	var sell_order4 *engine.Order = &engine.Order{}
	sell_order4.FromJSON([]byte(sell_data4))
	sell_data5 := `{
        "id" : "1",
        "side": "sell",
        "quantity": 10,
		"price" : 500,
		"timestamp" : 123
    }`

	var sell_order5 *engine.Order = &engine.Order{}
	sell_order5.FromJSON([]byte(sell_data5))

	book.AddSellOrder(*sell_order1)
	book.AddSellOrder(*sell_order2)
	book.AddSellOrder(*sell_order3)
	book.AddSellOrder(*sell_order4)
	book.AddSellOrder(*sell_order5)

	fmt.Println(book)

	var ab []engine.Trade
	ab = book.Process(*order3)
	fmt.Println(ab)

	//book.AddSellOrder(*sell_order1)

	//log.Println(*book)

	/*sl := []int{200, 300}

	sl = append(sl, 100)
	log.Println(sl)
	copy(sl[1:], sl[0:])
	log.Println(sl)
	sl[0] = 100
	log.Println(sl)*/

	/*var order *engine.Order = &engine.Order{}
	order.FromJSON([]byte(data))
	log.Println(order.Side)
	order.ID = "23"
	log.Println(*order)


	str := order.ToJSON()
	log.Println(string(str))
	*/
}
