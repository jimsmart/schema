package schema

var mysql = dialect{
	queries: [3]string{
		// tableNames
		pack(`
			SELECT table_name
			FROM
				information_schema.tables
			WHERE
				table_type='BASE TABLE' AND
				table_schema=database()
		`),
		// viewNames
		pack(`
			SELECT table_name
			FROM
				information_schema.tables
			WHERE
				table_type='VIEW' AND
				table_schema=database()
		`),
		// columnTypes
		"SELECT * FROM %s LIMIT 0",
	},
}
