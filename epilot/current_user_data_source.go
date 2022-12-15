package epilot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

type currentUserDataSourceModel struct {
	Email types.String `tfsdk:"email"`
	ID    types.String `tfsdk:"id"`
}

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource = &currentUserDataSource{}
)

// NewCurrentUserDataSource is a helper function to simplify the provider implementation.
func NewCurrentUserDataSource() datasource.DataSource {
	return &currentUserDataSource{}
}

// currentUserDataSource is the data source implementation.
type currentUserDataSource struct {
	client *epilotCommonContext
}

func (d *currentUserDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*epilotCommonContext)
}

// Metadata returns the data source type name.
func (d *currentUserDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_current_user"
}

// GetSchema defines the schema for the data source.
func (d *currentUserDataSource) GetSchema(_ context.Context) (tfsdk.Schema, diag.Diagnostics) {
	return tfsdk.Schema{
		Attributes: map[string]tfsdk.Attribute{
			"id": {
				Computed: true,
				Type:     types.StringType,
			},
			"email": {
				Type:     types.StringType,
				Computed: true,
			},
		},
	}, nil
}

// Read refreshes the Terraform state with the latest data.
func (d *currentUserDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	getMeResp, getMeErr := d.client.UserClient.GetMeV2WithResponse(ctx)
	if getMeErr != nil {
		panic(getMeErr)
	}

	if getMeResp.StatusCode() != 200 {
		panic(getMeResp.Status())
	}

	user := *getMeResp.JSON200
	state := currentUserDataSourceModel{
		Email: types.StringValue(string(*user.Email)),
		ID:    types.StringValue("placeholder"),
	}

	// Set state
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}