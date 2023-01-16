package main

import (
	"context"
	"terraform-provider-epilot/epilot"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Open Api client generation
//go:generate oapi-codegen --package=automation -generate=types,client -o ./epilot/automation/automation.gen.go http://localhost:3001/openapi.json
//go:generate oapi-codegen --package=user -generate=types,client -o ./epilot/user/user.gen.go http://localhost:3002/openapi.json

func main() {
	providerserver.Serve(context.Background(), epilot.New, providerserver.ServeOpts{
		// NOTE: This is not a typical Terraform Registry provider address,
		// such as registry.terraform.io/hashicorp/hashicups. This specific
		// provider address is used in these tutorials in conjunction with a
		// specific Terraform CLI configuration for manual development testing
		// of this provider.
		Address: "registry.terraform.io/epilot-dev/epilot",
	})
}
