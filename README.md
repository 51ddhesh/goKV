# goKV

goKV is a simple key-value store written in Go. This project is being made with an aim to get into concurrency and high performance systems.

goKV uses `sync` from Go for concurrency.

### Usage
```
go build -o kv
./kv
```

### Testing
The [`kv_test.go`](./kv_test.go) tests the basic functionalities like testing `Get`, `Set` and data persistence. Tests can be run using:
```
go test
```

