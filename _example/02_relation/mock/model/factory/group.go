// Code generated by eevee. DO NOT EDIT!

package factory

import (
	"relation/entity"
	"relation/model"
)

func DefaultGroup() *model.Group {
	value := &model.Group{Group: &entity.Group{
		ID:   0,
		Name: "",
	}}
	return value
}

func DefaultGroups() *model.Groups {
	values := &model.Groups{}
	{
		value := &model.Group{Group: &entity.Group{
			ID:   0,
			Name: "",
		}}
		values.Add(value)
	}
	return values
}
