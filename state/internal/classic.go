package internal

import "fmt"

type Switch struct {
	State State
}

func (s *Switch) On() {
	s.State.On(s)
}

func (s *Switch) Off() {
	s.State.Off(s)
}

func NewSwitch() *Switch {
	return &Switch{
		State: NewOffState(),
	}
}

type State interface {
	On(s *Switch)
	Off(s *Switch)
}

type BaseState struct{}

func (b *BaseState) On(s *Switch) {
	fmt.Println("Light is already on")
}

func (b *BaseState) Off(s *Switch) {
	fmt.Println("Light is already off")
}

type OnState struct {
	BaseState
}

func NewOnState() *OnState {
	fmt.Println("Light turned on")
	return &OnState{
		BaseState: BaseState{},
	}
}

func (o *OnState) Off(s *Switch) {
	fmt.Println("Turning the light off...")
	s.State = NewOffState()
}

type OffState struct {
	BaseState
}

func NewOffState() *OffState {
	fmt.Println("Light turned off")
	return &OffState{
		BaseState: BaseState{},
	}
}

func (o *OffState) On(s *Switch) {
	fmt.Println("Turning the light on...")
	s.State = NewOnState()
}

func RunClassic() {
	s := NewSwitch()
	s.On()
	s.Off()
	s.Off()
}
