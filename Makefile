all: clear install gen-example

install:
	protoc --proto_path=$(GOPATH)/src:. --go_out=. proto/ext.proto
	go install
gen-example:
	protoc --proto_path=$(GOPATH)/src:. --example_out=. --go_out=. example/example.proto
clear:
	rm -f proto/ext.pb.go
	rm -f example/example.ex.go
	rm -f example/example.pb.go

