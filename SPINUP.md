
# How this Tofu Provider Project was spinned up

## Prepare the source code

* In the `$HOME/terraform-provider-pesto` folder, We have git cloned the code of the provider template https://github.com/hashicorp/terraform-provider-scaffolding-framework, and executed :

```bash
export TF_SRC_CODE_FOLDER="$HOME/terraform-provider-pesto"

git clone https://github.com/hashicorp/terraform-provider-scaffolding-framework ${TF_SRC_CODE_FOLDER}


cd ${TF_SRC_CODE_FOLDER}

git checkout master

rm -fr ./git/

go mod edit -module terraform-provider-pesto

go mod tidy

go get github.com/3forges/pesto-api-client-go@v0.0.11

sed -i "s#github.com/hashicorp/terraform-provider-scaffolding-framework/internal/provider#terraform-provider-pesto/internal/provider#g" ./main.go


```

* Then change the `./internal/provider/provider.go` content so that its content is:

```Golang
package provider

import (
    "context"

    "github.com/hashicorp/terraform-plugin-framework/datasource"
    "github.com/hashicorp/terraform-plugin-framework/provider"
    "github.com/hashicorp/terraform-plugin-framework/provider/schema"
    "github.com/hashicorp/terraform-plugin-framework/resource"
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
    resp.Schema = schema.Schema{}
}

// Configure prepares a HashiCups API client for data sources and resources.
func (p *pestoProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

// DataSources defines the data sources implemented in the provider.
func (p *pestoProvider) DataSources(_ context.Context) []func() datasource.DataSource {
    return nil
}

// Resources defines the resources implemented in the provider.
func (p *pestoProvider) Resources(_ context.Context) []func() resource.Resource {
    return nil
}

```

* Then change the `main.go` content so that its content is:

```Golang
func main() {
    var debug bool

    flag.BoolVar(&debug, "debug", false, "set to true to run the provider with support for debuggers like delve")
    flag.Parse()

    opts := providerserver.ServeOpts{
        // NOTE: This is not a typical Terraform Registry provider address,
        // such as registry.terraform.io/hashicorp/pesto. This specific
        // provider address is used in these tutorials in conjunction with a
        // specific Terraform CLI configuration for manual development testing
        // of this provider.
        Address: "pesto-io.io/terraform/pesto",
        Debug:   debug,
    }

    err := providerserver.Serve(context.Background(), provider.New(version), opts)

    if err != nil {
        log.Fatal(err.Error())
    }
}

```

## Prepare testing the provider

### Windows

#### Step 1: Prepare the `terraform.rc`

* In git bash for windows, run:

```bash
export TF_RC_FILE_PATH=$(echo $APPDATA | sed 's#\\#/#g' | sed 's#C:#/c#g')/terraform.rc
echo "  TF_RC_FILE_PATH=[${TF_RC_FILE_PATH}]"


export MY_GO_BIN_FOLDER=$(go env GOBIN)
export MY_GO_PATH_FOLDER=$(go env GOPATH)

export MY_NONDEFAULT_GO_BIN_FOLDER=$(echo "$MY_GO_BIN_FOLDER/bin" | sed 's#\\#/#g' | sed 's#C:##g')

export MY_DEFAULT_GO_BIN_FOLDER=$(echo "$HOME/go/bin" | sed 's#\\#/#' | sed 's#/c##g')
export MY_DEFAULT_GO_BIN_FOLDER=$(echo "$MY_GO_PATH_FOLDER/bin" | sed 's#\\#/#g' | sed 's#C:##g')

echo "MY_GO_BIN_FOLDER=[${MY_GO_BIN_FOLDER}]"
echo "MY_GO_PATH_FOLDER=[${MY_GO_PATH_FOLDER}]"
echo "MY_GO_PATH_FOLDER/bin=[${MY_GO_PATH_FOLDER}/bin]"
echo "MY_DEFAULT_GO_BIN_FOLDER=[${MY_DEFAULT_GO_BIN_FOLDER}]"

# note:
# > if GOBIN is not empty, then the golang binary will be generated inside the bin subfolder of the 'GOBIN' folder.
# > if GOBIN is empty, then the golang binary will be generated inside the bin subfolder of the 'GOPATH' folder.

if [ "x${MY_GO_BIN_FOLDER}" == "x" ]; then
  echo "GO BIN is empty, so will use default"
  export PATH_FOR_DEV_OVERRIDES="${MY_NONDEFAULT_GO_BIN_FOLDER}"
else
  echo "GO BIN is not empty"
  export PATH_FOR_DEV_OVERRIDES="${MY_DEFAULT_GO_BIN_FOLDER}"
fi;


echo "  PATH_FOR_DEV_OVERRIDES=[${PATH_FOR_DEV_OVERRIDES}]"


cat <<EOF >${TF_RC_FILE_PATH}
provider_installation {

  dev_overrides {
      # "pesto-io.io/terraform/pesto" = "<PATH>"
      "pesto-io.io/terraform/pesto" = "${PATH_FOR_DEV_OVERRIDES}"
      }

  # For all other providers, install them directly from their origin provider
  # registries as normal. If you omit this, Terraform will _only_ use
  # the dev_overrides block, and so no other providers will be available.
  direct {}
}

EOF
```

Above, the `pesto-io.io/terraform/pesto` must match the value set in this block, in the `main.go` of your provider:

```Golang
    opts := providerserver.ServeOpts{
        // NOTE: This is not a typical Terraform Registry provider address,
        // such as registry.terraform.io/hashicorp/hashicups. This specific
        // provider address is used in these tutorials in conjunction with a
        // specific Terraform CLI configuration for manual development testing
        // of this provider.
        Address: "pesto-io.io/terraform/pesto",
        Debug:   debug,
    }
```

#### Step 2: go install your provider

To make our developed terraform provider executable, available to our terraform recipe, it needs to be placed into the folder path set above with the `PATH_FOR_DEV_OVERRIDES` env. var. in the `terraformrc`.

Here we above set the value of `PATH_FOR_DEV_OVERRIDES` to the path of:

* The bin subfolder of the `GOBIN` folder, if `GOBIN` is not empty.
* The bin subfolder of the `GOPATH` folder, if `GOBIN` is empty.

Now we also know that, wen you run `go install .` for a golang project, the built binary is placed in:

* The bin subfolder of the `GOBIN` folder, if `GOBIN` is not empty.
* The bin subfolder of the `GOPATH` folder, if `GOBIN` is empty.

This is why, because of the configuration we set in the `terraformrc`, to make our developed terraform provider executable, available to our terraform recipe, we only need to run, either in powershell or git bash for windows :

```bash
go install .
```

#### Step 3: run a TOFU/terraform recipe using the pesto provider

Now, we are ready to start running TOFU/terraform recipe.

From the root of the folder where you git cloned this repo, Create the `./examples/myTest/` folder, Cd into the `./examples/myTest/` folder, and create a `main.tf` file with the below content:

```Terraform
terraform {
  required_providers {
    pesto = {
      source = "pesto-io.io/terraform/pesto"
    }
  }
}

provider "pesto" {}

data "pesto_coffees" "example" { }
```

And in powershell, in the same folder, run:

```Powershell
$env:TF_LOG = "debug"

tofu init

tofu validate
tofu fmt

tofu plan
tofu apply -auto-approve

```

What tells us that the setup is successful, is that the DEBUG logs will confirm the pesto terraform plugin is running, and an error is thrown telling that the `pesto_coffees` datasource does not exist, which makes sense, as far as of now, we did not define any datasource of name `pesto_coffees` in our provider code:

```bash
2024-08-03T13:09:43.313+0200 [DEBUG] provider: using plugin: version=6
2024-08-03T13:09:43.373+0200 [ERROR] vertex "data.pesto_coffees.example" error: Invalid data source
╷
│ Error: Invalid data source
│
│   on main.tf line 11, in data "pesto_coffees" "example":
│   11: data "pesto_coffees" "example" {}
│
│ The provider pesto-io.io/terraform/pesto does not support data source "pesto_coffees".
╵
2024-08-03T13:09:43.381+0200 [DEBUG] provider.stdio: received EOF, stopping recv loop: err="rpc error: code = Unavailable desc = error reading from server: EOF"
```

But yet, it still proves that the golang excutable built for the terraform provider is indeed found by the terraform test recipe inside the `./examples/myTest/` folder.
