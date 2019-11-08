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

- Microsoft SQL Server
- MySQL
- Oracle
- Postgres
- SQLite

For a list of supported drivers, and their capabilities with regards to sql.ColumnType support, see [drivercaps](https://github.com/jimsmart/drivercaps)

## Installation
```bash
$ go get github.com/jimsmart/schema
```

```go
import "github.com/jimsmart/schema"
```

### Dependencies

- A [supported](https://github.com/jimsmart/drivercaps) database driver.
- Standard library.
- [Ginkgo](https://onsi.github.io/ginkgo/) and [Gomega](https://onsi.github.io/gomega/) if you wish to run the tests.
- Tests also require [Docker Compose](https://docs.docker.com/compose/install/) to be installed.

## Example

See GoDocs for usage examples.

## Documentation

GoDocs [https://godoc.org/github.com/jimsmart/schema](https://godoc.org/github.com/jimsmart/schema)

## Testing

First, bring up the various databases for the testing environment. 

Execute `docker-compose up` inside the project folder, wait until all of the Docker services have completed their startup (there is no further output in the terminal), and open a second terminal. (In future one may choose to use `docker-compose up -d` instead.)

To run the tests execute `go test` inside the project folder.

For a full coverage report, try:

```bash
$ go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

To shutdown the Docker services, execute `docker-compose down -v` inside the project folder.

## License

Package schema is copyright 2018-2019 by Jim Smart and released under the [BSD 3-Clause License](LICENSE.md)

## History

- v0.0.4: Test environment now uses Docker.
- v0.0.3: Minor code cleanups.
- v0.0.2: Added identifier escaping for methods that query sql.ColumnType.
- v0.0.1: Started using Go modules.
- 2019-11-04: Fix for renamed driver struct in github.com/mattn/go-oci8 (Oracle)
- 2019-11-04: Fix for renamed driver struct in github.com/denisenkom/go-mssqldb (MSSQL)
