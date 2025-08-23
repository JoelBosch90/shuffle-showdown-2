#!/bin/bash
################################################################################
#
#   Run
#
#       This bash file runs all commands to setup up a local development
#       environment for the client.
#
################################################################################

# Exit immediately if a command exits with a non-zero status.
set -e

current_directory=$(pwd)
client_root_directory="${current_directory%/*}"
app_directory="$client_root_directory/app"
infrastructure_directory="$client_root_directory/infrastructure"

echo "Deploying the Client CDK stack to LocalStack"
cd $infrastructure_directory

bucket_name="clientstack-website-000000000000-eu-central-1"

# Set dummy AWS credentials
export AWS_ACCESS_KEY_ID=test
export AWS_SECRET_ACCESS_KEY=test
export AWS_REGION=eu-central-1
export PRIVATE_SHUFFLE_SHOWDOWN_DOMAIN_CERTIFICATE_ARN=arn:aws:acm:us-east-1:123456789012:certificate/12345678-1234-1234-1234-123456789012
export PUBLIC_SHUFFLE_SHOWDOWN_DOMAIN="$bucket_name"

# Prepare and deploy the CDK stack to LocalStack
cdklocal bootstrap
cdklocal doctor
cdklocal synth
cdklocal deploy --require-approval never

# Sync local assets to the LocalStack S3 bucket. This is necessary because CDK won't upload correctly to LocalStack.
aws s3 sync $app_directory/build s3://$bucket_name --endpoint-url http://localhost:4566