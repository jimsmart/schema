alter session set "_ORACLE_SCRIPT"=true; 
CREATE USER test_user IDENTIFIED BY Password123;
GRANT CONNECT, RESOURCE, DBA TO test_user;
