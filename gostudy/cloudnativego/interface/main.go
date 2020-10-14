package main

import (
	"fmt"
)

type attacker struct {
	apower int
	hp     int
}

type sword struct {
	attacker
	color string
}

type gun struct {
	attacker
	short bool
}

func (at sword) wield() bool {
	fmt.Println("this is a sowrd!")
	return true
}

func (gu gun) wield() bool {
	fmt.Println("this is a gun")
	return true
}

type weapon interface {
	wield() bool
}

func wielder(we weapon) bool {
	fmt.Println("Wielding......")
	return we.wield()
}

func main() {
	gun1 := gun{attacker: attacker{apower: 10, hp: 100}, short: false}
	sword1 := sword{attacker: attacker{apower: 9,
		hp: 120}, color: "green"}
	wielder(gun1)
	isno := wielder(sword1)
	fmt.Println(isno)
}
