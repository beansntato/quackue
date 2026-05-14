package quackue

import "fmt"

type Quackue struct {
}

func NewQuackue() *Quackue {
	return &Quackue{}
}

func (q *Quackue) Run() {
	fmt.Println("running...")
}
