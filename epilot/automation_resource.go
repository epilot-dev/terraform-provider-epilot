package epilot

import (
	"context"
	openapi_types "github.com/deepmap/oapi-codegen/pkg/types"
	"github.com/google/uuid"
	automationapi "terraform-provider-epilot/epilot/automation-api"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &automationResource{}
	_ resource.ResourceWithConfigure = &automationResource{}
)

// automationResourceModel maps the resource schema data.
type automationResourceModel struct {
	FlowName string                   `tfsdk:"flow_name"`
	Triggers []automationTriggerModel `tfsdk:"triggers"`
	Actions  []automationActionsModel `tfsdk:"actions"`
}

// automationTriggerModel maps order item data.
type automationTriggerModel struct {
	Type          string             `tfsdk:"type"`
	Configuration triggerConfigModel `tfsdk:"configuration"`
}

// triggerConfigurationModel maps coffee order item data.
type triggerConfigModel struct {
	SourceID string `tfsdk:"source_id"`
}

// automationTriggerModel maps order item data.
type automationActionsModel struct {
	Type   string            `tfsdk:"type"`
	Config actionConfigModel `tfsdk:"config"`
}

// triggerConfigurationModel maps coffee order item data.
type actionConfigModel struct {
	EmailTemplateID string `tfsdk:"email_template_id"`
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
	var plan automationResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	var triggers []automationapi.AnyTrigger
	for _, trigger := range plan.Triggers {
		var anyTrigger automationapi.AnyTrigger

		anyTrigger.FromJourneySubmitTrigger(automationapi.JourneySubmitTrigger{
			Type: automationapi.JourneySubmitTriggerType(trigger.Type),
			Configuration: struct {
				SourceId openapi_types.UUID `json:"source_id"`
			}{
				SourceId: uuid.MustParse(trigger.Configuration.SourceID),
			},
		})

		triggers = append(triggers, anyTrigger)
	}

	var actions []automationapi.AnyActionConfig
	for _, action := range plan.Actions {
		var anyAction automationapi.AnyActionConfig
		var actionType interface{} = action.Type

		emailTemplateId := action.Config.EmailTemplateID

		anyAction.FromSendEmailActionConfig(automationapi.SendEmailActionConfig{
			Type: &actionType,
			Config: &automationapi.SendEmailConfig{
				EmailTemplateId: &emailTemplateId,
			},
		})

		actions = append(actions, anyAction)
	}

	createFlowData := automationapi.AutomationFlow{
		FlowName: plan.FlowName,
		Triggers: triggers,
		Actions:  actions,
	}

	res, err := r.client.AutomationClient.CreateFlowWithResponse(ctx, createFlowData)

	if err != nil {
		panic(err)
	}

	if res.StatusCode() != 201 {
		panic(res)
	}

	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
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
