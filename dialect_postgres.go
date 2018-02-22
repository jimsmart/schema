package schema

var postgres = dialect{
	queries: [3]string{
		// tableNames
		pack(`
			SELECT table_name
			FROM
				information_schema.tables
			WHERE
				table_type='BASE TABLE' AND
				table_schema=current_schema()
		`),
		// viewNames
		pack(`
			SELECT table_name
			FROM
				information_schema.tables
			WHERE
				table_type='VIEW' AND
				table_schema=current_schema()
		`),
		// columnTypes
		"SELECT * FROM %s LIMIT 0",
	},
}
