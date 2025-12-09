package notifier

import "fmt"

type User struct {
	ID        int
	Email     string
	Phone     string
	PushToken string
}

// просто 1 пользователь
func (u User) GetUserById(id int) (*User, error) {
	if id == 1 {
		return &User{
			ID:        1,
			Email:     "some@mail.ru",
			Phone:     "85342977532",
			PushToken: "fjhda",
		}, nil
	}
	return nil, fmt.Errorf("Пользователь не наиден")
}
