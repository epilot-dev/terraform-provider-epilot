package main

import (
	"context"
	"terraform-provider-epilot/epilot"

	"github.com/hashicorp/terraform-plugin-framework/providerserver"
)

// Open Api client generation
//go:generate oapi-codegen --package=main -generate=types,client -o ./automation.gen.go http://localhost:3001/openapi.json

// Provider documentation generation.
//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate --provider-name terraform-provider-epilot

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
