package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func TupleFreePercent(db *sql.DB, tableName string) (float64, error) {
	sql := "SELECT free_percent FROM pgstattuple($1)"

	var active_percent float64
	if err := db.QueryRow(sql, tableName).Scan(&active_percent); err != nil {
		return 0, err
	}

	return active_percent, nil
}

func TupleActivePercent(db *sql.DB, tableName string) (float64, error) {
	sql := "SELECT tuple_percent FROM pgstattuple($1)"

	var active_percent float64
	if err := db.QueryRow(sql, tableName).Scan(&active_percent); err != nil {
		return 0, err
	}

	return active_percent, nil
}

func TupleDeadPercent(db *sql.DB, tableName string) (float64, error) {
	sql := "SELECT dead_tuple_percent FROM pgstattuple($1)"

	var active_percent float64
	if err := db.QueryRow(sql, tableName).Scan(&active_percent); err != nil {
		return 0, err
	}

	return active_percent, nil
}
