terraform {
  required_providers {
    slugify = {
      source = "public-cloud-wl/slugify"
      version = "0.1.0"
    }
  }
}
provider "slugify" {
  # example configuration here
}

output "name" {
  value = provider::slugify::slug("Hello, World!") # "hello-world"
}