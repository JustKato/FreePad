#!/bin/bash

echo "Removing old";
rm dist/freepad;

source ../.env
# Yeah, this is my solution
export DOMAIN_BASE CACHE_MAP_LIMIT API_BAN_LIMIT DATABASE_DRIVER MYSQL_ROOT_PASSWORD MYSQL_DATABASE MYSQL_USER MYSQL_PASSWORD MYSQL_URL MYSQL_PORT IS_DEV

# Remember current path
MYDIR=`pwd`;
# Go into src
cd src;
# Build
echo "Building..."

go build -o ../dist/freepad .
# Go back!
cd $MYDIR;

cd dist

./freepad && cd $MYDIR;