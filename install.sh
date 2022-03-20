#!/bin/bash

curl -fsSL https://apt.releases.hashicorp.com/gpg | sudo apt-key add -

sudo apt-add-repository "deb [arch=amd64] https://apt.releases.hashicorp.com $(lsb_release -cs) main"

sudo apt-get update && sudo apt-get install terraform


curl -O https://dl.google.com/dl/cloudsdk/channels/rapid/downloads/google-cloud-sdk-377.0.0-linux-x86_64.tar.gz

tar -xf google-cloud-sdk-377.0.0-linux-x86_64.tar.gz 
./google-cloud-sdk/install.sh
./google-cloud-sdk/bin/gcloud init
