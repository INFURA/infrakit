terraform {
  required_version = "~> 1.3"

  backend "remote" {}

  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.53"
    }
  }
}

provider "aws" {
  region = "us-east-1"
}
