// Code generated by eevee. DO NOT EDIT!

package model

import (
	"daoplugin/entity"
	"daoplugin/model"
)

func DefaultWorld() *model.World {
	value := model.NewWorld(&entity.World{
		ID:   1,
		Name: "",
	}, nil)
	return value
}

func DefaultsWorld() *model.Worlds {
	values := model.NewWorlds(entity.Worlds{})
	values.Add(DefaultWorld())
	return values
}
