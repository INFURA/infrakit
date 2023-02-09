module "ecs" {
  source = "./modules/aws_ecs_cluster"
  name = "infrakit"
  environment = var.environment
}
