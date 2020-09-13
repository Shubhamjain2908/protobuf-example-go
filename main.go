package main

import (
	"fmt"
	"io/ioutil"
	"log"

	complexpb "github.com/shubhamjain2908/protobuf-example-go/src/complex"

	enumpb "github.com/shubhamjain2908/protobuf-example-go/src/enum_example"

	"github.com/golang/protobuf/jsonpb"

	"github.com/golang/protobuf/proto"

	simplepb "github.com/shubhamjain2908/protobuf-example-go/src/simple"
)

func main() {
	sm := doSimple()

	readAndWriteDeme(sm)
	jsonDemo(sm)

	doEnum()

	doComplex()
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

func readAndWriteDeme(sm proto.Message) {
	writeToFile("Simple.bin", sm)

	// creating a simple struct
	sm2 := &simplepb.SimpleMessage{}
	readFromFile("Simple.bin", sm2)
	fmt.Println("Read the contents", sm2)
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

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		log.Fatalln("Something went wrong when reading the file", err)
		return err
	}

	err2 := proto.Unmarshal(in, pb)
	if err2 != nil {
		log.Fatalln("Coudn't put the bytes into the protocol buffer struct", err)
		return err2
	}

	return nil
}

func jsonDemo(sm proto.Message) {
	smAsString := toJSON(sm)

	fmt.Println("Json string =>", smAsString)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(smAsString, sm2)
	fmt.Println("Successfully created protp Struct =>", sm2)
}

func toJSON(pb proto.Message) string {
	marshler := jsonpb.Marshaler{}
	out, err := marshler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}
	return out
}

func fromJSON(in string, pb proto.Message) {
	err := jsonpb.UnmarshalString(in, pb)
	if err != nil {
		log.Fatalln("Couldn't unmarshal the JSON into the pb struct", err)
	}
}

func doEnum() {
	ep := enumpb.EnumMessage{
		Id:           7,
		DayOfTheWeek: enumpb.DayOfTheWeek_SUNDAY,
	}

	fmt.Println("Day => ", ep.GetDayOfTheWeek())
}

func doComplex() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   1,
			Name: "First message",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			&complexpb.DummyMessage{
				Id:   2,
				Name: "Second message",
			},
			&complexpb.DummyMessage{
				Id:   3,
				Name: "Third message",
			},
		},
	}

	fmt.Println(cm)
}
