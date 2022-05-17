#!/bin/sh


echo "Removing old";
rm dist/freepad;

# Remember current path
MYDIR=`pwd`;
# Go into src
cd src;
# Build
echo "Building..."
export RELEASE_MODE=1;
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ../dist/freepad .
# Go back!
cd $MYDIR;