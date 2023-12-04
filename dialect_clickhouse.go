package schema

import (
	"database/sql"
)

// TODO(js) Should we be filtering out system tables, like we currently do?

const clickhouseAllColumns = `SELECT * FROM %s LIMIT 0`

const clickhouseTableNamesWithSchema = `
  SELECT database AS table_schema,
         name AS table_name
  FROM system.tables
  WHERE NOT is_temporary 
    AND engine NOT LIKE '%View'
    AND engine NOT LIKE 'System%'
    AND has_own_data != 0
    AND database NOT IN ('information_schema', 'system', 'INFORMATION_SCHEMA')
  ORDER BY database, name;
`

const clickhouseViewNamesWithSchema = `
SELECT database AS table_schema,
       name     AS table_name
FROM system.tables
WHERE engine LIKE '%View'
  AND database NOT IN ('information_schema', 'system', 'INFORMATION_SCHEMA')
ORDER BY table_schema,
         table_name
`

const clickhousePrimaryKey = `
	SELECT
		name
	FROM
		system.columns
	WHERE
		database = currentDatabase() AND
		table = $1 AND
	  is_in_primary_key
	ORDER BY
		position DESC
`

const clickhousePrimaryKeyWithSchema = `
	SELECT
		name
	FROM
		system.columns
	WHERE
		database = $1 AND
		table = $2 AND
	  is_in_primary_key
	ORDER BY
		position DESC
`

type clickhouseDialect struct{}

func (clickhouseDialect) escapeIdent(ident string) string {
	// "tablename"
	return escapeWithDoubleQuotes(ident)
}

func (d clickhouseDialect) ColumnTypes(db *sql.DB, schema, name string) ([]*sql.ColumnType, error) {
	return fetchColumnTypes(db, clickhouseAllColumns, schema, name, d.escapeIdent)
}

func (clickhouseDialect) PrimaryKey(db *sql.DB, schema, name string) ([]string, error) {
	if schema == "" {
		return fetchNames(db, clickhousePrimaryKey, "", name)
	}
	return fetchNames(db, clickhousePrimaryKeyWithSchema, schema, name)
}

func (clickhouseDialect) TableNames(db *sql.DB) ([][2]string, error) {
	return fetchObjectNames(db, clickhouseTableNamesWithSchema)
}

func (clickhouseDialect) ViewNames(db *sql.DB) ([][2]string, error) {
	return fetchObjectNames(db, clickhouseViewNamesWithSchema)
}
