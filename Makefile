all: clear install gen-example

install:
	protoc -I $(GOPATH)/src:. --go_out=. proto/ext.proto
	go install
gen-example:
	protoc -I $(GOPATH)/src:. --roles_out=. --go_out=. example/example.proto
clear:
	rm -f proto/ext.pb.go
	rm -f example/example.ex.go
	rm -f example/example.pb.go

