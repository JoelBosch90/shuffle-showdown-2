#!/bin/bash
################################################################################
#
#   Test
#
#       This bash file runs all tests for the client application, shows a
#       command-line report detailing any failing tests, and calculates the
#       overall test coverage. It will fail if any tests fail or if the test
#       coverage is below the threshold.
#
################################################################################

# Exit immediately if a command exits with a non-zero status.
set -e

TEST_COVERAGE_THRESHOLD=100

current_directory=$(pwd)
client_root_directory=${current_directory%/*}
app_directory="$client_root_directory/app"
infrastructure_directory="$client_root_directory/infrastructure"

cd $app_directory
npm run test

cd $infrastructure_directory
npm run test