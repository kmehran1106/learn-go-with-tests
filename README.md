* Commands to format go code:
  * `go fmt ./...`
  * `golangci-lint run`
* Commands to run tests:
  * To run inside package/directory: `go test`
  * To run recursively across all folders: `go test ./...`
  * To run a particular directory recursvely: `go test ./<folder>/...`
  * Flag for coverage: `-cover`
  * Flag for benchmarking: `-bench .`
  * Flag for benchmarking with memory: `-bechmem`
  * <b>Note:</b> It is best to run benchmark tests on a particular directory only
