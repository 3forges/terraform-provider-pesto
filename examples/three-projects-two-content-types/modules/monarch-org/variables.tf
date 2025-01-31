variable "github_org" {
  description = "The Github org for the URL of all pesto projects version controlled on Github.com"
  # type        = map(string)
  type        = string
  default     = "3forges"
  # default     = {}
}
variable "default_frontmatter_def" {
  description = "The default frontmatter definition for all content types of all pesto projects version controlled in that Github.com org"
  type        = map(string)
  default     = {}
}
