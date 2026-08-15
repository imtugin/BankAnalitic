package main

import "fmt"

type preparation interface {
	steeping()
	inATermos()
}

type Range struct {
	Min int
	Max int
}
type time struct {
	Min int
	Max int
}
type temp struct {
	Min int
	Max int
}
type OolungMilk struct {
	water        string
	rinsing      bool
	amount       Range
	temperature  temp
	steepingTime []time
} // Улун (молочный)
type GreenLongjing struct {
	water        string
	rinsing      bool
	amount       Range
	temperature  int
	steepingTime time
} // Зелёный чай (Лунцзин)
type RedJinJunMei struct {
	water        string
	rinsing      bool
	amount       Range
	temperature  int
	steepingTime time
} // Красный чай (Цзинь Цзюнь Мэй)
type Puer struct {
	water        string
	rinsing      bool
	amount       Range
	temperature  int
	steepingTime time
} // Пуэр (шу, обычный)
type PuerChaTouBigHeads struct {
	water        string
	rinsing      bool
	amount       Range
	temperature  int
	steepingTime time
} // Пуэр Ча Тоу "Большие головы"
type PuerTheMostInvigorating struct {
	water        string
	rinsing      bool
	amount       Range
	temperature  int
	steepingTime time
} // Пуэр "Бодрейший"
type WhiteTinSong struct {
	water        string
	rinsing      bool
	amount       Range
	temperature  int
	steepingTime time
} // Белый чай (Тын Сон)
type Herbal struct {
	water        string
	rinsing      bool
	amount       Range
	temperature  int
	steepingTime time
} // Травяной чай
func (tea OolungMilk) steeping() {
	rising := "не нужен"
	//var	numberOfSteepings int
	//for i := 1; i < len(tea.steepingTime); i++{
	//	numberOfSteepings += i
	//}
	if tea.rinsing == true {
		rising = "нужен"
	}
	fmt.Printf("Вода:%s\n Промывочный пролив:%s\n Закладка: %d - %d\n Температура: %d - %d\n", tea.water, rising, tea.amount.Min, tea.amount.Max, tea.temperature.Min, tea.temperature.Max)
	fmt.Println("Проливы")
	for intdex, value := range tea.steepingTime {
		fmt.Printf("%d. %d - %d\n", intdex, value.Min, value.Max)
	}
}
func (tea OolungMilk) inATermos() {
	fmt.Println("В разработке")
}
func main() {

	var oolung OolungMilk = OolungMilk{
		water:        "Святой источник, Архыз, Шишкин лес",
		rinsing:      false,
		amount:       Range{3, 4},
		temperature:  temp{70, 75},
		steepingTime: []time{{20, 25}, {20, 25}, {30, 35}},
	}
	var tea preparation = oolung
	tea.steeping()
	// var greenLongjing GreenLongjing = GreenLongjing{rinsing: false}
	// var rRedJinJunMei edJinJunMei = RedJinJunMei{rinsing: true}
	// var puer Puer = Puer{rinsing: true}
	// var puerChaTouBigHeads PuerChaTouBigHeads = PuerChaTouBigHeads{rinsing: true}
	// var puerTheMostInvigorating PuerTheMostInvigorating = PuerTheMostInvigorating{rinsing: true}
	// var whiteTinSong WhiteTinSong = WhiteTinSong{rinsing: false}
	// var herbal Herbal = Herbal{rinsing: false}
}
