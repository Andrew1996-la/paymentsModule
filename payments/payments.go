package payments

import "fmt"

type PaymentMethods interface {
	Pay(amount float64) int
	Cancel(id int)
}

type PaymentsModule struct {
	paymentsInfo   map[int]paymentInfo
	paymentMethods PaymentMethods
}

func NewPaymentModule(paymentMethods PaymentMethods) *PaymentsModule {
	return &PaymentsModule{
		paymentsInfo:   make(map[int]paymentInfo),
		paymentMethods: paymentMethods,
	}
}

/*
принимает:
- описание покупки
- сумму покупки
возвращает:
- id покупки
*/
func (p *PaymentsModule) Pay(description string, amount float64) int {
	id := p.paymentMethods.Pay(amount)

	info := paymentInfo{
		description: description,
		amount:      amount,
		canceled:    false,
	}

	p.paymentsInfo[id] = info

	return id
}

/*
принимает:
- id покупки для закрытия
возвращает:
- ошибку если не успешно закрыто
*/
func (p *PaymentsModule) Cancel(id int) {
	info, ok := p.paymentsInfo[id]
	
	if !ok {
		fmt.Println("Операция не наидена")
		return
	}

	p.paymentMethods.Cancel(id)

	info.canceled = true
	p.paymentsInfo[id] = info
}

/*
принимает:
- id покупки для закрытия
возвращает:
- информацию о проведенной операции
*/
func (p *PaymentsModule) Show(id int) paymentInfo {
	info, ok := p.paymentsInfo[id]
	
	if !ok {
		fmt.Println("Операция не была наидена")
		return paymentInfo{}
	}

	return info
}

/*
принимает:
- ничего
возвращает:
- информацию о проведенных оперциях
*/
func (p *PaymentsModule) ShowAll() map[int]paymentInfo {
	tempMap := make(map[int]paymentInfo, len(p.paymentsInfo))

	for index, value := range p.paymentsInfo {
		tempMap[index] = value
	}

	return tempMap
}
