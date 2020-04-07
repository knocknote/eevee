// Code generated by eevee. DO NOT EDIT!

package entity

import (
	"go.knocknote.io/rapidash"
	"golang.org/x/xerrors"
)

type Field struct {
	ID         uint64 `json:"id"`
	Name       string `json:"name"`
	LocationX  int    `json:"locationX"`
	LocationY  int    `json:"locationY"`
	ObjectNum  int    `json:"objectNum"`
	Level      int    `json:"level"`
	Difficulty int    `json:"difficulty"`
}

type Fields []*Field

func (e Fields) IDs() []uint64 {
	values := make([]uint64, 0, len(e))
	for _, value := range e {
		values = append(values, value.ID)
	}
	return values
}

func (e Fields) Names() []string {
	values := make([]string, 0, len(e))
	for _, value := range e {
		values = append(values, value.Name)
	}
	return values
}

func (e Fields) LocationXes() []int {
	values := make([]int, 0, len(e))
	for _, value := range e {
		values = append(values, value.LocationX)
	}
	return values
}

func (e Fields) LocationIes() []int {
	values := make([]int, 0, len(e))
	for _, value := range e {
		values = append(values, value.LocationY)
	}
	return values
}

func (e Fields) ObjectNums() []int {
	values := make([]int, 0, len(e))
	for _, value := range e {
		values = append(values, value.ObjectNum)
	}
	return values
}

func (e Fields) Levels() []int {
	values := make([]int, 0, len(e))
	for _, value := range e {
		values = append(values, value.Level)
	}
	return values
}

func (e Fields) Difficulties() []int {
	values := make([]int, 0, len(e))
	for _, value := range e {
		values = append(values, value.Difficulty)
	}
	return values
}

func (e *Field) Struct() *rapidash.Struct {
	return rapidash.NewStruct("fields").
		FieldUint64("id").
		FieldString("name").
		FieldInt("location_x").
		FieldInt("location_y").
		FieldInt("object_num").
		FieldInt("level").
		FieldInt("difficulty")
}

func (e *Field) EncodeRapidash(enc rapidash.Encoder) error {
	if e.ID != 0 {
		enc.Uint64("id", e.ID)
	}
	enc.String("name", e.Name)
	enc.Int("location_x", e.LocationX)
	enc.Int("location_y", e.LocationY)
	enc.Int("object_num", e.ObjectNum)
	enc.Int("level", e.Level)
	enc.Int("difficulty", e.Difficulty)
	if err := enc.Error(); err != nil {
		return xerrors.Errorf("failed to encode: %w", err)
	}
	return nil
}

func (e *Fields) EncodeRapidash(enc rapidash.Encoder) error {
	for _, v := range *e {
		if err := v.EncodeRapidash(enc.New()); err != nil {
			return xerrors.Errorf("failed to encode: %w", err)
		}
	}
	return nil
}

func (e *Field) DecodeRapidash(dec rapidash.Decoder) error {
	e.ID = dec.Uint64("id")
	e.Name = dec.String("name")
	e.LocationX = dec.Int("location_x")
	e.LocationY = dec.Int("location_y")
	e.ObjectNum = dec.Int("object_num")
	e.Level = dec.Int("level")
	e.Difficulty = dec.Int("difficulty")
	if err := dec.Error(); err != nil {
		return xerrors.Errorf("failed to decode: %w", err)
	}
	return nil
}

func (e *Fields) DecodeRapidash(dec rapidash.Decoder) error {
	decLen := dec.Len()
	values := make(Fields, decLen)
	for i := 0; i < decLen; i++ {
		var v Field
		if err := v.DecodeRapidash(dec.At(i)); err != nil {
			return xerrors.Errorf("failed to decode: %w", err)
		}
		values[i] = &v
	}
	*e = values
	return nil
}
