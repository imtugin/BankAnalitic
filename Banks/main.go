package main

import "fmt"

type Banks interface {
	Ozon
	Tinkoff
	VTB
}
type Ozon interface {
}
type Tinkoff interface {
}
type VTB interface {
}

type Deoisut struct { //Вклад на определённый срок
	Annual      float64 // годовая ставка
	Amount      float64
	MinTerm     int // В месяцах минимальный срок, на который можно открыть
	Description string
	// Term Срок в месяцах. Пока не знаю надо или нет
}
type ServingsAccount struct { // Накопительный счёт - снимаем в любой момент
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
	MonthlyIncome      float64
	Income             float64
}

// Изменяем или создаём цифровой актив с учётом срока
func BuyDigitalAsset() DigitalAsset {
	var choice int
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
	// Считаю ставку с учётом вычета. Доход за период вклада
	asset.MonthlyIncome = asset.Amount*(100/asset.Annual)/12 - asset.Tax/12
	asset.Income = asset.MonthlyIncome * float64(choice)
	return asset
}
func main() {
	Banks := make([]Banks, 0) // Тут создали срез пока без элементов
	// То же самое с остальными интерфейсами
	Ozon := make([]Ozon, 0)
	Tinkoff := make([]Tinkoff, 0)
	VTB := make([]VTB, 0)
	var ozonDailyBalance ServingsAccount = ServingsAccount{
		Annual:             11.5,
		AccrOnDailyBalance: true,
	}

	var ozonMonthly ServingsAccount = ServingsAccount{
		Annual:             12.0,
		AccrOnDailyBalance: false,
	}

	var ozonDepozit Deoisut = Deoisut{
		Annual:      11,
		MinTerm:     1,
		Description: "При увеличении срока вклада годовая ставка может доходить до 13.6 %",
	}

	var ozonDigitalAsset DigitalAsset = DigitalAsset{
		Annual:             15.2,
		AnnualIncludingTax: 0, // Это поле будет наверное для наглядности только, нужно придумать как считать
		MinTerm:            1,
		Description:        "При увеличении срока годовая ставка может доходить до 15.2 %. Нужно пересчитать с учётом вычета налогов",
	}
	switch Choice { // При выборе создаётся новый продукт в зависимости от выбора
	case 1:
		// Создаётся цифровой актив
		BuyDigitalAsset()
	}
}
