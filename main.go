package main

import "study/notifier"

func main() {
	emailNotifier := notifier.NewEmailNotification()
	service := notifier.NotificationService{}

	service.SetNotifire(emailNotifier)
	service.NotifyUser(1, "just email notifier")
}
