// Code generated by eevee. DO NOT EDIT!

package repository

import (
	"context"
	"log"
	"rapidashplugin/entity"
	"rapidashplugin/model"
	"reflect"

	"golang.org/x/xerrors"
)

type SkillMock struct {
	expect *SkillExpect
}

func (r *SkillMock) EXPECT() *SkillExpect {
	return r.expect
}

func NewSkillMock() *SkillMock {
	return &SkillMock{expect: NewSkillExpect()}
}

type SkillToModelExpect struct {
	expect        *SkillExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(value *entity.Skill)
	value         *entity.Skill
	r0            *model.Skill
}

func (r *SkillToModelExpect) Return(r0 *model.Skill) *SkillToModelExpect {
	r.r0 = r0
	return r
}

func (r *SkillToModelExpect) Do(action func(value *entity.Skill)) *SkillToModelExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *SkillToModelExpect) OutOfOrder() *SkillToModelExpect {
	r.isOutOfOrder = true
	return r
}

func (r *SkillToModelExpect) AnyTimes() *SkillToModelExpect {
	r.isAnyTimes = true
	return r
}

func (r *SkillToModelExpect) Times(n int) *SkillToModelExpect {
	r.requiredTimes = n
	return r
}

func (r *SkillMock) ToModel(value *entity.Skill) (r0 *model.Skill) {
	if len(r.expect.toModel) == 0 {
		log.Printf("cannot find mock method for Skill.ToModel")
		return
	}
	for _, exp := range r.expect.toModel {
		if !reflect.DeepEqual(exp.value, value) {
			continue
		}
		for _, action := range exp.actions {
			action(value)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			log.Printf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		return
	}
	log.Printf("invalid argument Skill value:[%+v]", value)
	return
}

func (r *SkillExpect) ToModel(value *entity.Skill) *SkillToModelExpect {
	exp := &SkillToModelExpect{
		actions: []func(value *entity.Skill){},
		expect:  r,
		value:   value,
	}
	r.toModel = append(r.toModel, exp)
	return exp
}

type SkillToModelsExpect struct {
	expect        *SkillExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(values entity.Skills)
	values        entity.Skills
	r0            *model.Skills
}

func (r *SkillToModelsExpect) Return(r0 *model.Skills) *SkillToModelsExpect {
	r.r0 = r0
	return r
}

func (r *SkillToModelsExpect) Do(action func(values entity.Skills)) *SkillToModelsExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *SkillToModelsExpect) OutOfOrder() *SkillToModelsExpect {
	r.isOutOfOrder = true
	return r
}

func (r *SkillToModelsExpect) AnyTimes() *SkillToModelsExpect {
	r.isAnyTimes = true
	return r
}

func (r *SkillToModelsExpect) Times(n int) *SkillToModelsExpect {
	r.requiredTimes = n
	return r
}

func (r *SkillMock) ToModels(values entity.Skills) (r0 *model.Skills) {
	if len(r.expect.toModels) == 0 {
		log.Printf("cannot find mock method for Skill.ToModels")
		return
	}
	for _, exp := range r.expect.toModels {
		if !reflect.DeepEqual(exp.values, values) {
			continue
		}
		for _, action := range exp.actions {
			action(values)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			log.Printf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		return
	}
	log.Printf("invalid argument Skill values:[%+v]", values)
	return
}

func (r *SkillExpect) ToModels(values entity.Skills) *SkillToModelsExpect {
	exp := &SkillToModelsExpect{
		actions: []func(values entity.Skills){},
		expect:  r,
		values:  values,
	}
	r.toModels = append(r.toModels, exp)
	return exp
}

type SkillCreateExpect struct {
	expect        *SkillExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(ctx context.Context, value *entity.Skill)
	ctx           context.Context
	value         *entity.Skill
	r0            *model.Skill
	r1            error
}

func (r *SkillCreateExpect) Return(r0 *model.Skill, r1 error) *SkillCreateExpect {
	r.r0 = r0
	r.r1 = r1
	return r
}

func (r *SkillCreateExpect) Do(action func(ctx context.Context, value *entity.Skill)) *SkillCreateExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *SkillCreateExpect) OutOfOrder() *SkillCreateExpect {
	r.isOutOfOrder = true
	return r
}

func (r *SkillCreateExpect) AnyTimes() *SkillCreateExpect {
	r.isAnyTimes = true
	return r
}

func (r *SkillCreateExpect) Times(n int) *SkillCreateExpect {
	r.requiredTimes = n
	return r
}

func (r *SkillMock) Create(ctx context.Context, value *entity.Skill) (r0 *model.Skill, r1 error) {
	if len(r.expect.create) == 0 {
		r1 = xerrors.New("cannot find mock method for Skill.Create")
		return
	}
	for _, exp := range r.expect.create {
		if !reflect.DeepEqual(exp.ctx, ctx) {
			continue
		}
		if !reflect.DeepEqual(exp.value, value) {
			continue
		}
		for _, action := range exp.actions {
			action(ctx, value)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			r1 = exp.r1
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			r1 = xerrors.Errorf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		r1 = exp.r1
		return
	}
	r1 = xerrors.Errorf("invalid argument Skill ctx:[%+v] value:[%+v]", ctx, value)
	return
}

func (r *SkillExpect) Create(ctx context.Context, value *entity.Skill) *SkillCreateExpect {
	exp := &SkillCreateExpect{
		actions: []func(ctx context.Context, value *entity.Skill){},
		ctx:     ctx,
		expect:  r,
		value:   value,
	}
	r.create = append(r.create, exp)
	return exp
}

type SkillCreatesExpect struct {
	expect        *SkillExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(ctx context.Context, entities entity.Skills)
	ctx           context.Context
	entities      entity.Skills
	r0            *model.Skills
	r1            error
}

func (r *SkillCreatesExpect) Return(r0 *model.Skills, r1 error) *SkillCreatesExpect {
	r.r0 = r0
	r.r1 = r1
	return r
}

func (r *SkillCreatesExpect) Do(action func(ctx context.Context, entities entity.Skills)) *SkillCreatesExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *SkillCreatesExpect) OutOfOrder() *SkillCreatesExpect {
	r.isOutOfOrder = true
	return r
}

func (r *SkillCreatesExpect) AnyTimes() *SkillCreatesExpect {
	r.isAnyTimes = true
	return r
}

func (r *SkillCreatesExpect) Times(n int) *SkillCreatesExpect {
	r.requiredTimes = n
	return r
}

func (r *SkillMock) Creates(ctx context.Context, entities entity.Skills) (r0 *model.Skills, r1 error) {
	if len(r.expect.creates) == 0 {
		r1 = xerrors.New("cannot find mock method for Skill.Creates")
		return
	}
	for _, exp := range r.expect.creates {
		if !reflect.DeepEqual(exp.ctx, ctx) {
			continue
		}
		if !reflect.DeepEqual(exp.entities, entities) {
			continue
		}
		for _, action := range exp.actions {
			action(ctx, entities)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			r1 = exp.r1
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			r1 = xerrors.Errorf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		r1 = exp.r1
		return
	}
	r1 = xerrors.Errorf("invalid argument Skill ctx:[%+v] entities:[%+v]", ctx, entities)
	return
}

func (r *SkillExpect) Creates(ctx context.Context, entities entity.Skills) *SkillCreatesExpect {
	exp := &SkillCreatesExpect{
		actions:  []func(ctx context.Context, entities entity.Skills){},
		ctx:      ctx,
		entities: entities,
		expect:   r,
	}
	r.creates = append(r.creates, exp)
	return exp
}

type SkillFindAllExpect struct {
	expect        *SkillExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context)
	a0            context.Context
	r0            *model.Skills
	r1            error
}

func (r *SkillFindAllExpect) Return(r0 *model.Skills, r1 error) *SkillFindAllExpect {
	r.r0 = r0
	r.r1 = r1
	return r
}

func (r *SkillFindAllExpect) Do(action func(a0 context.Context)) *SkillFindAllExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *SkillFindAllExpect) OutOfOrder() *SkillFindAllExpect {
	r.isOutOfOrder = true
	return r
}

func (r *SkillFindAllExpect) AnyTimes() *SkillFindAllExpect {
	r.isAnyTimes = true
	return r
}

func (r *SkillFindAllExpect) Times(n int) *SkillFindAllExpect {
	r.requiredTimes = n
	return r
}

func (r *SkillMock) FindAll(a0 context.Context) (r0 *model.Skills, r1 error) {
	if len(r.expect.findAll) == 0 {
		r1 = xerrors.New("cannot find mock method for Skill.FindAll")
		return
	}
	for _, exp := range r.expect.findAll {
		if !reflect.DeepEqual(exp.a0, a0) {
			continue
		}
		for _, action := range exp.actions {
			action(a0)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			r1 = exp.r1
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			r1 = xerrors.Errorf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		r1 = exp.r1
		return
	}
	r1 = xerrors.Errorf("invalid argument Skill a0:[%+v]", a0)
	return
}

func (r *SkillExpect) FindAll(a0 context.Context) *SkillFindAllExpect {
	exp := &SkillFindAllExpect{
		a0:      a0,
		actions: []func(a0 context.Context){},
		expect:  r,
	}
	r.findAll = append(r.findAll, exp)
	return exp
}

type SkillFindByIDExpect struct {
	expect        *SkillExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context, a1 uint64)
	a0            context.Context
	a1            uint64
	r0            *model.Skill
	r1            error
}

func (r *SkillFindByIDExpect) Return(r0 *model.Skill, r1 error) *SkillFindByIDExpect {
	r.r0 = r0
	r.r1 = r1
	return r
}

func (r *SkillFindByIDExpect) Do(action func(a0 context.Context, a1 uint64)) *SkillFindByIDExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *SkillFindByIDExpect) OutOfOrder() *SkillFindByIDExpect {
	r.isOutOfOrder = true
	return r
}

func (r *SkillFindByIDExpect) AnyTimes() *SkillFindByIDExpect {
	r.isAnyTimes = true
	return r
}

func (r *SkillFindByIDExpect) Times(n int) *SkillFindByIDExpect {
	r.requiredTimes = n
	return r
}

func (r *SkillMock) FindByID(a0 context.Context, a1 uint64) (r0 *model.Skill, r1 error) {
	if len(r.expect.findByID) == 0 {
		r1 = xerrors.New("cannot find mock method for Skill.FindByID")
		return
	}
	for _, exp := range r.expect.findByID {
		if !reflect.DeepEqual(exp.a0, a0) {
			continue
		}
		if !reflect.DeepEqual(exp.a1, a1) {
			continue
		}
		for _, action := range exp.actions {
			action(a0, a1)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			r1 = exp.r1
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			r1 = xerrors.Errorf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		r1 = exp.r1
		return
	}
	r1 = xerrors.Errorf("invalid argument Skill a0:[%+v] a1:[%+v]", a0, a1)
	return
}

func (r *SkillExpect) FindByID(a0 context.Context, a1 uint64) *SkillFindByIDExpect {
	exp := &SkillFindByIDExpect{
		a0:      a0,
		a1:      a1,
		actions: []func(a0 context.Context, a1 uint64){},
		expect:  r,
	}
	r.findByID = append(r.findByID, exp)
	return exp
}

type SkillFindByIDsExpect struct {
	expect        *SkillExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context, a1 []uint64)
	a0            context.Context
	a1            []uint64
	r0            *model.Skills
	r1            error
}

func (r *SkillFindByIDsExpect) Return(r0 *model.Skills, r1 error) *SkillFindByIDsExpect {
	r.r0 = r0
	r.r1 = r1
	return r
}

func (r *SkillFindByIDsExpect) Do(action func(a0 context.Context, a1 []uint64)) *SkillFindByIDsExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *SkillFindByIDsExpect) OutOfOrder() *SkillFindByIDsExpect {
	r.isOutOfOrder = true
	return r
}

func (r *SkillFindByIDsExpect) AnyTimes() *SkillFindByIDsExpect {
	r.isAnyTimes = true
	return r
}

func (r *SkillFindByIDsExpect) Times(n int) *SkillFindByIDsExpect {
	r.requiredTimes = n
	return r
}

func (r *SkillMock) FindByIDs(a0 context.Context, a1 []uint64) (r0 *model.Skills, r1 error) {
	if len(r.expect.findByIDs) == 0 {
		r1 = xerrors.New("cannot find mock method for Skill.FindByIDs")
		return
	}
	for _, exp := range r.expect.findByIDs {
		if !reflect.DeepEqual(exp.a0, a0) {
			continue
		}
		if len(exp.a1) != len(a1) {
			continue
		}
		if exp.isOutOfOrder {
			isMatched := func() bool {
				for _, exp := range exp.a1 {
					found := false
					for idx, act := range a1 {
						if exp != act {
							continue
						}
						a1 = append(a1[:idx], a1[idx+1:]...)
						found = true
						break
					}
					if !found {
						return false
					}
				}
				return true
			}()
			if !isMatched {
				continue
			}
		} else {
			if !reflect.DeepEqual(exp.a1, a1) {
				continue
			}
		}
		for _, action := range exp.actions {
			action(a0, a1)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			r1 = exp.r1
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			r1 = xerrors.Errorf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		r1 = exp.r1
		return
	}
	r1 = xerrors.Errorf("invalid argument Skill a0:[%+v] a1:[%+v]", a0, a1)
	return
}

func (r *SkillExpect) FindByIDs(a0 context.Context, a1 []uint64) *SkillFindByIDsExpect {
	exp := &SkillFindByIDsExpect{
		a0:      a0,
		a1:      a1,
		actions: []func(a0 context.Context, a1 []uint64){},
		expect:  r,
	}
	r.findByIDs = append(r.findByIDs, exp)
	return exp
}

type SkillUpdateByIDExpect struct {
	expect        *SkillExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context, a1 uint64, a2 map[string]interface{})
	a0            context.Context
	a1            uint64
	a2            map[string]interface{}
	r0            error
}

func (r *SkillUpdateByIDExpect) Return(r0 error) *SkillUpdateByIDExpect {
	r.r0 = r0
	return r
}

func (r *SkillUpdateByIDExpect) Do(action func(a0 context.Context, a1 uint64, a2 map[string]interface{})) *SkillUpdateByIDExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *SkillUpdateByIDExpect) OutOfOrder() *SkillUpdateByIDExpect {
	r.isOutOfOrder = true
	return r
}

func (r *SkillUpdateByIDExpect) AnyTimes() *SkillUpdateByIDExpect {
	r.isAnyTimes = true
	return r
}

func (r *SkillUpdateByIDExpect) Times(n int) *SkillUpdateByIDExpect {
	r.requiredTimes = n
	return r
}

func (r *SkillMock) UpdateByID(a0 context.Context, a1 uint64, a2 map[string]interface{}) (r0 error) {
	if len(r.expect.updateByID) == 0 {
		r0 = xerrors.New("cannot find mock method for Skill.UpdateByID")
		return
	}
	for _, exp := range r.expect.updateByID {
		if !reflect.DeepEqual(exp.a0, a0) {
			continue
		}
		if !reflect.DeepEqual(exp.a1, a1) {
			continue
		}
		if !reflect.DeepEqual(exp.a2, a2) {
			continue
		}
		for _, action := range exp.actions {
			action(a0, a1, a2)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			r0 = xerrors.Errorf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		return
	}
	r0 = xerrors.Errorf("invalid argument Skill a0:[%+v] a1:[%+v] a2:[%+v]", a0, a1, a2)
	return
}

func (r *SkillExpect) UpdateByID(a0 context.Context, a1 uint64, a2 map[string]interface{}) *SkillUpdateByIDExpect {
	exp := &SkillUpdateByIDExpect{
		a0:      a0,
		a1:      a1,
		a2:      a2,
		actions: []func(a0 context.Context, a1 uint64, a2 map[string]interface{}){},
		expect:  r,
	}
	r.updateByID = append(r.updateByID, exp)
	return exp
}

type SkillUpdateByIDsExpect struct {
	expect        *SkillExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context, a1 []uint64, a2 map[string]interface{})
	a0            context.Context
	a1            []uint64
	a2            map[string]interface{}
	r0            error
}

func (r *SkillUpdateByIDsExpect) Return(r0 error) *SkillUpdateByIDsExpect {
	r.r0 = r0
	return r
}

func (r *SkillUpdateByIDsExpect) Do(action func(a0 context.Context, a1 []uint64, a2 map[string]interface{})) *SkillUpdateByIDsExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *SkillUpdateByIDsExpect) OutOfOrder() *SkillUpdateByIDsExpect {
	r.isOutOfOrder = true
	return r
}

func (r *SkillUpdateByIDsExpect) AnyTimes() *SkillUpdateByIDsExpect {
	r.isAnyTimes = true
	return r
}

func (r *SkillUpdateByIDsExpect) Times(n int) *SkillUpdateByIDsExpect {
	r.requiredTimes = n
	return r
}

func (r *SkillMock) UpdateByIDs(a0 context.Context, a1 []uint64, a2 map[string]interface{}) (r0 error) {
	if len(r.expect.updateByIDs) == 0 {
		r0 = xerrors.New("cannot find mock method for Skill.UpdateByIDs")
		return
	}
	for _, exp := range r.expect.updateByIDs {
		if !reflect.DeepEqual(exp.a0, a0) {
			continue
		}
		if len(exp.a1) != len(a1) {
			continue
		}
		if exp.isOutOfOrder {
			isMatched := func() bool {
				for _, exp := range exp.a1 {
					found := false
					for idx, act := range a1 {
						if exp != act {
							continue
						}
						a1 = append(a1[:idx], a1[idx+1:]...)
						found = true
						break
					}
					if !found {
						return false
					}
				}
				return true
			}()
			if !isMatched {
				continue
			}
		} else {
			if !reflect.DeepEqual(exp.a1, a1) {
				continue
			}
		}
		if !reflect.DeepEqual(exp.a2, a2) {
			continue
		}
		for _, action := range exp.actions {
			action(a0, a1, a2)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			r0 = xerrors.Errorf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		return
	}
	r0 = xerrors.Errorf("invalid argument Skill a0:[%+v] a1:[%+v] a2:[%+v]", a0, a1, a2)
	return
}

func (r *SkillExpect) UpdateByIDs(a0 context.Context, a1 []uint64, a2 map[string]interface{}) *SkillUpdateByIDsExpect {
	exp := &SkillUpdateByIDsExpect{
		a0:      a0,
		a1:      a1,
		a2:      a2,
		actions: []func(a0 context.Context, a1 []uint64, a2 map[string]interface{}){},
		expect:  r,
	}
	r.updateByIDs = append(r.updateByIDs, exp)
	return exp
}

type SkillDeleteByIDExpect struct {
	expect        *SkillExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context, a1 uint64)
	a0            context.Context
	a1            uint64
	r0            error
}

func (r *SkillDeleteByIDExpect) Return(r0 error) *SkillDeleteByIDExpect {
	r.r0 = r0
	return r
}

func (r *SkillDeleteByIDExpect) Do(action func(a0 context.Context, a1 uint64)) *SkillDeleteByIDExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *SkillDeleteByIDExpect) OutOfOrder() *SkillDeleteByIDExpect {
	r.isOutOfOrder = true
	return r
}

func (r *SkillDeleteByIDExpect) AnyTimes() *SkillDeleteByIDExpect {
	r.isAnyTimes = true
	return r
}

func (r *SkillDeleteByIDExpect) Times(n int) *SkillDeleteByIDExpect {
	r.requiredTimes = n
	return r
}

func (r *SkillMock) DeleteByID(a0 context.Context, a1 uint64) (r0 error) {
	if len(r.expect.deleteByID) == 0 {
		r0 = xerrors.New("cannot find mock method for Skill.DeleteByID")
		return
	}
	for _, exp := range r.expect.deleteByID {
		if !reflect.DeepEqual(exp.a0, a0) {
			continue
		}
		if !reflect.DeepEqual(exp.a1, a1) {
			continue
		}
		for _, action := range exp.actions {
			action(a0, a1)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			r0 = xerrors.Errorf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		return
	}
	r0 = xerrors.Errorf("invalid argument Skill a0:[%+v] a1:[%+v]", a0, a1)
	return
}

func (r *SkillExpect) DeleteByID(a0 context.Context, a1 uint64) *SkillDeleteByIDExpect {
	exp := &SkillDeleteByIDExpect{
		a0:      a0,
		a1:      a1,
		actions: []func(a0 context.Context, a1 uint64){},
		expect:  r,
	}
	r.deleteByID = append(r.deleteByID, exp)
	return exp
}

type SkillDeleteByIDsExpect struct {
	expect        *SkillExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context, a1 []uint64)
	a0            context.Context
	a1            []uint64
	r0            error
}

func (r *SkillDeleteByIDsExpect) Return(r0 error) *SkillDeleteByIDsExpect {
	r.r0 = r0
	return r
}

func (r *SkillDeleteByIDsExpect) Do(action func(a0 context.Context, a1 []uint64)) *SkillDeleteByIDsExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *SkillDeleteByIDsExpect) OutOfOrder() *SkillDeleteByIDsExpect {
	r.isOutOfOrder = true
	return r
}

func (r *SkillDeleteByIDsExpect) AnyTimes() *SkillDeleteByIDsExpect {
	r.isAnyTimes = true
	return r
}

func (r *SkillDeleteByIDsExpect) Times(n int) *SkillDeleteByIDsExpect {
	r.requiredTimes = n
	return r
}

func (r *SkillMock) DeleteByIDs(a0 context.Context, a1 []uint64) (r0 error) {
	if len(r.expect.deleteByIDs) == 0 {
		r0 = xerrors.New("cannot find mock method for Skill.DeleteByIDs")
		return
	}
	for _, exp := range r.expect.deleteByIDs {
		if !reflect.DeepEqual(exp.a0, a0) {
			continue
		}
		if len(exp.a1) != len(a1) {
			continue
		}
		if exp.isOutOfOrder {
			isMatched := func() bool {
				for _, exp := range exp.a1 {
					found := false
					for idx, act := range a1 {
						if exp != act {
							continue
						}
						a1 = append(a1[:idx], a1[idx+1:]...)
						found = true
						break
					}
					if !found {
						return false
					}
				}
				return true
			}()
			if !isMatched {
				continue
			}
		} else {
			if !reflect.DeepEqual(exp.a1, a1) {
				continue
			}
		}
		for _, action := range exp.actions {
			action(a0, a1)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			r0 = xerrors.Errorf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		return
	}
	r0 = xerrors.Errorf("invalid argument Skill a0:[%+v] a1:[%+v]", a0, a1)
	return
}

func (r *SkillExpect) DeleteByIDs(a0 context.Context, a1 []uint64) *SkillDeleteByIDsExpect {
	exp := &SkillDeleteByIDsExpect{
		a0:      a0,
		a1:      a1,
		actions: []func(a0 context.Context, a1 []uint64){},
		expect:  r,
	}
	r.deleteByIDs = append(r.deleteByIDs, exp)
	return exp
}

type SkillCountExpect struct {
	expect        *SkillExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context)
	a0            context.Context
	r0            int64
	r1            error
}

func (r *SkillCountExpect) Return(r0 int64, r1 error) *SkillCountExpect {
	r.r0 = r0
	r.r1 = r1
	return r
}

func (r *SkillCountExpect) Do(action func(a0 context.Context)) *SkillCountExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *SkillCountExpect) OutOfOrder() *SkillCountExpect {
	r.isOutOfOrder = true
	return r
}

func (r *SkillCountExpect) AnyTimes() *SkillCountExpect {
	r.isAnyTimes = true
	return r
}

func (r *SkillCountExpect) Times(n int) *SkillCountExpect {
	r.requiredTimes = n
	return r
}

func (r *SkillMock) Count(a0 context.Context) (r0 int64, r1 error) {
	if len(r.expect.count) == 0 {
		r1 = xerrors.New("cannot find mock method for Skill.Count")
		return
	}
	for _, exp := range r.expect.count {
		if !reflect.DeepEqual(exp.a0, a0) {
			continue
		}
		for _, action := range exp.actions {
			action(a0)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			r1 = exp.r1
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			r1 = xerrors.Errorf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		r1 = exp.r1
		return
	}
	r1 = xerrors.Errorf("invalid argument Skill a0:[%+v]", a0)
	return
}

func (r *SkillExpect) Count(a0 context.Context) *SkillCountExpect {
	exp := &SkillCountExpect{
		a0:      a0,
		actions: []func(a0 context.Context){},
		expect:  r,
	}
	r.count = append(r.count, exp)
	return exp
}

type SkillDeleteExpect struct {
	expect        *SkillExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context, a1 *entity.Skill)
	a0            context.Context
	a1            *entity.Skill
	r0            error
}

func (r *SkillDeleteExpect) Return(r0 error) *SkillDeleteExpect {
	r.r0 = r0
	return r
}

func (r *SkillDeleteExpect) Do(action func(a0 context.Context, a1 *entity.Skill)) *SkillDeleteExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *SkillDeleteExpect) OutOfOrder() *SkillDeleteExpect {
	r.isOutOfOrder = true
	return r
}

func (r *SkillDeleteExpect) AnyTimes() *SkillDeleteExpect {
	r.isAnyTimes = true
	return r
}

func (r *SkillDeleteExpect) Times(n int) *SkillDeleteExpect {
	r.requiredTimes = n
	return r
}

func (r *SkillMock) Delete(a0 context.Context, a1 *entity.Skill) (r0 error) {
	if len(r.expect.delete) == 0 {
		r0 = xerrors.New("cannot find mock method for Skill.Delete")
		return
	}
	for _, exp := range r.expect.delete {
		if !reflect.DeepEqual(exp.a0, a0) {
			continue
		}
		if !reflect.DeepEqual(exp.a1, a1) {
			continue
		}
		for _, action := range exp.actions {
			action(a0, a1)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			r0 = xerrors.Errorf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		return
	}
	r0 = xerrors.Errorf("invalid argument Skill a0:[%+v] a1:[%+v]", a0, a1)
	return
}

func (r *SkillExpect) Delete(a0 context.Context, a1 *entity.Skill) *SkillDeleteExpect {
	exp := &SkillDeleteExpect{
		a0:      a0,
		a1:      a1,
		actions: []func(a0 context.Context, a1 *entity.Skill){},
		expect:  r,
	}
	r.delete = append(r.delete, exp)
	return exp
}

type SkillUpdateExpect struct {
	expect        *SkillExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context, a1 *entity.Skill)
	a0            context.Context
	a1            *entity.Skill
	r0            error
}

func (r *SkillUpdateExpect) Return(r0 error) *SkillUpdateExpect {
	r.r0 = r0
	return r
}

func (r *SkillUpdateExpect) Do(action func(a0 context.Context, a1 *entity.Skill)) *SkillUpdateExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *SkillUpdateExpect) OutOfOrder() *SkillUpdateExpect {
	r.isOutOfOrder = true
	return r
}

func (r *SkillUpdateExpect) AnyTimes() *SkillUpdateExpect {
	r.isAnyTimes = true
	return r
}

func (r *SkillUpdateExpect) Times(n int) *SkillUpdateExpect {
	r.requiredTimes = n
	return r
}

func (r *SkillMock) Update(a0 context.Context, a1 *entity.Skill) (r0 error) {
	if len(r.expect.update) == 0 {
		r0 = xerrors.New("cannot find mock method for Skill.Update")
		return
	}
	for _, exp := range r.expect.update {
		if !reflect.DeepEqual(exp.a0, a0) {
			continue
		}
		if !reflect.DeepEqual(exp.a1, a1) {
			continue
		}
		for _, action := range exp.actions {
			action(a0, a1)
		}
		if exp.isAnyTimes {
			r0 = exp.r0
			return
		}
		if exp.requiredTimes > 1 && exp.calledTimes > exp.requiredTimes {
			r0 = xerrors.Errorf("invalid call times. requiredTimes: [%d] calledTimes: [%d]", exp.requiredTimes, exp.calledTimes)
			return
		}
		exp.calledTimes++
		r0 = exp.r0
		return
	}
	r0 = xerrors.Errorf("invalid argument Skill a0:[%+v] a1:[%+v]", a0, a1)
	return
}

func (r *SkillExpect) Update(a0 context.Context, a1 *entity.Skill) *SkillUpdateExpect {
	exp := &SkillUpdateExpect{
		a0:      a0,
		a1:      a1,
		actions: []func(a0 context.Context, a1 *entity.Skill){},
		expect:  r,
	}
	r.update = append(r.update, exp)
	return exp
}

type SkillExpect struct {
	toModel     []*SkillToModelExpect
	toModels    []*SkillToModelsExpect
	create      []*SkillCreateExpect
	creates     []*SkillCreatesExpect
	findAll     []*SkillFindAllExpect
	findByID    []*SkillFindByIDExpect
	findByIDs   []*SkillFindByIDsExpect
	updateByID  []*SkillUpdateByIDExpect
	updateByIDs []*SkillUpdateByIDsExpect
	deleteByID  []*SkillDeleteByIDExpect
	deleteByIDs []*SkillDeleteByIDsExpect
	count       []*SkillCountExpect
	delete      []*SkillDeleteExpect
	update      []*SkillUpdateExpect
}

func NewSkillExpect() *SkillExpect {
	return &SkillExpect{
		count:       []*SkillCountExpect{},
		create:      []*SkillCreateExpect{},
		creates:     []*SkillCreatesExpect{},
		delete:      []*SkillDeleteExpect{},
		deleteByID:  []*SkillDeleteByIDExpect{},
		deleteByIDs: []*SkillDeleteByIDsExpect{},
		findAll:     []*SkillFindAllExpect{},
		findByID:    []*SkillFindByIDExpect{},
		findByIDs:   []*SkillFindByIDsExpect{},
		toModel:     []*SkillToModelExpect{},
		toModels:    []*SkillToModelsExpect{},
		update:      []*SkillUpdateExpect{},
		updateByID:  []*SkillUpdateByIDExpect{},
		updateByIDs: []*SkillUpdateByIDsExpect{},
	}
}
