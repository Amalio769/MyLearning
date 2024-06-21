# Debugging in Go
## Popular Debugging Methods
### Print Statement
**Pros**
* Part of `fmt` package
* Very simple to use

**Cons**
* Concurrency is not supported
* Can't handle complex scenarios

### Log Package
**Pros**
* Part of `log` package
* Relatively simple to use
* Supports concurrency
```go
import (
    "log"
    "os"
)

file, err := os.Openfile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
if err != nil {
    log.Fatal(err)
}
log.SetOutput(file)
...
log.Println( ... )
```

**Cons**
* Can't handle complex scenarios

### Delve Utility
**Pros**
* Command Line
* Integrates with IDE
* Full-featured debugging tool
* Supports remote-debug
* Preferred over [GDB](https://go.dev/doc/gdb)

**Cons**
* Lots of features can be overwhelming