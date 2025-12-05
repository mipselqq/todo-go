.PHONY: run test genswagger prereq

run:
	go run ./cmd/todo-go

test:
	go test ./...

genswagger:
	~/go/bin/swag init -g ./cmd/todo-go/main.go -o ./docs/swagger

prereq:
	go install github.com/air-verse/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest
