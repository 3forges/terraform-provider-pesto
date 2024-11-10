> Jean-Baptiste Lasselle: This terrform plugin soruce code was created by first git cloning https://github.com/hashicorp/terraform-provider-scaffolding-framework, see [`./README.SCAFFOLDING.md`](./README.SCAFFOLDING.md)

# Terraform Provider Pesto

This epository contains the terraform porvider for the [Pesto API](https://github.com/3forges/pesto-api), containing:

- A resource and a data source (`internal/provider/`),
- Examples (`examples/`) and generated documentation (`docs/`),
- Miscellaneous meta files.

## Requirements

- [Terraform](https://developer.hashicorp.com/terraform/downloads) >= 1.0
- [Go](https://golang.org/doc/install) >= 1.21

## Building The Provider

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command:

```shell
go install
```

## Using the provider

Fill this in for each provider

## Developing the Provider

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

<!--

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

To generate or update documentation, run `go generate`.

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```
-->

We will describe a 3 steps procedure youcan follow, to setup yoru dev envrionment.

You will end up with the source code of this provider inside a folder of your choice, and inside a `./examples/myTest/` subfolder that you will create, you will write terraform code (`*.tf`) on which you will run `tofu apply`, to test using the executble you built from the source code of the provider.

> N.B: In the below proceudre, the terraform code (`*.tf`) on which you will run `tofu apply`, to test using the executble you built from the source code of the provider, may be any folder on your machine, even if it is not a subfolder of the folder where you git cloned this repo.

### Windows

#### Step 1: Prepare the `terraform.rc`

- In git bash for windows, run:

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

- The bin subfolder of the `GOBIN` folder, if `GOBIN` is not empty.
- The bin subfolder of the `GOPATH` folder, if `GOBIN` is empty.

Now we also know that, wen you run `go install .` for a golang project, the built binary is placed in:

- The bin subfolder of the `GOBIN` folder, if `GOBIN` is not empty.
- The bin subfolder of the `GOPATH` folder, if `GOBIN` is empty.

This is why, because of the configuration we set in the `terraformrc`, to make our developed terraform provider executable, available to our terraform recipe, we only need to run, either in powershell or git bash for windows :

```bash
go install .
```

##### Step 3: run a TOFU/terraform recipe using the pesto provider

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

resource "pesto_project" "godzilla_project" {
  name                 = "godzillaRulesDemo"
  description          = "first example project to test creating a project with OpenTOFU the terraformation king. It also has been updated using the OpenTOFU King. A third test updating a project, to test if [stringplanmodifier.UseStateForUnknown()] works."
  git_service_provider = "giteaJapan"
  git_ssh_uri          = "git@github.com:3forges/godzillaRulesDemo.git"
}

data "pesto_projects" "all_pesto_projects" {
  depends_on = [
    pesto_project.godzilla_project
  ]
}
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

You can also run the above commands inside the `./examples/three-apis`

## ANNEX: run acceptance tests and generate the docs

To generate or update documentation, run `go generate`.

> Nota Bene: In the `main.go`, the comment line `//go:generate go run github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs generate -provider-name pesto` determines the exact command used to invoke the docs generation tool, here <https://github.com/hashicorp/terraform-plugin-docs/cmd/tfplugindocs>

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```shell
make testacc
```