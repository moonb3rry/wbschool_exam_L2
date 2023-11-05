package pattern

import (
	"errors"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/
/* «Фабричный метод» -- порождающий паттерн проектирования, который предлагает определить интерфейс для создания объекта,
но позволяет подклассам изменять тип создаваемого объекта, передавая создание объектов подклассам.
В Go "фабричный метод" может быть реализован через определение интерфейса для создания объекта и
структур, реализующих этот интерфейс.*/

// ProductType -- перечисление для разных типов продуктов
type ProductType int

const (
	ProductTypeA ProductType = iota
	ProductTypeB
)

// Product -- интерфейс для продуктов
type Product interface {
	Use() string
}

// Creator -- интерфейс с фабричным методом
type Creator interface {
	CreateProduct(pt ProductType) (Product, error)
}

// ConcreteCreator -- конкретная реализация интерфейса Creator
type ConcreteCreator struct{}

// CreateProduct -- фабричный метод для ConcreteCreator
func (cc *ConcreteCreator) CreateProduct(pt ProductType) (Product, error) {
	switch pt {
	case ProductTypeA:
		return &ConcreteProductA{}, nil
	case ProductTypeB:
		return &ConcreteProductB{}, nil
	default:
		return nil, errors.New("Тип продукта не поддерживается")
	}
}

// ConcreteProductA --конкретная реализация Product
type ConcreteProductA struct{}

func (pa *ConcreteProductA) Use() string {
	return "Используем продукт А"
}

// ConcreteProductB -- еще одна реализация Product
type ConcreteProductB struct{}

func (pb *ConcreteProductB) Use() string {
	return "Используем продукт B"
}

//func main() {
//	creator := &ConcreteCreator{}
//	productA, err := creator.CreateProduct(ProductTypeA)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(productA.Use()) // Используем продукт А
//
//	productB, err := creator.CreateProduct(ProductTypeB)
//	if err != nil {
//		fmt.Println(err)
//		return
//	}
//	fmt.Println(productB.Use()) // Используем продукт B
//
//	unsupportedProduct, err := creator.CreateProduct(99)
//	if err != nil {
//		fmt.Println(err) // Тип продукта не поддерживается
//		return
//	}
//	fmt.Println(unsupportedProduct.Use())
//}
