# TODO(js) This sleep is terrible, we need to find a better solution
# (I think it will cause problems on Travis)
# We should instead do something like https://github.com/Microsoft/mssql-docker/issues/11
sleep 30s

echo "running MSSQL setup script"
/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P 7kRZ4mUsSD4XedMq -i init.sql
