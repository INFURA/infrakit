resource "aws_apprunner_service" "service" {
  service_name = var.name

  source_configuration {
    image_repository {
      image_configuration {
        port = "8000"
        start_command = "proxy"
        runtime_environment_variables = {
          "INFRAKIT_PROXY__LISTEN_ADDR": ":8000"
          "INFRAKIT_PROXY__HEALTH_LISTEN_ADDR": ":9000"
        }
      }

      image_identifier = var.image
      image_repository_type = "ECR_PUBLIC"
    }

    auto_deployments_enabled = false
  }

  tags = {
    Terraform = "true"
    Environment = var.environment
  }
}

resource "aws_apprunner_custom_domain_association" "custom" {
  domain_name = var.custom_domain
  service_arn = aws_apprunner_service.service.arn
}

output "certificate_validation_records" {
  value = aws_apprunner_custom_domain_association.custom.certificate_validation_records
}

output "dns_target" {
  value = aws_apprunner_custom_domain_association.custom.dns_target
}
