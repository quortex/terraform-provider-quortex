resource "quortex_ott_pool" "my_pool" {
  name                = "my_pool"
  streaming_countries = ["FRA"]
  input_region        = "france"
  published           = true
  processing_type     = "advanced"
}
