// Package schema provides access to database schema metadata, for database/sql drivers.
//
// For further information about current driver support status, see https://github.com/jimsmart/schema
//
// Table Metadata
//
// The schema package works alongside database/sql and its underlying driver to provide schema metadata.
//
//  // Fetch names of all tables
//  tnames, err := schema.TableNames(db)
//  	...
//  // tnames is [][2]string
//  for i := range tnames {
//  	fmt.Println("Table:", tnames[i][1])
//  }
//
//  // Output:
//  // Table: employee_tbl
//  // Table: department_tbl
//  // Table: sales_tbl
//
// Both user permissions and current database/schema effect table visibility.
//
// Use schema.ColumnTypes() to query column type metadata for a single table:
//
//  // Fetch column metadata for given table
//  tcols, err := schema.ColumnTypes(db, "", "employee_tbl")
//  	...
//  // tcols is []*sql.ColumnType
//  for i := range tcols {
//  	fmt.Println("Column:", tcols[i].Name(), tcols[i].DatabaseTypeName())
//  }
//
//  // Output:
//  // Column: employee_id INTEGER
//  // Column: first_name TEXT
//  // Column: last_name TEXT
//  // Column: created_at TIMESTAMP
//
// To query table names and column type metadata for all tables, use schema.Tables().
//
// See also https://golang.org/pkg/database/sql/#ColumnType
//
// Note: underlying support for column type metadata is driver implementation specific and somewhat variable.
//
// View Metadata
//
// The same metadata can also be queried for views also:
//
//  // Fetch names of all views
//  vnames, err := schema.ViewNames(db)
//  	...
//  // Fetch column metadata for given view
//  vcols, err := schema.ColumnTypes(db, "", "monthly_sales_view")
//  	...
//  // Fetch column metadata for all views
//  views, err := schema.Views(db)
//  	...
//
// Primary Key Metadata
//
// To obtain a list of columns making up the primary key for a given table:
//
//  // Fetch primary key for given table
//  pks, err := schema.PrimaryKey(db, "", "employee_tbl")
//  	...
//  // pks is []string
//  for i := range pks {
//  	fmt.Println("Primary Key:", pks[i])
//  }
//
//  // Output:
//  // Primary Key: employee_id
//
package schema
