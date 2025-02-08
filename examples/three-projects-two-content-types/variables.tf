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
/*
variable "s3_backend_access_key" {
  type = any
  default = jsondecode(file("${path.module}/.secrets/minio.backend.credentials.json")).accessKey
  // default = try(jsondecode(file("./.secrets/minio.backend.credentials.json")).accessKey, {})
  // try(jsondecode(file("./.secrets/minio.backend.credentials.json")), {})
}
variable "s3_backend_secret_key" {
  type = any
  default = jsondecode(file("${path.module}/.secrets/minio.backend.credentials.json")).secretKey
  // default = try(jsondecode(file("./.secrets/minio.backend.credentials.json")).secretKey, {})
  // try(jsondecode(file("./.secrets/minio.backend.credentials.json")), {})
}
*/