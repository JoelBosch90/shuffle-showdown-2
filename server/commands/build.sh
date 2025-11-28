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

# Exit immediately if a command exits with a non-zero status.
set -e

# By default, Go will include the version control system in the build. This
# flag will prevent that from happening. This is necesary to avoid errors
# when running this command in a CI/CD pipeline.
export GOFLAGS=-buildvcs=false

# Set the environment variables for the ARM64 architecture that is used
# in AWS.
export GOOS=linux
export GOARCH=arm64
export CGO_ENABLED=0

current_directory=$(pwd)
server_root_directory=${current_directory%/*}

build_directory() {
  if [ -d $1 ]; then
    cd $1

    go mod tidy
    go build -o bootstrap
    echo Built $1
    
    cd - > /dev/null
  fi
}

# Build infrastructure.
build_infrastructure() {
  build_directory $server_root_directory/infrastructure/
}

# Build all controllers in the project.
build_controllers() {
  for controller_directory in $server_root_directory/source/controllers/*/; do
    build_directory $controller_directory
  done
}

projects_to_build=(build_controllers build_infrastructure)
for project in ${projects_to_build[@]}; do
  ${project}
done