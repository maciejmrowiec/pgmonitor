package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func TupleFreePercent(db *sql.DB, tableName string) (float64, error) {
	query := "SELECT free_percent FROM pgstattuple($1)"

	var percent sql.NullFloat64
	if err := db.QueryRow(query, tableName).Scan(&percent); err != nil {
		return 0, err
	}

	return percent.Float64, nil
}

func TupleActivePercent(db *sql.DB, tableName string) (float64, error) {
	query := "SELECT tuple_percent FROM pgstattuple($1)"

	var percent sql.NullFloat64
	if err := db.QueryRow(query, tableName).Scan(&percent); err != nil {
		return 0, err
	}

	return percent.Float64, nil
}

func TupleDeadPercent(db *sql.DB, tableName string) (float64, error) {
	query := "SELECT dead_tuple_percent FROM pgstattuple($1)"

	var percent sql.NullFloat64
	if err := db.QueryRow(query, tableName).Scan(&percent); err != nil {
		return 0, err
	}

	return percent.Float64, nil
}
