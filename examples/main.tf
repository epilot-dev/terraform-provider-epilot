terraform {
  required_providers {
    epilot = {
      version = "0.2"
      source  = "epilot/terraform-provider-epilot"
    }
  }
}

provider "epilot" {}

