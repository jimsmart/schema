if [ "$0" != "sqlite"]; then
    exit 0;
fi
docker-compose up -d $0
