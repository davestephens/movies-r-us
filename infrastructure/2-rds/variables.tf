variable "application_name" {
    description = "Application name"
    type = string
    default = "movies-r-us"
}

variable "aws_region" {
    description = "Region to build resources"
    type = string
    default = "eu-west-1"
}

variable "vpc_id" {
    description = "VPC id to create the database"
    type = string
}

variable "subnet_ids" {
    description = "Subnet IDs to attach database to"
    type = list(string)
}