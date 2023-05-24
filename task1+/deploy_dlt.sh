#!/bin/bash

# stop&dlt container
ssh ubvmaster@192.168.1.102 "docker stop webserv"
ssh ubvmaster@192.168.1.102 "docker rm webserv"

# dlt image
ssh ubvmaster@192.168.1.102 "docker rmi webserv"

# dlt app
ssh ubvmaster@192.168.1.102 "rm -rf /home/ubvmaster/devops_go"

echo "DONE."
