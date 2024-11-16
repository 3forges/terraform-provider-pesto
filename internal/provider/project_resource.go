package provider

import (
	"context"
	"fmt"
	"time"

	// "strconv".

	"github.com/3forges/pesto-api-client-go"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &projectResource{}
	_ resource.ResourceWithConfigure = &projectResource{}
)

// NewProjectResource is a helper function to simplify the provider implementation.
func NewProjectResource() resource.Resource {
	return &projectResource{}
}

// projectResource is the resource implementation.
type projectResource struct {
	client *pesto.Client
}

// Metadata returns the resource type name.
func (r *projectResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_project"
}

// Schema defines the schema for the resource.
func (r *projectResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
				Computed: false,
			},
			"description": schema.StringAttribute{
				Required: true,
				Computed: false,
			},
			"git_ssh_uri": schema.StringAttribute{
				Required: true,
				Computed: false,
			},
			"git_service_provider": schema.StringAttribute{
				Required: true,
				Computed: false,
			},

			///////////////////////////////////////////
			///////////////////////////////////////////
			///////////////////////////////////////////

			/*
				            "pestoContentTypes": schema.ListNestedAttribute{
				                Required: false,
				                NestedObject: schema.NestedAttributeObject{
				                    Attributes: map[string]schema.Attribute{
										"id": schema.StringAttribute{
											Computed: true,
										},
										"project_id": schema.StringAttribute{
											Required: true,
								            Computed: true,
										},
										"name": schema.StringAttribute{
											Required: true,
								            Computed: true,
										},
										"frontmatter_definition": schema.StringAttribute{
											Required: true,
								            Computed: true,
										},
										"description": schema.StringAttribute{
											Required: false,
								            Computed: true,
										},
				                    },
				                },
				            },

			*/

			///////////////////////////////////////////
			///////////////////////////////////////////
			///////////////////////////////////////////

		},
	}
}

// projectResourceModel maps the resource schema data.
type projectResourceModel struct {
	ID                   types.String `tfsdk:"id"`
	Name                 types.String `tfsdk:"name"`
	Description          types.String `tfsdk:"description"`
	Git_ssh_uri          types.String `tfsdk:"git_ssh_uri"`
	Git_service_provider types.String `tfsdk:"git_service_provider"`
	// PestoContentTypes []pestoContentTypeModel `tfsdk:"pesto_content_types"`
	LastUpdated types.String `tfsdk:"last_updated"`
}

/*
type pestoContentTypeModel struct {
	ID                     types.String `tfsdk:"id"`
	Project_id             types.String `tfsdk:"project_id"`
	Name                   types.String `tfsdk:"name"`
	Frontmatter_definition types.String `tfsdk:"frontmatter_definition"`
	Description            types.String `tfsdk:"description"`
	LastUpdated            types.String `tfsdk:"last_updated"`
}
*/
// Create creates the resource and sets the initial Terraform state.
func (r *projectResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan projectResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	apiRequestBody := pesto.CreatePestoProjectPayload{
		// ID:                   plan.ID.ValueString(),
		Name:                 plan.Name.ValueString(),
		Description:          plan.Description.ValueString(),
		Git_ssh_uri:          plan.Git_ssh_uri.ValueString(),
		Git_service_provider: plan.Git_service_provider.ValueString(),
	}
	// var projectsToCreate []pesto.PestoProject
	tflog.Info(ctx, fmt.Sprintf("PROJECT RESOURCE - CREATE - Creating pesto project of name : %v \n", plan.Name))
	//.Printf("CReating pesto project of name : %v \n",projectsToCreate[i].Name)
	// Create new project
	project, err := r.client.CreatePestoProject(ctx, apiRequestBody, nil)
	tflog.Debug(ctx, fmt.Sprintf("PROJECT RESOURCE - CREATE - here is the tfsdk response object: %v", resp))
	tflog.Debug(ctx, fmt.Sprintf("PROJECT RESOURCE - CREATE - here is the project returned from Pesto API: %v", project))

	var isProjectNil string

	if project != nil {
		isProjectNil = "NO pesto project object is not NIL"
	} else {
		isProjectNil = "YES pesto project object is NIL!"
	}
	tflog.Debug(ctx, fmt.Sprintf("PROJECT RESOURCE - CREATE - Is the project returned from Pesto API NIL ?: %v", isProjectNil))

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating pesto project",
			"Could not create pesto project, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, fmt.Sprintf("PROJECT RESOURCE - CREATE - Successfully created pesto project of name : %v \n", project.Name))
	plan.ID = types.StringValue(plan.ID.ValueString())
	plan = projectResourceModel{
		ID:                   types.StringValue(project.ID),
		Name:                 types.StringValue(project.Name),
		Description:          types.StringValue(project.Description),
		Git_ssh_uri:          types.StringValue(project.Git_ssh_uri),
		Git_service_provider: types.StringValue(project.Git_service_provider),
	}
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *projectResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Read resource information.
	// func (r *orderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state projectResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed project value from HashiCups
	project, err := r.client.GetPestoProject(state.ID.ValueString())
	tflog.Debug(ctx, fmt.Sprintf("PROJECT RESOURCE - READ - here is the tfsdk response object: %v", resp))
	tflog.Debug(ctx, fmt.Sprintf("PROJECT RESOURCE - CREATE - here is the project returned by GetPestoProject from Pesto API: %v", project))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Pesto Pesto Project",
			"Could not read Pesto Pesto Project ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state

	state = projectResourceModel{
		ID:                   types.StringValue(project.ID),
		Name:                 types.StringValue(project.Name),
		Description:          types.StringValue(project.Description),
		Git_ssh_uri:          types.StringValue(project.Git_ssh_uri),
		Git_service_provider: types.StringValue(project.Git_service_provider),
	}
	// - A read operation does not modify the state, so i don't set [state.LastUpdated]
	// state.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set refreshed state
	diags = resp.State.Set(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Update updates the resource and sets the updated Terraform state on success.
func (r *projectResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	//func (r *orderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values to update from plan (not state)
	var plan projectResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// retrieve project ID from state (not plan):
	// Not useful anymore, instead I use
	// [stringplanmodifier.UseStateForUnknown()] in
	// Schema! Neat!
	/*
		var state projectResourceModel
		stateDiags := req.State.Get(ctx, &state)
		resp.Diagnostics.Append(stateDiags...)
	*/

	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan

	apiRequestBody := pesto.UpdatePestoProjectPayload{
		ID: plan.ID.ValueString(),
		// ID:                   state.ID.ValueString(),
		Name:                 plan.Name.ValueString(),
		Description:          plan.Description.ValueString(),
		Git_ssh_uri:          plan.Git_ssh_uri.ValueString(),
		Git_service_provider: plan.Git_service_provider.ValueString(),
	}

	tflog.Info(ctx, fmt.Sprintf("PROJECT RESOURCE - UPDATE - Updating pesto project of (plan.)ID : %v \n", plan.ID.ValueString()))
	// tflog.Info(ctx, fmt.Sprintf("PROJECT RESOURCE - UPDATE - Updating pesto project of (state.)ID : %v \n", state.ID.ValueString()))
	tflog.Info(ctx, fmt.Sprintf("PROJECT RESOURCE - UPDATE - Updating pesto project of (plan.)name : %v \n", plan.Name))
	tflog.Info(ctx, fmt.Sprintf("PROJECT RESOURCE - UPDATE - Updating pesto project of (apiRequestBody.)ID : %v \n", apiRequestBody.ID))
	tflog.Info(ctx, fmt.Sprintf("PROJECT RESOURCE - UPDATE - Updating pesto project of (apiRequestBody.)Name : %v \n", apiRequestBody.Name))
	tflog.Info(ctx, fmt.Sprintf("PROJECT RESOURCE - UPDATE - Updating pesto project of (apiRequestBody.)Git_ssh_uri : %v \n", apiRequestBody.Git_ssh_uri))
	tflog.Info(ctx, fmt.Sprintf("PROJECT RESOURCE - UPDATE - Updating pesto project of (apiRequestBody.)Description : %v \n", apiRequestBody.Description))

	tflog.Info(ctx, fmt.Sprintf("PROJECT RESOURCE - UPDATE - Updating pesto project with payload : %v \n", apiRequestBody))
	// Update existing Pesto Project
	project, err := r.client.UpdatePestoProject(ctx, apiRequestBody, nil)

	tflog.Debug(ctx, fmt.Sprintf("PROJECT RESOURCE - UPDATE - here is the tfsdk response object: %v", resp))
	tflog.Debug(ctx, fmt.Sprintf("PROJECT RESOURCE - UPDATE - here is the project returned from Pesto API: %v", project))

	var isUpdatedProjectNil string

	if project != nil {
		isUpdatedProjectNil = "NO updated pesto project object is not NIL"
	} else {
		isUpdatedProjectNil = "YES updated pesto project object is NIL!"
	}
	tflog.Debug(ctx, fmt.Sprintf("PROJECT RESOURCE - UPDATE - Is the updated project returned from Pesto API NIL ?: %v", isUpdatedProjectNil))

	// _, err := r.client.UpdatePestoProject(plan.ID.ValueString(), hashicupsItems)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Pesto Project",
			"Could not update order, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, fmt.Sprintf("PROJECT RESOURCE - UPDATE - Successfully updated pesto project of name : %v \n", project.Name))

	// Update resource state with the Pesto Project returned by the API (mybe that one is not necessary ? I'm not sure, yet)
	plan.ID = types.StringValue(plan.ID.ValueString())
	plan = projectResourceModel{
		ID:                   types.StringValue(project.ID),
		Name:                 types.StringValue(project.Name),
		Description:          types.StringValue(project.Description),
		Git_ssh_uri:          types.StringValue(project.Git_ssh_uri),
		Git_service_provider: types.StringValue(project.Git_service_provider),
	}
	// Update resource state with updated sub-object if
	// there are some, because sub-objects are not populated.
	// -----------------------
	// /!\ /!\ /!\ SO HERE: THIS WOULD MEAN THAT
	// PESTO API WILL NEED AN ENDPOINT TO
	// RETRIEVE ALL CONTENT TYPES FROM THEIR
	// PROJECT ID  (a foreign key in the database of the Pesto API)
	// -----------------------
	// - this would also mean that we expect
	//   that a pesto project has a [content_types] List property,
	//   that allow creating and updating the content types of a given Pesto Project
	// >> I have reason to dislike that design
	// >> Why ?
	// >> Because I want that a single content-type entity can be reused in several pesto project,
	// >> so I would have to modify the Pesto API so that a the Content Type Entity has a [project_ids] property, that is an array (or a set) of project IDs.
	// >> that would mean having an (N <-> N) relation between Pesto Projects and Pesto Content Types.
	// >> ---
	// >> Yet, having an (N <-> N) relation between
	// >> Pesto Projects and Pesto Content Types, is
	// >> a bit too high a complexity to my taste, and
	// >> since we are talking about an
	// >> end user feature, I can think of another design
	// >> to bring that same reusability for end user:
	// >>
	// >> Using the App GUI, The user can "import" a
	// >> content type, from one project, to another,
	// >> which is easy to implement only in the GUI
	// >> presentation layer, and on the API side, it
	// >> only requires and API endpoint import(contentType PestoContentType, destinationProject PestoProject)
	// >> ---
	// >> And at the terraform level, the
	// >> "pesto_projects" datasource Read method
	// >> allows to fetch all the content-types of
	// >> a given project, using the API endpoint to
	// >> fetch all PestoContentTypes of
	// >> a given PestoProject's ID
	// >>
	// >>
	// -----------------------
	// >>
	// >> So we won't add this for
	// >> the Pesto Pesto Provider, yet, if
	// >> we did, it would look like this:
	// -----------------------

	/*
		retrievedPestoContentTypes, err := r.client.GetPestoContentTypes(plan.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError(
				"Error Fetching Pesto Content Types",
				"Could not read Pesto Content Types for Pesto Project ID "+plan.ID.ValueString()+": "+err.Error(),
			)
			return
		}
	*/

	// Update resource state with updated Pesto Content Types
	// -
	// In [projectResourceModel] you would
	// have a [PestoContentTypes] Field, lke you seee commented above.
	// -
	// Also a pestoContentTypeModel struct would be defined, like the one you can see in this source file, commented.
	// And also the resource Schema would have a Nested List like the one you can see in this source file, commented.

	/*
		plan.PestoContentTypes = []pestoContentTypeModel{}

		for _, contentTypeItem := range retrievedPestoContentTypes {
			plan.PestoContentTypes = append(plan.PestoContentTypes, pestoContentTypeModel{
				ID:                     types.StringValue(contentTypeItem.ID),
				Project_id:             types.StringValue(contentTypeItem.Name),
				Name:                   types.StringValue(contentTypeItem.Teaser),
				Frontmatter_definition: types.StringValue(contentTypeItem.Description),
				Description:            types.StringValue(contentTypeItem.Price),
			})
		}
	*/
	// And finally update last updated timestamp
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// -
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Delete deletes the resource and removes the Terraform state on success.
func (r *projectResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// func (r *orderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state projectResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing order
	project, err := r.client.DeletePestoProject(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Pesto Project",
			fmt.Sprintf("Could not delete Pesto Project of ID=[%v], name=[%v] unexpected error: %v", state.ID, state.Name, err.Error()),
		)
		return
	}
	tflog.Info(ctx, fmt.Sprintf("PROJECT RESOURCE - DELETE - Successfully deleted pesto project of name : %v \n", project.Name))

}

//////////////////////////////////////////////////////
/// CONFIGURE METHOD IMPLEMENTATION
//////////////////////////////////////////////////////

// Configure adds the provider configured client to the resource.
func (r *projectResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

	r.client = client
}
