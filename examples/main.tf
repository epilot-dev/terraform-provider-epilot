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
  token = "eyJraWQiOiJ2ZFR0MGQrK1RMc2FQZ2tsQ3AzMDVGbEMxc1lOUCtUOXpsaElzMkJ3WERrPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiIxNzEyMTkwMy1kM2JlLTRhZTktODZiZS04YjhkZDRmYzY0ZTYiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLmV1LWNlbnRyYWwtMS5hbWF6b25hd3MuY29tXC9ldS1jZW50cmFsLTFfaGh6MnVJQ2xIIiwicGhvbmVfbnVtYmVyX3ZlcmlmaWVkIjp0cnVlLCJjdXN0b206aXZ5X29yZ19pZCI6IjY2IiwiY29nbml0bzp1c2VybmFtZSI6Im4uZ29lbEBlcGlsb3QuY2xvdWQiLCJjdXN0b206aXZ5X3VzZXJfaWQiOiI4MjYwMiIsImF1ZCI6ImdqOXAwanJlaWh0cTAwY3JpNmEwZmUzMDYiLCJldmVudF9pZCI6IjcwY2M4OGMzLTNlNWUtNDQxZC1iMmUxLTYzN2U1Zjk3ODdjNCIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNjcyODg2MDY4LCJwaG9uZV9udW1iZXIiOiIrOTE5OTcxNTQzNDUzIiwiZXhwIjoxNjc0MDYwNTA2LCJpYXQiOjE2NzQwNTY5MDYsImVtYWlsIjoibi5nb2VsQGVwaWxvdC5jbG91ZCJ9.lFa1bFA2_xWNTQn7iLc1_9xwsOg6WmZVJhRYR8_KjVPGQ1_AfPSyqEoAdCS1SwWIZhGf36X0i1jG7xkDxo9wBAotxXvFnLDk0bkXzR-6Ic4zNvth1AvN5uPx81AQkYr1NgJLdEapTrmy02qUqi42QQ_IbivKEogKcxSCo3pziazKCPxvfykfKopLw62VMkdrIJf7XuQf2DSfYqwef_gwEhYP6odp_bQgQ0d7Ul-T8wKZ7LDbyxJqAh2qKO2ZTQ6CVxwG8oCZ6BIGp9kUh06UbCwYbHxEnKtpiSmAeAGy8OKUxAkqpWVW6JZeAttONw2M_xB6IqdmYzqs54v0DJXX6w"
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
 