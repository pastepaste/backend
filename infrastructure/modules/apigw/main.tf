resource "aws_api_gateway_rest_api" "main" {
  name = "${var.name}"
}

module "api" {
  source = "lambda_resource"

  rest_api_id = "${aws_api_gateway_rest_api.main.id}"
  parent_id   = "${aws_api_gateway_rest_api.main.root_resource_id}"
  path_part   = "{proxy+}"

  methods_count = 1

  methods = [
    {
      http_method  = "ANY"
      function_arn = "${var.api_arn}"
    },
  ]
}

resource "aws_api_gateway_deployment" "main" {
  depends_on = [
    "module.api",
  ]

  rest_api_id = "${aws_api_gateway_rest_api.main.id}"
  stage_name  = "${var.env}"
}
