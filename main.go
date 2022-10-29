package main

import (
	"log"
	"exchange/engine"
	"github.com/shopspring/decimal"
	//"fmt"
	//"encoding/json"
	//"time"
)

func CreateOrder(id string, buy_or_sell engine.Side, quantity decimal.Decimal, price decimal.Decimal, time int64) *engine.Order{
	var created_order engine.Order
	created_order.ID = id
	created_order.Side = buy_or_sell
	created_order.Quantity = quantity
	created_order.Price = price
	created_order.Timestamp = time

	return &created_order
}

func main(){

	data1 := `{
        "id" : "1",
        "side": "buy",
        "quantity": 10,
		"price" : 300,
		"timestamp" : 123
    }`

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
        "id" : "1",
        "side": "buy",
        "quantity": 10,
		"price" : 150,
		"timestamp" : 123
    }`
	var order4 *engine.Order = &engine.Order{}
	order4.FromJSON([]byte(data4))

	/*
	log.Println(order1)
	log.Println(order2)
	log.Println(order3)
	log.Println(order4)
	*/

	var book *engine.OrderBook = &engine.OrderBook{}
	book.AddBuyOrder(*order2)
	book.AddBuyOrder(*order3)
	book.AddBuyOrder(*order4)
	book.AddBuyOrder(*order1)

	//log.Println(*book)


	sl := []int {200, 300}

	sl = append(sl, 100)
	log.Println(sl)
	copy(sl[1:], sl[0:])
	log.Println(sl)
	sl[0] = 100
	log.Println(sl)







	/*var order *engine.Order = &engine.Order{}
	order.FromJSON([]byte(data))
	log.Println(order.Side)
	order.ID = "23"
	log.Println(*order)


	str := order.ToJSON()
	log.Println(string(str))
	*/
}