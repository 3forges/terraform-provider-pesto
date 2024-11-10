# Running a Tofu with the pesto provider

## How to use the Pesto Provider

* Create or update all by running:

```Powershell
$env:TF_LOG = "debug"

tofu init

tofu validate
tofu fmt

tofu plan
tofu apply -auto-approve

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
