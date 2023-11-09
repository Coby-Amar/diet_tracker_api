// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: reports.sql

package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createReport = `-- name: CreateReport :one
INSERT INTO _reports(
    _date,
    _amout_of_entries,
    _carbohydrates_total,
    _proteins_total,
    _fats_total,
    _user_id
)
VALUES ($1,$2,$3,$4,$5,$6)
RETURNING _id, _created_at, _updated_at, _date, _amout_of_entries, _carbohydrates_total, _proteins_total, _fats_total, _user_id
`

type CreateReportParams struct {
	Date               pgtype.Date    `json:"date"`
	AmoutOfEntries     int16          `json:"numberOfEntries"`
	CarbohydratesTotal pgtype.Numeric `json:"carbohydratesTotal"`
	ProteinsTotal      pgtype.Numeric `json:"proteinsTotal"`
	FatsTotal          pgtype.Numeric `json:"fatsTotal"`
	UserID             pgtype.UUID    `json:"-"`
}

func (q *Queries) CreateReport(ctx context.Context, arg CreateReportParams) (Report, error) {
	row := q.db.QueryRow(ctx, createReport,
		arg.Date,
		arg.AmoutOfEntries,
		arg.CarbohydratesTotal,
		arg.ProteinsTotal,
		arg.FatsTotal,
		arg.UserID,
	)
	var i Report
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Date,
		&i.AmoutOfEntries,
		&i.CarbohydratesTotal,
		&i.ProteinsTotal,
		&i.FatsTotal,
		&i.UserID,
	)
	return i, err
}

const deleteReport = `-- name: DeleteReport :exec
DELETE FROM _reports
WHERE _reports._id = $1
`

func (q *Queries) DeleteReport(ctx context.Context, ID pgtype.UUID) error {
	_, err := q.db.Exec(ctx, deleteReport, ID)
	return err
}

const getAllUserReports = `-- name: GetAllUserReports :many
SELECT _id, _created_at, _updated_at, _date, _amout_of_entries, _carbohydrates_total, _proteins_total, _fats_total, _user_id FROM _reports
WHERE _user_id = $1
`

func (q *Queries) GetAllUserReports(ctx context.Context, UserID pgtype.UUID) ([]Report, error) {
	rows, err := q.db.Query(ctx, getAllUserReports, UserID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Report
	for rows.Next() {
		var i Report
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Date,
			&i.AmoutOfEntries,
			&i.CarbohydratesTotal,
			&i.ProteinsTotal,
			&i.FatsTotal,
			&i.UserID,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getReportByID = `-- name: GetReportByID :one
SELECT _id, _created_at, _updated_at, _date, _amout_of_entries, _carbohydrates_total, _proteins_total, _fats_total, _user_id FROM _reports
WHERE _id = $1
`

func (q *Queries) GetReportByID(ctx context.Context, ID pgtype.UUID) (Report, error) {
	row := q.db.QueryRow(ctx, getReportByID, ID)
	var i Report
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Date,
		&i.AmoutOfEntries,
		&i.CarbohydratesTotal,
		&i.ProteinsTotal,
		&i.FatsTotal,
		&i.UserID,
	)
	return i, err
}

const updateReport = `-- name: UpdateReport :one
UPDATE _reports
SET
    _amout_of_entries = $2,
    _carbohydrates_total = $3,
    _proteins_total = $4,
    _fats_total = $5,
    _updated_at = NOW()
WHERE _reports._id = $1
RETURNING _id, _created_at, _updated_at, _date, _amout_of_entries, _carbohydrates_total, _proteins_total, _fats_total, _user_id
`

type UpdateReportParams struct {
	ID                 pgtype.UUID    `json:"id"`
	AmoutOfEntries     int16          `json:"numberOfEntries"`
	CarbohydratesTotal pgtype.Numeric `json:"carbohydratesTotal"`
	ProteinsTotal      pgtype.Numeric `json:"proteinsTotal"`
	FatsTotal          pgtype.Numeric `json:"fatsTotal"`
}

func (q *Queries) UpdateReport(ctx context.Context, arg UpdateReportParams) (Report, error) {
	row := q.db.QueryRow(ctx, updateReport,
		arg.ID,
		arg.AmoutOfEntries,
		arg.CarbohydratesTotal,
		arg.ProteinsTotal,
		arg.FatsTotal,
	)
	var i Report
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Date,
		&i.AmoutOfEntries,
		&i.CarbohydratesTotal,
		&i.ProteinsTotal,
		&i.FatsTotal,
		&i.UserID,
	)
	return i, err
}