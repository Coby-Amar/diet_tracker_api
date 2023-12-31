// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0
// source: copyfrom.go

package database

import (
	"context"
)

// iteratorForCreateReportEntries implements pgx.CopyFromSource.
type iteratorForCreateReportEntries struct {
	rows                 []CreateReportEntriesParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateReportEntries) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateReportEntries) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].ProductID,
		r.rows[0].ReportID,
		r.rows[0].Amount,
		r.rows[0].Carbohydrates,
		r.rows[0].Proteins,
		r.rows[0].Fats,
	}, nil
}

func (r iteratorForCreateReportEntries) Err() error {
	return nil
}

func (q *Queries) CreateReportEntries(ctx context.Context, arg []CreateReportEntriesParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"_report_entries"}, []string{"_product_id", "_report_id", "_amount", "_carbohydrates", "_proteins", "_fats"}, &iteratorForCreateReportEntries{rows: arg})
}
