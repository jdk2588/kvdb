test:
	go test -v . -cover

build:
	env GOOS=linux GARCH=amd64 CGO_ENABLED=1 go build -o kvdb -installsuffix cgo .

build_mac:
	env GOOS=darwin GARCH=amd64 CGO_ENABLED=0 go build -o kvdb -a -installsuffix cgo .
