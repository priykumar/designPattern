package main

type phonePeInterface interface {
	getBalance(accNumber string)
	sendMoney(from, to string, balance float64)
}
