terraform {
  required_providers {
    quortex = {
      version = "0.0.1"
      source  = "localhost/quortex/quortex"
    }
  }
}

provider "quortex" {
  host = "https://my_host"
  api_key {
    api_key = "my_api_key"
  }
}
