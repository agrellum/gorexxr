package lang

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type rexxio struct {
}

func rxio() (rcvr *rexxio) {
	rcvr = &rexxio{}
	return
}

// NetRexx: method Ask() static returns Rexx
//
// Differences: None
func Ask() (*Rexx, error) {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		return nil, RxException(7, "input error.")
	}
	line = strings.Replace(line, "\n", "", -1)
	return RxFromString(line), nil
}

// NetRexx: method AskOne() static returns Rexx
//
// Differences: None
func AskOne() (*Rexx, error) {
	one := rune(0)
	reader := bufio.NewReader(os.Stdin)
	one, _, err := reader.ReadRune()
	if err != nil {
		return nil, RxException(7, "input error.")
	}
	return RxFromRune(one), nil
}

// NetRexx: method Say(aline=char[]) static returns boolean
//
// Differences: None
func Say(aline []rune) bool {
	if aline == nil {
		fmt.Println()
	} else {
		if len(aline) == 0 {
			fmt.Println()
		} else if aline[len(aline)-1] != '\000' {
			fmt.Println(string(aline))
		} else {
			bline := make([]rune, len(aline)-1)
			copy(bline, aline)
			fmt.Printf("%v", string(bline))
		}
	}
	return false
}

// NetRexx: method Say(b=boolean) static returns boolean
//
// Differences: None
func SayBool(b bool) bool {
	return Say(RxFromBool(b).ToRunes())
}

// NetRexx: method Say(f=float) static returns boolean
//
// Differences: None
func SayFloat32(f float32) bool {
	return Say(RxFromFloat32(f).ToRunes())
}

// NetRexx: method Say(d=double) static returns boolean
//
// Differences: None
func SayFloat64(d float64) bool {
	return Say(RxFromFloat64(d).ToRunes())
}

// NetRexx: method Say(n=long) static returns boolean
//
// Differences: None
func SayInt64(n int64) bool {
	return Say([]rune(fmt.Sprintf("%v", n)))
}

// NetRexx: method Say(line=Rexx) static returns boolean
//
// Differences: If line == nil, Say(nil)
func SayRexx(line *Rexx) bool {
	if line == nil {
		// Prevents Go Panic
		return Say(nil)
	} else {
		return Say(line.ToRunes())
	}
}

// NetRexx: method Say(c=char) static returns boolean
//
// Differences: None
func SayRune(c rune) bool {
	ca := make([]rune, 1)
	ca[0] = c
	return Say(ca)
}

// NetRexx: method Say(str=String) static returns boolean
//
// Differences: None
func SayString(str string) bool {
	return Say([]rune(str))
}
