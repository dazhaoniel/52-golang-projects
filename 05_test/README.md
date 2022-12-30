# Unit Testing in Go



Run all tests
```
go test .
```

Run all tests in one test suite
```
go test -run Test_alpha
```

Run one test
```
go test -run Test_isPrime
```

Check test coverage
```
go test -cover .
go test -coverprofile=coverage.out && go tool cover -html=coverage.out
```

`-v` for verbose