build:
	go build -o bin/uman ./main.go

run: build
	./bin/uman test.um