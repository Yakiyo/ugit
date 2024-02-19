build:
	go build -o ugit .

fmt:
	go fmt ./...

tidy:
	go mod tidy