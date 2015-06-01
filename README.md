# Stacker

## quick stacktrace tool for figuring out deadlocks and stuff.

## Usage

Just import the library
```go
import (
  _ "github.com/cromega/stacker"
)
```
Now Ctrl+C-ing your application will cause it to dump a full stacktrace and exit.  
Slightly less destructively, a stacktrace can also be obtained by visiting `http://localhost:6000` while your process is running.

## Configuration
Not much.
The http server port can be configured via the `STACKER_PORT` environment variable.
