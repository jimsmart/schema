package schema

import (
	"database/sql"
)

// TODO(js) Should we be filtering out system tables, like we currently do?

const postgresAllColumns = `SELECT * FROM %s LIMIT 0`

const postgresTableNamesWithSchema = `
	SELECT
		table_schema,
		table_name
	FROM
		information_schema.tables
	WHERE
		table_type = 'BASE TABLE' AND
		table_schema NOT IN ('pg_catalog', 'information_schema')
	ORDER BY
		table_schema,
		table_name
`

const postgresViewNamesWithSchema = `
	SELECT
		table_schema,
		table_name
	FROM
		information_schema.tables
	WHERE
		table_type = 'VIEW' AND
		table_schema NOT IN ('pg_catalog', 'information_schema')
	ORDER BY
		table_schema,
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

const postgresPrimaryKeyWithSchema = `
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
		kcu.table_schema = $1 AND
		kcu.table_name = $2
	ORDER BY
		kcu.ordinal_position
`

type postgresDialect struct{}

func (postgresDialect) escapeIdent(ident string) string {
	// "tablename"
	return escapeWithDoubleQuotes(ident)
}

func (d postgresDialect) ColumnTypes(db *sql.DB, schema, name string) ([]*sql.ColumnType, error) {
	return fetchColumnTypes(db, postgresAllColumns, schema, name, d.escapeIdent)
}

func (postgresDialect) PrimaryKey(db *sql.DB, schema, name string) ([]string, error) {
	if schema == "" {
		return fetchNames(db, postgresPrimaryKey, "", name)
	}
	return fetchNames(db, postgresPrimaryKeyWithSchema, schema, name)
}

func (postgresDialect) TableNames(db *sql.DB) ([][2]string, error) {
	return fetchObjectNames(db, postgresTableNamesWithSchema)
}

func (postgresDialect) ViewNames(db *sql.DB) ([][2]string, error) {
	return fetchObjectNames(db, postgresViewNamesWithSchema)
}
