package main

import (
	. "github.com/agrellum/gorexxr/pkg/lang"
)

func main() {

	// Create a function shortcut, so we do not have to type long names.
	str := RxFromString

	SayString("Enter three or more items seperated by spaces.")

	arg, err := Ask()

	if err == nil {

		template := []rune{1, 10, 4, 0, 1, 2, 3, 0}
		results := make([]*Rexx, 4)

		err := Parse(arg, template, results)

		if err == nil {

			SayRexx(str("First entry: ").OpCc(nil, results[0]))
			SayRexx(str("Second entry: ").OpCc(nil, results[1]))
			SayRexx(str("Third entry: ").OpCc(nil, results[2]))
			SayRexx(str("All the rest: ").OpCc(nil, results[3]))

		} else {
			SayString(err.Error())
		}
	} else {
		SayString(err.Error())
	}

	SayString("\n--------------  SIMPLE PARSING INSTRUCTIONS  --------------")
	SayString("Create a one line NetRexx program with only your parse statement.")
	SayString("Example: parse arg a b c .")
	SayString("Run nrc -compile -keep on it.")
	SayString("Run a Java formatter on the java source.")
	SayString("[]rune{data} will be in $01 = {1, 10, 4, 0, 1, 2, 3, 0};")
	SayString("make([]*Rexx, data) will be in $1[] = new netrexx.lang.Rexx[4];")

}
