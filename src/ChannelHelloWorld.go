package main

import (
	"fmt"

	"./greeting"
)

func main() {

	salutations := greeting.Salutations{
		{Name: "Bob", Greeting: "Hello"},
		{Name: "Bill", Greeting: "Hi"},
		{Name: "John", Greeting: "Yo!"},
	}

	c := make(chan greeting.Salutation)
	c2 := make(chan greeting.Salutation)

	// call a coroutine that will fill the channel

	go salutations.ChannelGreeter(c)
	go salutations.ChannelGreeter(c2)

	/*
		for s := range c {
			fmt.Println(s.Name)
		}
	*/
	for {
		select {
		case s, ok := <-c:
			if ok {
				fmt.Println(s, ":1")
			} else {
				return
			}
		case s, ok := <-c2:
			if ok {
				fmt.Println(s, ":2")
			} else {
				return
			}
		default:
			fmt.Println("Waiting...")
		}
	}
}
