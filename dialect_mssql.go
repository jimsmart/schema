package schema

import (
	"database/sql"
	"fmt"
)

const mssqlAllColumns = `SELECT * FROM %s WHERE 1=0`

// See https://stackoverflow.com/questions/8774928/how-to-exclude-system-table-when-querying-sys-tables

const mssqlTableNames = `
	SELECT
		t.name
	FROM
		sys.tables t
	INNER JOIN
		sys.schemas s
	ON	s.schema_id = t.schema_id
	LEFT JOIN
		sys.extended_properties ep
	ON	ep.major_id = t.[object_id]
	WHERE
		t.schema_id = SCHEMA_ID() AND
		t.is_ms_shipped = 0 AND
		(ep.class_desc IS NULL OR (ep.class_desc <> 'OBJECT_OR_COLUMN' AND
			ep.[name] <> 'microsoft_database_tools_support'))
	ORDER BY
		t.name
`

const mssqlViewNames = `
	SELECT
		t.name
	FROM
		sys.views t
	INNER JOIN
		sys.schemas s
	ON	s.schema_id = t.schema_id
	LEFT JOIN
		sys.extended_properties ep
	ON	ep.major_id = t.[object_id]
	WHERE
		t.schema_id = SCHEMA_ID() AND
		t.is_ms_shipped = 0 AND
		(ep.class_desc IS NULL OR (ep.class_desc <> 'OBJECT_OR_COLUMN' AND
			ep.[name] <> 'microsoft_database_tools_support'))
	ORDER BY
		t.name
`

const mssqlPrimaryKey = `
	SELECT
		tc.name
	FROM
		sys.schemas s
	INNER JOIN
		sys.tables t
	ON	s.schema_id = t.schema_id
	INNER JOIN
		sys.indexes i
	ON	t.object_id = i.object_id
	INNER JOIN
		sys.index_columns ic
	ON	i.object_id = ic.object_id AND
		i.index_id = ic.index_id
	INNER JOIN
		sys.columns tc
	ON	ic.object_id = tc.object_id AND
		ic.column_id = tc.column_id
	WHERE
		i.is_primary_key = 1 AND
		s.schema_id = SCHEMA_ID() AND
		t.name = ?
	ORDER BY
		ic.key_ordinal
`

type mssqlDialect struct{}

func (mssqlDialect) escapeIdent(ident string) string {
	// [tablename]
	return escapeWithBrackets(ident)
}

func (mssqlDialect) PrimaryKey(db *sql.DB, name string) ([]string, error) {
	return fetchNames(db, mssqlPrimaryKey, name)
}

func (d mssqlDialect) Table(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	q := fmt.Sprintf(mssqlAllColumns, d.escapeIdent(name))
	return fetchColumnTypes(db, q)
}

func (mssqlDialect) TableNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, mssqlTableNames, "")
}

func (d mssqlDialect) Tables(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	return fetchColumnTypesAll(db, d.TableNames, d.Table)
}

func (d mssqlDialect) View(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	q := fmt.Sprintf(mssqlAllColumns, d.escapeIdent(name))
	return fetchColumnTypes(db, q)
}

func (mssqlDialect) ViewNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, mssqlViewNames, "")
}

func (d mssqlDialect) Views(db *sql.DB) (map[string][]*sql.ColumnType, error) {
	return fetchColumnTypesAll(db, d.ViewNames, d.View)
}
