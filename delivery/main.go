package main

import "fmt"

type Delivery interface {
	Deliver() string
	EstimateTime() int
}
type Car struct {
	Distance float64
}
type Bike struct {
	Distance float64
}
type Pedestrian struct {
	Distance float64
}

func (d Car) Deliver() string { // Возвращает строку с описанием доставки
	return fmt.Sprintf("Доставка на машине: %.2f км", d.Distance)
}
func (d Car) EstimateTime() int { //Возвращает время доставки в минутах
	return int(d.Distance / 40 * 60)
}
func (d Bike) Deliver() string { // Возвращает строку с описанием доставки
	return fmt.Sprintf("Доставка на велосипеде: %.2f км", d.Distance)
}
func (d Bike) EstimateTime() int { //Возвращает время доставки в минутах

	return int(d.Distance / 15 * 60)
}
func (d Pedestrian) Deliver() string { // Возвращает строку с описанием доставки
	return fmt.Sprintf("Доставка пешая: %.2f км", d.Distance)
}
func (d Pedestrian) EstimateTime() int { //Возвращает время доставки в минутах
	return int(d.Distance / 5 * 60)
}

func main() {
	var car Car = Car{12}
	var bike Bike = Bike{3}
	var per Pedestrian = Pedestrian{0.8}
	orders := make([]Delivery, 3)
	orders[0] = &car
	orders[1] = &bike
	orders[2] = &per
	for i := range orders {
		fmt.Printf("%s за %d минут\n", orders[i].Deliver(), orders[i].EstimateTime())
	}
}
