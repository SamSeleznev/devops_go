#!/bin/bash

# SSH credentials
ssh_user="ubvmaster"
ssh_host="192.168.1.102"

# Docker container details
container_name="webserv"
image_name="webserv"
image_path="/home/ubvmaster/devops_go"
port_mapping="8080:8080"

# Stop and remove container
ssh $ssh_user@$ssh_host "docker stop $container_name"
ssh $ssh_user@$ssh_host "docker rm $container_name"

# Remove image
ssh $ssh_user@$ssh_host "docker rmi $image_name"

# Copy application files to the server
scp -rq $image_path $ssh_user@$ssh_host:/home/$ssh_user/

# Build image from Dockerfile
ssh $ssh_user@$ssh_host "docker build -t $image_name /home/$ssh_user/devops_go"

# Run container
ssh $ssh_user@$ssh_host "docker run -d --restart=always -p $port_mapping --name $container_name $image_name"
