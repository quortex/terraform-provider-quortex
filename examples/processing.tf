

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
  }

  video {
    codec     = "h264"
    bitrate   = 4500000
    framerate = "25"
    resolution {
      width  = 1280
      height = 720
    }
  }

  video {
    codec     = "h264"
    bitrate   = 2000000
    framerate = "25"
    resolution {
      width  = 960
      height = 540
    }
  }

  video {
    codec     = "h264"
    bitrate   = 730000
    framerate = "25"
    resolution {
      width  = 768
      height = 432
    }
  }

  video {
    codec     = "h264"
    bitrate   = 145000
    framerate = "25"
    resolution {
      width  = 416
      height = 234
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
    track      = "eng"
  }

}