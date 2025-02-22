#!/bin/bash
################################################################################
#
#   Run
#
#       This bash file runs all commands to setup up a local development
#       environment for the client.
#
################################################################################

current_directory=$(pwd)
client_root_directory="${current_directory%/*}"
infrastructure_directory="$client_root_directory/infrastructure"

echo "Deploying the Client CDK stack to LocalStack"
cd $infrastructure_directory

# Set dummy AWS credentials
export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
export AWS_REGION=eu-central-1

# Prepare and deploy the CDK stack to LocalStack
cdklocal bootstrap
cdklocal synth
cdklocal deploy --require-approval never

# Sync local assets to the LocalStack S3 bucket
aws s3 sync $client_root_directory/source/assets s3://website --endpoint-url http://localhost:4566
aws s3 website s3://website --index-document index.html --error-document error.html --endpoint-url http://localhost:4566