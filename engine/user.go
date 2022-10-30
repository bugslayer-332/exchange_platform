package engine

//import "github.com/shopspring/decimal"

type User struct {
	Username string `json:"username"`
	Assets   int    `json:"assets"`
	Fiat     int    `json:"fiat"`
}
