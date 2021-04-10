package schema

// See https://stackoverflow.com/questions/8774928/how-to-exclude-system-table-when-querying-sys-tables

var mssql = dialect{
	escapeIdent: escapeWithBrackets, // [tablename]
	queries: [4]string{
		// columnTypes query.
		`SELECT * FROM %s WHERE 1=0`,
		// tableNames query.
		pack(`
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
		`),
		// viewNames query.
		pack(`
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
			`),
		// primaryKeyNames query.
		pack(`
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
		`),
	},
}
