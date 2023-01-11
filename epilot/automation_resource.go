package epilot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &automationResource{}
	_ resource.ResourceWithConfigure = &automationResource{}
)

// automationResourceModel maps the resource schema data.
type automationResourceModel struct {
	FlowName types.String `tfsdk:"flow_name"`
	Triggers []any        `tfsdk:"triggers"`
	Actions  []any        `tfsdk:"actions"`
}

// automationTriggerModel maps order item data.
type automationTriggerModel struct {
	Type          types.String       `tfsdk:"type"`
	Configuration triggerConfigModel `tfsdk:"configuration"`
}

// triggerConfigurationModel maps coffee order item data.
type triggerConfigModel struct {
	SourceID types.String `tfsdk:"source_id"`
}

// automationTriggerModel maps order item data.
type automationActionsModel struct {
	Type   types.String      `tfsdk:"type"`
	Config actionConfigModel `tfsdk:"configuration"`
}

// triggerConfigurationModel maps coffee order item data.
type actionConfigModel struct {
	EmailTemplateModel types.String `tfsdk:"email_template_id"`
}

// NewAutomationResource is a helper function to simplify the provider implementation.
func NewAutomationResource() resource.Resource {
	return &automationResource{}
}

// automationResource is the resource implementation.
type automationResource struct {
	client *epilotCommonContext
}

// Configure adds the provider configured client to the resource.
func (r *automationResource) Configure(_ context.Context, req resource.ConfigureRequest, _ *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	r.client = req.ProviderData.(*epilotCommonContext)
}

// Metadata returns the resource type name.
func (r *automationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_automation"
}

// Schema defines the schema for the resource.
func (r *automationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"flow_name": schema.StringAttribute{
				Required: true,
			},
			"triggers": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Required: true,
						},
						"configuration": schema.SingleNestedAttribute{ // journey submission
							Required: true,
							Attributes: map[string]schema.Attribute{
								"source_id": schema.StringAttribute{
									Required: true,
								},
							},
						},
					},
				},
			},
			"actions": schema.ListNestedAttribute{
				Required: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"type": schema.StringAttribute{
							Required: true,
						},
						"config": schema.SingleNestedAttribute{ // send email
							Required: true,
							Attributes: map[string]schema.Attribute{
								"email_template_id": schema.StringAttribute{
									Required: true,
								},
							},
						},
					},
				},
			},
		},
	}
}

// Create creates the resource and sets the initial Terraform state.
func (r *automationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
}

// Read refreshes the Terraform state with the latest data.
func (r *automationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *automationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *automationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
}
