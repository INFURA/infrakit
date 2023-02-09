module "vpc" {
  source = "./modules/aws_vpc"
  environment = var.environment
}
