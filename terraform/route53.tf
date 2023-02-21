module "route53_zone_din_dev" {
  count = var.environment == "prod" ? 1 : 0

  source = "./modules/aws_route53_zone"

  name = "din.dev"
  environment = var.environment
}

resource "aws_route53_record" "runner_custom_domain_record" {
  count = var.environment == "prod" ? 1 : 0

  allow_overwrite = true
  name            = "din.dev"
  alias {
    name                   = data.terraform_remote_state.dev.outputs.proxy_deployment_dns_target[0]
    zone_id                = "Z01915732ZBZKC8D32TPT"
    evaluate_target_health = true
  }
  
  type            = "A"
  zone_id         = module.route53_zone_din_dev[0].zone_id
}
resource "aws_route53_record" "runner_custom_domain_certificate_record" {
  count = var.environment == "prod" ? length(data.terraform_remote_state.dev.outputs.proxy_deployment_certificate_validation_records[0]) : 0

  allow_overwrite = true
  name            = data.terraform_remote_state.dev.outputs.proxy_deployment_certificate_validation_records[0].*.name[count.index]
  records         = [
    data.terraform_remote_state.dev.outputs.proxy_deployment_certificate_validation_records[0].*.value[count.index]
  ]
  ttl             = 60
  type            = "CNAME"
  zone_id         = module.route53_zone_din_dev[0].zone_id
}
