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

variable "movies_r_us_notify_api_url" {
    description = "URL of the movies-r-us notify API running on k8s"
    type = string
    default = "http://localhost/notify"
}