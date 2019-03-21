package schema

// TODO(js) Is there some way to filter system tables (like mssql)? Or should we always just be using our own schema?

var oracle = dialect{
	queries: [3]string{
		// columnTypes query.
		`SELECT * FROM %s WHERE 1=0`,
		// tableNames query.
		pack(`
			SELECT table_name
			FROM
				table_privileges
			WHERE
				grantee IN (SELECT sys_context('userenv', 'current_schema') from dual) 
				AND
				select_priv='Y'
			UNION
			SELECT table_name
			FROM
				user_tables
		`),
		// viewNames query.
		pack(`
			SELECT view_name
			FROM
				all_views
			WHERE
				owner IN (SELECT sys_context('userenv', 'current_schema') from dual)
		`),
	},
}
