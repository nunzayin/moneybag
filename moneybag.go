package main

import (
	"errors"
)

const (
	BOTTOMLESS = "_"
	TAG_PREFIX = "#"
)

type (
	State struct {
		bags map[string]uint
		tags map[string][]string
	}
	Amount interface {
		get() uint
	}
	Transfer struct {
		source string
		destination string
		amount Amount
		comment string
	}
	Tagging struct {
		tag string
		names []string
	}
	Modifier interface {
		eval(state *State)
	}
	Book []Modifier
)

func (transfer Transfer) eval(state *State) error {
	
}

func main() {

}
