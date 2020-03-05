package main

import (
	"fmt"

	"github.com/google/wire"
)

//Monster 定义了恶魔结构体
type Monster struct {
	Name string
}

//NewMonster  定义恶魔的构建函数
func NewMonster() Monster {
	return Monster{Name: "Dog-Danny"}
}

//Hero 定义了英雄结构体
type Hero struct {
	Name string
}

//NewHero Hero的构造函数
func NewHero() Hero {
	return Hero{Name: "Hunter"}
}

//Mission 定义了“使命”结构体
type Mission struct {
	Monster Monster
	Hero    Hero
}

//NewMission 定义了Mission的构造函数
func NewMission(h Hero, m Monster) Mission {
	return Mission{m, h}
}

//InitMisson 定义了Mission的构造函数,使用了wire工程
func InitMisson() Mission {
	wire.Build(NewHero, NewMonster, NewMission)
	return Mission{}
}

func (m Mission) start() {
	fmt.Printf("%s defeats %s ,the world peace ", m.Hero.Name, m.Monster.Name)
}

func main() {
	monster := NewMonster()
	hero := NewHero()
	misson := NewMission(hero, monster)
	misson.start()

	im := InitMisson()
	im.start()
}
