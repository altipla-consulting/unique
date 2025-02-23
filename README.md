
# unique

[![Go Reference](https://pkg.go.dev/badge/github.com/altipla-consulting/unique.svg)](https://pkg.go.dev/github.com/altipla-consulting/unique)

Small library to deduplicate Go values.

## Install

```shell
go get github.com/altipla-consulting/unique
```


## Usage

```go
original := []string{"foo", "bar", "foo", "baz"}
uq := unique.Slice[string](original)
```


## Contributing

You can make pull requests or create issues in GitHub. Any code you send should be formatted using `make gofmt`.


## License

[MIT License](LICENSE)
