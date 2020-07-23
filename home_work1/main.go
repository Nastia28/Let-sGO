package main

import "fmt"

type contact struct {
	email string
	phone string
}
type worker struct {
	name string
	age int
	contact
}

func main() {
	var anastasia = worker{
		name: "Анастасія Берещенко",
		age: 20,
		contact: contact{
			email: "nasnia686@gmail.com",
			phone: "+380501239999",
		},
	}
	var valet = worker{
		name:    "Валентин Янковский",
		age:     21,
		contact: contact{
			email: "val@gmail.com",
			phone: "+380988046332",
		},
	}

	fmt.Println("Ім'я:", anastasia.name)
	fmt.Println("Вік:", anastasia.age)
	fmt.Println("Email:", anastasia.contact.email)
	fmt.Println("Номер телефону:", anastasia.phone)
	fmt.Println("\nІм'я:", valet.name, "\nВік:", valet.age, "\nEmail:", valet.email, "\nНомер телефону:", valet.phone)

}
