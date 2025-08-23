#!/bin/bash
################################################################################
#
#   load_environment
#
#     Function to safely load the root folder environment.
#
################################################################################

# Get the .env where this script is located
environment_file="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/../../../.env"

if [ -f "$environment_file" ]; then
  echo "Loading environment variables from $environment_file"
  
  while IFS= read -r line || [ -n "$line" ]; do
    # Skip empty lines and comments
    [[ -z "$line" || "$line" =~ ^[[:space:]]*# ]] && continue

    # Check if the line is in the expected "VAR=value" format
    if [[ "$line" =~ ^[A-Za-z_][A-Za-z0-9_]*= ]]; then
      var_name="${line%%=*}"

      # Export the variable to handle special characters safely
      export "$var_name"="${line#*=}"
      echo "Exported: $var_name"
    fi
  done < "$environment_file"
else
  echo "Warning: $environment_file not found"
fi