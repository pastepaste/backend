module "lambda_api" {
  source = "role"

  name = "${var.name}_lambda_api"

  principal = "lambda.amazonaws.com"

  policy_arns_count = 1

  policy_arns = [
    "${aws_iam_policy.create_put_cw_logs.arn}",
  ]
}
