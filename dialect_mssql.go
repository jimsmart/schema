package schema

// See https://stackoverflow.com/questions/8774928/how-to-exclude-system-table-when-querying-sys-tables

var mssql = dialect{
	queries: [3]string{
		// columnTypes query.
		`SELECT * FROM %s WHERE 1=0`,
		// tableNames query.
		pack(`
			SELECT T.name as name
			FROM
				sys.tables AS T
				INNER JOIN sys.schemas AS S ON S.schema_id = T.schema_id
				LEFT JOIN sys.extended_properties AS EP ON EP.major_id = T.[object_id]
			WHERE
				T.is_ms_shipped = 0 AND
				(EP.class_desc IS NULL OR (EP.class_desc <> 'OBJECT_OR_COLUMN' AND
				EP.[name] <> 'microsoft_database_tools_support'))
		`),
		// viewNames query.
		pack(`
			SELECT T.name as name
			FROM
				sys.views AS T
				INNER JOIN sys.schemas AS S ON S.schema_id = T.schema_id
				LEFT JOIN sys.extended_properties AS EP ON EP.major_id = T.[object_id]
			WHERE
				T.is_ms_shipped = 0 AND
				(EP.class_desc IS NULL OR (EP.class_desc <> 'OBJECT_OR_COLUMN' AND
				EP.[name] <> 'microsoft_database_tools_support'))
		`),
	},
}
