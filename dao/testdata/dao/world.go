package dao

import (
	"context"
	"database/sql"
	"eevee/entity"
	"fmt"
	"strings"

	"golang.org/x/xerrors"
)

type World interface {
	Count(context.Context) (int64, error)
	Create(context.Context, *entity.World) error
	Delete(context.Context, *entity.World) error
	DeleteByID(context.Context, uint64) error
	DeleteByIDs(context.Context, []uint64) error
	FindAll(context.Context) (entity.Worlds, error)
	FindByID(context.Context, uint64) (*entity.World, error)
	FindByIDs(context.Context, []uint64) (entity.Worlds, error)
	Update(context.Context, *entity.World) error
	UpdateByID(context.Context, uint64, map[string]interface{}) error
	UpdateByIDs(context.Context, []uint64, map[string]interface{}) error
}

type WorldImpl struct {
	tx *sql.Tx
}

func NewWorld(ctx context.Context, tx *sql.Tx) World {
	return &WorldImpl{tx: tx}
}

// generated by eevee
func (d *WorldImpl) Count(ctx context.Context) (r int64, e error) {
	var value int64
	query := "COUNT(*) FROM `worlds`"
	if err := d.tx.QueryRowContext(ctx, query).Scan(&value); err != nil {
		if err == sql.ErrNoRows {
			return 0, nil
		}
		return 0, xerrors.Errorf("failure query %s: %w", query, err)
	}
	return value, nil
}

// generated by eevee
func (d *WorldImpl) Create(ctx context.Context, value *entity.World) (e error) {
	query := "INSERT INTO `worlds` (`id`, `name`) VALUES (?, ?)"
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
func (d *WorldImpl) Delete(ctx context.Context, value *entity.World) (e error) {
	query := "DELETE FROM `worlds` WHERE id = ?"
	if _, err := d.tx.ExecContext(ctx, query, value.ID); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *WorldImpl) DeleteByID(ctx context.Context, a0 uint64) (e error) {
	query := "DELETE FROM `worlds` WHERE `id` = ?"
	if _, err := d.tx.ExecContext(ctx, query, a0); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *WorldImpl) DeleteByIDs(ctx context.Context, a0 []uint64) (e error) {
	args := []interface{}{}
	placeholders := make([]string, 0, len(a0))
	for _, v := range a0 {
		args = append(args, v)
		placeholders = append(placeholders, "?")
	}
	query := fmt.Sprintf("DELETE FROM `worlds` WHERE `id` IN (%s)", strings.Join(placeholders, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *WorldImpl) FindAll(ctx context.Context) (r entity.Worlds, e error) {
	values := entity.Worlds{}
	query := "SELECT `id`, `name` FROM `worlds`"
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
		var value entity.World
		if err := rows.Scan(&value.ID, &value.Name); err != nil {
			return values, xerrors.Errorf("cannot scan value: %w", err)
		}
		values = append(values, &value)
	}
	return values, nil
}

// generated by eevee
func (d *WorldImpl) FindByID(ctx context.Context, a0 uint64) (r *entity.World, e error) {
	var value entity.World
	query := "SELECT `id`, `name` FROM `worlds` WHERE `id` = ?"
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
func (d *WorldImpl) FindByIDs(ctx context.Context, a0 []uint64) (r entity.Worlds, e error) {
	values := entity.Worlds{}
	query := "SELECT `id`, `name` FROM `worlds` WHERE `id` IN (%s)"
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
		var value entity.World
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
func (d *WorldImpl) Update(ctx context.Context, value *entity.World) (e error) {
	args := []interface{}{value.Name, value.ID}
	query := "UPDATE `worlds` SET `name` = ? WHERE id = ?"
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *WorldImpl) UpdateByID(ctx context.Context, a0 uint64, updateMap map[string]interface{}) (e error) {
	columns := []string{}
	args := []interface{}{}
	for column, v := range updateMap {
		columns = append(columns, fmt.Sprintf("`%s` = ?", column))
		args = append(args, v)
	}
	args = append(args, a0)
	query := fmt.Sprintf("UPDATE `worlds` SET %s WHERE `id` = ?", strings.Join(columns, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}

// generated by eevee
func (d *WorldImpl) UpdateByIDs(ctx context.Context, a0 []uint64, updateMap map[string]interface{}) (e error) {
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
	query := fmt.Sprintf("UPDATE `worlds` SET %s WHERE `id` IN (%s)", strings.Join(columns, ", "), strings.Join(placeholders, ", "))
	if _, err := d.tx.ExecContext(ctx, query, args...); err != nil {
		return xerrors.Errorf("failure query %s: %w", query, err)
	}
	return nil
}
