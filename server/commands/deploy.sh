#!/bin/bash
################################################################################
#
#   Deploy
#
#       This bash file runs all commands to deploy all infrastructure for the
#       server.
#
################################################################################

# Exit immediately if a command exits with a non-zero status.
set -e

# By default, Go will include the version control system in the build. This
# flag will prevent that from happening. This is necesary to avoid errors
# when running this command in a CI/CD pipeline.
export GOFLAGS=-buildvcs=false

current_directory=$(pwd)
server_root_directory="${current_directory%/*}"
infrastructure_directory="$server_root_directory/infrastructure"

cd $infrastructure_directory
cdk synth
cdk deploy --require-approval never