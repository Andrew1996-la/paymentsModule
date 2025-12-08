package methods

import (
	"fmt"
	"math/rand"
)

type PayPal struct{}

func (p PayPal) Pay(amount float64) int {
	fmt.Println("Оплата PayPal")
	fmt.Println("Размер оплаты - ", amount)

	return rand.Int()
}

func (p PayPal) Cancel(id int) {
	fmt.Println("Отменена операция с id ", id)
}