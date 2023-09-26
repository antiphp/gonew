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
   --help, -h                                     show help
   --log.ctx value [ --log.ctx value ]            A list of context field appended to every log. Format: key=value. [$LOG_CTX]
   --log.format value                             Specify the format of logs. Supported formats: 'logfmt', 'json', 'console' [$LOG_FORMAT]
   --log.level value                              Specify the log level. e.g. 'debug', 'info', 'error'. (default: "info") [$LOG_LEVEL]
   --stats.dsn value                              The DSN of a stats backend. [$STATS_DSN]
   --stats.interval value                         The frequency at which the stats are reported. (default: 1s) [$STATS_INTERVAL]
   --stats.prefix value                           The prefix of the measurements names. [$STATS_PREFIX]
   --stats.tags value [ --stats.tags value ]      A list of tags appended to every measurement. Format: key=value. [$STATS_TAGS]
   --tracing.endpoint value                       The tracing backend endpoint. [$TRACING_ENDPOINT]
   --tracing.endpoint-insecure                    Determines if the endpoint is insecure. (default: false) [$TRACING_ENDPOINT_INSECURE]
   --tracing.exporter value                       The tracing backend. Supported: 'zipkin', 'otlphttp', 'otlpgrpc. Depreciated: 'jaeger' [$TRACING_EXPORTER]
   --tracing.ratio value                          The ratio between 0 and 1 of sample traces to take. (default: 0.5) [$TRACING_RATIO]
   --tracing.tags value [ --tracing.tags value ]  A list of tags appended to every trace. Format: key=value. [$TRACING_TAGS]
   --version, -v                                  print the version

   Foobar

   --addr value  HTTP address to listen to (default: ":8080") [$ADDR]
```
