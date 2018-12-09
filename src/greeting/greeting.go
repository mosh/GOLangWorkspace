package greeting

import "fmt"

const (
	PI       = 3.14
	Language = "Go"
)

const (
	A = iota
	B = iota
	C = iota
)

type Salutation struct {
	Name     string
	Greeting string
}

type Salutations []Salutation

func (salutations Salutations) Greetings() {
	for _, s := range salutations {
		fmt.Println(s.Name)
	}
}

func (salutations Salutations) ChannelGreeter(c chan Salutation) {
	for _, s := range salutations {
		c <- s
	}
	close(c)
}

func (salutation *Salutation) Rename(newName string) {
	salutation.Name = newName
}

func Greet(s Salutation) {
	_, alternateMessage := CreateMessage(s.Name, s.Greeting)
	fmt.Println(alternateMessage)
}

func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "HEY !" + name
	return
}

func AnotherCreateMessage(name string, greeting ...string) (message string, alternate string) {
	x := len(greeting)
	fmt.Println(x)
	message = greeting[0] + " " + name
	alternate = ""
	return
}

type Printer func(string)

func SomethingElse(value string, do Printer) {
	do(value)
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func ReturnFunc(value string) Printer {
	return func(s string) {
		fmt.Println(value + s)
	}
}
