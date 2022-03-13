---
page_title: "quortex_processing Resource - terraform-provider-quortex"
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

resource "quortex_processing" "my_proc_hd" {
  pool_id    = quortex_pool.my_pool.id
  name       = "hd"
  published  = true
  identifier = "hd"

  video {
    codec     = "h264"
    bitrate   = 7800000
    framerate = "25"
    resolution {
      width  = 1920
      height = 1080
    }
    advanced {
    }
  }


  audio {
    codec      = "aac-lc"
    bitrate    = 96000
    channels   = "2.0"
    samplerate = "48000"
    track      = "eng"
    output     = "eng"
  }

  subtitle {
    track = "eng"
  }

}
```

## Argument reference

The following arguments are required:

- `pool_id` - (Required) The id of the pool on which the target will be created.
- `name` - (Required) This is the name of the processing.

The following arguments are optional:

- `identifier` - (Optional) This is the identifier of the processing configuration.
- `published` - (Optional) Specify if the processing configuration must be published.
- `video` - (Optional) Specify a video media for the processing configuration. Structure is [documented below](#nested_video).
- `audio` - (Optional) Specify an audio media for the processing configuration. Structure is [documented below](#nested_audio).
- `subtitle` - (Optional) Specify a subtitle media for the processing configuration. Structure is [documented below](#nested_subtitle).

<a name="nested_video"></a>The `video` block supports:

- `codec` - (Required) Specify the video codec for the media. Valid values are `h264`.
- `bitrate` - (Required) Specify the bitrate of the video media, in bps.
- `framerate` - (Required) Specify the framerate of the video media. Valid values are `12.5`, `24`, `25`, `50` and `60`.
- `resolution` - (Required) Specify the resolution of the video media. Structure is [documented below](#nested_resolution).
- `advanced` - (Optional) Advanced configuration of the video media. Structure is [documented below](#nested_advanced).

<a name="nested_resolution"></a>The `resolution` block supports:

- `width` - (Required) Width.
- `height` - (Required) Height.

<a name="nested_advanced"></a>The `advanced` block supports:

- `profile` - (Optional) Specify the encoder profile as defined by MPEG/ITU standards. Valid values are `baseline`, `main` and `high`.
- `level` - (Optional) Specify the encoder level as defined by MPEG/ITU standards. Valid values are `1`, `1.0`, `1.1`, `1.2`, `1.3`, `2`, `2.0`, `2.1`, `2.2`, `3`, `3.0`, `3.1`, `3.2`, `4`, `4.0`, `4.1`, `4.2`, `5`, `5.0`, `5.1` and `5.2`.
- `bframe` - (Optional) Specify whether B frames can reference each other.
- `bframe_number` - (Optional) Specify the maximum number of concecutive B frames (exclusive with bframe set to `false`).
- `encoding_mode` - (Optional) Specify the rate control mode, constant bitrate or variable bitrate capped to bitrate configured.
- `key_frame_interval` - (Optional) Specify the maximal distance between I frames (in ms).
- `maxrate` - (Optional) Specify a minimum tolerance for the output bitrate to be used (in bits per second). Default sets it to 10% over the bitrate.
- `quality` - (Optional) Specify the quality mode of the video encoder. Valid values are `standard`.

<a name="nested_audio"></a>The `audio` block supports:

- `track` - (Required) Select input audio track based on this language code. Supports wildcard (), eg: fr matches fre and fra.
- `output` - (Required) Language code written in the output stream.
- `codec` - (Required) Specify the audio codec for the media. Valid values are `aac-lc` and `he-aac`.
- `channels` - (Required) Specify the number of channels of the audio media. Valid values are `1.0` and `2.0`.
- `bitrate` - (Required) Specify the bitrate of the audio media, in bps. Allowed values depend on the codec and channel configuration. aac_lc stereo: `64000`, `80000`, `96000`, `112000`, `128000`, `144000`, `160000`, `176000`, `192000`, `208000`, `224000`, `240000`, `256000` aac_lc mono: `64000`, `80000`, `96000`, `112000`, `128000` he_aac stereo: `32000`, `48000`, `64000`, `80000`, `96000` he_aac mono: `32000`, `48000`
- `samplerate` - (Required) Specify the samplerate of the audio media. Valid values are `48000` and `44100`.
- `ad` - (Optional) Set 'audio description' flag for this audio track.

<a name="nested_subtitle"></a>The `subtitle` block supports:

- `track` - (Required) Specify the component tracking. This is the mapping of the subtitle processing on the subtitle component.
- `hoh` - (Optional) Set 'Subtitles for the Deaf and Hard of hearing' flag in this subtitle track.

## Attributes reference

- `id` - This is the universal unique identifier of the ressource.
