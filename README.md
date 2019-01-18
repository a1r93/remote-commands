# remote-commands

## Description

The remote commands project is a small server with a single route:

```
/commands
```

It can be accessed with a simple GET request with query parameters that should
be prefixed with `cmds`

The possible values for cmds are:

- **utc-time**: returns the current UTC time
- **cpu-usage**: returns the machine current CPU usage
- **available-ram**: returns the machine current available RAM
- **say-something**: makes the machine say a predefined sentence
- **capture**: makes the machine take a screenshot

Example request:

```
http:localhost:8080/commands?cmds=available-ram&cmds=utc-time
```

## Why

The purpose of this little project is to take a look at how Go works by using
goroutines, http, exec, structures, etc.
