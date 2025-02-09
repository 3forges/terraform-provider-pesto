> Jean-Baptiste Lasselle: This terrform plugin source code was created by first git cloning https://github.com/hashicorp/terraform-provider-scaffolding-framework, see [`./README.SCAFFOLDING.md`](./README.SCAFFOLDING.md)

# Terraform Provider Pesto

This repository contains the terraform provider for the [Pesto API](https://github.com/3forges/pesto-api), containing:

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

_TO upgrade version of the `pesto-api-client-go`_ I once had to run:

```bash
Utilisateur@Utilisateur-PC MINGW64 ~/terraform-provider-pesto (feature/add/circleci/pipeline)
$ export GOPRIVATE=github.com

Utilisateur@Utilisateur-PC MINGW64 ~/terraform-provider-pesto (feature/add/circleci/pipeline)
$ go get github.com/3forges/pesto-api-client-go@v0.0.12
go: downloading github.com/3forges/pesto-api-client-go v0.0.12
go: upgraded github.com/3forges/pesto-api-client-go v0.0.11 => v0.0.12

```


```bash
export GOPRIVATE=github.com
go get github.com/hashicorp/terraform-plugin-framework-validators@v0.16.0
go mod tidy
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

We will describe a 3 steps procedure you can follow, to setup your dev envrionment.

You will end up with the source code of this provider inside a folder of your choice, and inside a `./examples/myTest/` subfolder that you will create, you will write terraform code (`*.tf`) on which you will run `tofu apply`, to test using the executable you built from the source code of the provider.

> N.B: In the below proceudre, the terraform code (`*.tf`) on which you will run `tofu apply`, to test using the executble you built from the source code of the provider, may be any folder on your machine, even if it is not a subfolder of the folder where you git cloned this repo.

### Windows

#### Step 0: Instlal GNU make

> Keep in mind you can easy add `make`, but it doesn't come packaged with all the standard UNIX build toolchain--so you will have to ensure those are installed *and* on your PATH, or you will encounter endless error messages.

- Go to [ezwinports](https://sourceforge.net/projects/ezwinports/files/).
- Download `make-4.1-2-without-guile-w32-bin.zip` (get the version without guile).
- Extract zip.
- Copy the contents to your `Git\mingw64\` merging the folders, but do NOT overwrite/replace any existing files. 

Credits & thx : <https://gist.github.com/evanwill/0207876c3243bbb6863e65ec5dc3f058#make>

#### Step 1: Prepare the `terraform.rc`

- In git bash for windows, run:

```bash
# ---
# This make command will backup your terraformrc if 
# it existed, and generate a new terraformrc, with
# the desired dev configuration
. ./.make.env.win.sh; make dev.win.terraformrc
```

In the generated `terraformrc` file, the `pesto-io.io/terraform/pesto` value must match the value set in this block, in the `main.go` of your provider:

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

This is why, because of the configuration we set in the `terraformrc`, to make our developed terraform provider executable, available to our terraform recipe, we only need to run in git bash for windows :

```bash
# go install .
. ./.make.env.win.sh; make dev.win
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

## The last issue

I don't have any issue in the pipe for now.
