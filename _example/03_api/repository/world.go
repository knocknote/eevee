// Code generated by eevee. DO NOT EDIT!

package repository

import (
	"api/dao"
	"api/entity"
	"api/model"
	"context"
	"database/sql"

	"golang.org/x/xerrors"
)

type World interface {
	ToModel(*entity.World) *model.World
	ToModels(entity.Worlds) *model.Worlds
	Create(context.Context, *entity.World) (*model.World, error)
	Creates(context.Context, entity.Worlds) (*model.Worlds, error)
	FindAll(context.Context) (*model.Worlds, error)
	FindByID(context.Context, uint64) (*model.World, error)
	FindByIDs(context.Context, []uint64) (*model.Worlds, error)
	UpdateByID(context.Context, uint64, map[string]interface{}) error
	UpdateByIDs(context.Context, []uint64, map[string]interface{}) error
	DeleteByID(context.Context, uint64) error
	DeleteByIDs(context.Context, []uint64) error
	Count(context.Context) (int64, error)
	Delete(context.Context, *entity.World) error
	Update(context.Context, *entity.World) error
}

type WorldImpl struct {
	worldDAO dao.World
	repo     Repository
}

func NewWorld(ctx context.Context, tx *sql.Tx) *WorldImpl {
	return &WorldImpl{worldDAO: dao.NewWorld(ctx, tx)}
}

func (r *WorldImpl) ToModel(value *entity.World) *model.World {
	return r.createCollection(entity.Worlds{value}).First()
}

func (r *WorldImpl) ToModels(values entity.Worlds) *model.Worlds {
	return r.createCollection(values)
}

func (r *WorldImpl) Create(ctx context.Context, value *entity.World) (*model.World, error) {
	if err := r.worldDAO.Create(ctx, value); err != nil {
		return nil, xerrors.Errorf("cannot Create: %w", err)
	}
	v := r.ToModel(value)
	v.SetSavedValue(value)
	v.SetAlreadyCreated(true)
	return v, nil
}

func (r *WorldImpl) Creates(ctx context.Context, entities entity.Worlds) (*model.Worlds, error) {
	for _, v := range entities {
		if _, err := r.Create(ctx, v); err != nil {
			return nil, xerrors.Errorf("cannot Create: %w", err)
		}
	}
	values := r.ToModels(entities)
	values.Each(func(v *model.World) {
		v.SetSavedValue(v.World)
		v.SetAlreadyCreated(true)
	})
	return values, nil
}

func (r *WorldImpl) FindAll(a0 context.Context) (*model.Worlds, error) {
	values, err := r.worldDAO.FindAll(a0)
	if err != nil {
		return nil, xerrors.Errorf("failed to FindAll: %w", err)
	}
	collection := r.createCollection(values)
	collection.Each(func(v *model.World) {
		v.SetSavedValue(v.World)
		v.SetAlreadyCreated(true)
	})
	return collection, nil
}

func (r *WorldImpl) FindByID(a0 context.Context, a1 uint64) (*model.World, error) {
	value, err := r.worldDAO.FindByID(a0, a1)
	if err != nil {
		return nil, xerrors.Errorf("failed to FindByID: %w", err)
	}
	if value == nil {
		return nil, nil
	}
	v := r.createCollection(entity.Worlds{value}).First()
	v.SetSavedValue(v.World)
	v.SetAlreadyCreated(true)
	return v, nil
}

func (r *WorldImpl) FindByIDs(a0 context.Context, a1 []uint64) (*model.Worlds, error) {
	values, err := r.worldDAO.FindByIDs(a0, a1)
	if err != nil {
		return nil, xerrors.Errorf("failed to FindByIDs: %w", err)
	}
	collection := r.createCollection(values)
	collection.Each(func(v *model.World) {
		v.SetSavedValue(v.World)
		v.SetAlreadyCreated(true)
	})
	return collection, nil
}

func (r *WorldImpl) UpdateByID(a0 context.Context, a1 uint64, a2 map[string]interface{}) error {
	if err := r.worldDAO.UpdateByID(a0, a1, a2); err != nil {
		return xerrors.Errorf("failed to update: %w", err)
	}
	return nil
}

func (r *WorldImpl) UpdateByIDs(a0 context.Context, a1 []uint64, a2 map[string]interface{}) error {
	if err := r.worldDAO.UpdateByIDs(a0, a1, a2); err != nil {
		return xerrors.Errorf("failed to update: %w", err)
	}
	return nil
}

func (r *WorldImpl) DeleteByID(a0 context.Context, a1 uint64) error {
	if err := r.worldDAO.DeleteByID(a0, a1); err != nil {
		return xerrors.Errorf("failed to delete: %w", err)
	}
	return nil
}

func (r *WorldImpl) DeleteByIDs(a0 context.Context, a1 []uint64) error {
	if err := r.worldDAO.DeleteByIDs(a0, a1); err != nil {
		return xerrors.Errorf("failed to delete: %w", err)
	}
	return nil
}

func (r *WorldImpl) Count(a0 context.Context) (r0 int64, r1 error) {
	r0, r1 = r.worldDAO.Count(a0)
	if r1 != nil {
		r1 = xerrors.Errorf("failed to Count: %w", r1)
	}
	return
}

func (r *WorldImpl) Delete(a0 context.Context, a1 *entity.World) (r0 error) {
	r0 = r.worldDAO.Delete(a0, a1)
	if r0 != nil {
		r0 = xerrors.Errorf("failed to Delete: %w", r0)
	}
	return
}

func (r *WorldImpl) Update(a0 context.Context, a1 *entity.World) (r0 error) {
	r0 = r.worldDAO.Update(a0, a1)
	if r0 != nil {
		r0 = xerrors.Errorf("failed to Update: %w", r0)
	}
	return
}

func (r *WorldImpl) createCollection(entities entity.Worlds) *model.Worlds {
	values := model.NewWorlds(entities)
	for i := 0; i < len(entities); i += 1 {
		values.Add(r.create(entities[i], values))
	}
	return values
}

func (r *WorldImpl) create(entity *entity.World, values *model.Worlds) *model.World {
	value := model.NewWorld(entity, r.worldDAO)
	value.SetConverter(r.repo.(model.ModelConverter))
	return value
}
