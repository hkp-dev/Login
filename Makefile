# - Tạo file Makefile và đặt lệnh build trong đó, build bằng lệnh make (cMake)
BINARY_NAME = app
GOLANG = go

build:
	$(GOLANG) build -o $(BINARY_NAME) cmd/main.go
clean:
	rm -f $(BINARY_NAME)
run: build
	./$(BINARY_NAME)
	go run cmd/main.go
	