package main

import (
	"fmt"
	"study/payments"
	"study/payments/methods"
)

func main() {
	method := methods.Card{}
	paymentModule := payments.NewPaymentModule(method)

 	paymentModule.Pay("Iphone", 100)
	
	idMac:=paymentModule.Pay("Mac", 200)
	paymentModule.Cancel(idMac)

	allInfo := paymentModule.ShowAll()

	fmt.Println(allInfo)
}
