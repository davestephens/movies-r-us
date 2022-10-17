
module "security_group" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = var.application_name
  description = "Complete PostgreSQL example security group"
  vpc_id      = "vpc-045a898be8a796255"

  # ingress
  ingress_with_cidr_blocks = [
    {
      from_port   = 5432
      to_port     = 5432
      protocol    = "tcp"
      description = "PostgreSQL access from within VPC"
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
  port     = 5432

  multi_az               = true
  vpc_security_group_ids = [module.security_group.security_group_id]

  # DB subnet group
  create_db_subnet_group = true
  subnet_ids             = [
    "subnet-063b4847e097d8ac6", // eu-west-1a
    "subnet-02e29e23415f5737b", // eu-west-2a
    "subnet-00ad29ccc309a731d"  // eu-west-3a
    ]

}