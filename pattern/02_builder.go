package pattern

import (
	"errors"
)

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/* Паттерн "строитель" предоставляет способ создания составного объекта опосредованно. */

// DessertBuilder интерфейс для строителя десерта
type DessertBuilder interface {
	SetBase() error
	AddTopping() error
	GetDessert() (*Dessert, error)
}

// Dessert -- структура представляющая десерт
type Dessert struct {
	Base    string
	Topping string
}

// IceCreamBuilder -- конкретный строитель для мороженого
type IceCreamBuilder struct {
	dessert *Dessert
}

func NewIceCreamBuilder() DessertBuilder {
	return &IceCreamBuilder{&Dessert{}}
}

func (b *IceCreamBuilder) SetBase() error {
	b.dessert.Base = "Мороженое"
	return nil
}

func (b *IceCreamBuilder) AddTopping() error {
	b.dessert.Topping = "карамельным сиропом"
	return nil
}

func (b *IceCreamBuilder) GetDessert() (*Dessert, error) {
	if b.dessert.Base == "" || b.dessert.Topping == "" {
		return nil, errors.New("невозможно создать десерт: не все компоненты заданы")
	}
	return b.dessert, nil
}

// Director -- директор, который управляет процессом сборки десерта
type Director struct {
	builder DessertBuilder
}

func NewDirector(builder DessertBuilder) *Director {
	return &Director{builder}
}

func (d *Director) Construct() (*Dessert, error) {
	err := d.builder.SetBase()
	if err != nil {
		return nil, err
	}
	err = d.builder.AddTopping()
	if err != nil {
		return nil, err
	}
	return d.builder.GetDessert()
}

//func main() {
//	iceCreamBuilder := NewIceCreamBuilder()
//	director := NewDirector(iceCreamBuilder)
//
//	dessert, err := director.Construct()
//	if err != nil {
//		fmt.Printf("Ошибка при создании десерта: %v\n", err)
//		return
//	}
//
//	fmt.Printf("Десерт готов: %s с %s\n", dessert.Base, dessert.Topping)
//}
