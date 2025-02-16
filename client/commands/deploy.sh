#!/bin/bash
################################################################################
#
#   Deploy
#
#       This bash file runs all commands to deploy all infrastructure for the
#       client.
#
################################################################################

current_directory=$(pwd)
client_root_directory="${current_directory%/*}"
infrastructure_directory="$client_root_directory/infrastructure"

cd $infrastructure_directory
cdk synth
cdk deploy