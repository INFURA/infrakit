module "public_ecr" {
  count = var.environment == "dev" ? 1 : 0

  source = "terraform-aws-modules/ecr/aws"

  repository_name = "infrakit"
  repository_type = "public"

  public_repository_catalog_data = {
    description       = "Decentralized Infrastructure Network InfraKit"
    operating_systems = ["Linux"]
    architectures     = ["x86", "ARM"]
  }

  tags = {
    Terraform   = "true"
    Environment = var.environment
  }
}
