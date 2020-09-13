package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/proto"

	simplepb "github.com/shubhamjain2908/protobuf-example-go/src/simple"
)

func main() {
	sm := doSimple()

	writeToFile("Simple.bin", sm)
}

// returning reference to SimpleMessage (pass by reference)
func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My simple Message",
		SampleList: []int32{1, 2, 7, 8},
	}

	sm.Name = "I renamed you"

	fmt.Println("Id => ", sm.GetId())
	fmt.Println("Name => ", sm.GetName())

	return &sm
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialize to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written!")
	return nil
}
