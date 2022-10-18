module "s3_bucket" {
  source = "terraform-aws-modules/s3-bucket/aws"

  bucket = "${var.application_name}-movie-dump"
  acl    = "private"

  versioning = {
    enabled = true
  }

}