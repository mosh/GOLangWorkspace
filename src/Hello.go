package main

import (
	"./employee"
	"./greeting"
)

func main() {

	//message := "Hello Go World!" // declaration and initialization
	//a, b, c := 1,false,3

	//var messageRef *string = &message

	//*messageRef = "goodbye"

	//var s = greeting.Salutation { "Joe", "Hello"}
	//var t = greeting.Salutation { Greeting : "Hello", Name : "Bill"}

	//fmt.Println(message,a,b,c,*greeting)
	//fmt.Println(PI,Language)
	//fmt.Println(A,B,C)

	//greeting.Greet(s)
	//greeting.Greet(t)

	//SomethingElse("Frankie", PrintLine)

	//greeting.ReturnFunc("Print me")("2")

	slice := greeting.Salutations{
		{Name: "Bob", Greeting: "Hello"},
		{Name: "Joe", Greeting: "Hi"},
		{Name: "Bill", Greeting: "Yo"},
	}

	slice.Greetings()

	slice[0].Rename("Mary")

	slice.Greetings()

	e := employee.Employee{Id: 1, Name: "Frank"}
	e.Praise()

}
