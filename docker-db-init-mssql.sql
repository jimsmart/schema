CREATE LOGIN mssql_test_user WITH PASSWORD = 'Password-123';
GO
CREATE USER mssql_test_user FOR LOGIN mssql_test_user;
GO
CREATE SCHEMA test_db AUTHORIZATION mssql_test_user;
GO
ALTER USER mssql_test_user WITH default_schema = test_db;
GO
EXEC sp_addrolemember db_ddladmin, mssql_test_user;
GO

-- CREATE TABLE web_resource (
--             id				INTEGER NOT NULL,
--             url				NVARCHAR NOT NULL UNIQUE,
--             content			VARBINARY,
--             compressed_size	INTEGER NOT NULL,
--             content_length	INTEGER NOT NULL,
--             content_type	NVARCHAR NOT NULL,
--             etag			NVARCHAR NOT NULL,
--             last_modified	NVARCHAR NOT NULL,
--             created_at		DATETIME NOT NULL,
--             modified_at		DATETIME,
--             PRIMARY KEY (id)
-- );
-- GO
-- CREATE INDEX idx_web_resource_url ON web_resource(url);
-- GO
-- CREATE INDEX idx_web_resource_created_at ON web_resource (created_at);
-- GO
-- CREATE INDEX idx_web_resource_modified_at ON web_resource (modified_at);
-- GO
-- CREATE VIEW web_resource_view AS SELECT id, url FROM web_resource;
-- GO

-- -- Tests for correct identifer escaping.
-- CREATE TABLE [blanks in name] (id INTEGER, PRIMARY KEY (id));
-- GO
-- CREATE TABLE [[brackets]] in name] (id INTEGER, PRIMARY KEY (id));
-- GO
-- CREATE TABLE ["d.quotes" in name] (id INTEGER, PRIMARY KEY (id));
-- GO
-- CREATE TABLE ['s.quotes' in name] (id INTEGER, PRIMARY KEY (id));
-- GO
-- CREATE TABLE [{braces} in name] (id INTEGER, PRIMARY KEY (id));
-- GO
-- CREATE TABLE [`backticks` in name] (id INTEGER, PRIMARY KEY (id));
-- GO
-- CREATE TABLE [backslashes\in\name] (id INTEGER, PRIMARY KEY (id));
-- GO
