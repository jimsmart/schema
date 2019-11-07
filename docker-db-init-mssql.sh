# TODO(js) This sleep is terrible, we need to find a better solution
# (I think it will cause problems on Travis)
# We should instead do something like https://github.com/Microsoft/mssql-docker/issues/11
sleep 30s

echo "running MSSQL setup script"
/opt/mssql-tools/bin/sqlcmd -S localhost -U sa -P 7kRZ4mUsSD4XedMq -i init.sql

# TODO(js) For Travis, we should also:
# - install microsoft's sql cli tools
# - we should then invoke a local script that performs a similar wait to the above
#   (here we'd be waiting outside docker, for mssql to start, before running our tests)
# - We can maybe use travisretry here?
#   See https://docs.travis-ci.com/user/common-build-problems/#timeouts-installing-dependencies
