package schema

import (
	"database/sql"
	"fmt"
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

func (d sqliteDialect) Table(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	q := fmt.Sprintf(sqliteAllColumns, d.escapeIdent(name))
	return fetchColumnTypes(db, q)
}

func (sqliteDialect) TableNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, sqliteTableNames, "")
}

func (d sqliteDialect) Tables(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	return fetchColumnTypesAll(db, d.TableNames, d.Table)
}

func (sqliteDialect) PrimaryKey(db *sql.DB, name string) ([]string, error) {
	return fetchNames(db, sqlitePrimaryKey, name)
}

func (d sqliteDialect) View(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	q := fmt.Sprintf(sqliteAllColumns, d.escapeIdent(name))
	return fetchColumnTypes(db, q)
}

func (sqliteDialect) ViewNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, sqliteViewNames, "")
}

func (d sqliteDialect) Views(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	return fetchColumnTypesAll(db, d.ViewNames, d.View)
}
