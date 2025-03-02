//go:build wireinject
// +build wireinject

package main

import "github.com/google/wire"

// InitMission 注入器，用于创建并注入依赖
func InitMission(name string) (Mission, error) {
	wire.Build(NewMonster, NewPlayer, NewMission)
	return Mission{}, nil
}
