package textbasics

import (
	"fmt"
)

type UserHeight struct {
	feet   int
	inches int
}

type User struct {
	firstName string
	lastName  string
	age       int
	height    UserHeight
}

type Laughable interface {
	laugh(about string)
}

func NewUser(firstname, lastname string, age int, feet, inches int) *User {
	return &User{
		firstName: firstname,
		lastName:  lastname,
		age:       age,
		height: UserHeight{
			feet:   feet,
			inches: inches,
		},
	}
}

func (u *User) sayHello() {
	msg := fmt.Sprintf("Hello my nameis %s and I am %d foot %d inch(es) tall.", u.firstName, u.height.feet, u.height.inches)
	fmt.Println(msg)
}

func Run() {
	u := NewUser("Nnaemeka", "Doe", 23, 6, 1)
	u.sayHello()
}
