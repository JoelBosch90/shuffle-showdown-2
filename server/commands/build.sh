#!/bin/bash
################################################################################
#
#   Build
#
#       This bash file loops through all the controllers in the project and
#       builds them for the ARM64 architecture. This allows the project to be
#       deployed to the cloud.
#
################################################################################

current_directory=$(pwd)
server_root_directory="${current_directory%/*}"
controllers_directory="$server_root_directory/src/controllers"

for controller_directory in $controllers_directory/*/; do
  if [ -d "$controller_directory" ]; then
    cd "$controller_directory"

    echo "Building controller in $controller_directory"
    go mod tidy
    GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap
    
    cd - > /dev/null
  fi
done