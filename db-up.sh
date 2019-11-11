if [ "$1" = "sqlite" ]; then
    exit 0;
fi
if [ "$1" = "oracle" ]; then
    echo "Performing Docker login for Oracle Container Registry"
    echo "$OCR_PASSWORD" | docker login -u "$OCR_USERNAME" --password-stdin container-registry.oracle.com
fi

echo "Using Docker Compose to bring up service: $1"
docker-compose up -d $1
