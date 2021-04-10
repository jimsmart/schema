package schema

import (
	"database/sql"
	"fmt"
)

const mysqlAllColumns = `SELECT * FROM %s LIMIT 0`

const mysqlTableNames = `
	SELECT
		table_name
	FROM
		information_schema.tables
	WHERE
		table_type = 'BASE TABLE' AND
		table_schema = database()
	ORDER BY
		table_name
`

const mysqlViewNames = `
	SELECT
		table_name
	FROM
		information_schema.tables
	WHERE
		table_type = 'VIEW' AND
		table_schema = database()
	ORDER BY
		table_name
`

const mysqlPrimaryKey = `
	SELECT
		sta.column_name
	FROM
		information_schema.tables tab
	INNER JOIN
		information_schema.statistics sta
	ON	sta.table_schema = tab.table_schema AND
		sta.table_name = tab.table_name AND
		sta.index_name = 'primary'
	WHERE
		tab.table_schema = database() AND
		tab.table_type = 'BASE TABLE' AND
		tab.table_name = ?
	ORDER BY
		sta.seq_in_index
`

type mysqlDialect struct{}

func (mysqlDialect) escapeIdent(ident string) string {
	// `tablename`
	return escapeWithBackticks(ident)
}

func (d mysqlDialect) Table(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	q := fmt.Sprintf(mysqlAllColumns, d.escapeIdent(name))
	return fetchColumnTypes(db, q)
}

func (mysqlDialect) TableNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, mysqlTableNames, "")
}

func (d mysqlDialect) Tables(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	return fetchColumnTypesAll(db, d.TableNames, d.Table)
}

func (mysqlDialect) PrimaryKey(db *sql.DB, name string) ([]string, error) {
	return fetchNames(db, mysqlPrimaryKey, name)
}

func (d mysqlDialect) View(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	q := fmt.Sprintf(mysqlAllColumns, d.escapeIdent(name))
	return fetchColumnTypes(db, q)
}

func (mysqlDialect) ViewNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, mysqlViewNames, "")
}

func (d mysqlDialect) Views(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	return fetchColumnTypesAll(db, d.ViewNames, d.View)
}
