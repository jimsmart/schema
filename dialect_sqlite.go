package schema

var sqlite = dialect{
	escapeIdent: escapeWithDoubleQuotes, // "tablename"
	queries: [4]string{
		// columnTypes query.
		`SELECT * FROM %s LIMIT 0`,
		// tableNames query.
		`SELECT name FROM sqlite_master WHERE type = 'table' ORDER BY name`,
		// viewNames query.
		`SELECT name FROM sqlite_master WHERE type = 'view' ORDER BY name`,
		// primaryKeyNames query.
		`SELECT name FROM pragma_table_info(?) WHERE pk > 0 ORDER BY pk`,
	},
}
