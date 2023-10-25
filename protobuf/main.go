package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)
import "learngo/protobuf/generated"

func Encode() *generated.Person {
	typ := &generated.Person{
		Name:   "Niklas",
		Id:     1,
		Email:  "nyklas",
		Phones: nil,
	}

	raw, err := proto.Marshal(typ)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	parsed := &generated.Person{}
	if err := proto.Unmarshal(raw, parsed); err != nil {
		log.Fatalln("Failed to parse address book:", err)
	}

	return parsed
}

func main() {
	typ := Encode()

	// Write the new address book back to disk.
	out, err := proto.Marshal(typ)
	if err != nil {
		log.Fatalln("Failed to encode address book:", err)
	}

	if err := ioutil.WriteFile("Person.log", out, 0644); err != nil {
		log.Fatalln("Failed to write address book:", err)
	}

	fmt.Println(out)
}
