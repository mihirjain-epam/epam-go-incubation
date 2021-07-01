Run all tests using - `go test`
Run all tests for function `proxyOrCacheRequest` using - `go test -run ProxyOrCacheRequest`
Run all tests for ports using - `go test -run ProxyOrCacheRequest/Ports`
Run test for individual port using - `go test -run ProxyOrCacheRequest/Ports/Port=<port>` 
*Note* - testing supported for ports - [8080,9443,8081]
Run all tests for Unsafe methods using - `go test -run ProxyOrCacheRequest/UnsafeMethods`
Run tests for individual Unsafe methods using - `go test -run ProxyOrCacheRequest/UnsafeMethods/UnsafeMethod=<unsafe-method>`
*Note* - testing supported for ports - [POST,PUT,DELETE]
Run test for BadUrl using - `go test -run ProxyOrCacheRequest/BadUrl`

To check coverage - `go test -coverprofile="c.out"` `go tool cover -html="c.out"`