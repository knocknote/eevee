package dao

import (
	"context"
	"database/sql"
	"fmt"
	"simple/entity"
	"strings"

	"golang.org/x/xerrors"
)

type User interface {
	Count(context.Context) (int64, error)
	Create(context.Context, *entity.User) error
	Delete(context.Context, *entity.User) error
	DeleteByID(context.Context, uint64) error
	DeleteByIDs(context.Context, []uint64) error
	FindAll(context.Context) (entity.Users, error)
	FindByID(context.Context, uint64) (*entity.User, error)
	FindByIDs(context.Context, []uint64) (entity.Users, error)
	Update(context.Context, *entity.User) error
	UpdateByID(context.Context, uint64, map[string]interface{}) error
	UpdateByIDs(context.Context, []uint64, map[string]interface{}) error
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
	query := "INSERT INTO `users` (`id`, `name`) VALUES (?, ?)"
	result, err := d.tx.ExecContext(ctx, query, value.ID, value.Name)
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
func (d *UserImpl) FindAll(ctx context.Context) (r entity.Users, e error) {
	values := entity.Users{}
	query := "SELECT `id`, `name` FROM `users`"
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
		if err := rows.Scan(&value.ID, &value.Name); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *UserImpl) FindByID(ctx context.Context, a0 uint64) (r *entity.User, e error) {
	var value entity.User
	query := "SELECT `id`, `name` FROM `users` WHERE `id` = ?"
	if err := d.tx.QueryRowContext(ctx, query, a0).Scan(
		&value.ID,
		&value.Name,
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
	query := "SELECT `id`, `name` FROM `users` WHERE `id` IN (%s)"
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
		); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *UserImpl) Update(ctx context.Context, value *entity.User) (e error) {
	args := []interface{}{value.Name, value.ID}
	query := "UPDATE `users` SET `name` = ? WHERE id = ?"
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
