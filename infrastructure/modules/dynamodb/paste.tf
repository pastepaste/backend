resource "aws_dynamodb_table" "pastes" {
  name           = "${var.name}_pastes"
  read_capacity  = 5
  write_capacity = 5
  hash_key       = "Slug"

  attribute {
    name = "Slug"
    type = "S"
  }
}
