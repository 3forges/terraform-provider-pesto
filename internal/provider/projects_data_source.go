package provider

import (
	"context"
	"fmt"

	"github.com/3forges/pesto-api-client-go"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &projectsDataSource{}
	_ datasource.DataSourceWithConfigure = &projectsDataSource{}
)

// NewProjectsDataSource is a helper function to simplify the provider implementation.
func NewProjectsDataSource() datasource.DataSource {
	return &projectsDataSource{}
}

// projectsDataSource is the data source implementation.
type projectsDataSource struct {
	client *pesto.Client
}

// Metadata returns the data source type name.
func (d *projectsDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_projects"
}

// Schema defines the schema for the data source.
func (d *projectsDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"projects": schema.ListNestedAttribute{
				Computed: true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.StringAttribute{
							Computed: true,
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"git_ssh_uri": schema.StringAttribute{
							Computed: true,
						},
						"git_service_provider": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data
func (d *projectsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var state projectsDataSourceModel
	tflog.Info(ctx, "Pesto Terraform Provider - Here is an example log done just before calling [client.GetPestoProjects()] for [pesto_projects] datasource")

	projects, err := d.client.GetPestoProjects()
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Read Pesto Projects",
			err.Error(),
		)
		return
	}

	// Map response body to model
	for _, project := range projects {
		projectState := projectsModel{
			ID:                   types.StringValue(project.ID),
			Name:                 types.StringValue(project.Name),
			Description:          types.StringValue(project.Description),
			Git_ssh_uri:          types.StringValue(project.Git_ssh_uri),
			Git_service_provider: types.StringValue(project.Git_service_provider),
		}

		state.Projects = append(state.Projects, projectState)
	}

	// Set state
	tflog.Info(ctx, "Pesto Terraform Provider - Here is an example log done just before saving state for [pesto_projects] datasource")
	diags := resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Configure adds the provider configured client to the data source.
func (d *projectsDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Add a nil check when handling ProviderData because Terraform
	// sets that data after it calls the ConfigureProvider RPC.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*pesto.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *pesto.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}

//////////////////////////////////////
/// datasource model implementation

// projectsDataSourceModel maps the data source schema data.
type projectsDataSourceModel struct {
	Projects []projectsModel `tfsdk:"projects"`
}

// projectsModel maps projects schema data.
type projectsModel struct {
	ID                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	Git_ssh_uri          types.String `tfsdk:"git_ssh_uri"`
	Git_service_provider types.String `tfsdk:"git_service_provider"`
}
