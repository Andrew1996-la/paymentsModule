package notifier

import "fmt"

type EmailNotification struct{}

func (e EmailNotification) Send(user User, message string) error {
	fmt.Printf("[EMAIL NOTIFICATION]\nto %s,\nmessage: %s", user.Email, message)
	return nil
}

type PushNotification struct{}

func (p PushNotification) Send(user User, message string) error {
	fmt.Printf("[PUSH NOTIFICATION]\nto %s,\nmessage: %s", user.PushToken, message)
	return nil
}

type SMSNotification struct{}

func (s SMSNotification) Send(user User, message string) error {
	fmt.Printf("[SMS NOTIFICATION]\nto %s,\nmessage: %s", user.Phone, message)
	return nil
}

func NewEmailNotification()*EmailNotification {
	return &EmailNotification{}
}

func NewPushNotification()*PushNotification {
	return &PushNotification{}
}
func NewSmsNotification()*SMSNotification {
	return &SMSNotification{}
}