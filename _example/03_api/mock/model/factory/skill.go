// Code generated by eevee. DO NOT EDIT!

package factory

import (
	"api/entity"
	"api/model"
)

func DefaultSkill() *model.Skill {
	value := &model.Skill{Skill: &entity.Skill{
		ID:          0,
		SkillEffect: "",
	}}
	return value
}

func DefaultSkills() *model.Skills {
	values := &model.Skills{}
	{
		value := &model.Skill{Skill: &entity.Skill{
			ID:          0,
			SkillEffect: "",
		}}
		values.Add(value)
	}
	return values
}
