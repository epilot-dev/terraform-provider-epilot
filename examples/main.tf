terraform {
  required_providers {
    epilot = {
      version = "0.0.2"
      source  = "epilot-dev/epilot"
    }
  }
}

provider "epilot" {
  # token = "epilot-token"
}

data "epilot_current_user" "user" {

}

output "user_email" {
  value = data.epilot_current_user.user.email
}