#!/bin/bash
################################################################################
#
#   Build
#
#       This bash file locally builds the client.
#
################################################################################

current_directory=$(pwd)
client_root_directory="${current_directory%/*}"
infrastructure_directory="$client_root_directory/infrastructure"

# Build the infrastructure.
cd $infrastructure_directory
npm ci > /dev/null
echo "Built $infrastructure_directory"