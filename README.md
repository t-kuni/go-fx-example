# go fx example

Sample code using [uber-go/fx](https://github.com/uber-go/fx).

## Usage

```
go run .
curl -X POST -d 'hello' http://localhost:44444/echo
```

## Test

```
go install go.uber.org/mock/mockgen@latest
go generate ./...
```