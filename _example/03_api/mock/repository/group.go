// Code generated by eevee. DO NOT EDIT!

package repository

import (
	"api/entity"
	"api/model"
	"context"
	"log"
	"reflect"

	"golang.org/x/xerrors"
)

type GroupMock struct {
	expect *GroupExpect
}

func (r *GroupMock) EXPECT() *GroupExpect {
	return r.expect
}

func NewGroupMock() *GroupMock {
	return &GroupMock{expect: NewGroupExpect()}
}

type GroupToModelExpect struct {
	expect        *GroupExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(value *entity.Group)
	value         *entity.Group
	r0            *model.Group
}

func (r *GroupToModelExpect) Return(r0 *model.Group) *GroupToModelExpect {
	r.r0 = r0
	return r
}

func (r *GroupToModelExpect) Do(action func(value *entity.Group)) *GroupToModelExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *GroupToModelExpect) OutOfOrder() *GroupToModelExpect {
	r.isOutOfOrder = true
	return r
}

func (r *GroupToModelExpect) AnyTimes() *GroupToModelExpect {
	r.isAnyTimes = true
	return r
}

func (r *GroupToModelExpect) Times(n int) *GroupToModelExpect {
	r.requiredTimes = n
	return r
}

func (r *GroupMock) ToModel(value *entity.Group) (r0 *model.Group) {
	if len(r.expect.toModel) == 0 {
		log.Printf("cannot find mock method for Group.ToModel")
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
	log.Printf("invalid argument Group value:[%+v]", value)
	return
}

func (r *GroupExpect) ToModel(value *entity.Group) *GroupToModelExpect {
	exp := &GroupToModelExpect{
		actions: []func(value *entity.Group){},
		expect:  r,
		value:   value,
	}
	r.toModel = append(r.toModel, exp)
	return exp
}

type GroupToModelsExpect struct {
	expect        *GroupExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(values entity.Groups)
	values        entity.Groups
	r0            *model.Groups
}

func (r *GroupToModelsExpect) Return(r0 *model.Groups) *GroupToModelsExpect {
	r.r0 = r0
	return r
}

func (r *GroupToModelsExpect) Do(action func(values entity.Groups)) *GroupToModelsExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *GroupToModelsExpect) OutOfOrder() *GroupToModelsExpect {
	r.isOutOfOrder = true
	return r
}

func (r *GroupToModelsExpect) AnyTimes() *GroupToModelsExpect {
	r.isAnyTimes = true
	return r
}

func (r *GroupToModelsExpect) Times(n int) *GroupToModelsExpect {
	r.requiredTimes = n
	return r
}

func (r *GroupMock) ToModels(values entity.Groups) (r0 *model.Groups) {
	if len(r.expect.toModels) == 0 {
		log.Printf("cannot find mock method for Group.ToModels")
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
	log.Printf("invalid argument Group values:[%+v]", values)
	return
}

func (r *GroupExpect) ToModels(values entity.Groups) *GroupToModelsExpect {
	exp := &GroupToModelsExpect{
		actions: []func(values entity.Groups){},
		expect:  r,
		values:  values,
	}
	r.toModels = append(r.toModels, exp)
	return exp
}

type GroupCreateExpect struct {
	expect        *GroupExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(ctx context.Context, value *entity.Group)
	ctx           context.Context
	value         *entity.Group
	r0            *model.Group
	r1            error
}

func (r *GroupCreateExpect) Return(r0 *model.Group, r1 error) *GroupCreateExpect {
	r.r0 = r0
	r.r1 = r1
	return r
}

func (r *GroupCreateExpect) Do(action func(ctx context.Context, value *entity.Group)) *GroupCreateExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *GroupCreateExpect) OutOfOrder() *GroupCreateExpect {
	r.isOutOfOrder = true
	return r
}

func (r *GroupCreateExpect) AnyTimes() *GroupCreateExpect {
	r.isAnyTimes = true
	return r
}

func (r *GroupCreateExpect) Times(n int) *GroupCreateExpect {
	r.requiredTimes = n
	return r
}

func (r *GroupMock) Create(ctx context.Context, value *entity.Group) (r0 *model.Group, r1 error) {
	if len(r.expect.create) == 0 {
		r1 = xerrors.New("cannot find mock method for Group.Create")
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
	r1 = xerrors.Errorf("invalid argument Group ctx:[%+v] value:[%+v]", ctx, value)
	return
}

func (r *GroupExpect) Create(ctx context.Context, value *entity.Group) *GroupCreateExpect {
	exp := &GroupCreateExpect{
		actions: []func(ctx context.Context, value *entity.Group){},
		ctx:     ctx,
		expect:  r,
		value:   value,
	}
	r.create = append(r.create, exp)
	return exp
}

type GroupCreatesExpect struct {
	expect        *GroupExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(ctx context.Context, entities entity.Groups)
	ctx           context.Context
	entities      entity.Groups
	r0            *model.Groups
	r1            error
}

func (r *GroupCreatesExpect) Return(r0 *model.Groups, r1 error) *GroupCreatesExpect {
	r.r0 = r0
	r.r1 = r1
	return r
}

func (r *GroupCreatesExpect) Do(action func(ctx context.Context, entities entity.Groups)) *GroupCreatesExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *GroupCreatesExpect) OutOfOrder() *GroupCreatesExpect {
	r.isOutOfOrder = true
	return r
}

func (r *GroupCreatesExpect) AnyTimes() *GroupCreatesExpect {
	r.isAnyTimes = true
	return r
}

func (r *GroupCreatesExpect) Times(n int) *GroupCreatesExpect {
	r.requiredTimes = n
	return r
}

func (r *GroupMock) Creates(ctx context.Context, entities entity.Groups) (r0 *model.Groups, r1 error) {
	if len(r.expect.creates) == 0 {
		r1 = xerrors.New("cannot find mock method for Group.Creates")
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
	r1 = xerrors.Errorf("invalid argument Group ctx:[%+v] entities:[%+v]", ctx, entities)
	return
}

func (r *GroupExpect) Creates(ctx context.Context, entities entity.Groups) *GroupCreatesExpect {
	exp := &GroupCreatesExpect{
		actions:  []func(ctx context.Context, entities entity.Groups){},
		ctx:      ctx,
		entities: entities,
		expect:   r,
	}
	r.creates = append(r.creates, exp)
	return exp
}

type GroupFindAllExpect struct {
	expect        *GroupExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context)
	a0            context.Context
	r0            *model.Groups
	r1            error
}

func (r *GroupFindAllExpect) Return(r0 *model.Groups, r1 error) *GroupFindAllExpect {
	r.r0 = r0
	r.r1 = r1
	return r
}

func (r *GroupFindAllExpect) Do(action func(a0 context.Context)) *GroupFindAllExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *GroupFindAllExpect) OutOfOrder() *GroupFindAllExpect {
	r.isOutOfOrder = true
	return r
}

func (r *GroupFindAllExpect) AnyTimes() *GroupFindAllExpect {
	r.isAnyTimes = true
	return r
}

func (r *GroupFindAllExpect) Times(n int) *GroupFindAllExpect {
	r.requiredTimes = n
	return r
}

func (r *GroupMock) FindAll(a0 context.Context) (r0 *model.Groups, r1 error) {
	if len(r.expect.findAll) == 0 {
		r1 = xerrors.New("cannot find mock method for Group.FindAll")
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
	r1 = xerrors.Errorf("invalid argument Group a0:[%+v]", a0)
	return
}

func (r *GroupExpect) FindAll(a0 context.Context) *GroupFindAllExpect {
	exp := &GroupFindAllExpect{
		a0:      a0,
		actions: []func(a0 context.Context){},
		expect:  r,
	}
	r.findAll = append(r.findAll, exp)
	return exp
}

type GroupFindByIDExpect struct {
	expect        *GroupExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context, a1 uint64)
	a0            context.Context
	a1            uint64
	r0            *model.Group
	r1            error
}

func (r *GroupFindByIDExpect) Return(r0 *model.Group, r1 error) *GroupFindByIDExpect {
	r.r0 = r0
	r.r1 = r1
	return r
}

func (r *GroupFindByIDExpect) Do(action func(a0 context.Context, a1 uint64)) *GroupFindByIDExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *GroupFindByIDExpect) OutOfOrder() *GroupFindByIDExpect {
	r.isOutOfOrder = true
	return r
}

func (r *GroupFindByIDExpect) AnyTimes() *GroupFindByIDExpect {
	r.isAnyTimes = true
	return r
}

func (r *GroupFindByIDExpect) Times(n int) *GroupFindByIDExpect {
	r.requiredTimes = n
	return r
}

func (r *GroupMock) FindByID(a0 context.Context, a1 uint64) (r0 *model.Group, r1 error) {
	if len(r.expect.findByID) == 0 {
		r1 = xerrors.New("cannot find mock method for Group.FindByID")
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
	r1 = xerrors.Errorf("invalid argument Group a0:[%+v] a1:[%+v]", a0, a1)
	return
}

func (r *GroupExpect) FindByID(a0 context.Context, a1 uint64) *GroupFindByIDExpect {
	exp := &GroupFindByIDExpect{
		a0:      a0,
		a1:      a1,
		actions: []func(a0 context.Context, a1 uint64){},
		expect:  r,
	}
	r.findByID = append(r.findByID, exp)
	return exp
}

type GroupFindByIDsExpect struct {
	expect        *GroupExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context, a1 []uint64)
	a0            context.Context
	a1            []uint64
	r0            *model.Groups
	r1            error
}

func (r *GroupFindByIDsExpect) Return(r0 *model.Groups, r1 error) *GroupFindByIDsExpect {
	r.r0 = r0
	r.r1 = r1
	return r
}

func (r *GroupFindByIDsExpect) Do(action func(a0 context.Context, a1 []uint64)) *GroupFindByIDsExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *GroupFindByIDsExpect) OutOfOrder() *GroupFindByIDsExpect {
	r.isOutOfOrder = true
	return r
}

func (r *GroupFindByIDsExpect) AnyTimes() *GroupFindByIDsExpect {
	r.isAnyTimes = true
	return r
}

func (r *GroupFindByIDsExpect) Times(n int) *GroupFindByIDsExpect {
	r.requiredTimes = n
	return r
}

func (r *GroupMock) FindByIDs(a0 context.Context, a1 []uint64) (r0 *model.Groups, r1 error) {
	if len(r.expect.findByIDs) == 0 {
		r1 = xerrors.New("cannot find mock method for Group.FindByIDs")
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
	r1 = xerrors.Errorf("invalid argument Group a0:[%+v] a1:[%+v]", a0, a1)
	return
}

func (r *GroupExpect) FindByIDs(a0 context.Context, a1 []uint64) *GroupFindByIDsExpect {
	exp := &GroupFindByIDsExpect{
		a0:      a0,
		a1:      a1,
		actions: []func(a0 context.Context, a1 []uint64){},
		expect:  r,
	}
	r.findByIDs = append(r.findByIDs, exp)
	return exp
}

type GroupUpdateByIDExpect struct {
	expect        *GroupExpect
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

func (r *GroupUpdateByIDExpect) Return(r0 error) *GroupUpdateByIDExpect {
	r.r0 = r0
	return r
}

func (r *GroupUpdateByIDExpect) Do(action func(a0 context.Context, a1 uint64, a2 map[string]interface{})) *GroupUpdateByIDExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *GroupUpdateByIDExpect) OutOfOrder() *GroupUpdateByIDExpect {
	r.isOutOfOrder = true
	return r
}

func (r *GroupUpdateByIDExpect) AnyTimes() *GroupUpdateByIDExpect {
	r.isAnyTimes = true
	return r
}

func (r *GroupUpdateByIDExpect) Times(n int) *GroupUpdateByIDExpect {
	r.requiredTimes = n
	return r
}

func (r *GroupMock) UpdateByID(a0 context.Context, a1 uint64, a2 map[string]interface{}) (r0 error) {
	if len(r.expect.updateByID) == 0 {
		r0 = xerrors.New("cannot find mock method for Group.UpdateByID")
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
	r0 = xerrors.Errorf("invalid argument Group a0:[%+v] a1:[%+v] a2:[%+v]", a0, a1, a2)
	return
}

func (r *GroupExpect) UpdateByID(a0 context.Context, a1 uint64, a2 map[string]interface{}) *GroupUpdateByIDExpect {
	exp := &GroupUpdateByIDExpect{
		a0:      a0,
		a1:      a1,
		a2:      a2,
		actions: []func(a0 context.Context, a1 uint64, a2 map[string]interface{}){},
		expect:  r,
	}
	r.updateByID = append(r.updateByID, exp)
	return exp
}

type GroupUpdateByIDsExpect struct {
	expect        *GroupExpect
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

func (r *GroupUpdateByIDsExpect) Return(r0 error) *GroupUpdateByIDsExpect {
	r.r0 = r0
	return r
}

func (r *GroupUpdateByIDsExpect) Do(action func(a0 context.Context, a1 []uint64, a2 map[string]interface{})) *GroupUpdateByIDsExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *GroupUpdateByIDsExpect) OutOfOrder() *GroupUpdateByIDsExpect {
	r.isOutOfOrder = true
	return r
}

func (r *GroupUpdateByIDsExpect) AnyTimes() *GroupUpdateByIDsExpect {
	r.isAnyTimes = true
	return r
}

func (r *GroupUpdateByIDsExpect) Times(n int) *GroupUpdateByIDsExpect {
	r.requiredTimes = n
	return r
}

func (r *GroupMock) UpdateByIDs(a0 context.Context, a1 []uint64, a2 map[string]interface{}) (r0 error) {
	if len(r.expect.updateByIDs) == 0 {
		r0 = xerrors.New("cannot find mock method for Group.UpdateByIDs")
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
	r0 = xerrors.Errorf("invalid argument Group a0:[%+v] a1:[%+v] a2:[%+v]", a0, a1, a2)
	return
}

func (r *GroupExpect) UpdateByIDs(a0 context.Context, a1 []uint64, a2 map[string]interface{}) *GroupUpdateByIDsExpect {
	exp := &GroupUpdateByIDsExpect{
		a0:      a0,
		a1:      a1,
		a2:      a2,
		actions: []func(a0 context.Context, a1 []uint64, a2 map[string]interface{}){},
		expect:  r,
	}
	r.updateByIDs = append(r.updateByIDs, exp)
	return exp
}

type GroupDeleteByIDExpect struct {
	expect        *GroupExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context, a1 uint64)
	a0            context.Context
	a1            uint64
	r0            error
}

func (r *GroupDeleteByIDExpect) Return(r0 error) *GroupDeleteByIDExpect {
	r.r0 = r0
	return r
}

func (r *GroupDeleteByIDExpect) Do(action func(a0 context.Context, a1 uint64)) *GroupDeleteByIDExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *GroupDeleteByIDExpect) OutOfOrder() *GroupDeleteByIDExpect {
	r.isOutOfOrder = true
	return r
}

func (r *GroupDeleteByIDExpect) AnyTimes() *GroupDeleteByIDExpect {
	r.isAnyTimes = true
	return r
}

func (r *GroupDeleteByIDExpect) Times(n int) *GroupDeleteByIDExpect {
	r.requiredTimes = n
	return r
}

func (r *GroupMock) DeleteByID(a0 context.Context, a1 uint64) (r0 error) {
	if len(r.expect.deleteByID) == 0 {
		r0 = xerrors.New("cannot find mock method for Group.DeleteByID")
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
	r0 = xerrors.Errorf("invalid argument Group a0:[%+v] a1:[%+v]", a0, a1)
	return
}

func (r *GroupExpect) DeleteByID(a0 context.Context, a1 uint64) *GroupDeleteByIDExpect {
	exp := &GroupDeleteByIDExpect{
		a0:      a0,
		a1:      a1,
		actions: []func(a0 context.Context, a1 uint64){},
		expect:  r,
	}
	r.deleteByID = append(r.deleteByID, exp)
	return exp
}

type GroupDeleteByIDsExpect struct {
	expect        *GroupExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context, a1 []uint64)
	a0            context.Context
	a1            []uint64
	r0            error
}

func (r *GroupDeleteByIDsExpect) Return(r0 error) *GroupDeleteByIDsExpect {
	r.r0 = r0
	return r
}

func (r *GroupDeleteByIDsExpect) Do(action func(a0 context.Context, a1 []uint64)) *GroupDeleteByIDsExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *GroupDeleteByIDsExpect) OutOfOrder() *GroupDeleteByIDsExpect {
	r.isOutOfOrder = true
	return r
}

func (r *GroupDeleteByIDsExpect) AnyTimes() *GroupDeleteByIDsExpect {
	r.isAnyTimes = true
	return r
}

func (r *GroupDeleteByIDsExpect) Times(n int) *GroupDeleteByIDsExpect {
	r.requiredTimes = n
	return r
}

func (r *GroupMock) DeleteByIDs(a0 context.Context, a1 []uint64) (r0 error) {
	if len(r.expect.deleteByIDs) == 0 {
		r0 = xerrors.New("cannot find mock method for Group.DeleteByIDs")
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
	r0 = xerrors.Errorf("invalid argument Group a0:[%+v] a1:[%+v]", a0, a1)
	return
}

func (r *GroupExpect) DeleteByIDs(a0 context.Context, a1 []uint64) *GroupDeleteByIDsExpect {
	exp := &GroupDeleteByIDsExpect{
		a0:      a0,
		a1:      a1,
		actions: []func(a0 context.Context, a1 []uint64){},
		expect:  r,
	}
	r.deleteByIDs = append(r.deleteByIDs, exp)
	return exp
}

type GroupCountExpect struct {
	expect        *GroupExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context)
	a0            context.Context
	r0            int64
	r1            error
}

func (r *GroupCountExpect) Return(r0 int64, r1 error) *GroupCountExpect {
	r.r0 = r0
	r.r1 = r1
	return r
}

func (r *GroupCountExpect) Do(action func(a0 context.Context)) *GroupCountExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *GroupCountExpect) OutOfOrder() *GroupCountExpect {
	r.isOutOfOrder = true
	return r
}

func (r *GroupCountExpect) AnyTimes() *GroupCountExpect {
	r.isAnyTimes = true
	return r
}

func (r *GroupCountExpect) Times(n int) *GroupCountExpect {
	r.requiredTimes = n
	return r
}

func (r *GroupMock) Count(a0 context.Context) (r0 int64, r1 error) {
	if len(r.expect.count) == 0 {
		r1 = xerrors.New("cannot find mock method for Group.Count")
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
	r1 = xerrors.Errorf("invalid argument Group a0:[%+v]", a0)
	return
}

func (r *GroupExpect) Count(a0 context.Context) *GroupCountExpect {
	exp := &GroupCountExpect{
		a0:      a0,
		actions: []func(a0 context.Context){},
		expect:  r,
	}
	r.count = append(r.count, exp)
	return exp
}

type GroupDeleteExpect struct {
	expect        *GroupExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context, a1 *entity.Group)
	a0            context.Context
	a1            *entity.Group
	r0            error
}

func (r *GroupDeleteExpect) Return(r0 error) *GroupDeleteExpect {
	r.r0 = r0
	return r
}

func (r *GroupDeleteExpect) Do(action func(a0 context.Context, a1 *entity.Group)) *GroupDeleteExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *GroupDeleteExpect) OutOfOrder() *GroupDeleteExpect {
	r.isOutOfOrder = true
	return r
}

func (r *GroupDeleteExpect) AnyTimes() *GroupDeleteExpect {
	r.isAnyTimes = true
	return r
}

func (r *GroupDeleteExpect) Times(n int) *GroupDeleteExpect {
	r.requiredTimes = n
	return r
}

func (r *GroupMock) Delete(a0 context.Context, a1 *entity.Group) (r0 error) {
	if len(r.expect.delete) == 0 {
		r0 = xerrors.New("cannot find mock method for Group.Delete")
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
	r0 = xerrors.Errorf("invalid argument Group a0:[%+v] a1:[%+v]", a0, a1)
	return
}

func (r *GroupExpect) Delete(a0 context.Context, a1 *entity.Group) *GroupDeleteExpect {
	exp := &GroupDeleteExpect{
		a0:      a0,
		a1:      a1,
		actions: []func(a0 context.Context, a1 *entity.Group){},
		expect:  r,
	}
	r.delete = append(r.delete, exp)
	return exp
}

type GroupUpdateExpect struct {
	expect        *GroupExpect
	isOutOfOrder  bool
	isAnyTimes    bool
	requiredTimes int
	calledTimes   int
	actions       []func(a0 context.Context, a1 *entity.Group)
	a0            context.Context
	a1            *entity.Group
	r0            error
}

func (r *GroupUpdateExpect) Return(r0 error) *GroupUpdateExpect {
	r.r0 = r0
	return r
}

func (r *GroupUpdateExpect) Do(action func(a0 context.Context, a1 *entity.Group)) *GroupUpdateExpect {
	r.actions = append(r.actions, action)
	return r
}

func (r *GroupUpdateExpect) OutOfOrder() *GroupUpdateExpect {
	r.isOutOfOrder = true
	return r
}

func (r *GroupUpdateExpect) AnyTimes() *GroupUpdateExpect {
	r.isAnyTimes = true
	return r
}

func (r *GroupUpdateExpect) Times(n int) *GroupUpdateExpect {
	r.requiredTimes = n
	return r
}

func (r *GroupMock) Update(a0 context.Context, a1 *entity.Group) (r0 error) {
	if len(r.expect.update) == 0 {
		r0 = xerrors.New("cannot find mock method for Group.Update")
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
	r0 = xerrors.Errorf("invalid argument Group a0:[%+v] a1:[%+v]", a0, a1)
	return
}

func (r *GroupExpect) Update(a0 context.Context, a1 *entity.Group) *GroupUpdateExpect {
	exp := &GroupUpdateExpect{
		a0:      a0,
		a1:      a1,
		actions: []func(a0 context.Context, a1 *entity.Group){},
		expect:  r,
	}
	r.update = append(r.update, exp)
	return exp
}

type GroupExpect struct {
	toModel     []*GroupToModelExpect
	toModels    []*GroupToModelsExpect
	create      []*GroupCreateExpect
	creates     []*GroupCreatesExpect
	findAll     []*GroupFindAllExpect
	findByID    []*GroupFindByIDExpect
	findByIDs   []*GroupFindByIDsExpect
	updateByID  []*GroupUpdateByIDExpect
	updateByIDs []*GroupUpdateByIDsExpect
	deleteByID  []*GroupDeleteByIDExpect
	deleteByIDs []*GroupDeleteByIDsExpect
	count       []*GroupCountExpect
	delete      []*GroupDeleteExpect
	update      []*GroupUpdateExpect
}

func NewGroupExpect() *GroupExpect {
	return &GroupExpect{
		count:       []*GroupCountExpect{},
		create:      []*GroupCreateExpect{},
		creates:     []*GroupCreatesExpect{},
		delete:      []*GroupDeleteExpect{},
		deleteByID:  []*GroupDeleteByIDExpect{},
		deleteByIDs: []*GroupDeleteByIDsExpect{},
		findAll:     []*GroupFindAllExpect{},
		findByID:    []*GroupFindByIDExpect{},
		findByIDs:   []*GroupFindByIDsExpect{},
		toModel:     []*GroupToModelExpect{},
		toModels:    []*GroupToModelsExpect{},
		update:      []*GroupUpdateExpect{},
		updateByID:  []*GroupUpdateByIDExpect{},
		updateByIDs: []*GroupUpdateByIDsExpect{},
	}
}
