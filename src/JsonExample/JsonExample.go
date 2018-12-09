package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type response1 struct {
	Page   int
	Fruits []string
}

type FamilyMember struct {
	Name    string
	Age     int
	Parents []string
}

func main() {

	rep := response1{Page: 1, Fruits: []string{"apple", "peach", "pear"}}

	blob, _ := json.MarshalIndent(rep, "", " ")

	fmt.Println("playing with json")

	s := string(blob[:])

	fmt.Println(s)

	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)

	var m FamilyMember

	if err := json.Unmarshal(b, &m); err != nil {
		log.Println(err)
	}

	fmt.Println(m.Name)
	fmt.Println(m.Parents)

}
