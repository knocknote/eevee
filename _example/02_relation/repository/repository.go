// Code generated by eevee. DO NOT EDIT!

package repository

import (
	"context"
	"database/sql"
	"relation/entity"
	"relation/model"
)

type Repository interface {
	Field() Field
	ToField(*entity.Field) *model.Field
	Group() Group
	ToGroup(*entity.Group) *model.Group
	Skill() Skill
	ToSkill(*entity.Skill) *model.Skill
	User() User
	ToUser(*entity.User) *model.User
	UserField() UserField
	ToUserField(*entity.UserField) *model.UserField
	World() World
	ToWorld(*entity.World) *model.World
}

type RepositoryImpl struct {
	field     func() Field
	group     func() Group
	skill     func() Skill
	user      func() User
	userField func() UserField
	world     func() World
}

func (r *RepositoryImpl) Field() Field {
	return r.field()
}

func (r *RepositoryImpl) Group() Group {
	return r.group()
}

func (r *RepositoryImpl) Skill() Skill {
	return r.skill()
}

func (r *RepositoryImpl) User() User {
	return r.user()
}

func (r *RepositoryImpl) UserField() UserField {
	return r.userField()
}

func (r *RepositoryImpl) World() World {
	return r.world()
}

func New(ctx context.Context, tx *sql.Tx) *RepositoryImpl {
	var (
		repo      *RepositoryImpl
		field     *FieldImpl
		group     *GroupImpl
		skill     *SkillImpl
		user      *UserImpl
		userField *UserFieldImpl
		world     *WorldImpl
	)
	repo = &RepositoryImpl{
		field: func() Field {
			if field != nil {
				return field
			}
			field = NewField(ctx, tx)
			field.repo = repo
			return field
		},
		group: func() Group {
			if group != nil {
				return group
			}
			group = NewGroup(ctx, tx)
			group.repo = repo
			return group
		},
		skill: func() Skill {
			if skill != nil {
				return skill
			}
			skill = NewSkill(ctx, tx)
			skill.repo = repo
			return skill
		},
		user: func() User {
			if user != nil {
				return user
			}
			user = NewUser(ctx, tx)
			user.repo = repo
			return user
		},
		userField: func() UserField {
			if userField != nil {
				return userField
			}
			userField = NewUserField(ctx, tx)
			userField.repo = repo
			return userField
		},
		world: func() World {
			if world != nil {
				return world
			}
			world = NewWorld(ctx, tx)
			world.repo = repo
			return world
		},
	}
	return repo
}

func (r *RepositoryImpl) ToField(value *entity.Field) *model.Field {
	return r.Field().ToModel(value)
}

func (r *RepositoryImpl) ToGroup(value *entity.Group) *model.Group {
	return r.Group().ToModel(value)
}

func (r *RepositoryImpl) ToSkill(value *entity.Skill) *model.Skill {
	return r.Skill().ToModel(value)
}

func (r *RepositoryImpl) ToUser(value *entity.User) *model.User {
	return r.User().ToModel(value)
}

func (r *RepositoryImpl) ToUserField(value *entity.UserField) *model.UserField {
	return r.UserField().ToModel(value)
}

func (r *RepositoryImpl) ToWorld(value *entity.World) *model.World {
	return r.World().ToModel(value)
}
