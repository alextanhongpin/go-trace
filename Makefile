trace:
	@go run main.go 2> trace.out
	@go tool trace trace.out
