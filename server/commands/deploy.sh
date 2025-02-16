#!/bin/bash
################################################################################
#
#   Deploy
#
#       This bash file runs all commands to deploy all infrastructure for the
#       server.
#
################################################################################

current_directory=$(pwd)
server_root_directory="${current_directory%/*}"
infrastructure_directory="$server_root_directory/infrastructure"

cd $infrastructure_directory
cdk synth
cdk deploy