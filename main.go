package main

import (
	"study/payments"
	"study/payments/methods"

	"github.com/k0kubun/pp"
)

func main() {
	methods := methods.Cripto{}

	payments := payments.NewPaymentModule(methods)

	payments.Pay("Бургер", 2)
	payments.Pay("Монитор", 100)
	info := payments.ShowAll()
	pp.Println(info)
}
