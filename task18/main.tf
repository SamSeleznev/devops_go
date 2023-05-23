provider "aws" {
  region = "ap-northeast-2"
}

resource "aws_instance" "ubuntuserver" {
  ami           = "ami-04cebc8d6c4f297a3"
  instance_type = "t2.micro"

  tags = {
    Name    = "UbuntuServer"
    Owner   = "Semen Seleznev"
    Project = "Task18Terraform"
  }
}
output "instance_ip_addr" {
  description = "The public IP address of the EC2 instance"
  value       = aws_instance.ubuntuserver.public_ip
}
