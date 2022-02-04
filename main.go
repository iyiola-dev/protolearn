package main

import (
	"fmt"
	"io/ioutil"
	"log"

	complexpb "github.com/iyiola-dev/protobufex/src/complex"
	enumpb "github.com/iyiola-dev/protobufex/src/enum"
	simplepb "github.com/iyiola-dev/protobufex/src/simple"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func main() {

	sm := doSimple()
	readAndWriteDemo(sm)
	jsonDemo(sm)
	readEnum()
	doComplex()

}

func doComplex() {
	cm := complexpb.ComplexMessage{
		DummyMessage: &complexpb.DummyMessage{
			Id:   1,
			Name: "The first one",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			{
				Id:   2,
				Name: "The second one"},
			{
				Id:   3,
				Name: "The third one",
			},
			{
				Id:   4,
				Name: "The fourth one",
			},
		},
	}
	cmJson := toJson(&cm)
	fmt.Println(cmJson)
}
func jsonDemo(sm protoreflect.ProtoMessage) {
	pol := toJson(sm)
	fmt.Println(pol)
	data := []byte(pol)
	pb := &simplepb.SimpleMessage{}

	fromJson(data, pb)
	fmt.Println(">>>>>>>>>>>>>>", pb)
}
func toJson(pb protoreflect.ProtoMessage) string {
	marshaller := protojson.MarshalOptions{}

	out := marshaller.Format(pb)

	return out
}

func fromJson(s []byte, m protoreflect.ProtoMessage) error {
	err := protojson.Unmarshal(s, m)
	if err != nil {
		panic(err)
	}
	return nil
}
func readEnum() {
	em := enumpb.EnumMessage{
		Id:            12,
		DaysOfTheWeek: enumpb.DaysOfTheWeek_THURSDAY,
	}

	fmt.Println(">>>>>>>>>>>>>>>", &em)
}
func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)
	pb := &simplepb.SimpleMessage{}
	readFile("simple.bin", pb)
	fmt.Println(pb)
}

func writeToFile(fName string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("cannot marshal proto due to ", err)
		return err
	}
	err = ioutil.WriteFile(fName, out, 0644)
	if err != nil {
		log.Fatalln("cannot write due to ", err)
		return err
	}
	return nil
}
func readFile(fName string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Println("could not read file due to: ", err)
		return err
	}
	err = proto.Unmarshal(in, pb)

	if err != nil {
		log.Println("could not unmarshal data from protobuf due to:", err)
		return err
	}
	return nil
}
func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         1,
		SampleList: []int32{1, 2, 4, 3, 2, 4, 5},
		Name:       "A very simple name",
		IsSimple:   false,
	}
	fmt.Println(&sm)
	return &sm
}
