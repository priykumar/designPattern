package main

import "fmt"

type iciciBank struct {
}

func (i *iciciBank) fetchBalance(accNumber string) {
	fmt.Printf("Balance for acccount number %v is INR 10000\n\n", accNumber)

}

func (i *iciciBank) transferMoney(from, to string, balance float64) {
	fmt.Printf("Tranferring %v from %v to %v\n\n", balance, from, to)
}
