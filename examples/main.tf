terraform {
  required_providers {
    quortex = {
      version = "0.0.1"
      source  = "quortex/quortex"
    }
  }
}

provider "quortex" {
  host    = "https://api.dev.saas-dev.quortex.io"
  api_key = "my_api_key"
}
