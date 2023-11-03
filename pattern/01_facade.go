package pattern

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/* "Фасад" -- паттерн, предоставляюший простой интерфейс для взаимодействия со сложной системой.*/

import "fmt"

// Cat - структура представляющая котика
type Cat struct {
	name string
}

func NewCat(name string) *Cat {
	return &Cat{name}
}

func (c *Cat) Meow() {
	fmt.Printf("%s: Мяу-мяу!\n", c.name)
}

func (c *Cat) Sleep() {
	fmt.Printf("%s: Засыпает...\n", c.name)
}

func (c *Cat) WakeUp() {
	fmt.Printf("%s: Просыпается!\n", c.name)
}

// CatFacade - фасад для управления котиками
type CatFacade struct {
	cats []*Cat
}

func NewCatFacade() *CatFacade {
	return &CatFacade{}
}

func (cf *CatFacade) AddCat(name string) {
	cat := NewCat(name)
	cf.cats = append(cf.cats, cat)
}

func (cf *CatFacade) MakeAllCatsMeow() {
	fmt.Println("== Фасад: Все котики мяукают:")
	for _, cat := range cf.cats {
		cat.Meow()
	}
}

func (cf *CatFacade) PutAllCatsToSleep() {
	fmt.Println("== Фасад: Все котики идут спать:")
	for _, cat := range cf.cats {
		cat.Sleep()
	}
}

func (cf *CatFacade) WakeUpAllCats() {
	fmt.Println("== Фасад: Все котики просыпаются:")
	for _, cat := range cf.cats {
		cat.WakeUp()
	}
}

//func main() {
//	catFacade := NewCatFacade()
//
//	// Добавляем несколько котиков
//	catFacade.AddCat("Беня")
//	catFacade.AddCat("Нарния")
//	catFacade.AddCat("Муся")
//
//	// Используем фасад для управления котиками
//	catFacade.MakeAllCatsMeow()
//	catFacade.PutAllCatsToSleep()
//	catFacade.WakeUpAllCats()
//}
