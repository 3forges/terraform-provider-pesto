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