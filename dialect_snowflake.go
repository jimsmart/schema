package schema

var snowflake = dialect{
	escapeIdent: noEscape, // tablename
	queries: postgres.queries,
}
