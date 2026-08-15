package main

//Функциональность:
//- Хранить книги в срезе из строк
//- Добавлять книгу (ввод с клавиатуры)
//- Показывать все книги с красивым выводом - с указанием прочитана или нет
//- Искать книгу по названию и автору
//- Отметить книгу как прочитанную по индексу
//- Удалять книгу по индексу
//❌
//✅

import (
	"fmt"
	"strings"
)

type Book struct {
	ID      int
	Title   string
	Aujthor string
	Year    int
	Read    bool
}

func main() {
	var books []Book
	var nextID = 1

	for {
		fmt.Println("\n📚 Каталог книг")
		fmt.Println("1. Добавить книгу")
		fmt.Println("2. Показать все книги")
		fmt.Println("3. Найти книгу")
		fmt.Println("4. Отметить как прочитанную")
		fmt.Println("5. Удалить книгу")
		fmt.Println("6. Выйти")
		fmt.Print("Выберите действие: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			addBook(&books, &nextID)
		case 2:
			listBooks(books)
		case 3:
			findBook(books)
		case 4:
			markAsRead(&books)
		case 5:
			deleteBook(&books)
		case 6:
			fmt.Println("До свидания!")
			return
		default:
			fmt.Println("Неверный выбор.")
		}
	}
}

func bookString(book Book) string {
	status := "❌"
	if book.Read {
		status = "✅"
	}
	return fmt.Sprintf("%d. %s | %s | %d | %s", book.ID, book.Title, book.Author, book.Year, status)
}

func addBook(books *[]Book, nextID *int) {
	var title, author string
	var year int

	fmt.Print("Название: ")
	fmt.Scanln(&title)
	fmt.Print("Автор: ")
	fmt.Scanln(&author)
	fmt.Print("Год: ")
	fmt.Scanln(&year)

	book := Book{
		ID:     *nextID,
		Title:  title,
		Author: author,
		Year:   year,
		Read:   false,
	}

	*books = append(*books, book)
	*nextID++
	fmt.Printf("Книга '%s' добавлена с ID %d!\n", title, book.ID)
}

func listBooks(books []Book) {
	if len(books) == 0 {
		fmt.Println("Книг пока нет.")
		return
	}
	for _, book := range books {
		fmt.Println(bookString(book))
	}
}

func findBook(books []Book) {
	var enter string
	fmt.Print("Введи название, либо атора книги: ")
	fmt.Scanln(&enter)

	found := false
	for _, book := range books {
		if strings.Contains(strings.ToLower(book.Title), strings.ToLower(enter)) || strings.Contains(strings.ToLower(book.Author), strings.ToLower(enter)) {
			fmt.Println(bookString(book))
			found = true
		}
	}
	if !found {
		fmt.Println("Книга не найдена.")
	}
}

func markAsRead(books *[]Book) { // создали функцию, которая в качестве параметров принимает адрес производного типа, который я сам создал, который из себя представляет срез []Book
	//
	var readedID int
	fmt.Println("Какую книгу отметить прочитанной, введите её ID?")
	fmt.Scanln(&readedID)
	for i := range *books {
		if (*books)[i].ID == readedID {
			(*books)[i].Read = true
			fmt.Println("Книга отмечена прочитаной")
			return
		}
	}
	fmt.Println("Книга с таким ID не найдена.")
}

func deleteBook(books *[]Book) {
	var delID int
	fmt.Println("Какую книгу удалить? введите ID.")
	fmt.Scanln(&delID)
	for i := range *books {
		if (*books)[i].ID == delID {
			*books = append((*books)[:i], (*books)[i+1:]...)
			fmt.Println("Книга удалена")
			return
		}
	}
	fmt.Println("Книга с таким ID не найдена")
}
