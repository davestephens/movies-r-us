output "rds_instance_arn" {
  description = "The Amazon Resource Name (ARN) of the instance"
  value       = module.rds.db_instance_arn
}

output "rds_instance_address" {
  description = "RDS instance address"
  value       = module.rds.db_instance_address
}

output "rds_instance_database" {
  description = "RDS instance database name"
  value       = module.rds.db_instance_name
}

output "rds_instance_username" {
  description = "RDS instance username"
  value       = module.rds.db_instance_username
  sensitive   = true
}

output "rds_instance_password" {
  description = "RDS instance password"
  value       = module.rds.db_instance_password
  sensitive   = true
}