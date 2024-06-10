# URL:

sample: https://stable.release.flatcar-linux.net/arm64-usr/current/

variable: ${scheme}://${channel}.release.flatcar-linux.net/${arch}-usr/${release}

Options:
  scheme:
    - http
    - https

  channel:
    - alpha
    - beta
    - stable
    - lts

  arch:
    - amd64
    - arm64

  release:
    - current
    - <VERSION>

cmd: fltctl [FLAGS] [COMMANDS]

FLAGS:
  -s      URL scheme to use, can be HTTP or HTTPS default to HTTPS (To be used on insecure or dev environments, DO NOT USE HTTP ON PRODUCTION ENVIRONMENT).
  -c      Channel.
  -a      Architecture.
  -r      Release.
  -u      Custom upstream root URL, when not provided, default to official flatcar upstream.
  -f      Flatconfig file location, default to (~/.config/fltctl/config.yaml,/etc/fltctl/config.yaml,/opt/fltctl/config.yaml,/usr/local/fltctl/config.yaml) location, searched in that order, used to configure fltctl command using a global config file, cli arguments take precedence over flatconfig directives.
  -i      Files and folders to inject onto the Flatcar base image structure, all files and folders are injected as absolute path from the root (/) directory.

COMMANDS:
  get:    Retrieve flatcar image from upstream or custom server if an alternative url is provided.
  build:  Build your final image from upstream base and different options you provided either from flatconfig file and/or cli flags.
