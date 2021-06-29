# go-rest-dev
Go Rest Implementation Experimentation

Run on local environment
$ docker build -t go-rest-dev .
$ docker run -it -p 8080:8080 go-rest-dev 
$ docker ps
http://localhost:8080/posts

To install grpc support for golang
Prerequisitos
instalar proto buffer
$ brew install protobuf
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
$ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1

$ go get -u google.golang.org/grpc
$ go get -u github.com/golang/protobuf/proto
$ go get -u github.com/golang/protobuf/protoc-gen-go

https://github.com/golang-standards/project-layout