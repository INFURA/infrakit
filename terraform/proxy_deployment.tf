module "proxy_deployment" {
  count = var.environment == "dev" ? 1 : 0

  source = "./modules/proxy_deployment"

  name = "infrakit"
  image = "${module.public_ecr[0].repository_url}:1.6.0"
  custom_domain = "din.dev"
  environment = var.environment
  eth_bootstrap_endpoint = var.eth_bootstrap_endpoint
}

output "proxy_deployment_dns_target" {
  value = module.proxy_deployment.*.dns_target
}

output "proxy_deployment_certificate_validation_records" {
  value = module.proxy_deployment.*.certificate_validation_records
}
