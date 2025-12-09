package notifier

type Notifire interface {
	Send(user User, message string) error
}
