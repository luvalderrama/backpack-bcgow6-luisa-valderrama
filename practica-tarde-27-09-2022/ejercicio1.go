package main

import "fmt"

type User struct {
	name string
	lastname string
	age int
	email string
	password string
}

func (u *User) UpdateName(name, lastname string) {
	u.name = name
	u.lastname = lastname
}

func (u *User) UpdateAge(age int) {
	u.age = age
}

func (u *User) UpdateEmail(email string) {
	u.email = email
}

func (u *User) UpdatePassword(password string) {
	u.password = password
}

func main() {
	user := User{
		name: "luisa",
		lastname: "valderrama",
		age: 18,
		email: "luvalde@gmail.com",
		password: "jsjdkk",
	}
	fmt.Println(user)
	user.UpdateName("juan", "pablo")
	fmt.Println(user)
}