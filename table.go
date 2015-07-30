package main

import (
	"database/sql"
	_ "github.com/lib/pq"
)

func UserTables(db *sql.DB) ([]string, error) {
	sql := "SELECT relname FROM pg_catalog.pg_stat_user_tables"

	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tables []string
	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			return nil, err
		}

		tables = append(tables, table)
	}

	return tables, nil
}

func TableDiskSize(db *sql.DB, tableName string) (float64, error) {
	sql := "select pg_table_size($1);"

	var size float64
	if err := db.QueryRow(sql, tableName).Scan(&size); err != nil {
		return 0, err
	}

	return size / 1000, nil
}

func TableIndexesDiskSize(db *sql.DB, tableName string) (float64, error) {
	sql := "select pg_indexes_size($1);"

	var size float64
	if err := db.QueryRow(sql, tableName).Scan(&size); err != nil {
		return 0, err
	}

	return size, nil
}
