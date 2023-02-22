variable "environment" {
  description = <<EOF
      Environment variable used to tag resources created by this module.

      **Example values are:**
        - temp
        - dev
        - staging
        - prod
  EOF
  type        = string
}

variable "env0_org_id" {
  description = "The ID of the organization in env0"
  type        = string
}

variable "eth_bootstrap_endpoint" {
  description = "Ethereum bootstrap endpoint."
  type        = string
}
