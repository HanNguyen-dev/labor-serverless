#!/bin/sh

# build config for AWS Lambda
export GOOS=linux GOARCH=amd64 CGO_ENABLED=0;

# build the inflation lambda function

cd inflation;
go build;

# build the unemployment lambda function

cd ../unemployment;
go build;

# compress and pack builds
cd ..;

zip inflation.zip inflation/inflation;
zip unemployment.zip unemployment/unemployment;
