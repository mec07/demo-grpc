# grpc
An example grpc project. This repo was inspired by this article: https://medium.com/pantomath/how-we-use-grpc-to-build-a-client-server-system-in-go-dd20045fa1c2

## Setup
To run this you will need the `protobuf` package installed. If you're on a mac you can do:
```
brew install protobuf
```
This gives you access to the `protoc` command.
You will also require  the `protoc-gen-go` executable:
```
go get -u github.com/golang/protobuf/protoc-gen-go
```
Make sure that you have added `$GOPATH/bin` to your `PATH` environment variable.

To compile the protobuf file, `api/api.proto` execute the following command:
```
protoc -I api/ -I ${GOPATH}/src --go_out=plugins=grpc:api api/api.proto
```

To compile the server into a binary:
```
go build -i -v -o bin/server server/main.go
```
To compile the client into a binary:
```
go build -i -v -o bin/client client/main.go
```

## Run
First run the server:
```
bin/server
```
Then run the client (which talks to the server):
```
bin/client
```
