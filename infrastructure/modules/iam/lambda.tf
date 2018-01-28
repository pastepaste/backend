module "lambda_api" {
  source = "role"

  name = "${var.name}_lambda_api"

  principal = "lambda.amazonaws.com"

  policy_arns_count = 2

  policy_arns = [
    "${aws_iam_policy.lambda_cloudwatch_logs.arn}",
    "${aws_iam_policy.lambda_dynamodb.arn}",
  ]
}
