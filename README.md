# go-average-coverage
Returns the average coverage for multiple go packages. Suitable for use with Gitlab's test coverage parsing.

## Install
```bash
$ go install github.com/skagget77/go-average-coverage/cmd/go-average-coverage
```
Note that if you're current working directory contains a `go.mod` file `go-average-coverage` will be added to it as a dependency, which you probably don't want.

## With Gitlab
To use `go-average-coverage` in Gitlab's CI pipeline pipe the result of `go test -cover` to `go-average-coverage`.
```yml
test-my-app:
  stage: test
  script:
    - go test -cover -mod=readonly ./... | go-average-coverage
```
If you are using a standard golang docker image `go-average-coverage` can be installed on-the-fly with `go install` as described above.
