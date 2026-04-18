package main

import (
	. "github.com/agrellum/gorexxr/pkg/lang"
)

func main() {

	// Create a function shortcut, so we do not have to type long names.
	str := RxFromString

	indexed_string := str("indexed_string") // our base stem
	initial_value := str("opossum")         // it's initial value

	SayString("set indexed_string's initial value to \"opossum\"\n")

	node, err := indexed_string.GetNode(initial_value)
	node.Leaf = initial_value

	if err != nil {
		SayString(err.Error())
	}

	SayString("Add key first_entry to indexed_string\n and set the value to \"first little joey\"\n")

	node, err = indexed_string.GetNode(str("first_entry"))
	node.Leaf = str("first litte joey")

	if err != nil {
		SayString(err.Error())
	}

	SayString("Simulate returning the intial value,")
	SayString(" when a key has not been set.\n")

	if indexed_string.Exists(str("not_here")).ToString() == "0" {
		SayRexx(initial_value.OpCc(nil, str("\n")))
	} else {
		value, _ := indexed_string.Get(str("not_here"))
		SayRexx(value.OpCc(nil, str("\n")))
	}

	SayString("Get the value of first_entry\n")

	if indexed_string.Exists(str("first_entry")).ToString() == "0" {
		SayRexx(initial_value)
	} else {
		value, _ := indexed_string.Get(str("first_entry"))
		SayRexx(value)
	}

}
