package schema

var sqlite = dialect{
	queries: [3]string{
		// tableNames query.
		pack(`
			SELECT name
			FROM
				sqlite_master
			WHERE
				type = 'table'
		`),
		// viewNames query.
		pack(`
			SELECT name
			FROM
				sqlite_master
			WHERE
				type = 'view'
		`),
		// columnTypes query.
		`SELECT * FROM %s LIMIT 0`,
	},
}