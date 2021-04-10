package schema_test

import (
	"fmt"

	// _ "github.com/mattn/go-oci8" // oci8
	// _ "gopkg.in/rana/ora.v4" // ora

	// _ "github.com/godror/godror" // godror

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

// See README.md to learn how to set up Oracle for testing purposes.

var _ = XDescribe("schema", func() {
	Context("using github.com/godror/godror (Oracle)", func() {

		const (
			user = "test_user"
			pass = "Password123"
			host = "localhost"
			port = "41521"
			dbs  = "xe"
		)

		var oracle = &testParams{
			DriverName: "godror",
			// DriverName: "oci8",
			// DriverName: "ora",
			ConnStr: fmt.Sprintf("%s/%s@%s:%s/%s", user, pass, host, port, dbs),

			CreateDDL: []string{`
				CREATE TABLE web_resource (
					id				NUMBER NOT NULL,
					url				NVARCHAR2(1024) NOT NULL UNIQUE,
					content			BLOB,
					compressed_size	NUMBER NOT NULL,
					content_length	NUMBER NOT NULL,
					content_type	NVARCHAR2(128) NOT NULL,
					etag			NVARCHAR2(128) NOT NULL,
					last_modified	NVARCHAR2(128) NOT NULL,
					created_at		TIMESTAMP WITH TIME ZONE NOT NULL,
					modified_at		TIMESTAMP WITH TIME ZONE,
					PRIMARY KEY (id)
				)`,
				// `CREATE INDEX idx_web_resource_url ON web_resource(url)`,
				`CREATE INDEX idx_web_resource_created_at ON web_resource(created_at)`,
				`CREATE INDEX idx_web_resource_modified_at ON web_resource(modified_at)`,
				`CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource`,
				`CREATE TABLE person (
					given_name		NVARCHAR2(128) NOT NULL,
					family_name		NVARCHAR2(128) NOT NULL,
					PRIMARY KEY (family_name, given_name)
				)`,
			},
			DropDDL: []string{
				`DROP TABLE person`,
				`DROP VIEW web_resource_view`,
				`DROP INDEX idx_web_resource_modified_at`,
				`DROP INDEX idx_web_resource_created_at`,
				// `DROP INDEX idx_web_resource_url`,
				`DROP TABLE web_resource`,
			},

			TableExpRes: []string{
				"ID",
				"URL",
				"CONTENT",
				"COMPRESSED_SIZE",
				"CONTENT_LENGTH",
				"CONTENT_TYPE",
				"ETAG",
				"LAST_MODIFIED",
				"CREATED_AT",
				"MODIFIED_AT",
			},
			ViewExpRes: []string{
				"ID",
				"URL",
			},

			TableNamesExpRes: []string{"PERSON", "WEB_RESOURCE"},
			ViewNameExpRes:   "WEB_RESOURCE_VIEW",

			PrimaryKeysExpRes: []string{"FAMILY_NAME", "GIVEN_NAME"},
		}

		SchemaTestRunner(oracle)
	})
})

// func oraDump(db *sql.DB) error {

// 	//SELECT table_name FROM user_tables
// 	rows, err := db.Query(`
// 		SELECT *
// 		  FROM user_tables
//    `)
// 	if err != nil {
// 		return err
// 	}
// 	defer rows.Close()

// 	ci, err := rows.ColumnTypes()
// 	if err != nil {
// 		return err
// 	}
// 	for _, c := range ci {
// 		log.Printf("%v", c)
// 	}

// 	cols, err := rows.Columns()
// 	if err != nil {
// 		return err
// 	}
// 	vals := make([]interface{}, len(cols))
// 	for i, _ := range cols {
// 		vals[i] = new(sql.RawBytes)
// 	}

// 	for rows.Next() {
// 		err = rows.Scan(vals...)
// 		if err != nil {
// 			// return err
// 			log.Printf("%v", err)
// 		}
// 		s := ""
// 		for _, v := range vals {
// 			s = s + fmt.Sprintf("%s ", v)
// 		}
// 		log.Print(s)
// 	}
// 	return nil
// }
