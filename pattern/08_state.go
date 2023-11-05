package pattern

import (
	"errors"
	"fmt"
)

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/* Паттерн "состояние" позволяет менять поведение класса при изменении состояния.
В отличие от "стратегии", здесь автоматически меняется поведение, в то время как в
паттерне "стратегия" клиент выбирает требуемую релаизацию. Также состояния могут быть
взаимосвязаны, а стратегии независимы друг от друга.*/

// LightState интерфейс для различных состояний выключателя
type LightState interface {
	On(lightSwitch *LightSwitch) error
	Off(lightSwitch *LightSwitch) error
}

// LightSwitch структура выключателя с текущим состоянием
type LightSwitch struct {
	state LightState
}

// SetState устанавливает новое состояние выключателя
func (ls *LightSwitch) SetState(state LightState) {
	ls.state = state
}

// On включает выключатель
func (ls *LightSwitch) On() error {
	return ls.state.On(ls)
}

// Off выключает выключатель
func (ls *LightSwitch) Off() error {
	return ls.state.Off(ls)
}

// OnState состояние включенного выключателя
type OnState struct{}

func (on *OnState) On(lightSwitch *LightSwitch) error {
	return errors.New("Свет уже включен")
}

func (on *OnState) Off(lightSwitch *LightSwitch) error {
	lightSwitch.SetState(&OffState{})
	fmt.Println("Выключаю свет")
	return nil
}

// OffState состояние выключенного выключателя
type OffState struct{}

func (off *OffState) On(lightSwitch *LightSwitch) error {
	lightSwitch.SetState(&OnState{})
	fmt.Println("Включаю свет")
	return nil
}

func (off *OffState) Off(lightSwitch *LightSwitch) error {
	return errors.New("Свет уже выключен")
}

//func main() {
//	lightSwitch := &LightSwitch{state: &OffState{}}
//
//	err := lightSwitch.On()
//	if err != nil {
//		fmt.Println(err)
//	} // Включаю свет
//
//	err = lightSwitch.On()
//	if err != nil {
//		fmt.Println(err)
//	} // Свет уже включен
//
//	err = lightSwitch.Off()
//	if err != nil {
//		fmt.Println(err)
//	} // Выключаю свет
//}
