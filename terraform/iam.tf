resource "aws_iam_role" "infrastructure" {
  name = "infrastructure"
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
  policy_arn = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryFullAccess"
}

resource "aws_iam_openid_connect_provider" "github" {
  url = "https://token.actions.githubusercontent.com"
  thumbprint_list = ["6938fd4d98bab03faadb97b34396831e3780aea1"]
  client_id_list = ["sts.amazonaws.com"]

  tags = {
    Terraform   = "true"
    Environment = var.environment
  }
}

resource "aws_iam_role" "github" {
  name = "github"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = "sts:AssumeRoleWithWebIdentity"
        Principal = {
          Federated = aws_iam_openid_connect_provider.github.arn
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
      Action = "sts.AssumeRole"
      Resource = aws_iam_role.infrastructure.arn
    }]
  })
}
