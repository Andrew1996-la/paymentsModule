package payments

import "fmt"

//описание интерфейса. Все что умеет Pay и Cancel является PaymentMethods
type PaymentMethods interface {
	Pay(amount float64) int
	Cancel(id int)
}

// Описание структуры PaymentsInfo. 
// paymentsInfo это мапа ключ это id а значение платеж
// paymentMethod это любас структура которая умеет Pay Cancel
type PaymentsModule struct {
	paymentsInfo   map[int]paymentInfo
	paymentMethods PaymentMethods
}

// Коструктор создания платежного модуля. Принимает стурктуру которая умеет Pay Cancel
// Возвращает указатель на модуль с проинициализированной мапой и методом оплаты
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
