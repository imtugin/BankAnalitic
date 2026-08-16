package main

import (
	"fmt"
	"math/rand"
)

type Player interface {
	kickball()
}
type FootballPlayer struct {
	power   int
	stamina int
}
type Messi struct {
	power   int
	stamina int
	height  float64
}

func (f Messi) kickball() {
	shot := f.power + f.stamina
	shot *= 2
	fmt.Printf("Гном бьёт по мячу %d, его рост %.2f", shot, f.height)
}
func (f FootballPlayer) kickball() {
	shot := f.power + f.stamina
	fmt.Printf("Я бью по мячу %d\n", shot)
}
func main() {
	players := make([]Player, 11) // в этом срезе каждый элемент поинтер, указывающий на данные из структуры и так же данные о том, как эта структура реализует интерфейс
	for p := range players {
		players[p] = FootballPlayer{
			power:   rand.Intn(10),
			stamina: rand.Intn(10),
		}
		if p == len(players)-1 {
			players[p] = Messi{10, 10, 1.2}
		}
	}
	for p := range players {
		players[p].kickball()
	}
	// Messi.kickball(Messi{})
}
