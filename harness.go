package quackue

import (
	"fmt"

	wal "beansntato.dev/walrus"
)

type Quackue struct {
}

func NewQuackue() *Quackue {
	return &Quackue{}
}

func (q *Quackue) Run() {
	fmt.Println("running...")
	wal.New()
}
