# go-httpe - HTTP Error Package (echo)

The `httpe` package in Go provides utilities for handling HTTP errors in a structured way. It allows creation of HTTP error objects with a specific status code and message, including optional internal error details. The package also supports parsing error messages into structured formats.

## Installation

```bash
go get github.com/yourusername/httpe
```

## Usage

### Creating a New HTTP Error

```go
err := httpe.NewHTTPError(404, "Not Found")
```

### Setting an Internal Error

```go
internalErr := errors.New("Some internal error")
httpErr := err.SetInternal(internalErr)
```

### Unwrapping an Error

```go
unwrappedErr := err.Unwrap()
```

### Getting the HTTP Status Code

```go
code := err.Code()
```

### Parsing Error Messages

```go
e := httpe.NewHTTPError(400, "field1: message1; field2: message2")
e.ParseMessageToErrors()
```

### Returning HTTP Error Response

```go
resp := httpe.NewHTTPError(400, "field1: message1; field2: message2").Return()
```

### Adding Internal Error to HTTP Error Response

```go
internalError := errors.New("internal error")
resp := e.WithInternal(internalError)
```

## Testing

The package includes a comprehensive suite of tests to ensure functionality:

```go
func TestNewHTTPError(t *testing.T) {
    // ...
}

func TestSetInternal(t *testing.T) {
    // ...
}

// ... additional test functions ...
```

Refer to the test suite for more detailed examples of how to use the `httpe` package.

-----

## Versioning and license

Our version numbers follow the [semantic versioning specification](http://semver.org/). You can see the available versions by checking the [tags on this repository](https://github.com/thiagozs/go-httpe/tags). For more details about our license model, please take a look at the [LICENSE](LICENSE.md) file.

**2024**, thiagozs
