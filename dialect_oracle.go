package schema

// TODO(js) Is there some way to filter system tables (like mssql)? Or should we always just be using our own schema?

var oracle = dialect{
	escapeIdent: escapeWithDoubleQuotes, // "tablename"
	queries: [4]string{
		// columnTypes query.
		`SELECT * FROM %s WHERE 1=0`,
		// tableNames query.
		pack(`
			SELECT
				table_name
			FROM
				all_tables
			WHERE
				owner IN (SELECT sys_context('userenv', 'current_schema') from dual)
			ORDER BY
				table_name
		`),
		// viewNames query.
		pack(`
			SELECT
				view_name
			FROM
				all_views
			WHERE
				owner IN (SELECT sys_context('userenv', 'current_schema') from dual)
			ORDER BY
				view_name
		`),
		// primaryKeyNames query.
		pack(`
			SELECT
				cc.column_name
			FROM
				all_constraints c,
				all_cons_columns cc
			WHERE
				c.constraint_type = 'P' AND
				c.constraint_name = cc.constraint_name AND
				c.owner = cc.owner AND
				cc.owner IN (SELECT sys_context('userenv', 'current_schema') from dual) AND
				cc.table_name = :1
			ORDER BY
				cc.position
		`),
	},
}
