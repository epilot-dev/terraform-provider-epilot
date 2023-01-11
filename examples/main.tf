variable "token" {
  type = string
}

terraform {
  required_providers {
    epilot = {
      source  = "epilot-dev/epilot"
    }
  }
}

provider "epilot" {
  token = var.token
}

data "epilot_current_user" "user" {

}

output "user_email" {
  value = data.epilot_current_user.user.email
}