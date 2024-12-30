run:
	go run cmd/main.go
verify:
	go test -goblin.timeout 10s
