resource "quortex_pool" "my_pool" {
  name                = "my_pool"
  streaming_countries = ["FRA"]
  input_region        = "france"
  published           = true
}
