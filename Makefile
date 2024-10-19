build:
	@go build -o bin/passdiy ./main.go

run:
	@./bin/passdiy

install:
	@sudo mv ./bin/passdiy /usr/bin/

uninstall:
	@sudo rm /usr/bin/passdiy
