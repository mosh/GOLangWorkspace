package employee

import (
	"fmt"
)


type Employee struct {
	Id int
	Name string
}

type Employable interface {
	Praise()
}

func (e Employee) Praise() {
	fmt.Println(e.Name + " is awesome")
}
