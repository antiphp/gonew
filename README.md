# Foobar

## Getting started with `gonew`

Start by installing `gonew` using go install:

```
$ go install golang.org/x/tools/cmd/gonew@latest
```

To copy an existing template, run `gonew` in your new project's parent directory with two arguments:

- first, the path to the template you wish to copy, and
- second, the module name of the project you are creating.

For example:

```
$ gonew github.com/antiphp/gonew github.com/antiphp/foobar
$ cd ./foobar
```

## Getting started with `foobar`

```
$ go run ./cmd/foobar --help
```

```
NAME:
   Foobar - A new cli application

USAGE:
   Foobar [global options] command [command options] [arguments...]

VERSION:
   <unknown> @ 1970-01-01T01:00:00+01:00

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h                           show help
   --log.ctx value [ --log.ctx value ]  A list of context field appended to every log. Format: key=value. [$LOG_CTX]
   --log.format value                   Specify the format of logs. Supported formats: 'logfmt', 'json', 'console' [$LOG_FORMAT]
   --log.level value                    Specify the log level. e.g. 'debug', 'info', 'error'. (default: "info") [$LOG_LEVEL]
   --version, -v                        print the version

   Foobar

   --addr value  HTTP address to listen to (default: ":8080") [$ADDR]
```
