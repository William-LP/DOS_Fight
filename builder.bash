#!/usr/bin/bash

rm -rf ./bin/
appname=$(basename $( pwd ))
go build -o "./bin/$appname.exe"

