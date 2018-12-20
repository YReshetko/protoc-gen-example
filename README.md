# Roles generator for protobuf methods (GO only)
## Installation
```
go get -u github.com/YReshetko/protoc-gen-roles
```
This is used with proto3 so the dependency of project should be based on https://github.com/protocolbuffers/protobuf (master)

## How to use
```proto
service Service {
    rpc SomeMethod (Request) returns (Response) {
        option (gen.roles.roles) = "Admin,User";
    }
}
```

```gen.roles.roles``` - is method option which come with the plugin. It should have string role values separated by comma without spaces.
Then the proto is generated to map where key - rpc method which is used on middleware interceptors, value - array of string roles.
 
## Generate go code
Use regular proto command
```
protoc --roles_out=. --go_out=. proto/some_proto.proto
```
If there is not specified ``go_package`` option the the output file will be in the same folder as protofile and has the same package as proto file (dots are replaced with underscores).
If ``go_package`` is specified then the output file will go to corresponding folder(s) inside the folder with proto file


#### TODO
- [ ] Add deps
- [ ] Make sure that ``go_package``  option works correctly
- [ ] Current roleas are string, should be changed to custom type of client code