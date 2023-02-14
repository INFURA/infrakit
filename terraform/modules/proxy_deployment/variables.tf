variable "environment" {
  description = "Environment variable used to tag resources created by this module."
  type        = string
}

variable "name" {
  description = "Name of the proxy deployment."
  type        = string
}

variable "image" {
  description = "ECR public image identifier."
  type        = string
}
