This is a simple CLI tool that converts data between various bases I use often.

## Install

    go install github.com/joeshaw/basen@latest

## Usage

    basen <from base> <to base> <value>

For example:

    $ basen b62 b36 2OtJsTXODC2OhUQ0M7Wk0b
    4o649lo8xm7utcy1cmahxxtvx

You can create custom bases as well by specifying an alphabet:

    $ basen hex custom:0123456789 ff
    255

## Supported bases

- base36, b36, 36
- base62, b62, 62
- base64, b64, 64
- hex
- raw
- custom:<alphabet>

## License

MIT
