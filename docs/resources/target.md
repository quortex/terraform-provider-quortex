---
page_title: "quortex_target Resource - terraform-provider-quortex"
subcategory: ""
description: |-
---

# quortex_target

Manage a target resource. For more information see
[the official documentation](https://help.quortex.io/en/)
and [the API reference](https://web.quortex.io/documentation/ott).

## Example Usage - basic configuration

```hcl
resource "quortex_pool" "my_pool" {
  name                = "my_pool"
  streaming_countries = ["FRA"]
  input_region        = "ireland"
  published           = true
}

resource "quortex_target" "my_target_hls" {
  pool_id          = quortex_pool.my_pool.id
  name             = "hls"
  published        = true
  identifier       = "hls"
  type             = "hls"
  playlist_length  = 10
  segment_duration = 6
}
```

## Argument reference

The following arguments are required:

- `pool_id` - (Required) The id of the pool on which the target will be created.
- `name` - (Required) This is the name of the target.
- `type` - (Required) Specify the type of target. Valid values are `hls` and `dash`.
- `segment_duration` - (Required) Specify the duration of the segments, in seconds. Total length (segment duration x playlist length) cannot be higher than 300s or lower than 10s.
- `playlist_length` - (Required) Specify the depth of the playlist. Total length (segment duration x playlist length) cannot be higher than 300s or lower than 10s.

The following arguments are optional:

- `identifier` - (Optional) This is the identifier of the target configuration.
- `published` - (Optional) Specify if the target configuration must be published.
- `container` - (Optional) Specify the media container. Valid values are `ts` and `fmp4`.
- `scte_35` - (Optional) SCTE-35 (ad insertion messages) configuration. Structure is [documented below](#nested_scte_35).

<a name="nested_scte_35"></a>The `scte_35` block supports:

- `enabled` - (Optional) Enable ad insertion messages.
- `filter_type` - (Optional) Configure the behaviour of the event ID filter. We can keep only the specified events (keep mode), or throw away only those events (deny mode).
- `filter_list` - (Optional) The list of event ID to filter. The filter behaviour depends on the filter_type. The expected format is a list of elements which could be an event ID or a range of event IDs (n-m). Decimal and hexadecimal (with 0x prefix) notation accepted.

## Attributes reference

- `id` - This is the universal unique identifier of the ressource.
