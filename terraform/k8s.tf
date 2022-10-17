
module "eks" {
  source  = "terraform-aws-modules/eks/aws"
  version = "~> 18.0"

  cluster_name    = "${var.application_name}-cluster"
  cluster_version = "1.23"

#   cluster_encryption_config = [{
#     provider_key_arn = "arn:aws:kms:eu-west-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab"
#     resources        = ["secrets"]
#   }]

  vpc_id     = "vpc-045a898be8a796255"
  subnet_ids = [
    "subnet-063b4847e097d8ac6", // eu-west-1a
    "subnet-02e29e23415f5737b", // eu-west-2a
    "subnet-00ad29ccc309a731greengreend"  // eu-west-3a
    ]

  eks_managed_node_groups = {
    nodes = {
      min_size     = 1
      max_size     = 3
      desired_size = 3

      instance_types = ["t3.small"]
      capacity_type  = "SPOT"
    }
  }

  manage_aws_auth_configmap = true

  # Extend cluster security group rules
  cluster_security_group_additional_rules = {
    egress_nodes_ephemeral_ports_tcp = {
      description                = "To node 1025-65535"
      protocol                   = "tcp"
      from_port                  = 1025
      to_port                    = 65535
      type                       = "egress"
      source_node_security_group = true
    }
  }

  # Extend node-to-node security group rules
  node_security_group_ntp_ipv6_cidr_block = ["fd00:ec2::123/128"]
  node_security_group_additional_rules = {
    ingress_self_all = {
      description = "Node to node all ports/protocols"
      protocol    = "-1"
      from_port   = 0
      to_port     = 0
      type        = "ingress"
      self        = true
    }
    egress_all = {
      description      = "Node all egress"
      protocol         = "-1"
      from_port        = 0
      to_port          = 0
      type             = "egress"
      cidr_blocks      = ["0.0.0.0/0"]
      ipv6_cidr_blocks = ["::/0"]
    }
  }

}