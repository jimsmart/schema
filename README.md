# schema

[![BSD3](https://img.shields.io/badge/license-BSD3-blue.svg?style=flat)](LICENSE.md)
[![Build Status](https://img.shields.io/travis/jimsmart/schema/master.svg?style=flat)](https://travis-ci.org/jimsmart/schema)
[![codecov](https://codecov.io/gh/jimsmart/schema/branch/master/graph/badge.svg)](https://codecov.io/gh/jimsmart/schema)
[![Go Report Card](https://goreportcard.com/badge/github.com/jimsmart/schema)](https://goreportcard.com/report/github.com/jimsmart/schema)
[![Used By](https://img.shields.io/sourcegraph/rrc/github.com/jimsmart/schema.svg)](https://sourcegraph.com/github.com/jimsmart/schema)
[![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/jimsmart/schema)

schema is a [Go](https://golang.org) package providing access to database schema metadata, for database/sql drivers.

TODO more docs

Currently supporting the following database engines / SQL dialects:

- SQLite
- Postgres
- Microsoft SQL Server
- Oracle
- MySQL

The following drivers are supported and pass all tests:

- github.com/mattn/go-sqlite3 - SQLite
- github.com/gwenn/gosqlite - SQLite
- github.com/jackc/pgx - Postgres
- github.com/lib/pq - Postgres
- github.com/denisenkom/go-mssqldb - Microsoft SQL Server
- gopkg.in/goracle.v2 - Oracle
- github.com/go-sql-driver/mysql - MySQL

The following drivers are 'supported', but with issues:

- github.com/mxk/go-sqlite/sqlite3 - SQLite - Driver provides no datatype information for columns.
- github.com/jbarham/gopgsqldriver - Postgres - Driver provides no datatype information for columns.
- github.com/minus5/gofreetds - Microsoft SQL Server - Driver provides no datatype information for columns. Driver error during test when attempting to CREATE VIEW.
- gopkg.in/rana/ora.v4 - Oracle - Driver provides datatypes that do not match types used in create DDL (e.g. NVARCHAR2 becomes VARCHAR2).
- github.com/mattn/go-oci8 - Oracle - Driver provides datatypes that do not match types used in create DDL (e.g. NVARCHAR2 becomes SQLT_CHR - incomplete implementation?).
- github.com/ziutek/mymysql - MySQL - Driver provides no datatype information for columns.


If your favourite driver or database is not featured in either of the above lists, open an issue providing further details.

Pull requests welcome!

TODO driver capability testing is in the process of being moved to a separate package, see [github.com/jimsmart/drivercaps](https://github.com/jimsmart/drivercaps)

## Installation
```bash
$ go get github.com/jimsmart/schema
```

```go
import "github.com/jimsmart/schema"
```

### Dependencies

- Some supported database driver from above lists.
- Standard library.
- [Ginkgo](https://onsi.github.io/ginkgo/) and [Gomega](https://onsi.github.io/gomega/) if you wish to run the tests.
- Tests also require various database drivers to be installed and configured.

## Example

See GoDocs for usage examples.

## Documentation

GoDocs [https://godoc.org/github.com/jimsmart/schema](https://godoc.org/github.com/jimsmart/schema)

## Testing

Note that a moderate amount of database setup and configuration is required for successful execution of the tests.

To run the tests execute `go test` inside the project folder.

For a full coverage report, try:

```bash
$ go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

## License

Package schema is copyright 2018 by Jim Smart and released under the [BSD 3-Clause License](LICENSE.md)
