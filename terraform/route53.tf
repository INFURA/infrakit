module "route53_zone_din_dev" {
  count = var.environment == "prod" ? 1 : 0

  source = "./modules/aws_route53_zone"

  name = "din.dev"
  environment = var.environment
}

# resource "aws_route53_record" "runner_custom_domain_record" {
#   count = var.environment == "prod" ? 1 : 0

#   allow_overwrite = true
#   name            = "din.dev"
#   records         = [
#     data.terraform_remote_state.dev.outputs.proxy_deployment_dns_target[0]
#   ]
#   ttl             = 60
#   type            = "CNAME"
#   zone_id         = module.route53_zone_din_dev[0].zone_id
# }
