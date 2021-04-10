package schema

import (
	"database/sql"
)

const sqliteAllColumns = `SELECT * FROM %s LIMIT 0`

const sqliteTableNames = `SELECT name FROM sqlite_master WHERE type = 'table' ORDER BY name`

const sqliteViewNames = `SELECT name FROM sqlite_master WHERE type = 'view' ORDER BY name`

const sqlitePrimaryKey = `SELECT name FROM pragma_table_info(?) WHERE pk > 0 ORDER BY pk`

type sqliteDialect struct{}

func (sqliteDialect) escapeIdent(ident string) string {
	// "tablename"
	return escapeWithDoubleQuotes(ident)
}

func (sqliteDialect) PrimaryKey(db *sql.DB, name string) ([]string, error) {
	return fetchNames(db, sqlitePrimaryKey, name)
}

func (d sqliteDialect) Table(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	return fetchColumnTypes(db, sqliteAllColumns, name, d.escapeIdent)
}

func (sqliteDialect) TableNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, sqliteTableNames, "")
}

func (d sqliteDialect) View(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	return fetchColumnTypes(db, sqliteAllColumns, name, d.escapeIdent)
}

func (sqliteDialect) ViewNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, sqliteViewNames, "")
}
