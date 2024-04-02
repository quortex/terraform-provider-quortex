resource "quortex_ott_input" "my_input" {
  pool_id    = quortex_ott_pool.my_pool.id
  name       = "ryan"
  published  = true
  identifier = "ryan"

  stream {
    name    = "ryan"
    enabled = true
    srt {
      latency = 1000
      listener {
      }
    }
  }
}
