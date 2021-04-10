package schema

var mysql = dialect{
	escapeIdent: escapeWithBackticks, // `tablename`
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
				table_schema = database()
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
				table_schema = database()
			ORDER BY
				table_name
		`),
		// primaryKeyNames query.
		pack(`
			SELECT
				sta.column_name
			FROM
				information_schema.tables tab
			INNER JOIN
				information_schema.statistics sta
			ON	sta.table_schema = tab.table_schema AND
				sta.table_name = tab.table_name AND
				sta.index_name = 'primary'
			WHERE
				tab.table_schema = database() AND
				tab.table_type = 'BASE TABLE' AND
				tab.table_name = ?
			ORDER BY
				sta.seq_in_index
		`),
	},
}
