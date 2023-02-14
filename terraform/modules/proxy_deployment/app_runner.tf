resource "aws_apprunner_service" "example" {
  service_name = var.name

  source_configuration {
    image_repository {
      image_configuration {
        port = "8000"
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
