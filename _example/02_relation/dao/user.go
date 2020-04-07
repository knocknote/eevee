package dao

import (
	"context"
	"database/sql"
	"fmt"
	"relation/entity"
	"strings"

	"golang.org/x/xerrors"
)

type User interface {
	Count(context.Context) (int64, error)
	Create(context.Context, *entity.User) error
	Delete(context.Context, *entity.User) error
	DeleteByGroupID(context.Context, uint64) error
	DeleteByGroupIDs(context.Context, []uint64) error
	DeleteByID(context.Context, uint64) error
	DeleteByIDs(context.Context, []uint64) error
	DeleteByName(context.Context, string) error
	DeleteByNames(context.Context, []string) error
	DeleteBySkillIDAndSkillRank(context.Context, uint64, int) error
	DeleteByWorldIDAndFieldID(context.Context, uint64, uint64) error
	FindAll(context.Context) (entity.Users, error)
	FindByGroupID(context.Context, uint64) (entity.Users, error)
	FindByGroupIDs(context.Context, []uint64) (entity.Users, error)
	FindByID(context.Context, uint64) (*entity.User, error)
	FindByIDs(context.Context, []uint64) (entity.Users, error)
	FindByName(context.Context, string) (*entity.User, error)
	FindByNames(context.Context, []string) (entity.Users, error)
	FindBySkillID(context.Context, uint64) (entity.Users, error)
	FindBySkillIDAndSkillRank(context.Context, uint64, int) (*entity.User, error)
	FindBySkillIDs(context.Context, []uint64) (entity.Users, error)
	FindByWorldID(context.Context, uint64) (entity.Users, error)
	FindByWorldIDAndFieldID(context.Context, uint64, uint64) (entity.Users, error)
	FindByWorldIDs(context.Context, []uint64) (entity.Users, error)
	Update(context.Context, *entity.User) error
	UpdateByGroupID(context.Context, uint64, map[string]interface{}) error
	UpdateByGroupIDs(context.Context, []uint64, map[string]interface{}) error
	UpdateByID(context.Context, uint64, map[string]interface{}) error
	UpdateByIDs(context.Context, []uint64, map[string]interface{}) error
	UpdateByName(context.Context, string, map[string]interface{}) error
	UpdateByNames(context.Context, []string, map[string]interface{}) error
	UpdateBySkillIDAndSkillRank(context.Context, uint64, int, map[string]interface{}) error
	UpdateByWorldIDAndFieldID(context.Context, uint64, uint64, map[string]interface{}) error
}

type UserImpl struct {
	tx *sql.Tx
}

func NewUser(ctx context.Context, tx *sql.Tx) User {
	return &UserImpl{tx: tx}
}

// generated by eevee
func (d *UserImpl) Count(ctx context.Context) (r int64, e error) {
	var value int64
	query := "COUNT(*) FROM `users`"
	if err := d.tx.QueryRowContext(ctx, query).Scan(&value); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, xerrors.Errorf("failure query %s: %w", query, err)
	}
	return value, nil
}

// generated by eevee
func (d *UserImpl) Create(ctx context.Context, value *entity.User) (e error) {
	query := "INSERT INTO `users` (`id`, `name`, `sex`, `age`, `skill_id`, `skill_rank`, `group_id`, `world_id`, `field_id`) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	result, err := d.tx.ExecContext(ctx, query, value.ID, value.Name, value.Sex, value.Age, value.SkillID, value.SkillRank, value.GroupID, value.WorldID, value.FieldID)
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
func (d *UserImpl) Delete(ctx context.Context, value *entity.User) (e error) {
	query := "DELETE FROM `users` WHERE id = ?"
	if _, err := d.tx.ExecContext(ctx, query, value.ID); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) DeleteByGroupID(ctx context.Context, a0 uint64) (e error) {
	query := "DELETE FROM `users` WHERE `group_id` = ?"
	if _, err := d.tx.ExecContext(ctx, query, a0); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) DeleteByGroupIDs(ctx context.Context, a0 []uint64) (e error) {
	args := []interface{}{}
	placeholders := make([]string, 0, len(a0))
	for _, v := range a0 {
		args = append(args, v)
		placeholders = append(placeholders, "?")
	}
	query := fmt.Sprintf("DELETE FROM `users` WHERE `group_id` IN (%s)", strings.Join(placeholders, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) DeleteByID(ctx context.Context, a0 uint64) (e error) {
	query := "DELETE FROM `users` WHERE `id` = ?"
	if _, err := d.tx.ExecContext(ctx, query, a0); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) DeleteByIDs(ctx context.Context, a0 []uint64) (e error) {
	args := []interface{}{}
	placeholders := make([]string, 0, len(a0))
	for _, v := range a0 {
		args = append(args, v)
		placeholders = append(placeholders, "?")
	}
	query := fmt.Sprintf("DELETE FROM `users` WHERE `id` IN (%s)", strings.Join(placeholders, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) DeleteByName(ctx context.Context, a0 string) (e error) {
	query := "DELETE FROM `users` WHERE `name` = ?"
	if _, err := d.tx.ExecContext(ctx, query, a0); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) DeleteByNames(ctx context.Context, a0 []string) (e error) {
	args := []interface{}{}
	placeholders := make([]string, 0, len(a0))
	for _, v := range a0 {
		args = append(args, v)
		placeholders = append(placeholders, "?")
	}
	query := fmt.Sprintf("DELETE FROM `users` WHERE `name` IN (%s)", strings.Join(placeholders, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) DeleteBySkillIDAndSkillRank(ctx context.Context, a0 uint64, a1 int) (e error) {
	query := "DELETE FROM `users` WHERE `skill_id` = ? AND `skill_rank` = ?"
	if _, err := d.tx.ExecContext(ctx, query, a0, a1); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) DeleteByWorldIDAndFieldID(ctx context.Context, a0 uint64, a1 uint64) (e error) {
	query := "DELETE FROM `users` WHERE `world_id` = ? AND `field_id` = ?"
	if _, err := d.tx.ExecContext(ctx, query, a0, a1); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) FindAll(ctx context.Context) (r entity.Users, e error) {
	values := entity.Users{}
	query := "SELECT `id`, `name`, `sex`, `age`, `skill_id`, `skill_rank`, `group_id`, `world_id`, `field_id` FROM `users`"
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
		var value entity.User
		if err := rows.Scan(&value.ID, &value.Name, &value.Sex, &value.Age, &value.SkillID, &value.SkillRank, &value.GroupID, &value.WorldID, &value.FieldID); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *UserImpl) FindByGroupID(ctx context.Context, a0 uint64) (r entity.Users, e error) {
	values := entity.Users{}
	query := "SELECT `id`, `name`, `sex`, `age`, `skill_id`, `skill_rank`, `group_id`, `world_id`, `field_id` FROM `users` WHERE `group_id` = ?"
	rows, err := d.tx.QueryContext(ctx, query, a0)
	if err != nil {
		return values, xerrors.Errorf("failure query %s: %w", query, err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			e = xerrors.Errorf("cannot close rows: %w", err)
		}
	}()
	for rows.Next() {
		var value entity.User
		if err := rows.Scan(
			&value.ID,
			&value.Name,
			&value.Sex,
			&value.Age,
			&value.SkillID,
			&value.SkillRank,
			&value.GroupID,
			&value.WorldID,
			&value.FieldID,
		); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *UserImpl) FindByGroupIDs(ctx context.Context, a0 []uint64) (r entity.Users, e error) {
	values := entity.Users{}
	query := "SELECT `id`, `name`, `sex`, `age`, `skill_id`, `skill_rank`, `group_id`, `world_id`, `field_id` FROM `users` WHERE `group_id` IN (%s)"
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
		var value entity.User
		if err := rows.Scan(
			&value.ID,
			&value.Name,
			&value.Sex,
			&value.Age,
			&value.SkillID,
			&value.SkillRank,
			&value.GroupID,
			&value.WorldID,
			&value.FieldID,
		); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *UserImpl) FindByID(ctx context.Context, a0 uint64) (r *entity.User, e error) {
	var value entity.User
	query := "SELECT `id`, `name`, `sex`, `age`, `skill_id`, `skill_rank`, `group_id`, `world_id`, `field_id` FROM `users` WHERE `id` = ?"
	if err := d.tx.QueryRowContext(ctx, query, a0).Scan(
		&value.ID,
		&value.Name,
		&value.Sex,
		&value.Age,
		&value.SkillID,
		&value.SkillRank,
		&value.GroupID,
		&value.WorldID,
		&value.FieldID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, xerrors.Errorf("failure query %s: %w", query, err)
	}
	return &value, nil
}

// generated by eevee
func (d *UserImpl) FindByIDs(ctx context.Context, a0 []uint64) (r entity.Users, e error) {
	values := entity.Users{}
	query := "SELECT `id`, `name`, `sex`, `age`, `skill_id`, `skill_rank`, `group_id`, `world_id`, `field_id` FROM `users` WHERE `id` IN (%s)"
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
		var value entity.User
		if err := rows.Scan(
			&value.ID,
			&value.Name,
			&value.Sex,
			&value.Age,
			&value.SkillID,
			&value.SkillRank,
			&value.GroupID,
			&value.WorldID,
			&value.FieldID,
		); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *UserImpl) FindByName(ctx context.Context, a0 string) (r *entity.User, e error) {
	var value entity.User
	query := "SELECT `id`, `name`, `sex`, `age`, `skill_id`, `skill_rank`, `group_id`, `world_id`, `field_id` FROM `users` WHERE `name` = ?"
	if err := d.tx.QueryRowContext(ctx, query, a0).Scan(
		&value.ID,
		&value.Name,
		&value.Sex,
		&value.Age,
		&value.SkillID,
		&value.SkillRank,
		&value.GroupID,
		&value.WorldID,
		&value.FieldID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, xerrors.Errorf("failure query %s: %w", query, err)
	}
	return &value, nil
}

// generated by eevee
func (d *UserImpl) FindByNames(ctx context.Context, a0 []string) (r entity.Users, e error) {
	values := entity.Users{}
	query := "SELECT `id`, `name`, `sex`, `age`, `skill_id`, `skill_rank`, `group_id`, `world_id`, `field_id` FROM `users` WHERE `name` IN (%s)"
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
		var value entity.User
		if err := rows.Scan(
			&value.ID,
			&value.Name,
			&value.Sex,
			&value.Age,
			&value.SkillID,
			&value.SkillRank,
			&value.GroupID,
			&value.WorldID,
			&value.FieldID,
		); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *UserImpl) FindBySkillID(ctx context.Context, a0 uint64) (r entity.Users, e error) {
	values := entity.Users{}
	query := "SELECT `id`, `name`, `sex`, `age`, `skill_id`, `skill_rank`, `group_id`, `world_id`, `field_id` FROM `users` WHERE `skill_id` = ?"
	rows, err := d.tx.QueryContext(ctx, query, a0)
	if err != nil {
		return values, xerrors.Errorf("failure query %s: %w", query, err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			e = xerrors.Errorf("cannot close rows: %w", err)
		}
	}()
	for rows.Next() {
		var value entity.User
		if err := rows.Scan(
			&value.ID,
			&value.Name,
			&value.Sex,
			&value.Age,
			&value.SkillID,
			&value.SkillRank,
			&value.GroupID,
			&value.WorldID,
			&value.FieldID,
		); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *UserImpl) FindBySkillIDAndSkillRank(ctx context.Context, a0 uint64, a1 int) (r *entity.User, e error) {
	var value entity.User
	query := "SELECT `id`, `name`, `sex`, `age`, `skill_id`, `skill_rank`, `group_id`, `world_id`, `field_id` FROM `users` WHERE `skill_id` = ? AND `skill_rank` = ?"
	if err := d.tx.QueryRowContext(ctx, query, a0, a1).Scan(
		&value.ID,
		&value.Name,
		&value.Sex,
		&value.Age,
		&value.SkillID,
		&value.SkillRank,
		&value.GroupID,
		&value.WorldID,
		&value.FieldID,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, xerrors.Errorf("failure query %s: %w", query, err)
	}
	return &value, nil
}

// generated by eevee
func (d *UserImpl) FindBySkillIDs(ctx context.Context, a0 []uint64) (r entity.Users, e error) {
	values := entity.Users{}
	query := "SELECT `id`, `name`, `sex`, `age`, `skill_id`, `skill_rank`, `group_id`, `world_id`, `field_id` FROM `users` WHERE `skill_id` IN (%s)"
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
		var value entity.User
		if err := rows.Scan(
			&value.ID,
			&value.Name,
			&value.Sex,
			&value.Age,
			&value.SkillID,
			&value.SkillRank,
			&value.GroupID,
			&value.WorldID,
			&value.FieldID,
		); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *UserImpl) FindByWorldID(ctx context.Context, a0 uint64) (r entity.Users, e error) {
	values := entity.Users{}
	query := "SELECT `id`, `name`, `sex`, `age`, `skill_id`, `skill_rank`, `group_id`, `world_id`, `field_id` FROM `users` WHERE `world_id` = ?"
	rows, err := d.tx.QueryContext(ctx, query, a0)
	if err != nil {
		return values, xerrors.Errorf("failure query %s: %w", query, err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			e = xerrors.Errorf("cannot close rows: %w", err)
		}
	}()
	for rows.Next() {
		var value entity.User
		if err := rows.Scan(
			&value.ID,
			&value.Name,
			&value.Sex,
			&value.Age,
			&value.SkillID,
			&value.SkillRank,
			&value.GroupID,
			&value.WorldID,
			&value.FieldID,
		); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *UserImpl) FindByWorldIDAndFieldID(ctx context.Context, a0 uint64, a1 uint64) (r entity.Users, e error) {
	values := entity.Users{}
	query := "SELECT `id`, `name`, `sex`, `age`, `skill_id`, `skill_rank`, `group_id`, `world_id`, `field_id` FROM `users` WHERE `world_id` = ? AND `field_id` = ?"
	rows, err := d.tx.QueryContext(ctx, query, a0, a1)
	if err != nil {
		return values, xerrors.Errorf("failure query %s: %w", query, err)
	}
	defer func() {
		if err := rows.Close(); err != nil {
			e = xerrors.Errorf("cannot close rows: %w", err)
		}
	}()
	for rows.Next() {
		var value entity.User
		if err := rows.Scan(
			&value.ID,
			&value.Name,
			&value.Sex,
			&value.Age,
			&value.SkillID,
			&value.SkillRank,
			&value.GroupID,
			&value.WorldID,
			&value.FieldID,
		); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *UserImpl) FindByWorldIDs(ctx context.Context, a0 []uint64) (r entity.Users, e error) {
	values := entity.Users{}
	query := "SELECT `id`, `name`, `sex`, `age`, `skill_id`, `skill_rank`, `group_id`, `world_id`, `field_id` FROM `users` WHERE `world_id` IN (%s)"
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
		var value entity.User
		if err := rows.Scan(
			&value.ID,
			&value.Name,
			&value.Sex,
			&value.Age,
			&value.SkillID,
			&value.SkillRank,
			&value.GroupID,
			&value.WorldID,
			&value.FieldID,
		); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *UserImpl) Update(ctx context.Context, value *entity.User) (e error) {
	args := []interface{}{value.Name, value.Sex, value.Age, value.SkillID, value.SkillRank, value.GroupID, value.WorldID, value.FieldID, value.ID}
	query := "UPDATE `users` SET `name` = ?, `sex` = ?, `age` = ?, `skill_id` = ?, `skill_rank` = ?, `group_id` = ?, `world_id` = ?, `field_id` = ? WHERE id = ?"
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) UpdateByGroupID(ctx context.Context, a0 uint64, updateMap map[string]interface{}) (e error) {
	columns := []string{}
	args := []interface{}{}
	for column, v := range updateMap {
		columns = append(columns, fmt.Sprintf("`%s` = ?", column))
		args = append(args, v)
	}
	args = append(args, a0)
	query := fmt.Sprintf("UPDATE `users` SET %s WHERE `group_id` = ?", strings.Join(columns, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) UpdateByGroupIDs(ctx context.Context, a0 []uint64, updateMap map[string]interface{}) (e error) {
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
	query := fmt.Sprintf("UPDATE `users` SET %s WHERE `group_id` IN (%s)", strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) UpdateByID(ctx context.Context, a0 uint64, updateMap map[string]interface{}) (e error) {
	columns := []string{}
	args := []interface{}{}
	for column, v := range updateMap {
		columns = append(columns, fmt.Sprintf("`%s` = ?", column))
		args = append(args, v)
	}
	args = append(args, a0)
	query := fmt.Sprintf("UPDATE `users` SET %s WHERE `id` = ?", strings.Join(columns, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) UpdateByIDs(ctx context.Context, a0 []uint64, updateMap map[string]interface{}) (e error) {
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
	query := fmt.Sprintf("UPDATE `users` SET %s WHERE `id` IN (%s)", strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) UpdateByName(ctx context.Context, a0 string, updateMap map[string]interface{}) (e error) {
	columns := []string{}
	args := []interface{}{}
	for column, v := range updateMap {
		columns = append(columns, fmt.Sprintf("`%s` = ?", column))
		args = append(args, v)
	}
	args = append(args, a0)
	query := fmt.Sprintf("UPDATE `users` SET %s WHERE `name` = ?", strings.Join(columns, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) UpdateByNames(ctx context.Context, a0 []string, updateMap map[string]interface{}) (e error) {
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
	query := fmt.Sprintf("UPDATE `users` SET %s WHERE `name` IN (%s)", strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) UpdateBySkillIDAndSkillRank(ctx context.Context, a0 uint64, a1 int, updateMap map[string]interface{}) (e error) {
	columns := []string{}
	args := []interface{}{}
	for column, v := range updateMap {
		columns = append(columns, fmt.Sprintf("`%s` = ?", column))
		args = append(args, v)
	}
	args = append(args, a0)
	args = append(args, a1)
	query := fmt.Sprintf("UPDATE `users` SET %s WHERE `skill_id` = ? AND `skill_rank` = ?", strings.Join(columns, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *UserImpl) UpdateByWorldIDAndFieldID(ctx context.Context, a0 uint64, a1 uint64, updateMap map[string]interface{}) (e error) {
	columns := []string{}
	args := []interface{}{}
	for column, v := range updateMap {
		columns = append(columns, fmt.Sprintf("`%s` = ?", column))
		args = append(args, v)
	}
	args = append(args, a0)
	args = append(args, a1)
	query := fmt.Sprintf("UPDATE `users` SET %s WHERE `world_id` = ? AND `field_id` = ?", strings.Join(columns, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}
