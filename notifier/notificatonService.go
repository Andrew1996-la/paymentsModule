package notifier

import "fmt"

type NotificationService struct {
	notifire Notifire
}

func (n *NotificationService) SetNotifire(provider Notifire) {
	n.notifire = provider
}

func (n *NotificationService) NotifyUser(userId int, message string) {
	u := User{}

	user, err := u.GetUserById(userId)

	if err != nil {
		fmt.Println("Ошибка получения пользователя")
	}

	if n.notifire == nil {
		fmt.Println("Не выбран провайдер уведомлений!")
		return
	}


	err = n.notifire.Send(*user, message)

	if err != nil {
		fmt.Println("Ошибка отправки:", err)
	}
}
