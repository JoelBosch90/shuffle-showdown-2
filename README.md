# shuffle-showdown-2
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