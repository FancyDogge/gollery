package main

import (
	"html/template"
	"os"
)

type User struct {
	Name string
	Age  int
	Meta UserMeta
}

type UserMeta struct {
	Visits int
}

func main() {
	t, err := template.ParseFiles("hello.gohtml")
	if err != nil {
		panic(err)
	}

	user := User{
		Name: "Sichik",
		Age:  33,
		Meta: UserMeta{ //в темплейте мы достаем по .Meta.Visits, а не .Meta.UserMeta.Visits, потому что UserMeta - это просто тип стракта, а не поле
			Visits: 10,
		},
	}

	err = t.Execute(os.Stdout, user)
	if err != nil {
		panic(err)
	}
}
