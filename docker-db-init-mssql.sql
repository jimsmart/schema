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
