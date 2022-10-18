module "lambda_function" {
  source = "terraform-aws-modules/lambda/aws"

  function_name = "${var.application_name}-receive-s3-notification"
  description   = "${var.application_name} notification receiver"
  handler       = "function.lambda_handler"
  runtime       = "python3.7"
  publish       = true
  source_path   = "../../lambda/function.py"
  environment_variables = {
    MRU_API_URL = var.movies_r_us_notify_api_url
  }
}

resource "aws_iam_role" "iam_for_lambda" {
  name = "iam_for_lambda"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow"
    }
  ]
}
EOF
}

resource "aws_lambda_permission" "allow_bucket" {
  statement_id  = "AllowExecutionFromS3Bucket"
  action        = "lambda:InvokeFunction"
  function_name = module.lambda_function.lambda_function_qualified_arn
  principal     = "s3.amazonaws.com"
  source_arn    = module.s3_bucket.s3_bucket_arn
}

resource "aws_s3_bucket_notification" "bucket_notification" {
  bucket = module.s3_bucket.s3_bucket_id

  lambda_function {
    lambda_function_arn = module.lambda_function.lambda_function_qualified_arn
    events              = ["s3:ObjectCreated:*"]
  }

  depends_on = [aws_lambda_permission.allow_bucket]
}