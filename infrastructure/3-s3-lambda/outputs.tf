output "lambda_arn" {
    value = module.lambda_function.lambda_function_qualified_arn
}

output "s3_arn" {
    value = module.s3_bucket.s3_bucket_arn
}
