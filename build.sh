#!/bin/sh


echo "Removing old";
rm dist/freepad;

# Remember current path
MYDIR=`pwd`;
# Go into src
cd src;
# Build
echo "Building..."
GIN_MODE=release CGO_ENABLED=0 GOOS=linux GIN_MODE=release go build -a -installsuffix cgo -o ../dist/freepad .
# Go back!
cd $MYDIR;