#!/bin/bash
################################################################################
#
#   Run
#
#       This bash file runs all commands to setup up a local development
#       environment for the server.
#
################################################################################

current_directory=$(pwd)
server_root_directory="${current_directory%/*}"
infrastructure_directory="$server_root_directory/src/infrastructure"

start_localstack() {
  echo "Starting LocalStack"
  cd $infrastructure_directory
  docker-compose up
}

deploy_cdk_stack() {
  echo "Deploying the CDK stack to LocalStack"
  cd $infrastructure_directory

  # Set dummy AWS credentials
  export AWS_ACCESS_KEY_ID=test
  export AWS_SECRET_ACCESS_KEY=test
  export AWS_REGION=eu-central-1

  # Prepare and deploy the CDK stack to LocalStack
  cdklocal bootstrap
  cdklocal synth
  cdklocal deploy --require-approval never
}

# Wait a few seconds for LocalStack to start before deploying the CDK stack
start_localstack & sleep 5 && deploy_cdk_stack & wait