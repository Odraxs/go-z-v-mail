#!/bin/bash

# Update package repositories
sudo apt-get update -y

# Install dependencies
sudo apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    gnupg-agent \
    software-properties-common

# Add Docker's official GPG key
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -

# Add Docker's repository
sudo add-apt-repository "deb [arch=amd64] https://download.docker.com/linux/ubuntu $(lsb_release -cs) stable"

# Install Docker
sudo apt-get update -y
sudo apt-get install -y docker-ce docker-ce-cli containerd.io

# Add current user to the docker group (to avoid using sudo with Docker)
sudo usermod -aG docker $USER

# Install Docker Compose
sudo curl -L "https://github.com/docker/compose/releases/download/1.29.2/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
sudo chmod +x /usr/local/bin/docker-compose

#access user root
sudo su

#create project folder
mkdir -p ~/project
cd ~/project

#clone your project from git repository
git clone https://github.com/Odraxs/go-z-v-mail.git
cd go-z-v-mail
git checkout docker-compose-for-terraform-deploy
git pull
chmod +x envs.sh
chmod a+rwx ./data-embedding
. envs.sh

#run your docker compose file
# Running the full project docker compose is too expensive for the t2.mico instance also only 1GB of memory so is not possible to set all the data into zincsearch
cd docker
docker-compose up -d
