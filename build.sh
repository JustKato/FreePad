#!/bin/sh
echo "Building FreePad...\n";

echo "Removing old build file...";
rm    dist/freepad 2> /dev/null || true
rm -r dist/static 2> /dev/null || true
rm -r dist/templates 2> /dev/null || true
rm    dist/.env 2> /dev/null || true

# Build
echo "Building executable"
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./dist/freepad .

echo "Copying templates"
cp -r ./templates ./dist/templates
cp -r ./static ./dist/static

echo "Building Done";