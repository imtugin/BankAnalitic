package main

import "fmt"

// Разобрать определение типа интерфейса
type animal interface {
	say()
}
type cat struct{}
type dog struct{}
type snake struct{}

func (c cat) say() {
	fmt.Println("Мяу")
}
func (c dog) say() {
	fmt.Println("Гав")
}
func (c snake) say() {
	fmt.Println("шшш")
}
func main() {
	var Kusia cat = cat{}
	value, ok := Kusia.(animal)
	fmt.Println(ok)
	fmt.Println(value)
}
