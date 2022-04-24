all:
	go build -o calc
test:
	go test -v ./... -coverprofile=coverage.out && go tool cover -func=coverage.out
benchmark:
	go test -bench . ./...
