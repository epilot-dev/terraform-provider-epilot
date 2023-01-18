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

resource "epilot_automation" "testing_journey_automation" {

  flow_name = "Nishu's Automation created By Terraform Provider"
  triggers = [{
    type       = "journey_submission"
    configuration = {
     source_id = "0e60a7d0-8c93-11ed-b051-83abad297d53"
    }
  }]
  
  actions = [{
    type = "send-email"
    config = {
      email_template_id = "ea3b4979-306a-4368-b2ae-09456730fa84"
    }
  }]
}
 