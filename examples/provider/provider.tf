terraform {
  required_providers {
    tools = {
      source  = "public-cloud-wl/tools"
      version = "0.2.0"
    }
  }
}
provider "tools" {
  # example configuration here
}

output "name" {
  value = provider::tools::slug("Hello, World!") # "hello-world"
}