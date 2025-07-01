package main

import "fmt"

type User struct {
	name string
	age  int
}

func updateUser(u *User, newName string, newAge int) {
	u.name = newName
	u.age = newAge
}

func main() {
	user := User{name: "Lati", age: 18}
	fmt.Println("Before update:", user)

	updateUser(&user, "Zahra", 20)
	fmt.Println("After update:", user)
}
