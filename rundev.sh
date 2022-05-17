#!/bin/bash

echo "Removing old";
rm dist/freepad;

# Yeah, this is my solution
export DOMAIN_BASE CACHE_MAP_LIMIT API_BAN_LIMIT DATABASE_DRIVER MYSQL_ROOT_PASSWORD MYSQL_DATABASE MYSQL_USER MYSQL_PASSWORD MYSQL_URL MYSQL_PORT

# Remember current path
MYDIR=`pwd`;
# Go into src
cd src;
# Build
echo "Building..."
unset RELEASE_MODE;
go build -o ../dist/freepad .
# Go back!
cd $MYDIR;

MYPATH=`pwd`;

cd dist

source ../.env
./freepad && cd $MYPATH;