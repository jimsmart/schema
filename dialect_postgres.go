package schema

var postgres = dialect{
	escapeIdent: escapeWithDoubleQuotes, // "tablename"
	queries: [4]string{
		// columnTypes query.
		`SELECT * FROM %s LIMIT 0`,
		// tableNames query.
		pack(`
			SELECT
				table_name
			FROM
				information_schema.tables
			WHERE
				table_type = 'BASE TABLE' AND
				table_schema = current_schema()
			ORDER BY
				table_name
		`),
		// viewNames query.
		pack(`
			SELECT
				table_name
			FROM
				information_schema.tables
			WHERE
				table_type = 'VIEW' AND
				table_schema = current_schema()
			ORDER BY
				table_name
		`),
		// primaryKeyNames query.
		pack(`
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
		`),
	},
}
