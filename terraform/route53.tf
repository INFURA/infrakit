module "route53_zone_din_dev" {
  count = var.environment == "prod" ? 1 : 0

  source = "./modules/aws_route53_zone"

  name = "din.dev"
  environment = var.environment
}
