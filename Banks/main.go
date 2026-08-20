package main

import "fmt"

type Product interface {
	NewProduct()
}

type Deposit struct { //Вклад на определённый срок
	Annual      float64 // годовая ставка
	Amount      float64
	MinTerm     int // В месяцах минимальный срок, на который можно открыть
	Description string
	// Term Срок в месяцах. Пока не знаю надо или нет
}
type SavingsAccount struct { // Накопительный счёт - снимаем в любой момент
	Annual             float64
	Amount             float64
	AccrOnDailyBalance bool
}                          //накопительный счёт. со снятием или без будет в поле структуры. Вернее - начисление процентов на ежедневный останок или на минимальный остаток в месяце
type DigitalAsset struct { // Цифровой актив - пока есть только вроде на озон. Снять вообще не получится
	Annual             float64 // Ставка годовая
	Amount             float64 // Сумма в рублях
	AnnualIncludingTax float64 // Ставка с учётом налогового вычета для сравнения с другими
	Term               int     // Срок в месяцах
	Description        string
	Tax                float64
}

// Изменяем или создаём цифровой актив с учётом срока
func BuyDigitalAsset() []Product {
	var choice int
	var arrProduct []Product
	var asset DigitalAsset = DigitalAsset{
		Tax: 13.0,
	}

	fmt.Println("Покупка цифрового актива, в каком банке? 1 - Ozon(доступен только здесь), 2 - Tinkoff, 3 - VTB")
	fmt.Scanln(&choice)
	switch choice {
	case 1:
		fmt.Println("Покупка ЦА в Ozon")
		// fmt.Println("На текущий момент доступны следующие условия")
		fmt.Println("Выберите срок покупки ЦА")
		fmt.Scanln(&choice)
		fmt.Println("Выберите сумму покупки ЦА")
		fmt.Scanln(&asset.Amount)
		switch choice {
		case 1:
			asset.Annual = 15.2
			asset.Term = 1
		case 3:
			asset.Annual = 15.7
			asset.Term = choice
		case 6:
			asset.Annual = 17
			asset.Term = choice
		case 9:
			asset.Annual = 17.5
			asset.Term = choice
		case 12:
			asset.Annual = 18
			asset.Term = choice
		}

	}
	return append(arrProduct, asset)
}
func NewProduct() {

}
func main() {
	products := make([]Product, 0)

}
