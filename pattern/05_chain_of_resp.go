package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/* Шаблон «Цепочка ответственности» позволяет создавать цепочки объектов. Запрос входит с одного конца цепочки
и движется от объекта к объекту, пока не будет найден подходящий обработчик.*/

// Order представляет заказ на напиток.
type Order struct {
	name     string // Название напитка
	quantity int    // Количество порций
}

// Handler определяет интерфейс обработчика заказов.
type Handler interface {
	handleOrder(order Order) error
	setNext(handler Handler)
}

// Barista - обработчик для приготовления напитков.
type Barista struct {
	nextHandler Handler
}

func (b *Barista) handleOrder(order Order) error {
	if order.name == "Кофе" {
		fmt.Printf("Готовим %d порций %s\n", order.quantity, order.name)
		return b.nextHandler.handleOrder(order)
	}
	return errors.New("Заказ не может быть обработан")
}

func (b *Barista) setNext(handler Handler) {
	b.nextHandler = handler
}

// Cashier - обработчик для приема платежей.
type Cashier struct {
	nextHandler Handler
}

func (c *Cashier) handleOrder(order Order) error {
	if order.name == "Кофе" {
		price := order.quantity * 3
		fmt.Printf("Получена оплата за %d порций %s: $%d\n", order.quantity, order.name, price)
		return c.nextHandler.handleOrder(order)
	}
	return errors.New("Заказ не может быть обработан")
}

func (c *Cashier) setNext(handler Handler) {
	c.nextHandler = handler
}

// Manager - обработчик для управления ресурсами.
type Manager struct {
	nextHandler Handler
}

func (m *Manager) handleOrder(order Order) error {
	if order.name == "Кофе" {
		fmt.Printf("Обеспечиваем ресурсами для приготовления %s\n", order.name)
		return nil
	}
	return errors.New("Заказ не может быть обработан")
}

func (m *Manager) setNext(handler Handler) {
	m.nextHandler = handler
}

//func main() {
//	// Создаем обработчиков
//	barista := &Barista{}
//	cashier := &Cashier{}
//	manager := &Manager{}
//
//	// Устанавливаем порядок обработки заказов
//	barista.setNext(cashier)
//	cashier.setNext(manager)
//
//	// Создаем заказы
//	orders := []Order{
//		{"Чай", 2},
//		{"Кофе", 3},
//		{"Сок", 4},
//	}
//
//	// Обрабатываем заказы
//	for _, order := range orders {
//		if err := barista.handleOrder(order); err != nil {
//			fmt.Printf("Ошибка обработки заказа \"%s\": %s\n", order.name, err)
//		}
//	}
//}
