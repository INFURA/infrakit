locals {
  tags = {
    Name       = var.name
    Terraform = "true"
    Environment = var.environment
  }
}
