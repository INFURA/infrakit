variable "environment" {
  description = "Environment variable used to tag resources created by this module."
  type        = string
}

variable "name" {
  description = "Name of the Route53 zone."
  type        = string
}
