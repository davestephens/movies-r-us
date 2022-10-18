
module "security_group" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = var.application_name
  description = "Complete PostgreSQL example security group"
  vpc_id      = var.vpc_id

  # ingress
  ingress_with_cidr_blocks = [
    {
      from_port   = 5432
      to_port     = 5432
      protocol    = "tcp"
      description = "PostgreSQL access from anywhere"
      cidr_blocks = "0.0.0.0/0"
    },
  ]

}

################################################################################
# RDS Module
################################################################################



module "rds" {
  source  = "terraform-aws-modules/rds/aws"
  version = "5.1.0"
  identifier = "${var.application_name}-postgres"

  engine               = "postgres"
  engine_version       = "14.1"
  family               = "postgres14" # DB parameter group
  major_engine_version = "14"         # DB option group
  instance_class       = "db.t4g.micro"

  allocated_storage     = 20
  max_allocated_storage = 25

  # NOTE: Do NOT use 'user' as the value for 'username' as it throws:
  # "Error creating DB Instance: InvalidParameterValue: MasterUsername
  # user cannot be used as it is a reserved word used by the engine"
  db_name  = replace("${var.application_name}-postgres", "-", "")
  username = replace("${var.application_name}-postgres", "-", "")
  password = "super_secret999"
  port     = 5432

  multi_az               = true
  vpc_security_group_ids = [module.security_group.security_group_id]

  # DB subnet group
  create_db_subnet_group = true
  subnet_ids             = var.subnet_ids

}