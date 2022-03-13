terraform {
  required_providers {
    quortex = {
      version = "0.1"
      source  = "quortex.io/edu/quortex"
    }
  }
}

provider "quortex" {
  host = "https://api.dev.saas-dev.quortex.io"
  api_key = "my_api_key"
}

resource "quortex_pool" "my_pool" {
  name                = "my_pool"
  streaming_countries = ["FRA"]
  input_region        = "france"
  published           = true
}

resource "quortex_input" "my_input" {
  pool_id    = quortex_pool.my_pool.id
  name       = "ryan"
  published  = true
  identifier = "ryan"

  stream {
    name    = "ryan #1"
    enabled = true
    srt {
      latency         = 1000
      connection_type = "listener"
    }
  }
  stream {
    name    = "ryan #2"
    enabled = true
    srt {
      latency         = 1000
      connection_type = "listener"
    }
  }
}