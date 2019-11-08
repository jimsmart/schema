# Based upon comments here https://github.com/Microsoft/mssql-docker/issues/11

echo "db-init.sh waiting for SQL Server to start."
while [ true ]; do
    sleep 1s
    /opt/mssql-tools/bin/sqlcmd -l 30 -S localhost -h-1 -V1 -U sa -P "$SA_PASSWORD" -Q "select name from sys.databases where state_desc != 'ONLINE'" 2>/dev/null | grep --quiet '0 rows affected'
    if [ $? -eq 0 ]; then
        # All databases are online.
        break
    fi
    # Retry.
done

echo "SQL Server is up. db-init.sh now running init.sql script."
/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P "$SA_PASSWORD" -i init.sql

# TODO(js) Can we check whether this has already been run enforce 'only run once'?
# Or should that be the responsibility of the .sql script? Or?
