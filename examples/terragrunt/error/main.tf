variable "input" {}

output "output" {
  value = "${var.input} ${var.other_input}"
}

resource "local_file" "example" {
  content  = "Hello, world!"
  filename = "${path.module}/example.txt"
}