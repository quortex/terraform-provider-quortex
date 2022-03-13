resource "quortex_target" "my_target_hls" {
  pool_id          = quortex_pool.my_pool.id
  name             = "hls"
  published        = true
  identifier       = "hls"
  type             = "hls"
  playlist_length  = 10
  segment_duration = 6
}

resource "quortex_target" "my_target_dash" {
  pool_id          = quortex_pool.my_pool.id
  name             = "dash"
  published        = true
  identifier       = "dash"
  type             = "dash"
  playlist_length  = 10
  segment_duration = 6
  scte_35 {
    enabled = true
  }
}