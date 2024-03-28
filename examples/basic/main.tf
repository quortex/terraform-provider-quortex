terraform {
  required_providers {
    quortex = {
      version = "0.2.0"
      source  = "quortex/quortex"
    }
  }
}

provider "quortex" {
  host = "https://my_host"
  api_key {
    api_key = "my_api_key"
  }
}
