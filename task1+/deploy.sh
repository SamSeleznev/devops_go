#!/bin/bash

# SSH credentials
ssh_user="ubvmaster"
ssh_host="192.168.1.102"

# Docker container details
container_name="webserv"
image_name="webserv"
image_path="/home/ubvmaster/devops_go"
port_mapping="8080:8080"

echo "Connecting to $ssh_host as $ssh_user..."
# Stop and remove container
echo "Stopping and removing existing container: $container_name"
ssh $ssh_user@$ssh_host "docker stop $container_name" >/dev/null 2>&1
ssh $ssh_user@$ssh_host "docker rm $container_name" >/dev/null 2>&1

# Remove image
echo "Removing existing image: $image_name"
ssh $ssh_user@$ssh_host "docker rmi $image_name" >/dev/null 2>&1

# Copy application files to the server
echo "Copying application files to $ssh_host"
scp -q -r $image_path $ssh_user@$ssh_host:/home/$ssh_user/ >/dev/null 2>&1

# Build image from Dockerfile
echo "Building new Docker image: $image_name"
ssh $ssh_user@$ssh_host "docker build -t $image_name /home/$ssh_user/devops_go"

# Run container
echo "Starting new container: $container_name"
ssh $ssh_user@$ssh_host "docker run -d --restart=always -p $port_mapping --name $container_name $image_name"
