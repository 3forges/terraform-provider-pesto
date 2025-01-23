# Running a Tofu with the pesto provider

This example shows how to create a pesto content type using the terraform provider.

For the moment, there are still issues, when I create a new Pesto COntent Type I get the below error:

```bash
2025-01-04T23:15:38.616+0100 [DEBUG] provider.terraform-provider-pesto.exe: PESTO API CLIENT GO - CREATE PESTO CONTENT TYPE - here is the API Response Body returned from Pesto API: [] : tf_resource_type=pesto_content_type tf_rpc=ApplyResourceChange @caller=C:/Users/Utilisateur/go/pkg/mod/github.com/3forges/pesto-api-client-go@v0.0.12/pestocontenttypes.go:121 @module=pesto tf_provider_addr=pesto-io.io/terraform/pesto tf_req_id=cc8bc2a7-5ba2-7ace-96e9-b6f773f15868 timestamp="2025-01-04T23:15:38.616+0100"
2025-01-04T23:15:38.616+0100 [DEBUG] provider.terraform-provider-pesto.exe: PESTO API CLIENT GO - CREATE PESTO CONTENT TYPE - Is the API Response Body returned from Pesto API NIL ?: NO API Response Body object is not NIL: @caller=C:/Users/Utilisateur/go/pkg/mod/github.com/3forges/pesto-api-client-go@v0.0.12/pestocontenttypes.go:130 @module=pesto tf_provider_addr=pesto-io.io/terraform/pesto tf_req_id=cc8bc2a7-5ba2-7ace-96e9-b6f773f15868 tf_resource_type=pesto_content_type tf_rpc=ApplyResourceChange timestamp="2025-01-04T23:15:38.616+0100"
2025-01-04T23:15:38.616+0100 [DEBUG] provider.terraform-provider-pesto.exe: CONTENT TYPE RESOURCE - CREATE - here is the tfsdk response object: &{{tftypes.Object["description":tftypes.String, "frontmatter_definition":tftypes.String, "id":tftypes.String, "last_updated":tftypes.String, "name":tftypes.String, "project_id":tftypes.String]<null> {map[description:{<nil> true false false false    [] [] <nil>} frontmatter_definition:{<nil> true false false false    [] [] <nil>} id:{<nil> false false true false    [] [{}] <nil>} last_updated:{<nil> false false true false    [] [] <nil>} name:{<nil> true false false false    [] [] <nil>} project_id:{<nil> true false false false    [] [] <nil>}] map[]    0}} 0xc000412168 []}: tf_rpc=ApplyResourceChange @caller=C:/Users/Utilisateur/terraform-provider-pesto/internal/provider/content_type_resource.go:157 tf_provider_addr=pesto-io.io/terraform/pesto tf_req_id=cc8bc2a7-5ba2-7ace-96e9-b6f773f15868 tf_resource_type=pesto_content_type @module=pesto timestamp="2025-01-04T23:15:38.616+0100"
2025-01-04T23:15:38.616+0100 [DEBUG] provider.terraform-provider-pesto.exe: CONTENT TYPE RESOURCE - CREATE - here is the content type returned from Pesto API: <nil>: @caller=C:/Users/Utilisateur/terraform-provider-pesto/internal/provider/content_type_resource.go:158 @module=pesto tf_provider_addr=pesto-io.io/terraform/pesto tf_rpc=ApplyResourceChange tf_req_id=cc8bc2a7-5ba2-7ace-96e9-b6f773f15868 tf_resource_type=pesto_content_type timestamp="2025-01-04T23:15:38.616+0100"
2025-01-04T23:15:38.616+0100 [WARN]  unexpected data: pesto-io.io/terraform/pesto:stdout="PESTO API CLIENT GO - CREATE PESTO CONTENT TYPE - here is the API Response Body returned from Pesto API: []"
2025-01-04T23:15:38.616+0100 [WARN]  unexpected data: pesto-io.io/terraform/pesto:stdout="PESTO API CLIENT GO - CREATE PESTO CONTENT TYPE - Is the API Response Body returned from Pesto API NIL ?: NO API Response Body object is not NIL"
2025-01-04T23:15:38.616+0100 [DEBUG] provider.terraform-provider-pesto.exe: CONTENT TYPE RESOURCE - CREATE - Is the content type returned from Pesto API NIL ?: YES pesto content type object is NIL!: @caller=C:/Users/Utilisateur/terraform-provider-pesto/internal/provider/content_type_resource.go:167 tf_provider_addr=pesto-io.io/terraform/pesto tf_req_id=cc8bc2a7-5ba2-7ace-96e9-b6f773f15868 tf_rpc=ApplyResourceChange tf_resource_type=pesto_content_type @module=pesto timestamp="2025-01-04T23:15:38.616+0100"
2025-01-04T23:15:38.617+0100 [ERROR] provider.terraform-provider-pesto.exe: Response contains error diagnostic: tf_req_id=cc8bc2a7-5ba2-7ace-96e9-b6f773f15868 tf_proto_version=6.6 tf_resource_type=pesto_content_type diagnostic_detail="Could not create pesto content type, unexpected error: unexpected end of JSON input" diagnostic_severity=ERROR @caller=C:/Users/Utilisateur/go/pkg/mod/github.com/hashicorp/terraform-plugin-go@v0.23.0/tfprotov6/internal/diag/diagnostics.go:58 diagnostic_summary="Error creating pesto content type" tf_provider_addr=pesto-io.io/terraform/pesto tf_rpc=ApplyResourceChange @module=sdk.proto timestamp="2025-01-04T23:15:38.616+0100"
2025-01-04T23:15:38.617+0100 [DEBUG] State storage *statemgr.Filesystem declined to persist a state snapshot
2025-01-04T23:15:38.617+0100 [ERROR] vertex "pesto_content_type.contenttype1_with_tofu" error: Error creating pesto content type
╷
│ Error: Error creating pesto content type
│
│   with pesto_content_type.contenttype1_with_tofu,
│   on main.tf line 23, in resource "pesto_content_type" "contenttype1_with_tofu":
│   23: resource "pesto_content_type" "contenttype1_with_tofu" {
│
│ Could not create pesto content type, unexpected error: unexpected end of
│ JSON input
╵
2025-01-04T23:15:38.852+0100 [DEBUG] provider.stdio: received EOF, stopping recv loop: err="rpc error: code = Unavailable desc = error reading from server: EOF"
2025-01-04T23:15:38.871+0100 [DEBUG] provider: plugin process exited: path=/Users/Utilisateur/go/bin/terraform-provider-pesto.exe pid=7024
2025-01-04T23:15:38.871+0100 [DEBUG] provider: plugin exited

```

The error I get is about an "_unexpected end of JSON input_", and:
* The pesto content type gets successfully created on the Pesto API side, 
* but the created PEsto Content Tye is not properly added in the terraform state, because of that error, which causes afterwards that deleting or updating the Pesto Content Type resource is not possible,
* and well I did found some other terraform issues which seem similar:
  * <https://github.com/hashicorp/terraform-provider-external/issues/23>
  * <https://github.com/hashicorp/terraform-provider-google/issues/7449>


OK I fixed the issue, it's just that the creation endpoint returned HTTP code 204 instead of returning HTTP code 201 with the created object in the API response.

Now all creation, update, and delete work.

## How to use the Pesto Provider

* Create or update all by running:

```Powershell
$env:TF_LOG = "debug"
# export TF_LOG="debug"
tofu init

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

