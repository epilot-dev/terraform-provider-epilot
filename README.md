# epilot Terraform Provider

Manage epilot configuration with terraform

https://registry.terraform.io/providers/epilot-dev/epilot

## Usage

```hcl
terraform {
  required_providers {
    epilot = {
      source = "epilot-dev/epilot"
      version = "0.0.1"
    }
  }
}

provider "epilot" {
  token = "<epilot access token>"
}
```

```sh
terraform init
```
