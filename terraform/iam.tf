data "aws_iam_policy" "default_permissions_boundary" {
  name = "DefaultPermissionBoundary"
}
resource "aws_iam_role" "infrastructure" {
  name = "infrastructure"
  permissions_boundary = data.aws_iam_policy.default_permissions_boundary.arn
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = "sts:AssumeRole"
        Principal = {
          AWS = data.aws_caller_identity.current.account_id
        }
      },
    ]
  })

  tags = {
    Terraform   = "true"
    Environment = var.environment
  }
}

resource "aws_iam_role_policy_attachment" "ecr" {
  role       = aws_iam_role.infrastructure.name
  policy_arn = "arn:aws:iam::aws:policy/AmazonElasticContainerRegistryPublicFullAccess"
}

resource "aws_iam_role_policy_attachment" "apprunner" {
  role       = aws_iam_role.infrastructure.name
  policy_arn = "arn:aws:iam::aws:policy/AWSAppRunnerFullAccess"
}

data "aws_iam_openid_connect_provider" "github" {
  url = "https://token.actions.githubusercontent.com"
}

resource "aws_iam_role" "github" {
  name = "github"
  permissions_boundary = data.aws_iam_policy.default_permissions_boundary.arn
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = "sts:AssumeRoleWithWebIdentity"
        Principal = {
          Federated = data.aws_iam_openid_connect_provider.github.arn
        }
        Condition = {
          StringLike = {
            "token.actions.githubusercontent.com:sub": "repo:INFURA/infrakit:*",
          },
          "ForAllValues:StringLike" = {
            "token.actions.githubusercontent.com:iss": "https://token.actions.githubusercontent.com",
            "token.actions.githubusercontent.com:aud" = "sts.amazonaws.com"
          },
        },
      },
    ]
  })
}

resource "aws_iam_role_policy" "github" {
  role = aws_iam_role.github.name
  name = "github"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect = "Allow"
      Action = "sts:AssumeRole"
      Resource = aws_iam_role.infrastructure.arn
    }]
  })
}
