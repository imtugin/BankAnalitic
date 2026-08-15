package main

import "fmt"

// Функциональность:
// - Хранить книги в срезе из строк
// - Добавлять книгу (ввод с клавиатуры)
// - Показывать все книги с красивым выводом - с указанием прочитана или нет
// - Искать книгу по названию и автору
// - Отметить книгу как прочитанную по индексу
// - Удалять книгу по индексу
// ❌
// ✅
type Book struct {
	ID     int
	Title  string
	Author string
	Year   int
}

func main() {
	var action int
	var book []Book
	var book1 Book = Book{}
	fmt.Println("Добро пожаловать в библиотеку книг.")
	fmt.Println("Введите желаемое действие из предложенных")
	fmt.Println("1 - Добавить книгу")
	fmt.Println("2 - Показать книги")
	fmt.Println("3 - Отметить книгу прочитанной по ID")
	fmt.Println("4 - Удалить книгу по ID")
	fmt.Println("5 - Искать книгу по названию или автору")
	fmt.Println("6 - выйти")
	fmt.Scanln(&action)
	switch action {
	case 1:
		addBook()
	case 2:
		fmt.Println(book)
	}
}
func addBook() {

}
