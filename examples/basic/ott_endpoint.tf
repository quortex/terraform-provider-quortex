resource "quortex_ott_ott_endpoint" "my_ott_endpoint" {
  pool_id         = quortex_ott_pool.my_pool.id
  input_uuid      = quortex_ott_input.my_input.id
  processing_uuid = quortex_ott_processing.my_proc_hd.id
  target_uuid     = quortex_ott_target.my_target_dash.id
  enabled         = true
}
