# alb-log-parser

A Go library for parsing AWS Application Load Balancer (ALB) access logs. This library provides a simple and efficient way to parse ALB log entries into structured data.

## Features

- 🚀 Fast and efficient parsing using regular expressions
- 📦 Support for all ALB log fields (as of November 2024)
- 💪 Type-safe field access through structured Go types

## Installation

```bash
go get github.com/hacomono-lib/alb-log-parser
```

## Usage

Here's a simple example of how to use the library:

```go
package main

import (
    "fmt"
    "compress/gzip"
    "os"
    "bufio"
    albparser "github.com/hacomono-lib/alb-log-parser"
)

func main() {
    // Create a new parser instance
    parser := albparser.NewAlbLogParser()

    // Open gzipped ALB log file
    // ALB logs are often in gz format when fetched from S3.
    f, err := os.Open("alb-logs.gz")
    if err != nil {
        fmt.Printf("Error opening file: %v\n", err)
        return
    }
    defer f.Close()

    // Create gzip reader
    gz, err := gzip.NewReader(f)
    if err != nil {
        fmt.Printf("Error creating gzip reader: %v\n", err)
        return
    }
    defer gz.Close()

    // Read line by line
    scanner := bufio.NewScanner(gz)
    for scanner.Scan() {
        record, err := parser.ParseAlbLog(scanner.Text())
        if err != nil {
            fmt.Printf("Error parsing log: %v\n", err)
            continue
        }

        // Access the parsed fields
        fmt.Printf("Request Method: %s\n", record.HttpMethod)
        fmt.Printf("Status Code: %s\n", record.ElbStatusCode)
        fmt.Printf("Client IP: %s\n", record.ClientIP)
    }

    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading file: %v\n", err)
    }
}
```

others

```go
package main

import (
    "fmt"
    albparser "github.com/hacomono-lib/alb-log-parser"
)

func main() {
    // Create a new parser instance
    parser := albparser.NewAlbLogParser()

    // Example ALB log line
    logLine := `http 2024-01-01T12:00:00.000000Z app/my-loadbalancer/1234567890 172.16.1.1:1234 10.0.1.1:80 0.001 0.002 0.003 200 200 123 456 "GET https://example.com/ HTTP/1.1" ...`

    // Parse the log line
    record, err := parser.ParseAlbLog(logLine)
    if err != nil {
        fmt.Printf("Error parsing log: %v\n", err)
        return
    }

    // Access the parsed fields
    fmt.Printf("Request Method: %s\n", record.HttpMethod)
    fmt.Printf("Status Code: %s\n", record.ElbStatusCode)
    fmt.Printf("Client IP: %s\n", record.ClientIP)
}
```

## Supported Fields

The parser extracts all fields from ALB access logs.

For a complete list of supported fields, see the [AWS documentation](https://docs.aws.amazon.com/elasticloadbalancing/latest/application/load-balancer-access-logs.html).

## Contributing

Contributions are welcome! Please feel free to submit pull requests. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

Diwamoto ([@Diwamoto](https://github.com/Diwamoto))