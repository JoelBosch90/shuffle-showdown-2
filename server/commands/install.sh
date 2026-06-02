#!/bin/bash
################################################################################
#
#   Install
#
#	   This bash file runs all commands to install all initial dependencies for
#	   the server.
#
################################################################################

# Exit immediately if a command exits with a non-zero status.
set -e

LOCALSTACK_VERSION=2026.4.0
GOLANG_VERSION=1.24.0
ARCHITECTURE=linux-amd64

# https://docs.docker.com/engine/install/ubuntu/
install_docker() {
    if docker version &> /dev/null; then
        echo "Docker is already installed. Skipping installation."
        return
    fi

	# Add Docker's official GPG key:
	sudo apt update
	sudo apt -y install ca-certificates curl
	sudo install -m 0755 -d /etc/apt/keyrings
	sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
	sudo chmod a+r /etc/apt/keyrings/docker.asc

	# Add the repository to Apt sources:
	sudo tee /etc/apt/sources.list.d/docker.sources <<-EOF
	Types: deb
	URIs: https://download.docker.com/linux/ubuntu
	Suites: $(. /etc/os-release && echo "${UBUNTU_CODENAME:-$VERSION_CODENAME}")
	Components: stable
	Architectures: $(dpkg --print-architecture)
	Signed-By: /etc/apt/keyrings/docker.asc
	EOF

	sudo apt update
	sudo apt -y install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
	
    if sudo systemctl status docker &> /dev/null; then
        echo "Docker service is already running. Skipping starting the service."
    else
        sudo systemctl start docker
    fi
}

# https://docs.docker.com/compose/install/linux/
install_docker_compose() {
    if docker compose version &> /dev/null; then
        echo "Docker Compose is already installed. Skipping installation."
        return
    fi

    sudo apt-get update
    sudo apt-get install -y docker-compose-plugin
}

install_aws_dependencies() {
    if ! npm -v &> /dev/null; then
        sudo apt -y install npm
    fi

	sudo npm install -g aws-cdk-local aws-cdk
}

# https://docs.localstack.cloud/aws/getting-started/installation/
install_localstack() {
    if localstack -v &> /dev/null; then
        echo "LocalStack CLI is already installed. Skipping installation."
        return
    fi

	curl --output localstack-cli-${LOCALSTACK_VERSION}-${ARCHITECTURE}-onefile.tar.gz	 --location https://github.com/localstack/localstack-cli/releases/download/v${LOCALSTACK_VERSION}/localstack-cli-${LOCALSTACK_VERSION}-${ARCHITECTURE}-onefile.tar.gz
	sudo tar xvzf localstack-cli-${LOCALSTACK_VERSION}-${ARCHITECTURE}-onefile.tar.gz -C /usr/local/bin
}

# https://go.dev/doc/install
install_golang() {
    if go version &> /dev/null; then
        echo "Golang is already installed. Skipping installation."
        return
    fi

	sudo rm -rf /usr/local/go
	curl --output go${GOLANG_VERSION}.${ARCHITECTURE}.tar.gz --location https://go.dev/dl/go${GOLANG_VERSION}.${ARCHITECTURE}.tar.gz
	sudo tar -C /usr/local -xzf go${GOLANG_VERSION}.${ARCHITECTURE}.tar.gz
	echo -e "\n# Set Golang path. Set by Shuffle Showdown Installation\nexport PATH=\$PATH:/usr/local/go/bin" >> /home/$SUDO_USER/.bashrc
    sudo rm go${GOLANG_VERSION}.${ARCHITECTURE}.tar.gz
}

install_docker
install_docker_compose
install_aws_dependencies
install_localstack
install_golang