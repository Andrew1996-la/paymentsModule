package methods

import (
	"fmt"
	"math/rand"
)

type Card struct{}

func (c Card) Pay(amount float64) int {
	fmt.Println("Оплата банковской картой")
	fmt.Println("Размер оплаты - ", amount)

	return rand.Int()
}

func (c Card) Cancel(id int) {
	fmt.Println("Отменена операция с id ", id)
}