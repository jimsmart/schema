package schema_test

import (
	"fmt"

	_ "gopkg.in/goracle.v2"

	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = Describe("schema", func() {
	Context("using github.com/go-goracle/goracle (Oracle)", func() {

		const (
			user = "test_user"
			pass = "password"
			host = "localhost"
			port = "32786"
			dbs  = "xe"
		)

		var oracle = &testParams{
			DriverName: "goracle",
			ConnStr:    fmt.Sprintf("%s/%s@%s:%s/%s", user, pass, host, port, dbs),

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
			},
			DropDDL: []string{
				`DROP VIEW web_resource_view`,
				`DROP INDEX idx_web_resource_modified_at`,
				`DROP INDEX idx_web_resource_created_at`,
				// `DROP INDEX idx_web_resource_url`,
				`DROP TABLE web_resource`,
			},

			TableExpRes: []string{
				"ID NUMBER",
				"URL NVARCHAR2",
				"CONTENT BLOB",
				"COMPRESSED_SIZE NUMBER",
				"CONTENT_LENGTH NUMBER",
				"CONTENT_TYPE NVARCHAR2",
				"ETAG NVARCHAR2",
				"LAST_MODIFIED NVARCHAR2",
				"CREATED_AT TIMESTAMP WITH TIMEZONE",
				"MODIFIED_AT TIMESTAMP WITH TIMEZONE",
			},
			ViewExpRes: []string{
				"ID NUMBER",
				"URL NVARCHAR2",
			},

			TableNameExpRes: "WEB_RESOURCE",
			ViewNameExpRes:  "WEB_RESOURCE_VIEW",
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
