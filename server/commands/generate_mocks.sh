#!/bin/bash
################################################################################
#
#   GenerateMocks
#
#       This bash file loops through all interfaces in the project and generates
#       a mock for each one. It will fail if any mocks fail to generate.
#
################################################################################

TEST_COVERAGE_THRESHOLD=100
SKIP_FILE_TAG="skip_test"

current_directory=$(pwd)
server_root_directory="${current_directory%/*}"

mock_directory() {
  cd "$1"
  echo "Generating mocks in $1"

  for file_path in $1/interfaces/*; do
    if [[ "$file_path" =~ /([^/]*?).go$ ]]; then
      file_name=${BASH_REMATCH[1]}
      
      if grep -q "type $file_name interface" $file_path; then
        echo "Generating mock for $file_name in $1"
        mockgen -source=interfaces/$file_name.go -package=mocks -destination=mocks/$file_name.go $file_name
      fi
    fi
  done
  
  cd - > /dev/null
}

# Test infrastructure.
mock_infrastructure() {
  mock_directory "$server_root_directory/src/infrastructure"
}

# Test all controllers in the project.
mock_controllers() {
  for controller_directory in $server_root_directory/src/controllers/*/; do
    if [ -d "$controller_directory" ]; then
      mock_directory "$controller_directory"
    fi
  done
}

mocks_to_generate=(mock_infrastructure mock_controllers)
for mock in "${mocks_to_generate[@]}"; do
  ${mock}
done