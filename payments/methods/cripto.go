package methods

import (
	"fmt"
	"math/rand"
)

type Cripto struct{}

func (c Cripto) Pay(amount float64) int {
	fmt.Println("Оплата крипто кошельком")
	fmt.Println("Размер оплаты - ", amount)

	return rand.Int()
}

func (c Cripto) Cancel(id int) {
	fmt.Println("Отмена крипто операции - ", id)
}
