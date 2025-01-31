terraform {
  required_providers {
    pesto = {
      source = "pesto-io.io/terraform/pesto"
      # https://www.youtube.com/watch?v=M3IXCWdg0_M
      # ---
      # To declare that a module requires particular versions of a specific provider, use a required_providers block inside a terraform block:
      # version = ">= 2.7.0"
    }
  }
}

/*
provider "pesto" {
  // host = "http://api.pesto.io:3000"
  // username = "education"
  // password = "sdsddg"
  host     = "http://api.pesto.io:3000"
  username = "education"
  password = "test123"
}

*/
