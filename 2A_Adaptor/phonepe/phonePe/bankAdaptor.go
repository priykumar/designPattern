package main

// This code will be present in phonePe repo

type thirdPartyBankAdaptor struct {
	icici iciciBank
}

func (i *thirdPartyBankAdaptor) getBalance(accNumber string) {
	i.icici.fetchBalance(accNumber)

}
func (i *thirdPartyBankAdaptor) sendMoney(from, to string, balance float64) {
	i.icici.transferMoney(from, to, balance)
}
