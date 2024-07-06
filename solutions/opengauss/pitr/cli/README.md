# PITR/CLI

PITR cli client for openGauss.

## project description

### layout

* `internal`:For internal code, use golang internal directory to isolate references.
    - `cmd`: All subcommands.
        - `view`: Cli output structure.
        - `backup.go`: `backup` subcommand.
        - `restore.go`: `restore` subcommand.
        - `show.go`: `show` subcommand.
    - `pkg`: Third party dependencies.
        - `model`: Input/Output structure.
        - `agent-server.go`: Agent server API.
        - `sharding-sphere-proxy.go`: Sharding-sphere proxy API.
* `main.go`: Golang main function.

