---
page_title: "quortex_ott_input Resource - terraform-provider-quortex"
subcategory: ""
description: |-
---

# quortex_ott_input

Manage an input resource. For more information see
[the official documentation](https://help.quortex.io/en/)
and [the API reference](https://web.quortex.io/documentation/ott).

## Example Usage - basic configuration

```hcl
resource "quortex_ott_input" "ryan" {
  pool_id    = quortex_ott_pool.my_pool.id
  name       = "ryan"
  published  = true
  identifier = "ryan"

  stream {
    name    = "ryan #1"
    enabled = true
    srt {
      latency = 1000
      caller {
        address = "srt://111.222.111.222:5000"
      }
    }
  }

  stream {
    name    = "ryan #2"
    enabled = true
    srt {
      latency = 1000
      caller {
        address = "srt://111.222.111.222:6000"
      }
    }
  }

}
```

## Argument reference

The following arguments are required:

- `pool_id` - (Required) The id of the pool on which the input will be created.
- `name` - (Required) This is the name of the input.

The following arguments are optional:

- `identifier` - (Optional) This is the identifier of the input configuration.
- `published` - (Optional) Specify if the processing input must be published.
- `stream` - (Optional) Specify a stream for the input configuration. Structure is [documented below](#nested_stream).

<a name="nested_stream"></a>The `stream` block supports:

- `name` - (Required) This is the name of the stream.
- `enabled` - (Optional) Specify if the stream must be enabled.
- `srt` - (Optional) Specify the str configuration. Structure is [documented below](#nested_srt).
- `rtmp` - (Optional) Specify the rtmp configuration. Structure is [documented below](#nested_rtmp).

<a name="nested_srt"></a>The `srt` block supports:

- `latency` - (Optional) This is the latency of the SRT link, in ms.
- `listener` - (Optional) Specify the listener configuration. Structure is [documented below](#nested_listener).
- `caller` - (Optional) Specify the caller configuration. Structure is [documented below](#nested_caller).
- `overrides` - (Optional) Specify the components overrides configuration. Structure is [documented below](#nested_overrides).

<a name="nested_rtmp"></a>The `rtmp` block supports:

- `overrides` - (Optional) Specify the components overrides configuration. Structure is [documented below](#nested_overrides).

<a name="nested_listener"></a>The `listener` block supports:

- `cidr` - (Optional) Specify the cidr block allowed to access the listener.

<a name="nested_caller"></a>The `caller` block supports:

- `address` - (Required) Input stream address.
- `passphrase` - (Optional) This is the passphrase of the SRT link. Must be between 10 and 79 characters long. Leave empty to disable encryption.

<a name="nested_overrides"></a>The `overrides` block supports:

- `pid` - (Required) Specify pid of component to override. Required for srt, not required for rtmp.
- `type` - (Required) Type of component in the stream.
- `enabled` - (Optional) Specifies whether to discard the corresponding component or not. The default behaviour depends on the component type. SCTE-35 tracks are disabled by default, all other types of tracks are enabled.
- `audio` - (Optional) Used when `type` is `audio`. Structure is [documented below](#nested_overrides_audio).
- `teletext` - (Optional) Used when `type` is `teletext`. Structure is [documented below](#nested_overrides_teletext).

<a name="nested_overrides_audio"></a>The overrides `audio` block supports:

- `language` - (Optional) Overrides the language indicated in the stream.
- `ad` - (Optional) Overrides the 'audio description' flag.

<a name="nested_overrides_teletext"></a>The overrides `teletext` block supports:

- `page` - (Required) Teletext magazine and page identifying a subtitle stream.
- `language` - (Optional) Overrides the language indicated in the stream.
- `sdh` - (Optional) Overrides the 'subtitles for the deaf and hard of hearing' flag.

## Attributes reference

- `id` - This is the universal unique identifier of the ressource.
