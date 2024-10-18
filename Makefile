build:
	go build -o pass ./main.go

install:
	sudo mv ./pass /usr/bin/
