#!/bin/bash
################################################################################
#
#   Deploy
#
#       This bash file runs all commands to deploy all infrastructure for the
#       client.
#
################################################################################

# Exit immediately if a command exits with a non-zero status.
set -e

current_directory=$(pwd)
client_root_directory="${current_directory%/*}"
infrastructure_directory="$client_root_directory/infrastructure"

. ../../shared/commands/functions/load_environment.sh

cd $infrastructure_directory
cdk doctor
cdk synth
cdk deploy --require-approval never