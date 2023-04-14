#!/bin/bash

# Останавливаем и удаляем контейнер
ssh ubvmaster@192.168.1.102 "docker stop webserv"
ssh ubvmaster@192.168.1.102 "docker rm webserv"

# Удаляем образ
ssh ubvmaster@192.168.1.102 "docker rmi webserv"

# Удаляем приложение
ssh ubvmaster@192.168.1.102 "rm -rf /home/ubvmaster/devops_go"

echo "Все действия, выполненные предыдущим скриптом, отменены."
