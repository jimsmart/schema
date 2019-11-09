package schema_test

// Setup of test db is performed on Docker startup
// via environment variables.

var mysqlDialect = testParams{
	CreateDDL: []string{
		`CREATE TABLE IF NOT EXISTS web_resource (
					id				INTEGER NOT NULL,
					url				VARCHAR(255) NOT NULL UNIQUE,
					content			BLOB,
					compressed_size	INTEGER NOT NULL,
					content_length	INTEGER NOT NULL,
					content_type	VARCHAR(128) NOT NULL,
					etag			VARCHAR(128) NOT NULL,
					last_modified	VARCHAR(128) NOT NULL,
					created_at		TIMESTAMP NOT NULL,
					modified_at		TIMESTAMP NULL DEFAULT NULL,
					PRIMARY KEY (id),
					INDEX (url),
					INDEX (created_at),
					INDEX (modified_at)
		)`,
		// TODO(js) should this have NOT EXISTS...?
		"CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource",
		// Tests for correct identifer escaping.
		"CREATE TABLE IF NOT EXISTS `blanks in name` (id INTEGER, PRIMARY KEY (id))",
		"CREATE TABLE IF NOT EXISTS `[brackets] in name` (id INTEGER, PRIMARY KEY (id))",
		"CREATE TABLE IF NOT EXISTS `\"d.quotes\" in name` (id INTEGER, PRIMARY KEY (id))",
		"CREATE TABLE IF NOT EXISTS `'s.quotes' in name` (id INTEGER, PRIMARY KEY (id))",
		"CREATE TABLE IF NOT EXISTS `{braces} in name` (id INTEGER, PRIMARY KEY (id))",
		"CREATE TABLE IF NOT EXISTS ```backticks`` in name` (id INTEGER, PRIMARY KEY (id))",
		"CREATE TABLE IF NOT EXISTS `backslashes\\in\\name` (id INTEGER, PRIMARY KEY (id))",
	},
	DropDDL: []string{
		"DROP TABLE `backslashes\\in\\name`",
		"DROP TABLE ```backticks`` in name`",
		"DROP TABLE `{braces} in name`",
		"DROP TABLE `'s.quotes' in name`",
		"DROP TABLE `\"d.quotes\" in name`",
		"DROP TABLE `[brackets] in name`",
		"DROP TABLE `blanks in name`",
		"DROP VIEW web_resource_view",
		"DROP TABLE web_resource",
	},

	TableNamesExpRes: []string{
		"web_resource",
		// Tests for correct identifer escaping.
		"blanks in name",
		"[brackets] in name",
		`"d.quotes" in name`,
		"'s.quotes' in name",
		"{braces} in name",
		"`backticks` in name",
		`backslashes\in\name`,
	},
	ViewNamesExpRes: []string{
		"web_resource_view",
	},
}
