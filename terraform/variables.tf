variable "environment" {
  description = <<EOF
      Environment variable used to tag resources created by this module.

      **Example values are:**
        - temp
        - dev
        - staging
        - prod

      **Notes:**
        Put here your notes if there is any.
  EOF
  type        = string
}
