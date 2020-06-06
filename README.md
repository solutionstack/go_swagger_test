# go-swagger api demo

### Checkout code

`git clone git@github.com:solutionstack/go_swagger_test.git && cd go_swagger_test`


###To start HTTP server:
` go run ./cmd/main.go`


### To Regenerate Swagger files

`make gen`

### To re-generate mocks against our interfaces:
`go generate ./...`

### To run tests
`go test -v ./...`
