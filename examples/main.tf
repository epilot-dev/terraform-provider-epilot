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
  token = "eyJraWQiOiJ2ZFR0MGQrK1RMc2FQZ2tsQ3AzMDVGbEMxc1lOUCtUOXpsaElzMkJ3WERrPSIsImFsZyI6IlJTMjU2In0.eyJzdWIiOiIxNzEyMTkwMy1kM2JlLTRhZTktODZiZS04YjhkZDRmYzY0ZTYiLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiaXNzIjoiaHR0cHM6XC9cL2NvZ25pdG8taWRwLmV1LWNlbnRyYWwtMS5hbWF6b25hd3MuY29tXC9ldS1jZW50cmFsLTFfaGh6MnVJQ2xIIiwicGhvbmVfbnVtYmVyX3ZlcmlmaWVkIjp0cnVlLCJjdXN0b206aXZ5X29yZ19pZCI6IjY2IiwiY29nbml0bzp1c2VybmFtZSI6Im4uZ29lbEBlcGlsb3QuY2xvdWQiLCJjdXN0b206aXZ5X3VzZXJfaWQiOiI4MjYwMiIsImF1ZCI6ImdqOXAwanJlaWh0cTAwY3JpNmEwZmUzMDYiLCJldmVudF9pZCI6IjcwY2M4OGMzLTNlNWUtNDQxZC1iMmUxLTYzN2U1Zjk3ODdjNCIsInRva2VuX3VzZSI6ImlkIiwiYXV0aF90aW1lIjoxNjcyODg2MDY4LCJwaG9uZV9udW1iZXIiOiIrOTE5OTcxNTQzNDUzIiwiZXhwIjoxNjczODg5MTU0LCJpYXQiOjE2NzM4ODU1NTQsImVtYWlsIjoibi5nb2VsQGVwaWxvdC5jbG91ZCJ9.nvL_maT_AlFsZm4lJOkey5PvZhqwbPsdYUIWsqxKbbXTgkq4KKE1hQN2qMA39Vf5izRfaRaDaQKOm4Y2QcjHM5ScdOurorj8_aC1l0EYIno3iCBUMyd_zQuIJRCxCG5NwGxzJJyA2ph_161hXbavCCWy37pv4nRhFMPidULS0X-4bnk9YX-OAVmV-e0INM7FDoccRwIZ7uFI_TGYjJBFo3U5M-aYvZuy0cnkXkk4mH0XEPLSPGipdYUK-bwzcGdLNd-hNTbkcR34Ry3J84x8MgFKgYNKzyVZl7TRJD8Ywd54e0Hf-66m6ppesdzHW9p5kdD19U-SBWhYpMZfjpyvpw"
}

data "epilot_current_user" "user" {

}

output "user_email" {
  value = data.epilot_current_user.user.email
}

resource "epilot_automation" "sample_journey_automation" {

  flow_name = "Created By Terraform Provider"
  triggers = [{
    type       = "journey_submission"
    configuration = {
     source_id = "5f9f9b0b-1b1a-4b1a-9c1a-5b1a9b1a9b1a"
    }
  }]
  
  actions = [{
    type = "send-email"
    config = {
      email_template_id = "5f9f9b0b-1b1a-4b1a-9c1a-5b1a9b1a9b1a"
    }
  }]
}
 