package epilot

import (
	"context"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces
var (
	_ provider.Provider = &epilotProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New() provider.Provider {
	return &epilotProvider{}
}

type epilotCommonContext struct {
	//Token string
	UserClient       *ClientWithResponses
	AutomationClient *ClientWithResponses
}

// epilotProvider is the provider implementation.
type epilotProvider struct{}

// epilotProviderModel maps provider schema data to a Go type.
type epilotProviderModel struct {
	Token types.String `tfsdk:"token"`
}

// Metadata returns the provider type name.
func (p *epilotProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "epilot"
}

// GetSchema defines the provider-level schema for configuration data.
func (p *epilotProvider) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Description: "Configure epilot",
		Attributes: map[string]tfsdk.Attribute{
			"token": {
				Description: "epilot access token. May also be provided via EPILOT_TOKEN environment variable.",
				Type:        types.StringType,
				Optional:    true,
				Sensitive:   true,
			},
		},
	}, nil
}

// Configure prepares a HashiCups API client for data sources and resources.
func (p *epilotProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	tflog.Info(ctx, "Configuring epilot client")

	// Retrieve provider data from configuration
	var config epilotProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Token.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Unknown epilot token",
			"The provider cannot create the epilot API client as there is an unknown configuration value for the epilot token.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	token := os.Getenv("EPILOT_TOKEN")

	if !config.Token.IsNull() {
		token = config.Token.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if token == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("token"),
			"Missing epilot API token",
			"The provider cannot create the epilot API client as there is an unknown configuration value for the epilot token."+
				"Set the token value in the configuration or use the EPILOT_TOKEN environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	ctx = tflog.SetField(ctx, "epilot_token", token)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "epilot_token")

	// TODO: Create a new epilot client using the configuration values
	bearerTokenProvider, bearerTokenProviderErr := securityprovider.NewSecurityProviderBearerToken(token)
	if bearerTokenProviderErr != nil {
		panic(bearerTokenProviderErr)
	}

	userClient, err := NewClientWithResponses("https://user.sls.epilot.io/", WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		panic(err)
	}

	automationClient, err := NewClientWithResponses("https://automation.sls.epilot.io/", WithRequestEditorFn(bearerTokenProvider.Intercept))
	if err != nil {
		panic(err)
	}

	// client := &epilotCommonContext{
	// 	Token: token,
	// 	UserClient: userClient,
	// }
	client := &epilotCommonContext{UserClient: userClient, AutomationClient: automationClient}

	resp.DataSourceData = client
	resp.ResourceData = client
}

// DataSources defines the data sources implemented in the provider.
func (p *epilotProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewCurrentUserDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *epilotProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewAutomationResource,
	}
}
