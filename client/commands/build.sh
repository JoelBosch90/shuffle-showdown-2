#!/bin/bash
################################################################################
#
#   Build
#
#       This bash file builds the cdk stack for the client.
#
################################################################################

current_directory=$(pwd)
client_root_directory="${current_directory%/*}"
infrastructure_directory="$client_root_directory/src/infrastructure"

cd $infrastructure_directory
# TDB: Add build command here