package main

import (
	"fmt"
	"os"
	"text/template"
	"math/rand"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	t, err := template.ParseFiles("./index.html")
	if err != nil {
		fmt.Println("parse file err:", err)
		return
	}
	var personList []Person
	for i := 0; i < 10; i++ {
		p := Person{Name: fmt.Sprintf("Mary %d", i), Age: int(rand.Int31())%100}
		personList = append(personList, p)
	}
	if err := t.Execute(os.Stdout, personList); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
