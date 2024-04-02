resource "quortex_ott_target" "my_target_hls" {
  pool_id          = quortex_ott_pool.my_pool.id
  name             = "hls"
  published        = true
  identifier       = "hls"
  type             = "hls"
  playlist_length  = 10
  segment_duration = 6
  container        = "ts"
  latency          = "standard"
}

resource "quortex_ott_target" "my_target_dash" {
  pool_id          = quortex_ott_pool.my_pool.id
  name             = "dash"
  published        = true
  identifier       = "dash"
  type             = "dash"
  playlist_length  = 10
  segment_duration = 6
  container        = "fmp4"
  latency          = "low"
}
