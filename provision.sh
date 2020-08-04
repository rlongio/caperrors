#!/usr/bin/env bash

### Update package repo
sudo yum update -y && sudo yum upgrade -y

### Install yum utils

# https://linux.die.net/man/1/yum-utils
sudo yum install yum utils -y

# Go grab the EPEL repo
sudo yum install epel-release -y

### Install development tools group
sudo yum groupinstall 'Development Tools' -y

### Install Python3, this includes pip and virtualenv
sudo yum install python3 python3-setuptools vim nano zsh wget -y

# Install virtualenv
pip3 install --user virtualenv

### Install Java JDK
sudo yum install java-1.8.0-openjdk-devel -y

### Install Go
GO_TAR_LOCATION='https://dl.google.com/go/go1.14.4.linux-amd64.tar.gz'
GO_TAR_FILENAME='go1.14.4.linux-amd64.tar.gz'

# Fetch from server
echo "Downloading go from $GO_TAR_LOCATION"
wget -q "$GO_TAR_LOCATION"

# Verify SHA256 Checksum
EXPECTED_SIGNATURE='aed845e4185a0b2a3c3d5e1d0a35491702c55889192bb9c30e67a3de6849c067'

echo "Verifying checksum of download"
echo "$EXPECTED_SIGNATURE $GO_TAR_FILENAME" | sha256sum -c
CHECKSUM_RESULT=$?
if [[ $CHECKSUM_RESULT -ne 0 ]]
then
    echo "Checksums do not match for $GO_TAR_FILENAME - $EXPECTED_SIGNATURE"
    exit 1
fi

# Extract to /usr/local
echo "Extracting tar to /usr/local"
sudo tar -C /usr/local -xzf $GO_TAR_FILENAME
if [[ $? -ne 0 ]]
then
    echo "Unable to extract from $GO_TAR_FILENAME"
    exit 2
fi

# Cleanup
echo "Removing $GO_TAR_FILENAME"
rm "$GO_TAR_FILENAME"

# Add to path and reload /etc/profile
echo "Adding GO to path"
echo "export PATH=$PATH:/usr/local/go/bin" | sudo tee -a /etc/profile.d/go_lang.sh
if [[ $? -ne 0 ]]
then
    echo "Unable to add GO to path"
    exit 3
fi
source /etc/profile

# Test successfully installed
echo "Testing if installation successful......"
go version > /dev/null 2&>1
if [[ $? -ne 0 ]]
then
    echo "Go failed to install"
    exit 4
else 
    echo "Go installation successful"
    exit 0
fi






