terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 4.0"
    }
  }
}

moved {
  from = aws_apprunner_service.example
  to = aws_apprunner_service.service
}
