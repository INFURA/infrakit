resource "aws_route53_zone" "din_dev" {
  count = var.environment == "prod" ? 1 : 0

  name = "din.dev"

  tags = {
    Terraform   = "true"
    Environment = var.environment
  }
}

output "din_dev_nameservers" {
  value = aws_route53_zone.din_dev.*.name_servers
}
