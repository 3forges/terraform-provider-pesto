# The Monarh Module

## ANNEX: Terraform Modules and providers' config

According https://developer.hashicorp.com/terraform/language/modules/develop/providers :

> Provider configurations can be defined only in a root Terraform module.
> Providers can be passed down to descendant modules in two ways: either implicitly through inheritance, or explicitly via the providers argument within a module block.


To declare that a module requires particular versions of a specific provider, use a required_providers block inside a terraform block:

