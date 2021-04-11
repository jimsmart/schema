package schema

import (
	"database/sql"
)

// TODO(js) Is there some way to filter system tables (like mssql)? Or should we always just be using our own schema?

const oracleAllColumns = `SELECT * FROM %s WHERE 1=0`

const oracleTableNames = `
	SELECT
		table_name
	FROM
		all_tables
	WHERE
		owner IN (SELECT sys_context('userenv', 'current_schema') from dual)
	ORDER BY
		table_name
`

const oracleTableNamesWithSchema = `
	SELECT
		owner,
		table_name
	FROM
		all_tables
	WHERE
		owner IN (SELECT sys_context('userenv', 'current_schema') from dual)
	ORDER BY
		owner,
		table_name
`

const oracleViewNames = `
	SELECT
		view_name
	FROM
		all_views
	WHERE
		owner IN (SELECT sys_context('userenv', 'current_schema') from dual)
	ORDER BY
		view_name
`

const oracleViewNamesWithSchema = `
	SELECT
		owner,
		view_name
	FROM
		all_views
	WHERE
		owner IN (SELECT sys_context('userenv', 'current_schema') from dual)
	ORDER BY
		owner,
		view_name
`

const oraclePrimaryKey = `
	SELECT
		cc.column_name
	FROM
		all_constraints c,
		all_cons_columns cc
	WHERE
		c.constraint_type = 'P' AND
		c.constraint_name = cc.constraint_name AND
		c.owner = cc.owner AND
		cc.owner IN (SELECT sys_context('userenv', 'current_schema') from dual) AND
		cc.table_name = :1
	ORDER BY
		cc.position
`

const oraclePrimaryKeyWithSchema = `
	SELECT
		cc.column_name
	FROM
		all_constraints c,
		all_cons_columns cc
	WHERE
		c.constraint_type = 'P' AND
		c.constraint_name = cc.constraint_name AND
		c.owner = cc.owner AND
		cc.owner = :1 AND
		cc.table_name = :2
	ORDER BY
		cc.position
`

type oracleDialect struct{}

func (oracleDialect) escapeIdent(ident string) string {
	// "tablename"
	return escapeWithDoubleQuotes(ident)
}

func (oracleDialect) PrimaryKey(db *sql.DB, schema, name string) ([]string, error) {
	if schema == "" {
		return fetchNames(db, oraclePrimaryKey, "", name)
	}
	return fetchNames(db, oraclePrimaryKeyWithSchema, schema, name)
}

func (d oracleDialect) Table(db *sql.DB, schema, name string) ([]*sql.ColumnType, error) {
	if schema == "" {
		return fetchColumnTypes(db, oracleAllColumns, name, d.escapeIdent)
	}
	return fetchColumnTypesWithSchema(db, oracleAllColumns, schema, name, d.escapeIdent)
}

func (oracleDialect) TableNames(db *sql.DB) ([][2]string, error) {
	return fetchNamesWithSchema(db, oracleTableNamesWithSchema, "", "")
}

func (d oracleDialect) View(db *sql.DB, schema, name string) ([]*sql.ColumnType, error) {
	if schema == "" {
		return fetchColumnTypes(db, oracleAllColumns, name, d.escapeIdent)
	}
	return fetchColumnTypesWithSchema(db, oracleAllColumns, schema, name, d.escapeIdent)
}

func (oracleDialect) ViewNames(db *sql.DB) ([][2]string, error) {
	return fetchNamesWithSchema(db, oracleViewNamesWithSchema, "", "")
}
