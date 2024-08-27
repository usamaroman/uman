build:
	go build -o bin/uman ./cmd/uman/main.go

run: build
	./bin/uman test.um