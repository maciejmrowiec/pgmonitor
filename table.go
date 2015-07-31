package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func UserTables(db *sql.DB) ([]string, error) {
	query := "SELECT relname FROM pg_catalog.pg_stat_user_tables"

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table sql.NullString
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}

		if table.Valid {
			tables = append(tables, table.String)
		}
	}

	return tables, nil
}

func TableDiskSize(db *sql.DB, tableName string) (float64, error) {
	query := "select pg_table_size($1);"

	var size sql.NullFloat64
	if err := db.QueryRow(query, tableName).Scan(&size); err != nil {
		return 0, err
	}

	return size.Float64, nil
}

func TableIndexesDiskSize(db *sql.DB, tableName string) (float64, error) {
	query := "select pg_indexes_size($1);"

	var size sql.NullFloat64
	if err := db.QueryRow(query, tableName).Scan(&size); err != nil {
		return 0, err
	}

	return size.Float64, nil
}
