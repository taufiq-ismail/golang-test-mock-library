package user

import "example/mock-library/activity"

type User struct {
	Activity activity.Activity
}

func (u *User) Use() error {
	return u.Activity.DoSomething(1, "Hello World")
}
