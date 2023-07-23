run:
	echo "/home/pedram/Downloads/111.tar.gz" |  go run main.go

build:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o updatepackage main.go