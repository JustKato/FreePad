#!/bin/bash

echo "Removing old";
rm dist/freepad;

# Remember current path
MYDIR=`pwd`;
# Go into src
cd src;
# Build
echo "Building..."
export GIN_MODE=debug;
go build -o ../dist/freepad .
# Go back!
cd $MYDIR;

MYPATH=`pwd`;

cd dist

./freepad && cd $MYPATH;