package main

type phonePe struct {
	adptr thirdPartyBankAdaptor
}

func (p *phonePe) getBalance(accNumber string) {
	p.adptr.getBalance(accNumber)
}

func (p *phonePe) sendMoney(from, to string, balance float64) {
	p.adptr.sendMoney(from, to, balance)
}
