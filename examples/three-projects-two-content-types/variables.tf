variable "monarch_github_org" {
  description = "The Github org of the Monarch Organization"
  # type        = map(string)
  type    = string
  default = "TheShield"
  # default     = {}
}
variable "common_default_frontmatter_def" {
  description = "The default frontmatter definition for all content types of all Orgs"
  type        = map(string)
  default     = {}
}
