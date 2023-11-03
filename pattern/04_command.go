package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*Паттерн "Команда" используется для инкапсуляции запроса в виде объекта, позволяя параметризовать
клиентские объекты с операциями, ставить запросы в очередь и поддерживать отмену операций*/

// Command - интерфейс команды
type Command interface {
	Execute() error
}

// Receiver - интерфейс получателя, который выполняет команду
type Receiver interface {
	Cook(food string) error
}

// ConcreteReceiver - конкретный получатель (ресторан)
type ConcreteReceiver struct {
	Name string
}

func NewConcreteReceiver(name string) *ConcreteReceiver {
	return &ConcreteReceiver{Name: name}
}

func (r *ConcreteReceiver) Cook(food string) error {
	fmt.Printf("%s готовит заказ: %s\n", r.Name, food)
	return nil
}

// ConcreteCommand - конкретная команда заказа еды
type ConcreteCommand struct {
	receiver Receiver
	food     string
}

func NewConcreteCommand(receiver Receiver, food string) *ConcreteCommand {
	return &ConcreteCommand{receiver, food}
}

func (c *ConcreteCommand) Execute() error {
	return c.receiver.Cook(c.food)
}

// Invoker - инициатор, который выполняет команды
type Invoker struct {
	commands []Command
}

func NewInvoker() *Invoker {
	return &Invoker{}
}

func (i *Invoker) AddCommand(command Command) {
	i.commands = append(i.commands, command)
}

func (i *Invoker) ExecuteCommands() {
	for _, command := range i.commands {
		err := command.Execute()
		if err != nil {
			fmt.Printf("Ошибка при выполнении команды: %v\n", err)
		}
	}
}

//func main() {
//	// Создаем рестораны
//	restaurant1 := NewConcreteReceiver("Ресторан 1")
//	restaurant2 := NewConcreteReceiver("Ресторан 2")
//
//	// Создаем команды заказа еды
//	command1 := NewConcreteCommand(restaurant1, "Пицца")
//	command2 := NewConcreteCommand(restaurant2, "Суши")
//	command3 := NewConcreteCommand(restaurant1, "Сувлак")
//
//	// Создаем инициатор
//	invoker := NewInvoker()
//
//	// Добавляем команды в очередь
//	invoker.AddCommand(command1)
//	invoker.AddCommand(command2)
//	invoker.AddCommand(command3)
//
//	// Выполняем команды
//	invoker.ExecuteCommands()
//}
