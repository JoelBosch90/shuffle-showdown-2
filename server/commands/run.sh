#!/bin/bash
################################################################################
#
#   Run
#
#       This bash file runs all commands to setup up a local development
#       environment for the server.
#
################################################################################

# Exit immediately if a command exits with a non-zero status.
set -e

current_directory=$(pwd)
server_root_directory="${current_directory%/*}"
infrastructure_directory="$server_root_directory/infrastructure"

echo "Deploying the Server CDK stack to LocalStack"
cd $infrastructure_directory

# Set dummy AWS credentials
export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
export AWS_REGION=eu-central-1

# Prepare and deploy the CDK stack to LocalStack
cdklocal bootstrap
cdklocal synth
cdklocal deploy --require-approval never