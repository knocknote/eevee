// Code generated by eevee. DO NOT EDIT!

package factory

import (
	"api/entity"
	"api/model"
	"context"
)

func DefaultUser() *model.User {
	userFields := DefaultUserFields()
	skill := DefaultSkill()
	world := DefaultWorld()
	value := &model.User{User: &entity.User{
		Age:       0,
		FieldID:   0,
		GroupID:   0,
		ID:        0,
		Name:      "",
		Sex:       "",
		SkillID:   0,
		SkillRank: 0,
		WorldID:   0,
	}}
	value.UserFields = func(context.Context) (*model.UserFields, error) {
		return userFields, nil
	}
	value.Skill = func(context.Context) (*model.Skill, error) {
		return skill, nil
	}
	value.World = func(context.Context) (*model.World, error) {
		return world, nil
	}
	return value
}

func DefaultUsers() *model.Users {
	values := &model.Users{}
	{
		userFields := DefaultUserFields()
		skill := DefaultSkill()
		world := DefaultWorld()
		value := &model.User{User: &entity.User{
			Age:       0,
			FieldID:   0,
			GroupID:   0,
			ID:        0,
			Name:      "",
			Sex:       "",
			SkillID:   0,
			SkillRank: 0,
			WorldID:   0,
		}}
		value.UserFields = func(context.Context) (*model.UserFields, error) {
			return userFields, nil
		}
		value.Skill = func(context.Context) (*model.Skill, error) {
			return skill, nil
		}
		value.World = func(context.Context) (*model.World, error) {
			return world, nil
		}
		values.Add(value)
	}
	return values
}
