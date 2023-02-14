module "proxy_deployment" {
  count = var.environment == "dev" ? 1 : 0

  source = "./modules/proxy_deployment"

  name = "infrakit"
  image = "${module.public_ecr[0].repository_url}:1.5.1"
  environment = var.environment
}
