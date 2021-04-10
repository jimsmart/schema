#!/bin/bash

# Downloads Oracle XE 18c, clones the official image scripts,
# and builds image 'oracle/database:18.4.0-xe'

# Based upon https://www.petefreitag.com/item/886.cfm

mkdir -p ./test/docker-oracle
cd ./test/docker-oracle

git clone https://github.com/oracle/docker-images.git

cd ./docker-images/OracleDatabase/SingleInstance/dockerfiles

curl -L https://download.oracle.com/otn-pub/otn_software/db-express/oracle-database-xe-18c-1.0-1.x86_64.rpm --output ./18.4.0/oracle-database-xe-18c-1.0-1.x86_64.rpm

./buildContainerImage.sh -x -v 18.4.0

echo "\nDocker image built ok"
echo "\nIt is safe to remove the work folder with: rm -rf ./test/docker-oracle"
