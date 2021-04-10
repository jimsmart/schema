package schema

import (
	"database/sql"
	"fmt"
)

const postgresAllColumns = `SELECT * FROM %s LIMIT 0`

const postgresTableNames = `
	SELECT
		table_name
	FROM
		information_schema.tables
	WHERE
		table_type = 'BASE TABLE' AND
		table_schema = current_schema()
	ORDER BY
		table_name
`

const postgresViewNames = `
	SELECT
		table_name
	FROM
		information_schema.tables
	WHERE
		table_type = 'VIEW' AND
		table_schema = current_schema()
	ORDER BY
		table_name
`

const postgresPrimaryKey = `
	SELECT
		kcu.column_name
	FROM
		information_schema.table_constraints tco
	JOIN
		information_schema.key_column_usage kcu
	ON	kcu.constraint_name = tco.constraint_name AND
		kcu.constraint_schema = tco.constraint_schema AND
		kcu.constraint_name = tco.constraint_name
	WHERE
		tco.constraint_type = 'PRIMARY KEY' AND
		kcu.table_schema = current_schema() AND
		kcu.table_name = $1
	ORDER BY
		kcu.ordinal_position
`

type postgresDialect struct{}

func (postgresDialect) escapeIdent(ident string) string {
	// "tablename"
	return escapeWithDoubleQuotes(ident)
}

func (d postgresDialect) Table(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	q := fmt.Sprintf(postgresAllColumns, d.escapeIdent(name))
	return fetchColumnTypes(db, q)
}

func (postgresDialect) TableNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, postgresTableNames, "")
}

func (d postgresDialect) Tables(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	return fetchColumnTypesAll(db, d.TableNames, d.Table)
}

func (postgresDialect) PrimaryKey(db *sql.DB, name string) ([]string, error) {
	return fetchNames(db, postgresPrimaryKey, name)
}

func (d postgresDialect) View(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	q := fmt.Sprintf(postgresAllColumns, d.escapeIdent(name))
	return fetchColumnTypes(db, q)
}

func (postgresDialect) ViewNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, postgresViewNames, "")
}

func (d postgresDialect) Views(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	return fetchColumnTypesAll(db, d.ViewNames, d.View)
}
