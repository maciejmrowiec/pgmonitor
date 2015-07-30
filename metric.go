package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"strings"
)

// Path: use '$' as place holder for tablename.

type TableMetric struct {
	db        *sql.DB
	unit      string
	path      string
	valueFunc func(*sql.DB, string) (float64, error)
}

func NewTableMetric(
	db *sql.DB,
	path string,
	valueFunc func(*sql.DB, string) (float64, error),
	unit string,
) *TableMetric {
	return &TableMetric{
		db:        db,
		unit:      unit,
		path:      path,
		valueFunc: valueFunc,
	}
}

func (t *TableMetric) GetUnits() string {
	return t.unit
}

func (t *TableMetric) GetName(table_name string) string {
	if strings.Contains(t.path, "$") {
		fmtString := strings.Replace(t.path, "$", "%s", 1)
		return fmt.Sprintf(fmtString, table_name)
	}

	return t.path
}

func (t *TableMetric) GetIdList() []string {
	tables, err := UserTables(t.db)

	if err != nil {
		log.Println(err)
		return nil
	}

	return tables
}

func (t *TableMetric) GetValue(table_name string) (float64, error) {
	return t.valueFunc(t.db, table_name)
}

type TableAverageSummaryMetric struct {
	db        *sql.DB
	unit      string
	path      string
	valueFunc func(*sql.DB, string) (float64, error)
}

func NewTableAverageSummaryMetric(
	db *sql.DB,
	path string,
	valueFunc func(*sql.DB, string) (float64, error),
	unit string,
) *TableAverageSummaryMetric {
	return &TableAverageSummaryMetric{
		db:        db,
		path:      path,
		valueFunc: valueFunc,
		unit:      unit,
	}
}

func (t *TableAverageSummaryMetric) GetName() string {
	return t.path
}

func (t *TableAverageSummaryMetric) GetUnits() string {
	return t.unit
}

func (t *TableAverageSummaryMetric) GetValue() (float64, error) {

	tables, err := UserTables(t.db)
	if err != nil {
		return 0, err
	}

	var sum float64
	var count float64

	for _, table := range tables {
		percent, err := t.valueFunc(t.db, table)
		if err != nil {
			continue
		}
		sum += percent
		count++
	}

	return float64(sum / count), nil
}
