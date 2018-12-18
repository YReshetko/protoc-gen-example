all: clear install gen

install:
	go install
gen:
	protoc --proto_path=$(GOPATH)/src:. --example_out=. --go_out=. proto/*.proto
gen-clean:
	protoc --proto_path=$(GOPATH)/src:. --go_out=. proto/example.proto

clear:
	rm -f proto/example.ex.go
	rm -f proto/example.pb.go

