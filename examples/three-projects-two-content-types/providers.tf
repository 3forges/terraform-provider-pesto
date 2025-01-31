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
