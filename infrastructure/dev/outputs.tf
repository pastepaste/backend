output "lambda_api_role_arn" {
  value = "${module.iam.lambda_api_arn}"
}

output "apigw_api_endpoint" {
  value = "${module.apigw.api_endpoint}"
}
