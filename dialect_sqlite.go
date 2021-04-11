package schema

import (
	"database/sql"
)

const sqliteAllColumns = `SELECT * FROM %s LIMIT 0`

const sqliteTableNamesWithSchema = `
	SELECT
		"" AS schema,
		name 
	FROM
		sqlite_master
	WHERE
		type = 'table'
	ORDER BY
		name
`

const sqliteViewNamesWithSchema = `
	SELECT
		"" AS schema,
		name
	FROM
		sqlite_master
	WHERE
		type = 'view'
	ORDER BY
		name
`

const sqlitePrimaryKey = `
	SELECT
		name
	FROM
		pragma_table_info(?)
	WHERE
		pk > 0
	ORDER BY
		pk
`

type sqliteDialect struct{}

func (sqliteDialect) escapeIdent(ident string) string {
	// "tablename"
	return escapeWithDoubleQuotes(ident)
}

func (d sqliteDialect) ColumnTypes(db *sql.DB, schema, name string) ([]*sql.ColumnType, error) {
	return fetchColumnTypes(db, sqliteAllColumns, schema, name, d.escapeIdent)
}

func (sqliteDialect) PrimaryKey(db *sql.DB, schema, name string) ([]string, error) {
	// if schema == "" {
	// 	return fetchNames(db, sqlitePrimaryKey, "", name)
	// }
	return fetchNames(db, sqlitePrimaryKey, "", name)
}

func (sqliteDialect) TableNames(db *sql.DB) ([][2]string, error) {
	return fetchNamesWithSchema(db, sqliteTableNamesWithSchema, "", "")
}

func (sqliteDialect) ViewNames(db *sql.DB) ([][2]string, error) {
	return fetchNamesWithSchema(db, sqliteViewNamesWithSchema, "", "")
}
