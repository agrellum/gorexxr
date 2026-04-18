package main

import (
	. "github.com/agrellum/gorexxr/pkg/lang"
)

func main() {

	// Create a function shortcut, so we do not have to type long names.
	str := RxFromString

	SayString("Create and display a string 999999999 characters long? Y/N")

	arg, err := AskOne()

	if err == nil {

		template := []rune{1, 10, 2, 0, 1, 0}
		results := make([]*Rexx, 2)

		err := Parse(arg, template, results)

		if err == nil {

			if results[0].ToString() == "Y" || results[0].ToString() == "y" {

				huge, err := str("123").Copies(str("333333333"))

				if err == nil {
					leng := huge.Length()
					stop, _ := leng.ToInt32()

					// This is a mathematical cheat - just call SubStr and then OpCc one time
					// our string repeats "123" - 1001001 is divisible by 3
					// The loop will run exactly 999 times. 999*1001001 = 999999999
					chunk, _ := huge.SubStr(RxFromInt32(1), RxFromInt32(1001001), RxFromRune(rune(0)))

					// Add the null character to the end of chunk.
					// It will be automatically be removed to make one long continuous print
					data := chunk.OpCc(nil, str("\000"))

					for i := int32(1); i <= stop; i += 1001001 {

						SayRexx(data)

					}
				} else {
					SayString(err.Error())
				}
			} else {
				SayRexx(str("Goodbye!!!"))
			}
		} else {
			SayString(err.Error())
		}
	} else {
		SayString(err.Error())
	}

}
