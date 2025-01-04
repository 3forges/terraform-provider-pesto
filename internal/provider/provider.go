package provider

import (
	"context"
	"os"

	"github.com/3forges/pesto-api-client-go"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-log/tflog"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ provider.Provider = &pestoProvider{}
)

// New is a helper function to simplify provider server and testing implementation.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &pestoProvider{
			version: version,
		}
	}
}

// pestoProvider is the provider implementation.
type pestoProvider struct {
	// version is set to the provider version on release, "dev" when the
	// provider is built and ran locally, and "test" when running acceptance
	// testing.
	version string
}

// Metadata returns the provider type name.
func (p *pestoProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "pesto"
	resp.Version = p.version
}

// Schema defines the provider-level schema for configuration data.
func (p *pestoProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"host": schema.StringAttribute{
				Optional: true,
			},
			"username": schema.StringAttribute{
				Optional: true,
			},
			"password": schema.StringAttribute{
				Optional:  true,
				Sensitive: true,
			},
		},
	}
}

//////////////////////////////////////////////
//////////////////////////////////////////////
//////////////////////////////////////////////
//// BEGIN CONFIGURE PART
//////////////////////////////////////////////
//////////////////////////////////////////////
//////////////////////////////////////////////

// // Configure prepares a Pesto API client for data sources and resources.
// func (p *pestoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
// }
// pestoProviderModel maps provider schema data to a Go type.
type pestoProviderModel struct {
	Host     types.String `tfsdk:"host"`
	Username types.String `tfsdk:"username"`
	Password types.String `tfsdk:"password"`
}

func (p *pestoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	// Retrieve provider data from configuration
	var config pestoProviderModel
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)

	if resp.Diagnostics.HasError() {
		return
	}

	// If practitioner provided a configuration value for any of the
	// attributes, it must be a known value.

	if config.Host.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Unknown HashiCups API Host",
			"The provider cannot create the HashiCups API client as there is an unknown configuration value for the HashiCups API host. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the PESTO_HOST environment variable.",
		)
	}

	if config.Username.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Unknown HashiCups API Username",
			"The provider cannot create the HashiCups API client as there is an unknown configuration value for the HashiCups API username. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the PESTO_USERNAME environment variable.",
		)
	}

	if config.Password.IsUnknown() {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Unknown HashiCups API Password",
			"The provider cannot create the HashiCups API client as there is an unknown configuration value for the HashiCups API password. "+
				"Either target apply the source of the value first, set the value statically in the configuration, or use the PESTO_PASSWORD environment variable.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Default values to environment variables, but override
	// with Terraform configuration value if set.

	host := os.Getenv("PESTO_HOST")
	username := os.Getenv("PESTO_USERNAME")
	password := os.Getenv("PESTO_PASSWORD")

	if !config.Host.IsNull() {
		host = config.Host.ValueString()
	}

	if !config.Username.IsNull() {
		username = config.Username.ValueString()
	}

	if !config.Password.IsNull() {
		password = config.Password.ValueString()
	}

	// If any of the expected configurations are missing, return
	// errors with provider-specific guidance.

	if host == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("host"),
			"Missing Pesto API Host",
			"The provider cannot create the Pesto API client as there is a missing or empty value for the Pesto API host. "+
				"Set the host value in the configuration or use the PESTO_HOST environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if username == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("username"),
			"Missing Pesto API Username",
			"The provider cannot create the Pesto API client as there is a missing or empty value for the Pesto API username. "+
				"Set the username value in the configuration or use the PESTO_USERNAME environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if password == "" {
		resp.Diagnostics.AddAttributeError(
			path.Root("password"),
			"Missing Pesto API Password",
			"The provider cannot create the Pesto API client as there is a missing or empty value for the Pesto API password. "+
				"Set the password value in the configuration or use the PESTO_PASSWORD environment variable. "+
				"If either is already set, ensure the value is not empty.",
		)
	}

	if resp.Diagnostics.HasError() {
		return
	}
	ctx = tflog.SetField(ctx, "pesto_host", host)
	ctx = tflog.SetField(ctx, "pesto_username", username)
	ctx = tflog.SetField(ctx, "pesto_password", password)
	ctx = tflog.MaskFieldValuesWithFieldKeys(ctx, "pesto_password")

	tflog.Debug(ctx, "PROVIDER.GO: Creating Pesto API client")
	// Create a new PestoAPI client using the configuration values
	client, err := pesto.NewClient(&host, &username, &password)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to Create Pesto API Client",
			"An unexpected error occurred when creating the Pesto API client. "+
				"If the error is not clear, please contact the provider developers.\n\n"+
				"Pesto Client Error: "+err.Error(),
		)
		return
	}

	// Make the Pesto client available during DataSource and Resource
	// type Configure methods.
	resp.DataSourceData = client
	resp.ResourceData = client
}

//////////////////////////////////////////////
//////////////////////////////////////////////
//////////////////////////////////////////////
//// END CONFIGURE PART
//////////////////////////////////////////////
//////////////////////////////////////////////
//////////////////////////////////////////////

//////////////////////////////////////////////

// DataSources defines the data sources implemented in the provider.
/*
func (p *pestoProvider) DataSources(_ context.Context) []func() datasource.DataSource {
    return nil
}
*/
// DataSources defines the data sources implemented in the provider.
func (p *pestoProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{
		NewProjectsDataSource,
	}
}

// Resources defines the resources implemented in the provider.
func (p *pestoProvider) Resources(_ context.Context) []func() resource.Resource {
	return []func() resource.Resource{
		NewProjectResource,
		NewContentTypeResource,
	}
}
