#!/bin/sh
echo "Building FreePad...\n";

echo "Removing old build file...";
rm    dist/freepad* 2> /dev/null || true
rm -r dist/static 2> /dev/null || true
rm -r dist/templates 2> /dev/null || true
rm    dist/.env 2> /dev/null || true

# Build
echo "Building executable"
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./dist/freepad .
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o ./dist/freepad-arm64 .
CGO_ENABLED=0 GOOS=windows go build -a -installsuffix cgo -o ./dist/freepad.exe .
CGO_ENABLED=0 GOOS=darwin go build -a -installsuffix cgo -o ./dist/freepad-darwin .
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -installsuffix cgo -o ./dist/freepad-darwin-64 .

echo "Copying templates"
cp -r ./templates ./dist/templates
cp -r ./static ./dist/static

echo "Building Done";