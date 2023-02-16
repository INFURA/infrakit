locals {
  tags = {
    Name       = var.name
    Terraform = "true"
    Environment = var.environment
  }
}

data "aws_caller_identity" "current" {}
