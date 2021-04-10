package schema

import (
	"database/sql"
	"fmt"
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

type oracleDialect struct{}

func (oracleDialect) escapeIdent(ident string) string {
	// "tablename"
	return escapeWithDoubleQuotes(ident)
}

func (oracleDialect) PrimaryKey(db *sql.DB, name string) ([]string, error) {
	return fetchNames(db, oraclePrimaryKey, name)
}

func (d oracleDialect) Table(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	q := fmt.Sprintf(oracleAllColumns, d.escapeIdent(name))
	return fetchColumnTypes(db, q)
}

func (oracleDialect) TableNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, oracleTableNames, "")
}

func (d oracleDialect) Tables(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	return fetchColumnTypesAll(db, d.TableNames, d.Table)
}

func (d oracleDialect) View(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	q := fmt.Sprintf(oracleAllColumns, d.escapeIdent(name))
	return fetchColumnTypes(db, q)
}

func (oracleDialect) ViewNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, oracleViewNames, "")
}

func (d oracleDialect) Views(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	return fetchColumnTypesAll(db, d.ViewNames, d.View)
}
