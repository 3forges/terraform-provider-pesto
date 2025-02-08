/**
 * The monarch module requires the 
 * 'pesto-io.io/terraform/pesto' terraform provider, so
 * the root module calling the monarch module MUST have that provider in its configuration, period.
 */
terraform {
  required_providers {
    pesto = {
      source = "pesto-io.io/terraform/pesto"
    }
  }
}



provider "pesto" {
  // host = "http://api.pesto.io:3000"
  // username = "education"
  // password = "sdsddg"
  host     = "http://api.pesto.io:3000"
  username = "education"
  password = "test123"
}



terraform {
  backend "s3" {
    bucket = "pesto-terraform-state" # Name of the S3 bucket
    endpoints = {
      s3 = "http://minio.pesto.io:9000" # Minio endpoint
    }
    key = "terraform.tfstate" # Name of the tfstate file

    access_key = "UJTblahcQE7Bunc0pnSL" # Access and secret keys
    secret_key = "EvAGVObs7cEUE22FgkmMTqUZnBHijUd8CXWxDwSU"

    region                      = "main" # Region validation will be skipped
    skip_credentials_validation = true   # Skip AWS related checks and validations
    skip_requesting_account_id  = true
    skip_metadata_api_check     = true
    skip_region_validation      = true
    use_path_style              = true # Enable path-style S3 URLs (https://<HOST>/<BUCKET> https://developer.hashicorp.com/terraform/language/settings/backends/s3#use_path_style
  }
}
/*
*/