// Code generated by eevee. DO NOT EDIT!

package entity

import "time"

type User struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Sex       string    `json:"sex"`
	Age       int       `json:"age"`
	SkillID   uint64    `json:"skillID"`
	SkillRank int       `json:"skillRank"`
	GroupID   uint64    `json:"groupID"`
	WorldID   uint64    `json:"worldID"`
	FieldID   uint64    `json:"fieldID"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Users []*User

func (e Users) IDs() []uint64 {
	values := make([]uint64, 0, len(e))
	for _, value := range e {
		values = append(values, value.ID)
	}
	return values
}

func (e Users) Names() []string {
	values := make([]string, 0, len(e))
	for _, value := range e {
		values = append(values, value.Name)
	}
	return values
}

func (e Users) Sexes() []string {
	values := make([]string, 0, len(e))
	for _, value := range e {
		values = append(values, value.Sex)
	}
	return values
}

func (e Users) Ages() []int {
	values := make([]int, 0, len(e))
	for _, value := range e {
		values = append(values, value.Age)
	}
	return values
}

func (e Users) SkillIDs() []uint64 {
	values := make([]uint64, 0, len(e))
	for _, value := range e {
		values = append(values, value.SkillID)
	}
	return values
}

func (e Users) SkillRanks() []int {
	values := make([]int, 0, len(e))
	for _, value := range e {
		values = append(values, value.SkillRank)
	}
	return values
}

func (e Users) GroupIDs() []uint64 {
	values := make([]uint64, 0, len(e))
	for _, value := range e {
		values = append(values, value.GroupID)
	}
	return values
}

func (e Users) WorldIDs() []uint64 {
	values := make([]uint64, 0, len(e))
	for _, value := range e {
		values = append(values, value.WorldID)
	}
	return values
}

func (e Users) FieldIDs() []uint64 {
	values := make([]uint64, 0, len(e))
	for _, value := range e {
		values = append(values, value.FieldID)
	}
	return values
}

func (e Users) CreatedAts() []time.Time {
	values := make([]time.Time, 0, len(e))
	for _, value := range e {
		values = append(values, value.CreatedAt)
	}
	return values
}

func (e Users) UpdatedAts() []time.Time {
	values := make([]time.Time, 0, len(e))
	for _, value := range e {
		values = append(values, value.UpdatedAt)
	}
	return values
}
