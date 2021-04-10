CREATE LOGIN test_user WITH PASSWORD = 'Password-123';
GO
CREATE USER test_user FOR LOGIN test_user;
GO
CREATE SCHEMA test_db AUTHORIZATION test_user;
GO
ALTER USER test_user WITH default_schema = test_db;
GO
EXEC sp_addrolemember db_ddladmin, test_user;
GO
