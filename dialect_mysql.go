package schema

import (
	"database/sql"
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

const mysqlTableNamesWithSchema = `
	SELECT
		table_schema,
		table_name
	FROM
		information_schema.tables
	WHERE
		table_type = 'BASE TABLE' 
	ORDER BY
		table_schema,
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

const mysqlViewNamesWithSchema = `
	SELECT
		table_schema,
		table_name
	FROM
		information_schema.tables
	WHERE
		table_type = 'VIEW' AND
		table_schema = database()
	ORDER BY
		table_schema,
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
		tab.table_type = 'BASE TABLE' AND
		tab.table_schema = database() AND
		tab.table_name = ?
	ORDER BY
		sta.seq_in_index
`

const mysqlPrimaryKeyWithSchema = `
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
		tab.table_type = 'BASE TABLE' AND
		tab.table_schema = ? AND
		tab.table_name = ?
	ORDER BY
		sta.seq_in_index
`

type mysqlDialect struct{}

func (mysqlDialect) escapeIdent(ident string) string {
	// `tablename`
	return escapeWithBackticks(ident)
}

func (mysqlDialect) PrimaryKey(db *sql.DB, schema, name string) ([]string, error) {
	if schema == "" {
		return fetchNames(db, mysqlPrimaryKey, "", name)
	}
	return fetchNames(db, mysqlPrimaryKeyWithSchema, schema, name)
}

func (d mysqlDialect) Table(db *sql.DB, schema, name string) ([]*sql.ColumnType, error) {
	if schema == "" {
		return fetchColumnTypes(db, mysqlAllColumns, name, d.escapeIdent)
	}
	return fetchColumnTypesWithSchema(db, mysqlAllColumns, schema, name, d.escapeIdent)
}

func (mysqlDialect) TableNames(db *sql.DB) ([][2]string, error) {
	return fetchNamesWithSchema(db, mysqlTableNamesWithSchema, "", "")
}

func (d mysqlDialect) View(db *sql.DB, schema, name string) ([]*sql.ColumnType, error) {
	if schema == "" {
		return fetchColumnTypes(db, mysqlAllColumns, name, d.escapeIdent)
	}
	return fetchColumnTypesWithSchema(db, mysqlAllColumns, schema, name, d.escapeIdent)
}

func (mysqlDialect) ViewNames(db *sql.DB) ([][2]string, error) {
	return fetchNamesWithSchema(db, mysqlViewNamesWithSchema, "", "")
}
