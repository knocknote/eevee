// Code generated by eevee. DO NOT EDIT!

package model

import (
	"api/entity"
	"api/model"
	"context"
)

func DefaultUserField() *model.UserField {
	value := model.NewUserField(&entity.UserField{
		FieldID: uint64(0x0),
		ID:      uint64(0x0),
		UserID:  uint64(0x0),
	}, nil)
	value.Field = func(context.Context) (*model.Field, error) {
		return DefaultField(), nil
	}
	return value
}

func DefaultsUserField() *model.UserFields {
	values := model.NewUserFields(entity.UserFields{})
	values.Add(DefaultUserField())
	return values
}
