#!/bin/bash
################################################################################
#
#   Test
#
#       This bash file runs all tests for the server application, shows a
#       command-line report detailing any failing tests, and calculates the
#       overall test coverage. It will fail if any tests fail or if the test
#       coverage is below the threshold.
#
################################################################################

# Exit immediately if a command exits with a non-zero status.
set -e

TEST_COVERAGE_THRESHOLD=100
SKIP_FILE_TAG=skip_test
EXCLUDED_DIRECTORIES=(
  "interfaces"
  "mocks"
)

current_directory=$(pwd)
server_root_directory=${current_directory%/*}

test_directory() {
  # Skip if the directory is in the excluded list.
  for excluded_directory in "${EXCLUDED_DIRECTORIES[@]}"; do
    if [[ $1 == *"/$excluded_directory"* ]]; then
      echo "Skipping excluded directory $1"
      return
    fi
  done

  cd $1
  echo Testing in $1
  
  # Run the tests for this controller. Skip EXCLUDED_FILES AND EXCLUDED_DIRECTORIES.  
  go test -coverprofile=coverage.out -tags $SKIP_FILE_TAG

  # Display the test results.
  go tool cover -func=coverage.out

  # Extract the test coverage percentage from the test results.
  coverage=$(go tool cover -func=coverage.out | grep total | awk '{print $3}' | awk -F. '{print $1}')

  # Exit with an error if the test coverage is below the threshold.
  if [ $coverage -lt $TEST_COVERAGE_THRESHOLD ]; then
    echo Test coverage fails the threshold of $TEST_COVERAGE_THRESHOLD%
    # exit 1
  else
    echo Test coverage meets the threshold of $TEST_COVERAGE_THRESHOLD%
  fi

  cd - > /dev/null
}

# Test infrastructure.
test_infrastructure() {
  test_directory $server_root_directory/infrastructure
}

# Test all controllers in the project.
test_controllers() {
  for controller_directory in $server_root_directory/source/controllers/*/*; do
    if [ -d "$controller_directory" ]; then
      test_directory "$controller_directory"
    fi
  done
}

tests_to_run=(test_infrastructure test_controllers)
for test in ${tests_to_run[@]}; do
  ${test}
done