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
  -f      Flatconfig file location, default to (~/.config/fltctl/config.yaml,/etc/fltctl/config.yaml,/opt/fltctl/config.yaml,/usr/local/fltctl/config.yaml) location, searched in that order, used to configure fltctl command using a global config file, cli arguments take precedence over flatconfig directives that take precedence over default.
  -i      Butane configuration to transpil and inject onto the Flatcar base image structure, all files and folders injected are relative to the butane config path.
  -v      Verify image SBOM and signature. If used with -u flag, needs a -ss flag too in order to check custom signature and SBOM.
  -ss     Custom SBOM and image signature path.
  -l      Control logs verbosity from info to trace, default to info.

COMMANDS:
  get:    Retrieve flatcar image from upstream or custom server if an alternative url is provided.
  build:  Build your final image from upstream base and different options you provided either from flatconfig file and/or cli flags.

DEFAULT COMPORTMENT:

init:
  * Compute the current working directory.
  * Look for a flatconfig file, if not found, use default values.
  * Look for an existing local Butane YAML file, if not found exit with an error message.
    * Validate it is a butane file.
  * Compute the get command.
  * Compute the build command.

get:
  * Retrieve the vanilla image depending on previous configuration onto a temporary directory (mktemp).

build:
  * Use
