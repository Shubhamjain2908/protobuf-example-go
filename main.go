package main

import (
	"fmt"

	simplepb "github.com/shubhamjain2908/protobuf-example-go/src/simple"
)

func main() {
	fmt.Println("Hello Go!")

	doSimple()
}

// returning reference to SimpleMessage (pass by reference)
func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "My simple Message",
		SampleList: []int32{1, 2, 7, 8},
	}

	fmt.Println(sm)

	sm.Name = "I renamed you"

	fmt.Println("Id => ", sm.GetId())
	fmt.Println("Name => ", sm.GetName())

	return &sm
}
