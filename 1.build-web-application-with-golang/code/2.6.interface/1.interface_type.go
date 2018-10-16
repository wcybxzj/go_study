package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
	loan   float32
}

type Employee struct {
	Human
	company string
	money   float32
}

//Human
func (h *Human) SayHi() {
	fmt.Printf("I am %s, my phone:%d\n", h.name, h.phone)
}

func (h *Human) Sing(word string) {
	fmt.Println("la,la,la,la,la", word)
}

func (h *Human) Guzzle(beer string) {
	fmt.Println("Guzzle Guzzle", beer)
}

//Employee
func (e *Employee) SayHi() {
	fmt.Printf("Hi I am %s, I work at %s, My phone %s\n",
		e.name, e.company, e.phone)
}

func (e *Employee) SpendSalary(amount float32) {
	e.money -= amount
}

//Student
func (s *Student) BorrowMoney(amount float32) {
	s.loan += amount
}

type Men interface {
	SayHi()
	Sing(word string)
	Guzzle(beer string)
}

type YoungChap interface {
	SayHi()
	Sing(word string)
	BorrowMoney(amount float32)
}

type ElderGent interface {
	SayHi()
	Sing(word string)
	SpendSalary(amount float32)
}

func main() {

}
