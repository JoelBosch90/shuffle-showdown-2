#!/bin/bash
################################################################################
#
#   Build
#
#       This bash file locally builds the client.
#
################################################################################

# Exit immediately if a command exits with a non-zero status.
set -e

current_directory=$(pwd)
client_root_directory="${current_directory%/*}"
app_directory="$client_root_directory/app"
infrastructure_directory="$client_root_directory/infrastructure"

# Build the app.
cd $app_directory
npm ci > /dev/null
npx playwright install-deps
npx playwright install
npm run build > /dev/null
echo "Built $app_directory"

# Build the infrastructure.
cd $infrastructure_directory
npm ci > /dev/null
echo "Built $infrastructure_directory"