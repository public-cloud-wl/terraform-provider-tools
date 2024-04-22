terraform {
  required_providers {
    hashicups = {
      source = "worldline/slugify"
      version = "0.1.0"
    }
  }
}
provider "slugify" {
  # example configuration here
}
