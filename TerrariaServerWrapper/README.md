# Terraria Server Wrapper

This is a simple go program designed to provide a save on SIGTERM or SIGINT
for the vanilla and tsock Terraria server. The SIGTERM or SIGINT save enables
the server to be run headless in a docker container; `docker stop` sends a
SIGTERM to the PID 1 on stop, but will kill the container if it continues to
run. The Terraria server does not respond to SIGTERM or SIGINT, but this
wrapper will cause the server to save and exit gracefully. Docker is the
intended use-case for this wrapper, and this, combined with go's excellent
concurrency, is why the wrapper is written in go. Go is statically-linked by
default, so this wrapper can be used in a docker container without requiring
any additional libraries.

The wrapper passes all command line arguments to the server
executable, and can effectively be called as though it were the server
executable. The server can still be controlled via the command line
UI; all stdin to the wrapper is passed through to the server.

## Compiling ##

Compiling this program should be a one-line solution. Simply `cd` to
the source directory and execute:

```
go build -o TerrariaServerWrapper cmd/wrapper.go
```

This will compile the wrapper to `TerrariaServerWrapper` for use in your server.
