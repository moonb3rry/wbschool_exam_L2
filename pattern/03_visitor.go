package pattern

import "fmt"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/* "Посетитель" позволяет добавлять функционал без необходимости изменять структуру. */

// Cat -- структура представляющая котика
type Cat struct {
	Name string
}

func NewCat(name string) *Cat {
	return &Cat{Name: name}
}

// Visitor -- интерфейс посетителя
type Visitor interface {
	VisitCat(cat *Cat) error
}

// CatHouse -- структура представляющая дом с котиками
type CatHouse struct {
	Cats []*Cat
}

func NewCatHouse() *CatHouse {
	return &CatHouse{}
}

func (ch *CatHouse) AddCat(cat *Cat) {
	ch.Cats = append(ch.Cats, cat)
}

// Accept -- метод для посещения дома котиками
func (ch *CatHouse) Accept(visitor Visitor) error {
	for _, cat := range ch.Cats {
		if err := visitor.VisitCat(cat); err != nil {
			return err
		}
	}
	return nil
}

// Feeder -- конкретный посетитель для кормления котиков
type Feeder struct{}

func NewFeeder() *Feeder {
	return &Feeder{}
}

func (f *Feeder) VisitCat(cat *Cat) error {
	fmt.Printf("Кормим котика %s\n", cat.Name)
	return nil
}

// Vet -- конкретный посетитель для посещения ветеринара
type Vet struct{}

func NewVet() *Vet {
	return &Vet{}
}

func (v *Vet) VisitCat(cat *Cat) error {
	fmt.Printf("Ведем котика %s к ветеринару\n", cat.Name)
	return nil
}

//func main() {
//	catHouse := NewCatHouse()
//	catHouse.AddCat(NewCat("Беня"))
//	catHouse.AddCat(NewCat("Нарния"))
//
//	feeder := NewFeeder()
//	vet := NewVet()
//
//	// Посещаем котиков разными посетителями
//	if err := catHouse.Accept(feeder); err != nil {
//		fmt.Printf("Ошибка при кормлении котиков: %v\n", err)
//	}
//
//	if err := catHouse.Accept(vet); err != nil {
//		fmt.Printf("Ошибка при посещении ветеринара: %v\n", err)
//	}
//}
