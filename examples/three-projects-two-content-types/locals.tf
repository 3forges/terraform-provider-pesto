locals {
  s3_backend_credentials = try(jsondecode(file("${path.module}/.secrets/minio.backend.credentials.json")), {})
  // s3_backend_access_key = try(jsondecode(file("${path.module}/.secrets/minio.backend.credentials.json")).accessKey, {})
  // s3_backend_secret_key = try(jsondecode(file("${path.module}/.secrets/minio.backend.credentials.json")).secretKey, {})
  // default = try(jsondecode(file("./.secrets/minio.backend.credentials.json")).accessKey, {})
  // try(jsondecode(file("./.secrets/minio.backend.credentials.json")), {})
}