package main

import "fmt"

type Activity interface {
	Name() string
	Calories() int
	Duration() string
	Intensity() string
}

type Running struct {
	Distance float64 // км
	TimeMin  int
}
type Swimming struct {
	Distance float64 // км
	TimeMin  int
}
type Strength struct {
	Exercises []string
	TimeMin   int
}

func (r Running) Name() string {
	return "Бег"
}
func (s Swimming) Name() string {
	return "Плавание"
}

func (s Strength) Name() string {
	return "Растяжка"
}

func (r Running) Calories() int {
	return int(r.Distance * 60)
}
func (s Swimming) Calories() int {
	return int(s.Distance * 400)

}
func (s Strength) Calories() int {
	return s.TimeMin * 8
}
func (r Running) Duration() string {
	return fmt.Sprintf("%d мин", r.TimeMin)
}
func (s Swimming) Duration() string {
	return fmt.Sprintf("%d мин", s.TimeMin)
}
func (s Strength) Duration() string {
	return fmt.Sprintf("%d мин", s.TimeMin)
}
func (r Running) Intensity() string {
	// если темп меньше 6 мин на км, то высокая. Иначе - средняя
	var temp float64 = float64(r.TimeMin) / r.Distance // за соклько времени я пробежал один км в среднем
	if temp < 6.0 {
		return "высокая"
	} else {
		return "средняя"
	}
}
func (s Swimming) Intensity() string {
	if s.Distance >= 2 {
		return "высокая"
	} else {
		return "средняя"
	}

}
func (s Strength) Intensity() string {
	if len(s.Exercises) >= 5 {
		return "высокая"
	} else {
		return "средняя"
	}
}
func TotalCalories(activities []Activity) int { // Суммарно сожжённые калории

}
func FilterByIntensity(activites []Activity, intensity string) []Activity { //Только тренировки с укаазанной интенсивностью
	var filter []Activity
	for i := range activites {
		if activites[i].Intensity() == intensity {
			filter = append(filter, activites[i])
		}
	}
	return filter
}
func LongestDuration(activites []Activity) Activity { // Тренировка с максимальным временем
	longest := activites[0]
	for i := range activites {
		if longest.Duration() < activites[i].Duration() {
			longest = activites[i]
		}
	}
	return longest
}
func PrintSummary(a Activity) { // Выводит: название, длительность, калории, интенсивность
	fmt.Println("Название тренировки:", a.Name())
	fmt.Println("Длительность:", a.Duration())
	fmt.Println("Расход калорий:", a.Calories())
	fmt.Println("Интерсивность:", a.Intensity())
}

func main() {
	// Вывести
	// Все тренировки и их показатели
	// Суммарные калории
	// Только высокоинтенсивные тренировки
	// Самая долгая тренировка
	// PrintSummary для каждой
	//	strength := []string{"шея", "плечи", "спина", "таз", "колени", "лучезапястные суставы", "локтевые суставы", "голеностоп", "пресс", "поза стола"}
	var intens int
	var intensity string
	activ := []Activity{
		Running{5, 24},
		Running{3, 15},
		Running{10, 50},
		Swimming{1, 40},
		Swimming{0.5, 25},
		Strength{[]string{"шея", "плечи", "спина", "колени"}, 10},
		Strength{[]string{"шея", "плечи", "спина", "таз", "колени", "лучезапястные суставы", "локтевые суставы", "голеностоп", "пресс", "поза стола"}, 30},
	}
	fmt.Println("Вот все тренировки:")
	for i := range activ {
		PrintSummary(activ[i])
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("Какая интенсивность тренировки интересует: 1 - Средняя, 2 - Высокая")
	fmt.Scanln(&intens)
	if intens != 1 || intens != 2 {
		fmt.Println("Введите 1 или 2")
	} else {
		if intens == 1 {
			intensity = "средняя"
		} else {
			intensity = "высокая"
		}
	}
	filtr := FilterByIntensity(activ, intensity)
	fmt.Println("Вот тренировки с уровнем активности", intensity, ":")
	for i := range filtr {
		fmt.Print("=======================", i, "=======================")
		PrintSummary(filtr[i])
	}
	fmt.Println()
	fmt.Println()
	fmt.Println("Вот самая долгая тернировка:")
	fmt.Println(LongestDuration(activ))
	fmt.Println()
	fmt.Println()
}
