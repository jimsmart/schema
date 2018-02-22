package schema

var sqlite = dialect{
	queries: [3]string{
		// tableNames
		pack(`
			SELECT name
			FROM
				sqlite_master
			WHERE
				type = 'table'
		`),
		// viewNames
		pack(`
			SELECT name
			FROM
				sqlite_master
			WHERE
				type = 'view'
		`),
		// columnTypes
		"SELECT * FROM %s LIMIT 0",
	},
}
