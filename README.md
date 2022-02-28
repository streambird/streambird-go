# Streambird Go API Client

This repository contains the open source Go client for Streambird API. Documentation can be found at: https://docs.streambird.com.

[![PkgGoDev](https://pkg.go.dev/badge/github.com/streambird/streambird-go/v1)](https://pkg.go.dev/github.com/streambird/streambird-go/v1)

## Requirements

- [Sign up](https://www.streambird.io) for a free Streambird account
- Create a new access key in the [Console](https://app.streambird.com/).
- An application written in Go to make use of this API

## Installation

The easiest way to use the Streambird API in your Go project is to install it using *go get*:

```
$ go get github.com/streambird/streambird-go/v1
```

### Examples

Here is a quick example on how to get started. Assuming the **go get** installation worked, you can import the streambird package like this:

```go
import "github.com/streambird/streambird-go/v1"
import streambirdClient "github.com/streambird/streambird-go/v1/client"
```

Then, create an instance of **streambirdClient**. It can be used to access the Streambird API.

```go
apiKey := "YOUR_SECRET_API_KEY"
// Create a client.
client := streambirdClient.New(apiKey)

```

This will give you something like:

```bash
$ go run example.go
```

Please see the other examples for a complete overview of all the available API calls.

## Errors

All other packages return `streambird.ErrorResponse` structs that contain a error message in `ErrorResponse.ErrorMessage`, error type in `ErrorResponse.ErrorType`, and HTTP status code in `ErrorResponse.StatusCode`.

## API Reference

Complete API documentation, instructions, and examples are available at:
[https://developers.streambird.io](https://developers.streambird.io).

## License

The Streambird Go API Client is licensed under [MIT License](http://opensource.org/licenses/MIT). Copyright (c) Streambird
