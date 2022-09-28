package main

import "fmt"

type Event interface {
	Register(EventListener)
}

type EventListener interface {
	OnFire()
}

type Mail struct {
	listener EventListener
}

func (m *Mail) Register(listener EventListener) {
	m.listener = listener
}

func (m *Mail) OnRecv() {
	m.listener.OnFire()
}

type Alarm struct {
}

func (a *Alarm) OnFire() {
	// 알람
	fmt.Println("알람! 알람!")
}

var mail = &Mail{}
var listener EventListener = &Alarm{}

func main() {
	mail.Register(listener)
	mail.OnRecv() // 4. 알람이 울리게 됩니다.
}
