# logger

Simple log library written in golang.

## Install

```bash
go get -u "github.com/sh-miyoshi/logger"
```

## Usage

```go
package main

import (
    "github.com/sh-miyoshi/logger"
)

func main() {
  // Show debug message, and out to STDOUT
  logger.Init(true, "")

  logger.Debug("debug message")
  logger.Info("info message")
  logger.Error("error message")
}
```

## Author

Shunsuke Miyoshi
