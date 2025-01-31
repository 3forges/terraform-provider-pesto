/**
 * https://developer.hashicorp.com/terraform/tutorials/modules/module-create
 **/
module "monarch" {
  source = "./modules/monarch-org/"

  github_org = var.monarch_github_org

  default_frontmatter_def = var.common_default_frontmatter_def
}
