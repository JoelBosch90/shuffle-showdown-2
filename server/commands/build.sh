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
server_root_directory=${current_directory%/*}

build_directory() {
  if [ -d $1 ]; then
    cd $1

    go mod tidy
    GOOS=linux GOARCH=arm64 CGO_ENABLED=0 go build -o bootstrap
    echo Built $1
    
    cd - > /dev/null
  fi
}

# Test infrastructure.
build_infrastructure() {
  build_directory $server_root_directory/src/infrastructure/
}

# Test all controllers in the project.
build_controllers() {
  for controller_directory in $server_root_directory/src/controllers/*/; do
    build_directory $controller_directory
  done
}

projects_to_build=(build_controllers build_infrastructure)
for project in ${projects_to_build[@]}; do
  ${project}
done