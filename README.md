# Shuffle Showdown 2
Welcome to Shuffle Showdown, the ultimate music guessing game! Shuffle Showdown is a web-based game where players compete to guess songs from a shuffled Spotify playlist. Each team takes turns listening to songs in random order, and they must guess the release year and place the song in order with their previously won songs. With its blend of music trivia and playlist shuffling, Shuffle Showdown offers endless fun for music lovers of all ages.

Features:
 - Use any Spotify playlist and start the game instantly.
 - Play with friends or family in teams and compete to be the first to gather 10 songs.

Get ready to put your music knowledge to the test and embark on a musical journey with Shuffle Showdown!

# Why version 2?
After Spotify updated their API and discontinued the preview links that the first version of
Shuffle Showdown relied on in November 2024, the entire project was restarted in this repository.
This follow-up version aims to add the following features to Shuffle Showdown:
 - Full infrastructure as code
 - Full test coverage
 - Deeper Spotify integration (will require login to play full songs)
 - More iterations on the user interface
 - New features

## Prerequisites
 - Install docker (https://docs.docker.com/engine/install/)
 - Install docker-compose (https://docs.docker.com/compose/install/linux/)
 - Install aws-cdk-local & aws-cdk (`npm install -g aws-cdk-local aws-cdk`)
 - Install LocalStack (https://docs.localstack.cloud/getting-started/installation/)

## Useful commands

 * `aws sso login`        authenticate with AWS
 * `./commands/build.sh`  pre-build the executables for all handlers
 * `cdk deploy`           deploy this stack to your default AWS account/region
 * `cdk diff`             compare deployed stack with current state
 * `cdk synth`            emits the synthesized CloudFormation template
 * `go test`              run unit tests