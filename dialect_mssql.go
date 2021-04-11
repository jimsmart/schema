package schema

import (
	"database/sql"
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

const mssqlTableNamesWithSchema = `
	SELECT
		schema_name(t.schema_id),
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
		t.is_ms_shipped = 0 AND
		(ep.class_desc IS NULL OR (ep.class_desc <> 'OBJECT_OR_COLUMN' AND
			ep.[name] <> 'microsoft_database_tools_support'))
	ORDER BY
		schema_name(t.schema_id),
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
		schema_name(t.schema_id),
		t.name
`

const mssqlViewNamesWithSchema = `
	SELECT
		schema_name(t.schema_id),
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
		t.is_ms_shipped = 0 AND
		(ep.class_desc IS NULL OR (ep.class_desc <> 'OBJECT_OR_COLUMN' AND
			ep.[name] <> 'microsoft_database_tools_support'))
	ORDER BY
		schema_name(t.schema_id),
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

const mssqlPrimaryKeyWithSchema = `
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
		s.schema_id = ? AND
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
	return fetchNames(db, mssqlPrimaryKey, "", name)
}

func (mssqlDialect) PrimaryKeyWithSchema(db *sql.DB, schema, name string) ([]string, error) {
	if schema == "" {
		return fetchNames(db, mssqlPrimaryKey, "", name)
	}
	return fetchNames(db, mssqlPrimaryKeyWithSchema, schema, name)
}

func (d mssqlDialect) Table(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	return fetchColumnTypes(db, mssqlAllColumns, name, d.escapeIdent)
}

func (d mssqlDialect) TableWithSchema(db *sql.DB, schema, name string) ([]*sql.ColumnType, error) {
	if schema == "" {
		return fetchColumnTypes(db, mssqlAllColumns, name, d.escapeIdent)
	}
	return fetchColumnTypesWithSchema(db, mssqlAllColumns, schema, name, d.escapeIdent)
}

func (mssqlDialect) TableNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, mssqlTableNames, "", "")
}

func (mssqlDialect) TableNamesWithSchema(db *sql.DB) ([][2]string, error) {
	return fetchNamesWithSchema(db, mssqlTableNamesWithSchema, "", "")
}

func (d mssqlDialect) View(db *sql.DB, name string) ([]*sql.ColumnType, error) {
	return fetchColumnTypes(db, mssqlAllColumns, name, d.escapeIdent)
}

func (d mssqlDialect) ViewWithSchema(db *sql.DB, schema, name string) ([]*sql.ColumnType, error) {
	if schema == "" {
		return fetchColumnTypes(db, mssqlAllColumns, name, d.escapeIdent)
	}
	return fetchColumnTypesWithSchema(db, mssqlAllColumns, schema, name, d.escapeIdent)
}

func (mssqlDialect) ViewNames(db *sql.DB) ([]string, error) {
	return fetchNames(db, mssqlViewNames, "", "")
}

func (mssqlDialect) ViewNamesWithSchema(db *sql.DB) ([][2]string, error) {
	return fetchNamesWithSchema(db, mssqlViewNamesWithSchema, "", "")
}
