package main

import "fmt"

type animal interface {
	say()
}
type iDontKnow interface {
	Eat()
}
type cat struct{}
type dog struct{}

func (c cat) say() {
	fmt.Println("Мяу")
}
func (d dog) say() {
	fmt.Println("Гав")
}
func (c cat) Eat() {
	fmt.Println("Ест сметану")
}
func (d dog) Eat() {
	fmt.Println("Грызёт кость")
}
func main() {
	var eating = []iDontKnow{cat{}, dog{}}
	var anim = []animal{cat{}, dog{}}
}
