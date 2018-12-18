package main

import (
	"github.com/golang/protobuf/proto"
	plugin "github.com/golang/protobuf/protoc-gen-go/plugin"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	rq := extractRequest()
	rs := processRequest(rq)
	save(rs)
}

func extractRequest() *plugin.CodeGeneratorRequest {
	log.Println("Run as gen plugin")
	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic("reading input")
	}
	//panic(fmt.Sprintf("%+v", string(data)))
	rq := new(plugin.CodeGeneratorRequest)
	if err := proto.Unmarshal(data, rq); err != nil {
		panic("parsing input proto")
	}

	if len(rq.FileToGenerate) == 0 {
		panic("no files to generate")
	}
	return rq
}



func save(rs *plugin.CodeGeneratorResponse)  {
	data, err := proto.Marshal(rs)
	if err != nil {
		panic("failed to marshal output proto")
	}
	_, err = os.Stdout.Write(data)
	if err != nil {
		panic("failed to write output proto")
	}
	//log.Println("Stdout written bytes:", n)
}


