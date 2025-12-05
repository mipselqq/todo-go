run:
	go run ./cmd/todo-go

genswagger:
	~/go/bin/swag init -g ./cmd/todo-go/main.go

prereq:
	go install github.com/air-verse/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest
