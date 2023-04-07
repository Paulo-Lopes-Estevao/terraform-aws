variable "aws_region" {
  description = "AWS region"
  default     = "us-east-1"
}

variable "instance_type" {
  description = "AWS instance type"
  default     = "t2.micro"
}

variable "instance_name" {
  description = "AWS instance name"
  type        = string
  default     = "my-app"
}

variable "instance_count" {
  description = "Number of instances to launch"
  type        = number
  default     = 1
}

variable "ami" {
  description = "ID of AMI to use for the instance"
  type        = string
  default     = "ami-007855ac798b5175e"
}

variable "key_name" {
  description = "Key name of the Key Pair to use for the instance; which can be managed using the `aws_key_pair` resource"
  type        = string
  default     = "key-par-my-app"
}

variable "security_group" {
  description = "Security group to associate"
  type        = string
  default     = "sg-0186f21a8eb2c28d5"
}

variable "subnet_id" {
  description = "Subnet ID to launch in"
  type        = string
  default     = "subnet-0cedc27cd1a3ee2c2"
}

variable "vpc_security_group_ids" {
  description = "A list of security group IDs to associate"
  type        = list(string)
  default     = ["sg-0186f21a8eb2c28d5"]
}

variable "associate_public_ip_address" {
  description = "Associate a public ip address with an instance in a VPC"
  type        = bool
  default     = true
}

variable "availability_zone" {
  description = "The AZ to start the instance in"
  type        = string
  default     = "us-east-1a"
}