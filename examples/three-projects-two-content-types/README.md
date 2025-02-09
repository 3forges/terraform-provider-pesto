# Running a Tofu with the pesto provider

This example shows how to create, update, delete, import a pesto content type using the Pesto terraform provider.

It assumes that you use an S3 backend, typically a minio service, with a backend configuration file located at `./.secrets/s3.backend.conf`, which content could be, for example:

```conf
bucket = "pesto-terraform-state" # Name of the S3 bucket
endpoints = {
s3 = "http://minio.pesto.io:9000" # Minio endpoint
}
key = "terraform.tfstate" # Name of the tfstate file

access_key = "<your S3 access Key>"
secret_key = "<your S3 secret Key>"

region                      = "main" # Region validation will be skipped
skip_credentials_validation = true   # Skip AWS related checks and validations
skip_requesting_account_id  = true
skip_metadata_api_check     = true
skip_region_validation      = true
use_path_style              = true # Enable path-style S3 URLs (https://<HOST>/<BUCKET> https://developer.hashicorp.com/terraform/language/settings/backends/s3#use_path_style
```

Before running any command, to initialize your terraform backend, you will typically run:

```bash
tofu init -backend-config=./.secrets/s3.backend.conf
# ---
# Or with the [-reconfigure] option if you need to
# force re-initializing your backend config.
# 
# tofu init -reconfigure -backend-config=./.secrets/s3.backend.conf
```

## How to use the Pesto Provider

### Windows

#### Git bash for windows

* Create or update all by running:

```bash

# tofu init
tofu init -backend-config=./.secrets/s3.backend.conf
# tofu init -reconfigure -backend-config=./.secrets/s3.backend.conf

tofu validate
tofu fmt

tofu plan
tofu apply -auto-approve

# ---
# To test that when I run 
# tofu refresh, the READ method is
# called, I ran:
# rm tflogs.tosee.logs
# tofu refresh >> tflogs.tosee.logs 2>&1

rm tflogs.tosee.logs
tofu validate
tofu fmt

tofu plan
tofu apply -auto-approve >> tflogs.tosee.logs 2>&1
```

* Delete all by running:

```bash
# ---
# Destroy it all:
tofu plan -destroy -out="my.first.destroy.plan.tfplan"; tofu apply "my.first.destroy.plan.tfplan";

```

## Pesto Provider Development environment

In this part, we show the pesto provider contributor, how to run this example, with the freshly built provider.

Note that if you build the provider, you need to setup your environment (among which the `terraformrc`), using the make commands and instrcutions in [the root README.md](../../README.md)

### Windows

#### Git bash for windows

* Create or update all by running:

```bash
# export TF_LOG="debug"
export TF_LOG="trace"

######
# ---
# Dev mode: After every new build of the terraform provider, we
# must update the checksums of the dependency lock file.
tofu providers lock  -fs-mirror="$(go env GOPATH | sed 's#C:##g')\bin"
######
# ---
# Dev mode: For the tofu init 
# command to successfully find the 
# terraform provider executable, even in dev overrides mode:
# becaue I pass to the [-plugin-dir=], the path of the folder
# I configured for the dev_overrides in the "terraformrc".
# 
tofu init -plugin-dir="$(go env GOPATH | sed 's#C:##g')\bin" -reconfigure -backend-config=./.secrets/s3.backend.conf

tofu validate
tofu fmt

tofu plan
tofu apply -auto-approve

# ---
# To test that when I run 
# tofu refresh, the READ method is
# called, I ran:
# rm tflogs.tosee.logs
# tofu refresh >> tflogs.tosee.logs 2>&1

rm tflogs.tosee.logs
tofu validate
tofu fmt

tofu plan
tofu apply -auto-approve >> tflogs.tosee.logs 2>&1
```

* Delete all by running:

```bash
# ---
# Destroy it all:
tofu plan -destroy -out="my.first.destroy.plan.tfplan"; tofu apply "my.first.destroy.plan.tfplan";

```

* Test creating a pesto project and importing it:

```bash
# ---
# Use the Pesto Web UI, a curl request, or the GraphQL Appollo client of the pesto API to create a new project without terraform

```

<!--
#### Powershell

TODO: the commands to use the [-plugin-dir=] tofu init command option, need to be adapted, e.g. how to run a sed command in powershell

* Create or update all by running:

```Powershell
$env:TF_LOG = "debug"
# export TF_LOG="debug"

tofu init

######
# crazy thing i found for, even in dev overrides mode, 
# forcing the tofu init command to successfully find the 
# terraform provider:

# works in git bash for windows

export WHERE_GO_INSTALL_PUTS_EXE="$(go env GOPATH | sed 's#C:#/c#g' | sed 's#\\#/#g')/bin"
export BUILT_PROVIDER_EXE_FILEPATH=${WHERE_GO_INSTALL_PUTS_EXE}/terraform-provider-pesto.exe

mkdir -p ${WHERE_GO_INSTALL_PUTS_EXE}/pesto-io.io/terraform/pesto/0.0.1/windows_386/

mkdir -p ${WHERE_GO_INSTALL_PUTS_EXE}/pesto-io.io/terraform/pesto/0.0.1/windows_amd64/

cp ${BUILT_PROVIDER_EXE_FILEPATH} ${WHERE_GO_INSTALL_PUTS_EXE}/pesto-io.io/terraform/pesto/0.0.1/windows_386/terraform-provider-pesto_v0.0.1_windows_386.exe

cp ${BUILT_PROVIDER_EXE_FILEPATH} ${WHERE_GO_INSTALL_PUTS_EXE}/pesto-io.io/terraform/pesto/0.0.1/windows_amd64/terraform-provider-pesto_v0.0.1_windows_amd64.exe

tofu init -plugin-dir="$(go env GOPATH | sed 's#C:##g')\bin" -reconfigure -backend-config=./.secrets/s3.backend.conf

tofu validate
tofu fmt

tofu plan
tofu apply -auto-approve

# ---
# To test tht when I run 
# tofu refresh, the READ method is
# called, I ran:
# rm tflogs.tosee.logs
# tofu refresh >> tflogs.tosee.logs 2>&1

rm tflogs.tosee.logs
tofu validate
tofu fmt

tofu plan
tofu apply -auto-approve >> tflogs.tosee.logs 2>&1
```

* Delete all by running:

```Powershell

# destroy it all, and recreate it all:

tofu plan -destroy -out="my.first.destroy.plan.tfplan"; tofu apply "my.first.destroy.plan.tfplan";

```

* Delete all and recreate everything from scratch:

```Powershell

# destroy it all, and recreate it all:

tofu plan -destroy -out="my.first.destroy.plan.tfplan"; tofu apply "my.first.destroy.plan.tfplan"; tofu plan -out="my.first.plan.tfplan"; tofu apply -auto-approve "my.first.plan.tfplan"

```

-->

<!--

### GNU/Linux Distributions

* In a bash shell:

-->