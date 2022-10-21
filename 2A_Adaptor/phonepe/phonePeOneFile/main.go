package main

import "fmt"

type phonePeInterface interface {
	getBalance(accNumber string)
	sendMoney(from, to string, balance float64)
}

type phonePe struct {
	adptr thirdPartyBankAdaptor
}

func (p *phonePe) getBalance(accNumber string) {
	p.adptr.getBalance(accNumber)
}

func (p *phonePe) sendMoney(from, to string, balance float64) {
	p.adptr.sendMoney(from, to, balance)
}

type iciciBank struct {
}

func (i *iciciBank) fetchBalance(accNumber string) {
	fmt.Printf("Balance for acccount number %v is INR 10000\n\n", accNumber)

}

func (i *iciciBank) transferMoney(from, to string, balance float64) {
	fmt.Printf("Tranferring %v from %v to %v\n\n", balance, from, to)
}

type thirdPartyBankAdaptor struct {
	icici iciciBank
}

func (i *thirdPartyBankAdaptor) getBalance(accNumber string) {
	i.icici.fetchBalance(accNumber)

}
func (i *thirdPartyBankAdaptor) sendMoney(from, to string, balance float64) {
	i.icici.transferMoney(from, to, balance)
}
