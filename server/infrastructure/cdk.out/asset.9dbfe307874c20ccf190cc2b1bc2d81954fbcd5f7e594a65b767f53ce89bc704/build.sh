#!/bin/bash
GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap
# chmod +x bootstrap