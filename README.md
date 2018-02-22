# schema

[![BSD3](https://img.shields.io/badge/license-BSD3-blue.svg?style=flat)](LICENSE.md) [![Build Status](https://img.shields.io/travis/jimsmart/schema/master.svg?style=flat)](https://travis-ci.org/jimsmart/schema) [![codecov](https://codecov.io/gh/jimsmart/schema/branch/master/graph/badge.svg)](https://codecov.io/gh/jimsmart/schema) [![Go Report Card](https://goreportcard.com/badge/github.com/jimsmart/schema)](https://goreportcard.com/report/github.com/jimsmart/schema) [![Godoc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/jimsmart/schema)

schema is a [Go](https://golang.org) package providing database schema metadata for database/sql drivers.

TODO more docs

The following drivers are supported and fully tested:

- github.com/mattn/go-sqlite3 - SQLite
- github.com/lib/pq - Postgres
- github.com/denisenkom/go-mssqldb - Microsoft SQL Server
- gopkg.in/goracle.v2 - Oracle
- github.com/go-sql-driver/mysql - MySQL

An effort has been made to support the following drivers, but they are currently untested:

- github.com/gwenn/gosqlite - SQLite
- github.com/mxk/go-sqlite - SQLite
- github.com/jackc/pgx/stdlib - Postgres
- github.com/jbarham/gopgsqldriver - Postgres
- github.com/minus5/gofreetds - Microsoft SQL Server
- gopkg.in/rana/ora.v4 - Oracle
- github.com/mattn/go-oci8 - Oracle
- github.com/ziutek/mymysql - MySQL

If you use the schema package with any of these drivers, please open an issue commenting whether it worked or not, to keep this list up-to-date.

If your favourite driver is not featured in either of the above lists, open an issue detailing which driver you use.

Pull requests welcome!


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

Package schema is copyright 2018 by Jim Smart and released under the [3-Clause BSD License](LICENSE.md)
