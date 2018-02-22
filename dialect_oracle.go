package schema

var oracle = dialect{
	queries: [3]string{
		// tableNames
		pack(`
			SELECT table_name
			FROM
				all_tables
			WHERE
				owner IN (SELECT sys_context('userenv', 'current_schema') from dual)
		`),
		// viewNames
		pack(`
			SELECT view_name
			FROM
				all_views
			WHERE
				owner IN (SELECT sys_context('userenv', 'current_schema') from dual)
		`),
		// columnTypes
		"SELECT * FROM %s WHERE 1=0",
	},
}
