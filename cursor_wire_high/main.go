package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

// Monster 怪物结构
type Monster struct {
	Name string
}

// NewMonster 创建一个新的怪物
func NewMonster() Monster {
	return Monster{Name: "kitty"}
}

// Player 勇士结构
type Player struct {
	Name string
}

// NewPlayer 创建一个新的勇士
func NewPlayer(name string) (Player, error) {
	if time.Now().Unix()%2 == 0 {
		return Player{}, errors.New("勇士已死亡")
	}
	return Player{Name: name}, nil
}

// Mission 任务结构
type Mission struct {
	Player  Player
	Monster Monster
}

// NewMission 创建一个新的任务
func NewMission(p Player, m Monster) Mission {
	return Mission{p, m}
}

// Start 开始任务
func (m Mission) Start() {
	fmt.Printf("%s 击败了 %s，世界和平了！\n", m.Player.Name, m.Monster.Name)
}

func main() {
	// 使用wire自动注入依赖
	mission, err := InitMission("dj")
	if err != nil {
		log.Fatal(err)
	}
	mission.Start()
}
