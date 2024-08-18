#!/bin/bash

GOOS=darwin GOARCH=amd64 go build -o gokubectl
mv gokubectl builds/macos