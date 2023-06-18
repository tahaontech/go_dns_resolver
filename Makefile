build:
	@go build -o bin/output cmd/main.go

run: build
	@bin/output

testwindows:
	nslookup - 127.0.0.1