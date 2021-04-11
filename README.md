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
- Snowflake
- SQLite

For a list of supported drivers, and their capabilities with regards to sql.ColumnType support, see [drivercaps](https://github.com/jimsmart/drivercaps)

## Installation

```bash
go get github.com/jimsmart/schema
```

```go
import "github.com/jimsmart/schema"
```

### Dependencies

- A [supported](https://github.com/jimsmart/drivercaps) database driver.
- Standard library.
- [Ginkgo](https://onsi.github.io/ginkgo/) and [Gomega](https://onsi.github.io/gomega/) are used in the tests.
- Tests also require [Docker Compose](https://docs.docker.com/compose/install/) to be installed.

## Example

See GoDocs for usage examples.

## Documentation

GoDocs [https://godoc.org/github.com/jimsmart/schema](https://godoc.org/github.com/jimsmart/schema)

## Testing

Database services for testing against are hosted in Docker.

To bring up the database services: execute `docker-compose up` inside the project folder, and wait until all of the Docker services have completed their startup (i.e. there is no further output in the terminal), then open a second terminal. (In future one may choose to use `docker-compose up -d` instead)

To run the tests execute `go test` inside the project folder.

For a full coverage report, try:

```bash
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

To shutdown the Docker services, execute `docker-compose down -v` inside the project folder.

### Oracle Setup Checklist

#### Build Docker Image

Build a Docker image for Oracle, by executing script:

```bash
./build_oracle_docker_image.sh
```

#### Increase Docker's RAM limits

By default, Docker allocates 2gb RAM to each container. To prevent out-of-memory errors when running Oracle, increase Docker's RAM limits.

Docker -> Preferences -> Resources -> Advanced -> Memory, change to 4gb, click Apply & Restart.

#### Install Oracle Client

Oracle database drivers require dynamic libraries that are part of the Oracle Client installation.

##### Mac

```bash
brew tap InstantClientTap/instantclient
brew install instantclient-basic
```

## License

Package schema is copyright 2018-2021 by Jim Smart and released under the [BSD 3-Clause License](LICENSE.md)

## History

- v0.2.0: Clean up front-facing API.
- v0.1.0: Added schema name to methods and results.
- v0.0.8: Disabled Oracle tests on Travis.
- v0.0.7: Added PrimaryKey method. TableNames and ViewNames are now sorted. Improved Oracle testing. Refactored dialect handling.
- v0.0.6: Fix Oracle quoting strategy. Added support for driver github.com/godror/godror.
- v0.0.5: Added dialect alias for Snowflake driver github.com/snowflakedb/gosnowflake.
- v0.0.4: Improved error handling for unknown DB driver types. Test environment now uses Docker.
- v0.0.3: Minor code cleanups.
- v0.0.2: Added identifier escaping for methods that query sql.ColumnType.
- v0.0.1: Started using Go modules.
- 2019-11-04: Fix for renamed driver struct in github.com/mattn/go-oci8 (Oracle)
- 2019-11-04: Fix for renamed driver struct in github.com/denisenkom/go-mssqldb (MSSQL)
