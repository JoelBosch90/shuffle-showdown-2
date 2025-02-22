#!/bin/bash
################################################################################
#
#   Shuffle Showdown setup
#
#       This bash file processes some basic actions for the Shuffle Showdown
#       setup project. It currently supports the following commands:
#
#           update        Updates the local repository to the latest version.
#           release       Release local changes to the development branch and
#                         pushes them to the live version.
#           development   Spins up a local development environment.
#           production    Spins up the current production version.
#           commit        Will commit and push the local changes on the current
#                         branch to the remote repository.
#                         (Takes an extra argument as the commit message.)
#
#   Usage example
#
#       You can use this file by executing this file and adding the commands
#       above seperated by spaces. These commands will be executed in order.
#       For example, to update the local repository and then run a local
#       development environment, you can run the following:
#
#           shuffle update dev
#
#       Some command require an extra argument. You can call these like this:
#
#           shuffle commit "Commit message"
#
#   Requirements
#
#       To properly use this file, there are two requirements that must be met:
#
#           Execute the following command to give this file permission to be
#           executed:
#
#               sudo chmod +x shortcuts.sh
#
#           Add a shortcut so that you can execute this file from anywhere and
#           no longer need to write the extension:
#
#               sudo ln -s $(pwd)/shortcuts.sh /usr/bin/shuffle
#
################################################################################

# Get access to the project's working directory.
WORKING_DIRECTORY="$(dirname "$(readlink -f "$0")")";
SERVER_COMMANDS_DIRECTORY="$WORKING_DIRECTORY/server/commands";
CLIENT_COMMANDS_DIRECTORY="$WORKING_DIRECTORY/client/commands";

################################################################################
#
#   mock_server
#       Function to generate mocks for the server application.
#
################################################################################
mock_server () {
  cd $SERVER_COMMANDS_DIRECTORY;

  ./mock.sh;
}

################################################################################
#
#   mock
#       Function to generate mocks for all applications.
#
################################################################################
mock () {
  mock_server;
}

################################################################################
#
#   build_server
#       Function to locally build the server application to make it ready to
#       deploy.
#
################################################################################
build_server () {
  cd $SERVER_COMMANDS_DIRECTORY;

  ./build.sh;
}

################################################################################
#
#   build
#       Function to locally build both server and client applications to make
#       them ready to deploy.
#
################################################################################
build () {
  build_server;
}

################################################################################
#
#   testClient
#       Function to locally run the full test suite for the client application.
#
################################################################################
test_client () {
  cd $CLIENT_COMMANDS_DIRECTORY;

  ./test.sh;
}

################################################################################
#
#   testServer
#       Function to locally run the full test suite for the server application.
#
################################################################################
test_server () {
  cd $SERVER_COMMANDS_DIRECTORY;

  ./test.sh;
}

################################################################################
#
#   test
#       Function to locally run the full test suite.
#
################################################################################
test () {
  test_server & test_client & wait;
}

################################################################################
#
#   run_localstack
#       Function to spin up a local version of the AWS cloud.
#
################################################################################
run_localstack () {
  cd $WORKING_DIRECTORY;
  docker-compose up
}

################################################################################
#
#   run_client
#       Function to spin up a local client development environment.
#
################################################################################
run_client () {
  cd $CLIENT_COMMANDS_DIRECTORY;

  ./run.sh;
}

################################################################################
#
#   run_server
#       Function to spin up a local server development environment.
#
################################################################################
run_server () {
  cd $SERVER_COMMANDS_DIRECTORY;

  ./run.sh;
}

################################################################################
#
#   run
#       Function to spin up a local development environment.
#
################################################################################
run () {
  # Allow for time to start up the localstack first.
  # Then deploy the server and client to the localstack.
  run_localstack & sleep 5 && run_server && run_client & wait;
}

################################################################################
#
#   deploy_server
#       Function to deploy the infrastructure for the server application to the
#       cloud.
#
################################################################################
deploy_server () {
  cd $SERVER_COMMANDS_DIRECTORY;

  ./deploy.sh;
}

################################################################################
#
#   deploy_client
#       Function to deploy the infrastructure for the client application to the
#       cloud.
#
################################################################################
deploy_client () {
  cd $CLIENT_COMMANDS_DIRECTORY;

  ./deploy.sh;
}

################################################################################
#
#   deploy
#       Function to deploy the infrastructure for the server and client
#       applications to the cloud.
#
################################################################################
deploy () {
  deploy_server & deploy_client & wait;
}

################################################################################
#
#   release
#       Function to release all development changes to production. GitHub
#       actions will pick this up and automatically release them on the live
#       server as well.
#
################################################################################
# release () {

# }

# Loop through the command line arguments.
while [[ $# -gt 0 ]]; do

  # Give more meaningful names to the command line arguments.
  command="$1"
  argument="$2"

  # Determine per command what to do.
  case "$command" in

    # Run `shuffle run` to spin up a local development environment.
    run)
      run
      shift # Get ready to process the next command.
      ;;

    # Run `shuffle mock` to generate all mocks for the server application.
    mock)
      mock
      shift # Get ready to process the next command.
      ;;

    # Run `shuffle test` to run the full test suite for all application.
    test)
      test
      shift # Get ready to process the next command.
      ;;

    # Run `shuffle build` to build all applications.
    build)
      build
      shift # Get ready to process the next command.
      ;;

    # Run `shuffle deploy` to deploy all applications to the cloud.
    deploy)
      deploy
      shift # Get ready to process the next command.
      ;;

    # Run `shuffle release` to release the current version of the
    # development branch and roll those changes out to the live version.
    release)
      release
      shift # Get ready to process the next command.
      ;;
  esac
done