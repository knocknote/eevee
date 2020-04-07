package dao

import (
	"context"
	"database/sql"
	"fmt"
	"relation/entity"
	"strings"

	"golang.org/x/xerrors"
)

type Skill interface {
	Count(context.Context) (int64, error)
	Create(context.Context, *entity.Skill) error
	Delete(context.Context, *entity.Skill) error
	DeleteByID(context.Context, uint64) error
	DeleteByIDs(context.Context, []uint64) error
	FindAll(context.Context) (entity.Skills, error)
	FindByID(context.Context, uint64) (*entity.Skill, error)
	FindByIDs(context.Context, []uint64) (entity.Skills, error)
	Update(context.Context, *entity.Skill) error
	UpdateByID(context.Context, uint64, map[string]interface{}) error
	UpdateByIDs(context.Context, []uint64, map[string]interface{}) error
}

type SkillImpl struct {
	tx *sql.Tx
}

func NewSkill(ctx context.Context, tx *sql.Tx) Skill {
	return &SkillImpl{tx: tx}
}

// generated by eevee
func (d *SkillImpl) Count(ctx context.Context) (r int64, e error) {
	var value int64
	query := "COUNT(*) FROM `skills`"
	if err := d.tx.QueryRowContext(ctx, query).Scan(&value); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, xerrors.Errorf("failure query %s: %w", query, err)
	}
	return value, nil
}

// generated by eevee
func (d *SkillImpl) Create(ctx context.Context, value *entity.Skill) (e error) {
	query := "INSERT INTO `skills` (`id`, `skill_effect`) VALUES (?, ?)"
	result, err := d.tx.ExecContext(ctx, query, value.ID, value.SkillEffect)
	if err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return xerrors.Errorf("cannot get LastInsertId: %w", err)
	}
	value.ID = uint64(id)
	return nil
}

// generated by eevee
func (d *SkillImpl) Delete(ctx context.Context, value *entity.Skill) (e error) {
	query := "DELETE FROM `skills` WHERE id = ?"
	if _, err := d.tx.ExecContext(ctx, query, value.ID); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *SkillImpl) DeleteByID(ctx context.Context, a0 uint64) (e error) {
	query := "DELETE FROM `skills` WHERE `id` = ?"
	if _, err := d.tx.ExecContext(ctx, query, a0); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *SkillImpl) DeleteByIDs(ctx context.Context, a0 []uint64) (e error) {
	args := []interface{}{}
	placeholders := make([]string, 0, len(a0))
	for _, v := range a0 {
		args = append(args, v)
		placeholders = append(placeholders, "?")
	}
	query := fmt.Sprintf("DELETE FROM `skills` WHERE `id` IN (%s)", strings.Join(placeholders, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *SkillImpl) FindAll(ctx context.Context) (r entity.Skills, e error) {
	values := entity.Skills{}
	query := "SELECT `id`, `skill_effect` FROM `skills`"
	rows, err := d.tx.QueryContext(ctx, query)
	if err != nil {
		return values, xerrors.Errorf("falure query %s: %w", query, err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			e = xerrors.Errorf("cannot close rows: %w", err)
		}
	}()
	for rows.Next() {
		var value entity.Skill
		if err := rows.Scan(&value.ID, &value.SkillEffect); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *SkillImpl) FindByID(ctx context.Context, a0 uint64) (r *entity.Skill, e error) {
	var value entity.Skill
	query := "SELECT `id`, `skill_effect` FROM `skills` WHERE `id` = ?"
	if err := d.tx.QueryRowContext(ctx, query, a0).Scan(
		&value.ID,
		&value.SkillEffect,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, xerrors.Errorf("failure query %s: %w", query, err)
	}
	return &value, nil
}

// generated by eevee
func (d *SkillImpl) FindByIDs(ctx context.Context, a0 []uint64) (r entity.Skills, e error) {
	values := entity.Skills{}
	query := "SELECT `id`, `skill_effect` FROM `skills` WHERE `id` IN (%s)"
	args := []interface{}{}
	placeholders := make([]string, 0, len(a0))
	for _, v := range a0 {
		args = append(args, v)
		placeholders = append(placeholders, "?")
	}
	selectQuery := fmt.Sprintf(query, strings.Join(placeholders, ", "))
	rows, err := d.tx.QueryContext(ctx, selectQuery, args...)
	if err != nil {
		return values, xerrors.Errorf("failure query %s: %w", query, err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			e = xerrors.Errorf("cannot close rows: %w", err)
		}
	}()
	for rows.Next() {
		var value entity.Skill
		if err := rows.Scan(
			&value.ID,
			&value.SkillEffect,
		); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *SkillImpl) Update(ctx context.Context, value *entity.Skill) (e error) {
	args := []interface{}{value.SkillEffect, value.ID}
	query := "UPDATE `skills` SET `skill_effect` = ? WHERE id = ?"
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *SkillImpl) UpdateByID(ctx context.Context, a0 uint64, updateMap map[string]interface{}) (e error) {
	columns := []string{}
	args := []interface{}{}
	for column, v := range updateMap {
		columns = append(columns, fmt.Sprintf("`%s` = ?", column))
		args = append(args, v)
	}
	args = append(args, a0)
	query := fmt.Sprintf("UPDATE `skills` SET %s WHERE `id` = ?", strings.Join(columns, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *SkillImpl) UpdateByIDs(ctx context.Context, a0 []uint64, updateMap map[string]interface{}) (e error) {
	columns := []string{}
	args := []interface{}{}
	for column, v := range updateMap {
		columns = append(columns, fmt.Sprintf("`%s` = ?", column))
		args = append(args, v)
	}
	for _, v := range a0 {
		args = append(args, v)
	}
	placeholders := make([]string, 0, len(a0))
	for range a0 {
		placeholders = append(placeholders, "?")
	}
	query := fmt.Sprintf("UPDATE `skills` SET %s WHERE `id` IN (%s)", strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}
