Go-Logger is a simple, lightweight logging library for Go, built on top of [logrus](https://github.com/sirupsen/logrus). It provides pre-configured loggers for different log levels, making it easy to add structured logging to your Go applications.

## Features

- Pre-configured loggers for Debug, Info, Warning, and Error levels
- Simple API with `Println`, `Printf`, `Fatalln`, and `Fatalf` methods
- Built on top of logrus for powerful and flexible logging
- Easy to use and integrate into existing Go projects
- Thread-safe initialization of loggers

## Installation

To install go-logger, use `go get`:

```bash
go get gitea.heheda.vip/root/go-logger
```

## Usage

Here's a quick example of how to use logger in your Go application:

```go
package main

import (
    "gitea.heheda.vip/root/go-logger"
)

func main() {
    // Using preset loggers
    logger.Debug.Println("This is a debug message")
    logger.Info.Println("This is an info message")
    logger.Warning.Println("This is a warning message")
    logger.Error.Println("This is an error message")

    // Using formatted logging
    logger.Info.Printf("Formatted message: %s", "Hello, World!")

    // Getting logger level
    level := logger.Debug.GetLevel()
    logger.Info.Printf("Debug logger level: %v", level)

    // Note: The following will exit the program
    // logger.Error.Fatalln("This is a fatal error message")
}
```

## Available Loggers

- `logger.Debug`: For debug-level messages
- `logger.Info`: For informational messages
- `logger.Warning`: For warning messages
- `logger.Error`: For error messages

## Methods

Each logger provides the following methods:

- `Println(v ...interface{})`: Log a message at the logger's level
- `Printf(format string, v ...interface{})`: Log a formatted message at the logger's level
- `Fatalln(v ...interface{})`: Log a message at Fatal level and exit the program
- `Fatalf(format string, v ...interface{})`: Log a formatted message at Fatal level and exit the program
- `GetLevel() logrus.Level`: Get the current logger level
- `SetOutput(output *os.File)`: Set the output destination for the logger

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- [Logrus](https://github.com/sirupsen/logrus) - The powerful logging library that logger is built upon