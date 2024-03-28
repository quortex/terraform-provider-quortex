resource "quortex_ott_input" "my_input" {
  pool_id    = quortex_ott_pool.my_pool.id
  name       = "ryan"
  published  = true
  identifier = "ryan"

  stream {
    name    = "ryan #1"
    enabled = true
    srt {
      latency = 1000
      listener {
      }
    }
  }

  stream {
    name    = "ryan #2"
    enabled = true
    srt {
      latency = 1000
      listener {
      }
    }
  }
}