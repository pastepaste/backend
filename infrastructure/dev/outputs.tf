output "lambda_api_role_arn" {
  value = "${module.iam.lambda_api_arn}"
}

output "apigw_api_endpoint" {
  value = "${module.apigw.api_endpoint}"
}

output "dynamodb_pastes_name" {
  value = "${module.dynamodb.pastes_name}"
}
