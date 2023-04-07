provider "aws" {
  region = var.aws_region
}

resource "aws_instance" "ec2_instance" {
  ami                         = var.ami
  instance_type               = var.instance_type
  key_name                    = var.key_name
  subnet_id                   = var.subnet_id
  vpc_security_group_ids      = var.vpc_security_group_ids
  associate_public_ip_address = var.associate_public_ip_address
  availability_zone           = var.availability_zone
  tags = {
    Name = var.instance_name
  }
  count = var.instance_count
}