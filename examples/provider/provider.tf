terraform {
  required_providers {
    tools = {
      source  = "public-cloud-wl/tools"
      version = "0.1.0"
    }
  }
}
provider "tools" {
  # example configuration here
}

output "name" {
  value = provider::tools::slug("Hello, World!") # "hello-world"
}