# Based upon comments here https://github.com/Microsoft/mssql-docker/issues/11

# TODO(js) This sleep is terrible, we need to find a better solution
# (I think it will cause problems on Travis)
# We should instead do something like https://github.com/Microsoft/mssql-docker/issues/11
sleep 30s

echo "running MSSQL setup script"
/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P $SA_PASSWORD -i init.sql
