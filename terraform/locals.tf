locals {
}

data "aws_caller_identity" "current" {}

data "terraform_remote_state" "dev" {
  backend = "remote"

  config = {
    hostname     = "backend.api.env0.com"
    organization = var.env0_org_id
    workspaces = {
      name = "infrakit-main-dev"
    }
  }
}
