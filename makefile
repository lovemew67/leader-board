all:
	go mod tidy
	go mod verify
	go run main.go