package provider

import (
	"context"
	"fmt"
	"regexp"
	"strings"
	"time"

	// "strconv".

	"github.com/3forges/pesto-api-client-go"
	"github.com/hashicorp/terraform-plugin-framework-validators/mapvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ resource.Resource              = &contentTypeResource{}
	_ resource.ResourceWithConfigure = &contentTypeResource{}
)

// NewContentTypeResource is a helper function to simplify the provider implementation.
func NewContentTypeResource() resource.Resource {
	return &contentTypeResource{}
}

// contentTypeResource is the resource implementation.
type contentTypeResource struct {
	client *pesto.Client
}

// Metadata returns the resource type name.
func (r *contentTypeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_content_type"
}

// Schema defines the schema for the resource.
func (r *contentTypeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"project_id": schema.StringAttribute{
				Required: true,
				Computed: false,
			},
			"name": schema.StringAttribute{
				Required: true,
				Computed: false,
				Validators: []validator.String{
					// mapvalidator.SizeAtLeast(1),
					stringvalidator.LengthAtLeast(1),
					stringvalidator.RegexMatches(
						regexp.MustCompile(`^[a-zA-Z]([a-zA-Z0-9\-\_]*)?[a-zA-Z0-9]$`),
						"Must be a valid Typescript interface name: must start with a letter, must not start with a number or a dash, or an underscore, must not contain any special character, must not end with a dash, or an underscore.",
					),
				},
			},
			/*
				"frontmatter_definition": schema.StringAttribute{
					Required: true,
					Computed: false,
				},
			*/
			"frontmatter_definition": schema.MapAttribute{
				Description: "The properties of the definied frontmatter, for this content type. They are provided as a map, where the keys are the properties " +
					"names and the values represent their TypeScript type, which can only be either of: `string`, `integer`, `boolean` " +
					"(`date` is not supported yet, but will soon).",
				Required:    true,
				ElementType: types.StringType,
				Validators: []validator.Map{
					// mapvalidator.SizeAtLeast(1),
					mapvalidator.SizeAtLeast(0),
					mapvalidator.KeysAre(
						stringvalidator.LengthAtLeast(1),
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^[a-zA-Z]([a-zA-Z0-9\-\_]*)?[a-zA-Z0-9]$`),
							"Must be a vlaid Typescript interface field name: must start with a letter, must not start with a number or a dash, or an underscore, must not contain any special character, must not end with a dash, or an underscore.",
						),
					),
					mapvalidator.ValueStringsAre(
						stringvalidator.RegexMatches(
							// regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9\-]*[a-zA-Z0-9])?$`),
							regexp.MustCompile(`^(string|number|boolean)$`),
							"must be either string, number, or boolean",
						),
					),
				},
				// ... potentially other fields ...
			},
			/* Below, an example  of a map attribute, with validators, from the proxmox terraform provider:
			"nodes": schema.MapAttribute{
				Description: "The member nodes for this group. They are provided as a map, where the keys are the node " +
					"names and the values represent their priority: integers for known priorities or `null` for unset " +
					"priorities.",
				Required:    true,
				ElementType: types.Int64Type,
				Validators: []validator.Map{
					mapvalidator.SizeAtLeast(1),
					mapvalidator.KeysAre(
						stringvalidator.RegexMatches(
							regexp.MustCompile(`^[a-zA-Z0-9]([a-zA-Z0-9\-]*[a-zA-Z0-9])?$`),
							"must be a valid Proxmox node name",
						),
					),
					mapvalidator.ValueInt64sAre(int64validator.Between(0, 1000)),
				},
			},
			*/
			"description": schema.StringAttribute{
				Required: true,
				Computed: false,
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
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

// contentTypeResourceModel maps the resource schema data.
type contentTypeResourceModel struct {
	ID         types.String `tfsdk:"id"`
	Project_id types.String `tfsdk:"project_id"`
	Name       types.String `tfsdk:"name"`
	// Frontmatter_definition types.String `tfsdk:"frontmatter_definition"`
	//Frontmatter_definition types.MapType `tfsdk:"frontmatter_definition"`
	Frontmatter_definition types.Map    `tfsdk:"frontmatter_definition"`
	Description            types.String `tfsdk:"description"`
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
func (r *contentTypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	// Retrieve values from plan
	var plan contentTypeResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan
	apiRequestBody := pesto.CreatePestoContentTypePayload{
		// ID:                   plan.ID.ValueString(),
		Project_id:             plan.Project_id.ValueString(),
		Name:                   plan.Name.ValueString(),
		Frontmatter_definition: r.bakeFrontmatterDefFieldsToStrTsInterface(plan.Frontmatter_definition, plan.Name.ValueString()), // plan.Frontmatter_definition.ValueString(),
		Description:            plan.Description.ValueString(),
	}

	// var projectsToCreate []pesto.PestoContentType
	tflog.Info(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - CREATE - Creating pesto content type of name : %v \n", plan.Name))
	//.Printf("CReating pesto content type of name : %v \n",projectsToCreate[i].Name)
	// Create new project
	contentType, err := r.client.CreatePestoContentType(ctx, apiRequestBody, nil)
	tflog.Debug(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - CREATE - here is the tfsdk response object: %v", resp))
	tflog.Debug(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - CREATE - here is the content type returned from Pesto API: %v", contentType))

	var isContentTypeNil string

	if contentType != nil {
		isContentTypeNil = "NO pesto content type object is not NIL"
	} else {
		isContentTypeNil = "YES pesto content type object is NIL!"
	}
	tflog.Debug(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - CREATE - Is the content type returned from Pesto API NIL ?: %v", isContentTypeNil))

	if err != nil {
		resp.Diagnostics.AddError(
			"Error creating pesto content type",
			"Could not create pesto content type, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - CREATE - Successfully created pesto content type of name : %v \n", contentType.Name))
	// plan.ID = types.StringValue(plan.ID.ValueString())
	plan.ID = types.StringValue(contentType.ID)

	/*
		// We don't need to update the
		// plan after the resource has been created
		// we only need to update :
		// -> the plan.ID for the Resource ID
		// -> the plan.LAstUpdated for the Resource LastUpdated filed
		// and that becasue we are going to
		// update the Terraform State with the plan values, after successfully creating the resource

			plan = contentTypeResourceModel{
				ID:                     types.StringValue(contentType.ID),
				Project_id:             types.StringValue(contentType.Project_id),
				Name:                   types.StringValue(contentType.Name),
				Frontmatter_definition: plan.Frontmatter_definition, // r.bakeFrontmatterDefFieldsToStrTsInterface(, plan.Name.ValueString()), //types.StringValue(contentType.Frontmatter_definition),
				Description:            types.StringValue(contentType.Description),
			}
	*/
	plan.LastUpdated = types.StringValue(time.Now().Format(time.RFC850))

	// Set state to fully populated data
	diags = resp.State.Set(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
}

// Read refreshes the Terraform state with the latest data.
func (r *contentTypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Read resource information.
	// func (r *orderResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	// Get current state
	var state contentTypeResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Get refreshed content type value from HashiCups
	contentType, err := r.client.GetPestoContentType(state.ID.ValueString())
	tflog.Debug(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - READ - here is the tfsdk response object: %v", resp))
	tflog.Debug(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - CREATE - here is the content type returned by GetPestoContentType from Pesto API: %v", contentType))
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Reading Pesto Pesto Content Type",
			"Could not read Pesto Pesto Content Type ID "+state.ID.ValueString()+": "+err.Error(),
		)
		return
	}

	// Overwrite items with refreshed state
	// fronmatter_def_tsInterfaceStr := r.bakeFrontmatterDefFieldsToStrTsInterface(contentType.Frontmatter_definition, contentType.Name)

	state = contentTypeResourceModel{
		ID:         types.StringValue(contentType.ID),
		Project_id: types.StringValue(contentType.Project_id),
		Name:       types.StringValue(contentType.Name),
		// That's my next TODO: I need to turn the string frontmatter into a Map
		Frontmatter_definition: r.convertStrToTsInterface(contentType.Frontmatter_definition), // types.StringValue(fronmatter_def_tsInterfaceStr), // plan.Frontmatter_definition.ValueString(), types.StringValue(contentType.Frontmatter_definition),
		Description:            types.StringValue(contentType.Description),
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
func (r *contentTypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	//func (r *orderResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	// Retrieve values to update from plan (not state)
	var plan contentTypeResourceModel
	diags := req.Plan.Get(ctx, &plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// retrieve content type ID from state (not plan):
	// Not useful anymore, instead I use
	// [stringplanmodifier.UseStateForUnknown()] in
	// Schema! Neat!
	/*
		var state contentTypeResourceModel
		stateDiags := req.State.Get(ctx, &state)
		resp.Diagnostics.Append(stateDiags...)
	*/

	if resp.Diagnostics.HasError() {
		return
	}

	// Generate API request body from plan

	apiRequestBody := pesto.UpdatePestoContentTypePayload{
		ID: plan.ID.ValueString(),
		// ID:                   state.ID.ValueString(),
		Project_id:             plan.Project_id.ValueString(),
		Name:                   plan.Name.ValueString(),
		Frontmatter_definition: plan.Frontmatter_definition.ValueString(),
		Description:            plan.Description.ValueString(),
	}

	tflog.Info(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - UPDATE - Updating pesto content type of (plan.)ID : %v \n", plan.ID.ValueString()))
	// tflog.Info(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - UPDATE - Updating pesto content type of (state.)ID : %v \n", state.ID.ValueString()))
	tflog.Info(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - UPDATE - Updating pesto content type of (plan.)name : %v \n", plan.Name))
	tflog.Info(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - UPDATE - Updating pesto content type of (apiRequestBody.)ID : %v \n", apiRequestBody.ID))
	tflog.Info(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - UPDATE - Updating pesto content type of (apiRequestBody.)Name : %v \n", apiRequestBody.Name))
	tflog.Info(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - UPDATE - Updating pesto content type of (apiRequestBody.)Project_id : %v \n", apiRequestBody.Project_id))
	tflog.Info(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - UPDATE - Updating pesto content type of (apiRequestBody.)Frontmatter_definition : %v \n", apiRequestBody.Frontmatter_definition))
	tflog.Info(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - UPDATE - Updating pesto content type of (apiRequestBody.)Description : %v \n", apiRequestBody.Description))

	tflog.Info(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - UPDATE - Updating pesto content type with payload : %v \n", apiRequestBody))
	// Update existing Pesto Content Type
	contentType, err := r.client.UpdatePestoContentType(ctx, apiRequestBody, nil)

	tflog.Debug(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - UPDATE - here is the tfsdk response object: %v", resp))
	tflog.Debug(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - UPDATE - here is the content type returned from Pesto API: %v", contentType))

	var isUpdatedContentTypeNil string

	if contentType != nil {
		isUpdatedContentTypeNil = "NO updated pesto content type object is not NIL"
	} else {
		isUpdatedContentTypeNil = "YES updated pesto content type object is NIL!"
	}
	tflog.Debug(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - UPDATE - Is the updated content type returned from Pesto API NIL ?: %v", isUpdatedContentTypeNil))

	// _, err := r.client.UpdatePestoContentType(plan.ID.ValueString(), hashicupsItems)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Updating Pesto Content Type",
			"Could not update order, unexpected error: "+err.Error(),
		)
		return
	}
	tflog.Info(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - UPDATE - Successfully updated pesto content type of name : %v \n", contentType.Name))

	// Update resource state with the Pesto Content Type returned by the API (mybe that one is not necessary ? I'm not sure, yet)
	plan.ID = types.StringValue(plan.ID.ValueString())
	plan = contentTypeResourceModel{
		ID:                     types.StringValue(contentType.ID),
		Project_id:             types.StringValue(contentType.Project_id),
		Name:                   types.StringValue(contentType.Name),
		Frontmatter_definition: types.StringValue(contentType.Frontmatter_definition),
		Description:            types.StringValue(contentType.Description),
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
	//   that a pesto content type has a [content_types] List property,
	//   that allow creating and updating the content types of a given Pesto Project
	// >> I have reason to dislike that design
	// >> Why ?
	// >> Because I want that a single content-type entity can be reused in several pesto project,
	// >> so I would have to modify the Pesto API so that a the Content Type Entity has a [project_ids] property, that is an array (or a set) of content type IDs.
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
	// >> "pesto_content_types" datasource Read method
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
	// In [contentTypeResourceModel] you would
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
func (r *contentTypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// func (r *orderResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Retrieve values from state
	var state contentTypeResourceModel
	diags := req.State.Get(ctx, &state)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Delete existing Content Type
	contentType, err := r.client.DeletePestoContentType(ctx, state.ID.ValueString(), nil)
	if err != nil {
		resp.Diagnostics.AddError(
			"Error Deleting Pesto Content Type",
			fmt.Sprintf("Could not delete Pesto Content Type of ID=[%v], name=[%v] unexpected error: %v", state.ID, state.Name, err.Error()),
		)
		return
	}
	tflog.Info(ctx, fmt.Sprintf("CONTENT TYPE RESOURCE - DELETE - Successfully deleted pesto content type of name : %v \n", contentType.Name))

}

//////////////////////////////////////////////////////
/// CONFIGURE METHOD IMPLEMENTATION
//////////////////////////////////////////////////////

// Configure adds the provider configured client to the resource.
func (r *contentTypeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

/////////////////////////////////////////////////////
/// UTILITY METHODS (not implemnting any framework component)
/////////////////////////////////////////////////////

// ////////
// / inspired by:
// /  > MapAttribute: https://github.com/bpg/terraform-provider-proxmox/blob/6f657892c0a29d6677ef6d72690dbfb991a67ad1/fwprovider/ha/resource_hagroup.go#L88
// /  > utility function: https://github.com/bpg/terraform-provider-proxmox/blob/6f657892c0a29d6677ef6d72690dbfb991a67ad1/fwprovider/ha/resource_hagroup.go#L325
// /  > Create method: https://github.com/bpg/terraform-provider-proxmox/blob/6f657892c0a29d6677ef6d72690dbfb991a67ad1/fwprovider/ha/resource_hagroup.go#L160
// bakeFrontmatterDefFieldsToStrTsInterface converts the map of frontmatter_definition fields into a string, which is a TypeScript Interface.
func (r *contentTypeResource) bakeFrontmatterDefFieldsToStrTsInterface(frontmatter_definition types.Map, contentTypeName string) string {
	fmFields := frontmatter_definition.Elements()
	fmFieldsArray := make([]string, len(fmFields))
	i := 0

	for name, value := range fmFields {
		if value.IsNull() {
			fmFieldsArray[i] = name
		} else {
			fmFieldsArray[i] = fmt.Sprintf("\n %s : %s ", name, value.(types.String).ValueString())
		}

		i++
	}

	// return strings.Join(fmFieldsArray, ",")
	tsInterfaceFields := strings.Join(fmFieldsArray, ",")
	return fmt.Sprintf(`export interface `+contentTypeName+`_frontmatter_def { 
%s }`, tsInterfaceFields)
}
