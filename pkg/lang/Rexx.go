// A port of the NetRexx Runtime 3.0.X to Go.
//
// Please read the NetRexx documentation or look at its source code.
//
// Not all methods were ported.
//
// There are some changes because Go strings cannot be nil.
//
// There are more error conditions set than NetRexx signals.
//
// Some functions prevent or adjust nil, 0 or negative parameters.
//
// Some functions allow -1 or return -1 for robustness.
//
// Use export GOMEMLIMIT=XXXXMiB when dealing with large data.
//
// Note these differences when using this module and mixing Go code:
//
//	1 - Bools
//		Use ToString() to match Netrexx's "0" or "1"
//
//	2 - Plain, Scientific and Engineering notation
//		Use the Format function to achieve the desired output.
//
//	3 - Rounding
//		Operating within this module only will avoid problems.
package lang

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

const DefaultDigits = DEFAULT_DIGITS
const DefaultForm = DEFAULT_FORM
const Lowers = "abcdefghijklmnopqrstuvwxyz"
const Uppers = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const Digits09 = "0123456789"

var Hexes = []rune("0123456789ABCDEFabcdef")

const ispos int8 = 1
const iszero int8 = 0
const isneg int8 = -1
const NotaNum int8 = -2

var MinExp int32 = -999999999
var MaxExp int32 = 999999999
var MinArg int32 = -999999999
var MaxArg int32 = 999999999

type Rexx struct {
	chars []rune
	ind   int8
	form  int8
	mant  []rune
	exp   int32
	dig   int32
	coll  map[*Rexx]*RexxNode
	mux   sync.Mutex
}

// NetRexx: method Rexx(s=char[],trynum=boolean) shared
//
// # This method is not accessible in NetRexx
//
// Differences: If s is nil, rcvr.chars is set to an empty rune array
func rx(s []rune, trynum bool) (rcvr *Rexx) {
	rcvr = &Rexx{}
	var dvalue int32 = 0
	eneg := false
	var j int32 = 0
	var c int32 = 0

	// added for go nil passed in - fixes bad layout
	if s == nil {
		// Prevents Go Panic later on.
		rcvr.chars = []rune{}
	} else {
		rcvr.chars = s
	}

	rcvr.ind = NotaNum
	if len(s) == 0 {
		return
	}
	if s[0] > '9' {
		if s[0] <= '\177' {
			return
		}
		if !unicode.IsNumber(s[0]) {
			return
		}
	}
	if !trynum {
		return
	}
	leng := int32(len(s))
	i := leng - 1
	for ; i >= 0; i-- {
		if s[i] != ' ' {
			break
		}
		leng--
		if leng == 0 {
			return
		}
	}
	var insign int32 = 0
	var start int32 = -1
	_1 := leng - 1
	i = 0
	for ; i <= _1; i++ {
		if s[i] == ' ' {
			continue
		}
		if s[i] == '-' {
			if insign != 0 {
				return
			}
			insign = -1
			continue
		}
		if s[i] != '+' {
			start = i
			break
		}
		if insign != 0 {
			return
		}
		insign = 1
	}
	if start < 0 {
		return
	}
	exotic := false
	var d int32 = 0
	var e int32 = -1
	var last int32 = -1
	_2 := leng - 1
	i = start
	for ; i <= _2; i++ {
		if s[i] >= '0' {
			if s[i] <= '9' {
				last = i
				d++
				continue
			}
		}
		if s[i] == '.' {
			if e >= 0 {
				return
			}
			e = i - start
			continue
		}
		if s[i] != 'e' {
			if s[i] != 'E' {
				if !unicode.IsNumber(s[i]) {
					return
				}
				exotic = true
				last = i
				d++
				continue
			}
		}
		if i > leng-3 {
			return
		}
		// if leng-i > 11 {
		// 	return
		// }
		if s[i+1] == '-' {
			eneg = true
		} else if s[i+1] == '+' {
			eneg = false
		} else {
			return
		}
		if leng-i-2 > 9 {
			return
		}
		_3 := leng - 1
		j = i + 2
		for ; j <= _3; j++ {
			if s[j] < '0' {
				return
			}
			if s[j] > '9' {
				if !unicode.IsNumber(s[j]) {
					return
				}
				dvalue = getdigit(s[j])
				if dvalue < 0 {
					return
				}
			} else {
				dvalue = int32(s[j]) - int32('0')
			}
			rcvr.exp = rcvr.exp*10 + dvalue
		}
		if eneg {
			rcvr.exp = int32(-rcvr.exp)
		}
		break
	}
	if d == 0 {
		return
	}
	if e >= 0 {
		rcvr.exp = rcvr.exp + e - d
	}
	_4 := last - 1
	i = start
	for ; i <= _4; i++ {
		if s[i] == '.' {
			start++
		} else if s[i] == '0' {
			start++
			d--
		} else if s[i] <= '9' {
			break
		} else {
			if getdigit(s[i]) != 0 {
				break
			}
			start++
			d--
		}
	}
	for {
		if exotic {
			rcvr.mant = make([]rune, d)
			j = 0
			_5 := d - 1
			i = 0
			for ; i <= _5; i++ {
				if s[start+i] == '.' {
					j = 1
				}
				c = int32(s[start+i+j])
				if c <= '9' {
					rcvr.mant[i] = rune(c)
				} else {
					dvalue = getdigit(rune(c))
					if dvalue < 0 {
						return
					}
					rcvr.mant[i] = rune(dvalue + int32('0'))
				}
			}
		} else if d == int32(len(rcvr.chars)) {
			rcvr.mant = rcvr.chars
		} else {
			rcvr.mant = make([]rune, d)
			j = 0
			_6 := d - 1
			i = 0
			for ; i <= _6; i++ {
				if s[start+i] == '.' {
					j = 1
				}
				rcvr.mant[i] = s[start+i+j]
			}
		}
		if !false {
			break
		}
	}
	if rcvr.mant[0] == '0' {
		rcvr.ind = iszero
		rcvr.exp = 0
	} else if insign < 0 {
		rcvr.ind = isneg
	} else {
		rcvr.ind = ispos
	}
	return
}

// NetRexx: method Rexx(flag=boolean)
//
// Differences: None
func RxFromBool(flag bool) (rcvr *Rexx) {
	rcvr = &Rexx{}
	rcvr.exp = 0
	rcvr.dig = 9
	rcvr.form = DefaultForm
	rcvr.mant = make([]rune, 1)
	if flag {
		rcvr.mant[0] = '1'
		rcvr.ind = ispos
	} else {
		rcvr.mant[0] = '0'
		rcvr.ind = iszero
	}
	rcvr.chars = rcvr.mant
	return
}

// NetRexx: method Rexx(in=Rexx)
//
// Differences: If in is nil, it is set to rx(nil, false)
func rxfromclone(in *Rexx) (rcvr *Rexx) {
	if in == nil {
		// Prevents Go Panic later on.
		in = rx(nil, false)
	}
	rcvr = &Rexx{}
	rcvr.chars = in.chars
	rcvr.ind = in.ind
	rcvr.mant = in.mant
	rcvr.exp = in.exp
	rcvr.dig = in.dig
	rcvr.form = in.form
	rcvr.coll = nil
	return
}

// NetRexx: method Rexx() shared
//
// Differences: Private
func rxfromempty() (rcvr *Rexx) {
	rcvr = &Rexx{}
	return
}

// NetRexx: method Rexx(num=float)
//
// Differences: None
func RxFromFloat32(num float32) (rcvr *Rexx) {
	value, _ := torxfromfloat64(float64(num), 7)
	rcvr = rxfromclone(value)
	return
}

// NetRexx: method Rexx(num=double)
//
// Differences: None
func RxFromFloat64(num float64) (rcvr *Rexx) {
	value, _ := torxfromfloat64(num, 16)
	rcvr = rxfromclone(value)
	return
}

// NetRexx: method Rexx(num=short)
//
// Differences: None
func RxFromInt16(num int16) (rcvr *Rexx) {
	rcvr = RxFromInt32(int32(num))
	return
}

// NetRexx: method Rexx(num=int)
//
// Differences: None
func RxFromInt32(num int32) (rcvr *Rexx) {
	rcvr = &Rexx{}
	rcvr.exp = 0
	rcvr.form = DefaultForm
	if num <= 9 {
		if num >= -9 {
			rcvr.mant = make([]rune, 1)
			if num > 0 {
				rcvr.mant[0] = rune(int32('0') + num)
				rcvr.chars = rcvr.mant
				rcvr.ind = ispos
			} else if num == 0 {
				rcvr.mant[0] = '0'
				rcvr.chars = rcvr.mant
				rcvr.ind = iszero
			} else {
				rcvr.chars = make([]rune, 2)
				rcvr.chars[0] = '-'
				rcvr.chars[1] = rune(int32('0') - num)
				rcvr.mant[0] = rcvr.chars[1]
				rcvr.ind = isneg
			}
			return
		}
	}
	if num > 0 {
		rcvr.ind = ispos
		rcvr.mant = []rune(strconv.FormatInt(int64(num), 10))
		rcvr.chars = rcvr.mant
		return
	}
	rcvr.ind = isneg
	rcvr.chars = nil
	rcvr.dig = 10
	// match the Java code
	if num == int32(int64(-2147483648)) {
		rcvr.mant = []rune("2147483648")
		return
	}
	num = int32(-num)
	rcvr.mant = []rune(strconv.FormatInt(int64(num), 10))
	return
}

// NetRexx: method Rexx(num=long)
//
// Differences: None
func RxFromInt64(num int64) (rcvr *Rexx) {
	rcvr = rx([]rune(strconv.FormatInt(num, 10)), true)
	return
}

// NetRexx: method Rexx(num=byte)
//
// Differences: None
func RxFromInt8(num int8) (rcvr *Rexx) {
	rcvr = RxFromInt32(int32(num))
	return
}

// NetRexx: method Rexx(inchar=char)
//
// Differences: None
func RxFromRune(inchar rune) (rcvr *Rexx) {
	rcvr = &Rexx{}
	_new := make([]rune, 1)
	_new[0] = inchar
	r := rx(_new, true)
	rcvr.chars = r.chars
	rcvr.ind = r.ind
	rcvr.mant = r.mant
	rcvr.exp = r.exp
	rcvr.dig = r.dig
	rcvr.form = r.form
	return
}

// NetRexx: method Rexx(in=char[])
//
// Differences: None
func RxFromRunes(in []rune) (rcvr *Rexx) {
	rcvr = &Rexx{}
	_new := make([]rune, len(in))
	copy(_new, in)
	r := rx(_new, true)
	rcvr.chars = r.chars
	rcvr.ind = r.ind
	rcvr.mant = r.mant
	rcvr.exp = r.exp
	rcvr.dig = r.dig
	rcvr.form = r.form
	return
}

// NetRexx: method Rexx(string=java.lang.String)
//
// Differences: None
func RxFromString(str string) (rcvr *Rexx) {
	rcvr = rx([]rune(str), true)
	return
}

// NetRexx: method Rexx(strings=java.lang.String[])
//
// Differences: None
func RxFromStrings(strings []string) (rcvr *Rexx) {
	rcvr = rx(sa2ca(strings), true)
	return
}

// NetRexx: method tochar(s=char[]) static returns char signals NotCharacterException
//
// Differences: None
func ToRuneFromRunes(s []rune) (rune, error) {
	if len(s) != 1 {
		return 0, RxException(7, string(s))
	}
	return s[0], nil
}

// NetRexx: method tochar(s=java.lang.String) static returns char signals NotCharacterException
//
// Differences: None
func ToRuneFromString(s string) (rune, error) {
	tmp := []rune(s)
	if len(tmp) != 1 {
		return 0, RxException(7, s)
	}
	return tmp[0], nil
}

// NetRexx: method tochararray(r=Rexx) constant returns char[]
//
// Differences: if r is nil, returns nil
func ToRunesFromRexx(r *Rexx) []rune {
	if r == nil {
		// Prevents Go Panic
		return nil
	}
	if r.chars == nil {
		r.chars = r.layout()
	}
	res := make([]rune, len(r.chars))
	copy(res, r.chars)
	return res
}

// NetRexx: method tochararray(c=char) static returns char[]
//
// Differences: None
func ToRunesFromRune(c rune) []rune {
	ca := make([]rune, 1)
	ca[0] = c
	return ca
}

// NetRexx: method toRexx(ca=char[]) constant returns Rexx
//
// Differences: Netrexx can return null
// rx(s []rune, trynum bool) will return "" to prevent Go panic on nil strings.
func ToRxFromRunes(ca []rune) *Rexx {
	_new := make([]rune, len(ca))
	copy(_new, ca)
	return rx(_new, true)
}

// NetRexx: method toRexx(s=String) constant returns Rexx
//
// Differences: Netrexx can return null
// Go will not allow s to be nil
func ToRxFromString(s string) *Rexx {
	return rx([]rune(s), true)
}

// NetRexx: method toString(r=Rexx) constant returns String
//
// Differences: If r is nil, returns ""
func ToString(r *Rexx) string {
	if r == nil {
		// golang does not allow nil strings
		return ""
	}
	if r.chars == nil {
		r.chars = r.layout()
	}
	value := string(r.chars)
	return value
}

// NetRexx: method abbrev(b=Rexx,len=Rexx b.intlength()) returns Rexx
//
// Differences: neither argument b nor leng can be nil
func (rcvr *Rexx) Abbrev(b *Rexx, leng *Rexx) (*Rexx, error) {
	if b == nil || leng == nil {
		// Prevents Go Panic
		return nil, RxException(0, "neither argument b nor leng can be nil")
	}
	n, err := leng.intcheck(0, MaxArg)
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if b.chars == nil {
		b.chars = b.layout()
	}
	return RxFromBool(abbrev(rcvr.chars, b.chars, n)), nil
}

// NetRexx: method abs() returns Rexx signals NumberFormatException
//
// Differences: None
func (rcvr *Rexx) Abs() (*Rexx, error) {
	var set *RexxSet
	if rcvr.ind == NotaNum {
		return nil, RxException(9, string(rcvr.chars))
	}
	if int32(len(rcvr.mant)) > DefaultDigits {
		set = RxSetWithDigit(int32(len(rcvr.mant)))
	} else {
		set = nil
	}
	if rcvr.ind >= 0 {
		return rcvr.OpPlus(set)
	}
	return rcvr.OpMinus(set)
}

// NetRexx: method b2d(bil=Rexx-1) returns Rexx
//
// Differences: argument bil cannot be nil
func (rcvr *Rexx) B2d(bil *Rexx) (*Rexx, error) {
	if bil == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument bil cannot be nil")
	}
	var firstd string
	var ble int32 = 0
	bl, err := bil.ToInt32()
	if err != nil {
		return nil, err
	}
	if bl == 0 {
		return RxFromRune('0'), nil
	}
	if bl < 0 {
		bl = rcvr.intlength() + 1
	}
	if bl > rcvr.intlength() {
		firstd = "0"
	} else {
		rx01, err := rcvr.Right(RxFromInt32(bl), ToRxFromRunes([]rune(" ")))
		if err != nil {
			return nil, err
		}
		rx02, err := rx01.Left(RxFromInt8(int8(1)), ToRxFromRunes([]rune(" ")))
		if err != nil {
			return nil, err
		}
		firstd = rx02.ToString()
	}
	rx03 := ToRxFromString(firstd)
	rx04, err := rcvr.Right(RxFromInt32(bl), rx03)
	if err != nil {
		return nil, err
	}
	if bl%4 != 0 {
		ble = bl + 4 - bl%4
	} else {
		ble = bl
	}
	rx05, err := rx04.Right(RxFromInt32(ble), rx03)
	if err != nil {
		return nil, err
	}
	rx06, err := rx05.B2x()
	if err != nil {
		return nil, err
	}
	rx07, err := rx06.X2d(RxFromInt32(ble / 4))
	if err != nil {
		return nil, err
	}
	return rx07, nil
}

// NetRexx: method b2x returns Rexx
//
// Differences: None
func (rcvr *Rexx) B2x() (*Rexx, error) {
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if len(rcvr.chars) == 0 {
		return nil, RxException(1, "No digits")
	}
	res := make([]rune, (len(rcvr.chars)+3)/4)
	j := len(res) - 1
	acc := 0
	mask := 1
	i := len(rcvr.chars) - 1
	for ; i >= 0; i-- {
		if rcvr.chars[i] == '0' {
		} else if rcvr.chars[i] == '1' {
			acc = acc + mask
		} else {
			return nil, RxException(1, ToRxFromRunes([]rune("Bad binary")).OpCcBlank(nil, rcvr).ToString())
		}
		mask = mask + mask
		if mask == 16 || i == 0 {
			res[j] = Hexes[acc]
			j--
			acc = 0
			mask = 1
		}
	}
	return RxFromRunes(res), nil
}

// NetRexx: method c2d returns Rexx
//
// Differences: None
func (rcvr *Rexx) C2d() (*Rexx, error) {
	rcvrchar, err := rcvr.padcheck()
	if err != nil {
		return nil, err
	}
	return RxFromInt64(int64(rcvrchar)), nil
}

// NetRexx: method c2x returns Rexx
//
// Differences: C2x char range is limited from 0 to 65535
func (rcvr *Rexx) C2x() (*Rexx, error) {
	rcvrchar, err := rcvr.padcheck()
	if err != nil {
		return nil, err
	}
	enc := int(rcvrchar)
	if enc < 0 || enc > 65535 {
		// function cannot handle runes > 65535
		return nil, RxException(1, "C2x char range is limited from 0 to 65535")
	}
	res := []rune("0000")
	res[3] = Hexes[enc%16]
	enc = enc / 16
	if enc == 0 {
		rx01, err := RxFromRunes(res).Right(RxFromInt8(int8(1)), ToRxFromRunes([]rune(" ")))
		if err != nil {
			return nil, err
		}
		return rx01, nil
	}
	res[2] = Hexes[enc%16]
	enc = enc / 16
	if enc == 0 {
		rx02, err := RxFromRunes(res).Right(RxFromInt8(int8(2)), ToRxFromRunes([]rune(" ")))
		if err != nil {
			return nil, err
		}
		return rx02, nil
	}
	res[1] = Hexes[enc%16]
	enc = enc / 16
	if enc == 0 {
		rx03, err := RxFromRunes(res).Right(RxFromInt8(int8(3)), ToRxFromRunes([]rune(" ")))
		if err != nil {
			return nil, err
		}
		return rx03, nil
	}
	res[0] = Hexes[enc%16]
	return RxFromRunes(res), nil
}

// NetRexx: method center(wid=Rexx,pad=Rexx " ") returns Rexx
//
// Differences: neither argument wid nor pad can be nil
func (rcvr *Rexx) Center(wid *Rexx, pad *Rexx) (*Rexx, error) {
	if wid == nil || pad == nil {
		// Prevents Go Panic
		return nil, RxException(0, "neither argument wid nor pad can be nil")
	}
	rx01, err := rcvr.Centre(wid, pad)
	if err != nil {
		return nil, err
	}
	return rx01, nil
}

// NetRexx: method centre(wid=Rexx,pad=Rexx " ") returns Rexx
//
// Differences: neither argument wid nor pad can be nil
func (rcvr *Rexx) Centre(wid *Rexx, pad *Rexx) (*Rexx, error) {
	if wid == nil || pad == nil {
		// Prevents Go Panic
		return nil, RxException(0, "neither argument wid nor pad can be nil")
	}
	width, err := wid.intcheck(0, MaxArg)
	if err != nil {
		return nil, err
	}
	padchar, err := pad.padcheck()
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	return ToRxFromRunes(centre(rcvr.chars, width, padchar)), nil
}

// NetRexx: method changestr(old=Rexx,new=Rexx) returns Rexx
//
// Differences: If old or _new is nil, no change occurs.
func (rcvr *Rexx) ChangeStr(old *Rexx, _new *Rexx) *Rexx {
	if old == nil || _new == nil {
		// Prevents Go Panic
		return rcvr
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if old.chars == nil {
		old.chars = old.layout()
	}
	if _new.chars == nil {
		_new.chars = _new.layout()
	}
	return ToRxFromRunes(changestr(old.chars, rcvr.chars, _new.chars))
}

// NetRexx: method charAt(index=int) returns char
//
// Differences: None
func (rcvr *Rexx) CharAt(index int32) (rune, error) {
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if index > int32(len(rcvr.chars))-1 || index < 0 {
		return 0, RxException(1, "index out of range")
	}
	return rcvr.chars[index], nil
}

// NetRexx: method compare(target=Rexx,pad=Rexx " ") returns Rexx
//
// Differences: neither argument target nor pad can be nil
func (rcvr *Rexx) Compare(target *Rexx, pad *Rexx) (*Rexx, error) {
	if target == nil || pad == nil {
		// Prevents Go Panic
		return nil, RxException(0, "neither argument target nor pad can be nil")
	}
	padchar, err := pad.padcheck()
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if target.chars == nil {
		target.chars = target.layout()
	}
	return RxFromInt32(compare(rcvr.chars, target.chars, padchar)), nil
}

// NetRexx: method copies(n=Rexx) returns Rexx signals BadArgumentException
//
// Differences: argument n cannot be nil
func (rcvr *Rexx) Copies(n *Rexx) (*Rexx, error) {
	if n == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument n cannot be nil")
	}
	rep, err := n.intcheck(0, MaxArg)
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	xlen := int32(len(rcvr.chars))
	res := make([]rune, rep*xlen)
	var start int32 = 0
	_15 := rep
	for ; _15 > 0; _15-- {
		copy(res[start:], rcvr.chars[:xlen])
		start = start + xlen
	}
	return rx(res, true), nil
}

// NetRexx: method countstr(b=Rexx) returns Rexx
//
// Differences: None
func (rcvr *Rexx) CountStr(b *Rexx) *Rexx {
	if b == nil {
		return RxFromInt32(0)
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if b.chars == nil {
		b.chars = b.layout()
	}
	return RxFromInt32(countstr(b.chars, rcvr.chars))
}

// NetRexx: method d2b(dil=Rexx "zip") returns Rexx
//
// Differences: argument dil cannot be nil
// NOTE : Use "zip" for default argument like NetRexx
func (rcvr *Rexx) D2b(dil *Rexx) (*Rexx, error) {
	if dil == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument dil cannot be nil")
	}
	var dl int32 = 0
	var dle int32 = 0
	if dil.OpEqS(nil, ToRxFromRunes([]rune("0"))) {
		return ToRxFromRunes([]rune("")), nil
	}
	if dil.OpEqS(nil, ToRxFromRunes([]rune("zip"))) {
		rx01, err := rcvr.D2x(RxFromInt32(-1))
		if err != nil {
			return nil, err
		}
		rx02 := rx01.Length()
		num, err := rx02.ToInt32()
		if err != nil {
			return nil, err
		}
		dl = num * 4
	} else {
		num, err := dil.ToInt32()
		if err != nil {
			return nil, err
		}
		dl = num
	}
	if dl%4 != 0 {
		dle = dl + 4 - dl%4
	} else {
		dle = dl
	}
	rx03, err := rcvr.D2x(RxFromInt32(dle / 4))
	if err != nil {
		return nil, err
	}
	rx04, err := rx03.X2b()
	if err != nil {
		return nil, err
	}
	rx05, err := rx04.Strip(RxFromRune('l'), RxFromRune('0'))
	if err != nil {
		return nil, err
	}
	if dil.OpEqS(nil, ToRxFromRunes([]rune("zip"))) {
		if rx05.OpEqS(nil, ToRxFromRunes([]rune(""))) {
			return RxFromRune('0'), nil
		}
		return rx05, nil
	}
	rx06, err := rx04.Right(RxFromInt32(dl), ToRxFromRunes([]rune(" ")))
	if err != nil {
		return nil, err
	}
	return rx06, nil
}

// NetRexx: method d2c returns Rexx
//
// Differences: None
func (rcvr *Rexx) D2c() (*Rexx, error) {
	num, err := rcvr.ToInt32()
	if err != nil {
		return nil, err
	}
	if num < 0 || num > 65535 {
		return nil, RxException(9, ToRxFromRunes([]rune("Encoding bad")).OpCcBlank(nil, rcvr).ToString())
	}
	return RxFromRune(rune(num)), nil
}

// NetRexx: method d2x(n=Rexx) returns Rexx signals BadArgumentException,NumberFormatException
//
// Differences: argument n cannot be nil, allows -1 to simulate D2x()
func (rcvr *Rexx) D2x(n *Rexx) (*Rexx, error) {
	if n == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument n cannot be nil")
	}
	// allow -1
	num, err := n.intcheck(-1, MaxArg)
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	value, err := d2x(rcvr, num)
	if err != nil {
		return nil, err
	}
	return ToRxFromRunes(value), nil
}

// NetRexx: method datatype(opt=Rexx) returns Rexx signals BadArgumentException
//
// Differences: argument opt cannot be nil
func (rcvr *Rexx) DataType(opt *Rexx) (*Rexx, error) {
	if opt == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument opt cannot be nil")
	}
	var ok int32 = 0
	var set *RexxSet
	ochar, err := opt.optioncheck("ABDLMNSUWX")
	if err != nil {
		return nil, err
	}
	for {
		if rcvr.intlength() == 0 {
			ok = 0
		} else if ochar == 'A' {
			rx01 := ToRxFromString(Lowers + Uppers + Digits09)
			rx02, err := rcvr.Verify(rx01, ToRxFromRunes([]rune("N")), RxFromInt32(1))
			if err != nil {
				return nil, err
			}
			if rx02.OpEqS(nil, zero) {
				ok = 1
			} else {
				ok = 0
			}
		} else if ochar == 'B' {
			rx01 := ToRxFromString("01")
			rx02, err := rcvr.Verify(rx01, ToRxFromRunes([]rune("N")), RxFromInt32(1))
			if err != nil {
				return nil, err
			}
			if rx02.OpEqS(nil, zero) {
				ok = 1
			} else {
				ok = 0
			}
		} else if ochar == 'D' {
			rx01 := ToRxFromString(Digits09)
			rx02, err := rcvr.Verify(rx01, ToRxFromRunes([]rune("N")), RxFromInt32(1))
			if err != nil {
				return nil, err
			}
			if rx02.OpEqS(nil, zero) {
				ok = 1
			} else {
				ok = 0
			}
		} else if ochar == 'L' {
			rx01 := ToRxFromString(Lowers)
			rx02, err := rcvr.Verify(rx01, ToRxFromRunes([]rune("N")), RxFromInt32(1))
			if err != nil {
				return nil, err
			}
			if rx02.OpEqS(nil, zero) {
				ok = 1
			} else {
				ok = 0
			}
		} else if ochar == 'M' {
			rx01 := ToRxFromString(Lowers + Uppers)
			rx02, err := rcvr.Verify(rx01, ToRxFromRunes([]rune("N")), RxFromInt32(1))
			if err != nil {
				return nil, err
			}
			if rx02.OpEqS(nil, zero) {
				ok = 1
			} else {
				ok = 0
			}
		} else if ochar == 'N' {
			if !(rcvr.ind == NotaNum) {
				ok = 1
			} else {
				ok = 0
			}
		} else if ochar == 'S' {
			rx01 := ToRxFromString("_" + Lowers + Uppers + Digits09)
			rx02, err := rcvr.Verify(rx01, ToRxFromRunes([]rune("N")), RxFromInt32(1))
			if err != nil {
				return nil, err
			}
			rx03 := rx02.OpEqS(nil, zero)
			rx04, err := rcvr.Left(RxFromInt8(int8(1)), RxFromString(" "))
			if err != nil {
				return nil, err
			}
			rx05 := ToRxFromString(Digits09)
			rx06, err := rx04.Verify(rx05, ToRxFromRunes([]rune("N")), RxFromInt32(1))
			rx07 := rx06.OpNotEqS(nil, zero)
			if rx03 && rx07 {
				ok = 1
			} else {
				ok = 0
			}
		} else if ochar == 'U' {
			rx01 := ToRxFromString(Uppers)
			rx02, err := rcvr.Verify(rx01, ToRxFromRunes([]rune("N")), RxFromInt32(1))
			if err != nil {
				return nil, err
			}
			if rx02.OpEqS(nil, zero) {
				ok = 1
			} else {
				ok = 0
			}
		} else if ochar == 'W' {
			if rcvr.ind == NotaNum {
				ok = 0
			} else {
				if int32(len(rcvr.mant)) > DefaultDigits {
					set = RxSetWithDigit(int32(len(rcvr.mant)))
				} else {
					set = nil
				}
				rx01, err := rcvr.OpDiv(set, RxFromRune('1'))
				if err != nil {
					return nil, err
				}
				rx02, err := rx01.Pos(RxFromRune('.'), RxFromInt32(1))
				if err != nil {
					return nil, err
				}
				if rx02.OpEqS(nil, zero) {
					ok = 1
				} else {
					ok = 0
				}
			}
		} else if ochar == 'X' {
			rx01 := ToRxFromRunes(Hexes)
			rx02, err := rcvr.Verify(rx01, ToRxFromRunes([]rune("N")), RxFromInt32(1))
			if err != nil {
				return nil, err
			}
			if rx02.OpEqS(nil, zero) {
				ok = 1
			} else {
				ok = 0
			}
		} else {
			// Not possible, opt.optioncheck("ABDLMNSUWX") would have returned an error
			return nil, RxException(6, "")
		}
		if !false {
			break
		}
	}
	return RxFromInt32(ok), nil
}

// NetRexx: method delstr(n=Rexx,length=Rexx this.intlength()) signals BadArgumentException returns Rexx
//
// Differences: neither argument n nor length can be nil
func (rcvr *Rexx) DelStr(n *Rexx, length *Rexx) (*Rexx, error) {
	if n == nil || length == nil {
		// Prevents Go Panic
		return nil, RxException(0, "neither argument n nor length can be nil")
	}
	start, err := n.intcheck(1, MaxArg)
	if err != nil {
		return nil, err
	}
	leng, err := length.intcheck(0, MaxArg)
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	return ToRxFromRunes(delstr(rcvr.chars, start, leng)), nil
}

// NetRexx: method delword(n=Rexx,length=Rexx this.intwords()) signals BadArgumentException returns Rexx
//
// Differences: neither argument n nor length can be nil
func (rcvr *Rexx) DelWord(n *Rexx, length *Rexx) (*Rexx, error) {
	if n == nil || length == nil {
		// Prevents Go Panic
		return nil, RxException(0, "neither argument n nor length can be nil")
	}
	start, err := n.intcheck(1, MaxArg)
	if err != nil {
		return nil, err
	}
	leng, err := length.intcheck(0, MaxArg)
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	return ToRxFromRunes(delword(rcvr.chars, start, leng)), nil
}

// NetRexx: method format(before=Rexx null,after=Rexx null,explaces=Rexx null,exdigits=Rexx null,exform=Rexx null) returns Rexx signals BadArgumentException,NumberFormatException
//
// Differences: None
func (rcvr *Rexx) Format(before *Rexx, after *Rexx, explaces *Rexx, exdigits *Rexx, exform *Rexx) (*Rexx, error) {
	var b int32 = 0
	var a int32 = 0
	var p int32 = 0
	var d int32 = 0
	var f rune
	if rcvr.ind == NotaNum {
		return nil, RxException(9, string(rcvr.chars))
	}
	if before == nil {
		b = -1
	} else {
		value, err := before.intcheck(1, MaxArg)
		if err != nil {
			return nil, err
		}
		b = value
	}
	if after == nil {
		a = -1
	} else {
		value, err := after.intcheck(0, MaxArg)
		if err != nil {
			return nil, err
		}
		a = value
	}
	if explaces == nil {
		p = -1
	} else {
		value, err := explaces.intcheck(1, MaxArg)
		if err != nil {
			return nil, err
		}
		p = value
	}
	if exdigits == nil {
		d = -1
	} else {
		value, err := exdigits.intcheck(0, MaxArg)
		if err != nil {
			return nil, err
		}
		d = value
	}
	if exform == nil {
		f = 'S'
	} else {
		_1, err := exform.optioncheck("SE")
		if err != nil {
			return nil, err
		} else {
			f = _1
		}
	}
	_02, err := format(rcvr, b, a, p, d, f)
	if err != nil {
		return nil, err
	}
	return rx(_02, true), nil
}

// NetRexx: method hashCode returns int
//
// Differences: None
func (rcvr *Rexx) HashCode() int32 {
	over := 0
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if len(rcvr.chars) < 14 {
		over = (len(rcvr.chars) + 1) / 2
	} else {
		over = 7
	}
	hash := 0
	_7 := over - 1
	i := 0
	for ; i <= _7; i++ {
		hash = hash*7 + int(rcvr.chars[i])*2 + int(rcvr.chars[len(rcvr.chars)-i-1])
	}
	return int32(hash)
}

// NetRexx: method insert(new=Rexx,n=Rexx 0,length=Rexx new.intlength(),pad=Rexx " ") signals BadArgumentException returns Rexx
//
// Differences: arguments _new_, n, length and pad cannot be nil
func (rcvr *Rexx) Insert(_new *Rexx, n *Rexx, length *Rexx, pad *Rexx) (*Rexx, error) {
	if n == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument n cannot be nil")
	}
	num, err := n.intcheck(0, MaxArg)
	if err != nil {
		return nil, err
	}
	if length == nil {
		return nil, RxException(0, "argument length cannot be nil")
	}
	leng, err := length.intcheck(0, MaxArg)
	if err != nil {
		return nil, err
	}
	if pad == nil {
		return nil, RxException(0, "argument pad cannot be nil")
	}
	padchar, err := pad.padcheck()
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if _new == nil {
		return nil, RxException(0, "argument _new cannot be nil")
	}
	if _new.chars == nil {
		_new.chars = _new.layout()
	}
	return rx(insert(rcvr.chars, _new.chars, num, leng, padchar), true), nil
}

// NetRexx: method lastpos(needle=Rexx,start=Rexx) returns Rexx signals BadArgumentException
//
// Differences: neither argument needle nor start can be nil
func (rcvr *Rexx) LastPos(needle *Rexx, start *Rexx) (*Rexx, error) {
	var j int32 = 0
	if needle == nil || start == nil {
		// Prevents Go Panic
		return nil, RxException(0, "neither argument needle nor start can be nil")
	}
	num, err := start.intcheck(1, MaxArg)
	if err != nil {
		return nil, err
	}
	startoff := num - 1
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if startoff >= int32(len(rcvr.chars)) {
		startoff = int32(len(rcvr.chars)) - 1
	}
	if needle.chars == nil {
		needle.chars = needle.layout()
	}
	nlength := int32(len(needle.chars))
	if nlength == 0 {
		return RxFromInt32(0), nil
	}
	startoff = startoff - nlength + 1
	i := startoff
i:
	for ; i >= 0; i-- {
		_16 := nlength - 1
		j = 0
		for ; j <= _16; j++ {
			if needle.chars[j] != rcvr.chars[i+j] {
				continue i
			}
		}
		return RxFromInt32(i + 1), nil
	}
	return RxFromInt32(0), nil
}

// NetRexx: method left(length=Rexx,pad=Rexx " ") returns Rexx
//
// Differences: neither argument length nor pad can be nil
func (rcvr *Rexx) Left(length *Rexx, pad *Rexx) (*Rexx, error) {
	if length == nil || pad == nil {
		// Prevents Go Panic
		return nil, RxException(0, "neither argument length nor pad can be nil")
	}
	value, err := rcvr.SubStr(RxFromInt32(1), length, pad)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// NetRexx: method length() returns Rexx
//
// Differences: None
func (rcvr *Rexx) Length() *Rexx {
	return RxFromInt32(rcvr.intlength())
}

// NetRexx: method lower(n=Rexx 1,length=this.length()) signals BadArgumentException returns Rexx
//
// Differences: neither argument n nor pad can be nil
func (rcvr *Rexx) Lower(n *Rexx, length *Rexx) (*Rexx, error) {
	if n == nil || length == nil {
		// Prevents Go Panic
		return nil, RxException(0, "neither argument n nor pad can be nil")
	}
	num, err := n.intcheck(1, MaxArg)
	if err != nil {
		return nil, err
	}
	startoff := num - 1
	leng, err := length.intcheck(0, MaxArg)
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	j := int32(len(rcvr.chars))
	if j == 0 {
		return rxfromempty(), nil
	}
	res := make([]rune, j)
	if leng < j || startoff > 0 {
		copy(res, rcvr.chars)
	}
	_17 := j - 1
	_18 := leng
	i := startoff
	for ; i <= _17 && _18 > 0; i++ {
		res[i] = unicode.ToLower(rcvr.chars[i])
		_18--
	}
	return rx(res, true), nil
}

// NetRexx: method max(rhs=Rexx) returns Rexx signals NumberFormatException
//
// Differences: argument rhs cannot be nil
func (rcvr *Rexx) Max(rhs *Rexx) (*Rexx, error) {
	if rhs == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument rhs cannot be nil")
	}
	var ret *Rexx
	if rcvr.ind == NotaNum {
		return nil, RxException(9, string(rcvr.chars))
	}
	if rhs.ind == NotaNum {
		return nil, RxException(9, string(rcvr.chars))
	}
	leng := DefaultDigits
	if int32(len(rcvr.mant)) > leng {
		leng = int32(len(rcvr.mant))
	}
	if int32(len(rhs.mant)) > leng {
		leng = int32(len(rhs.mant))
	}
	value, err := rcvr.docompare(RxSetWithDigit(leng), rhs)
	if err != nil {
		return nil, err
	}
	if value < 0 {
		ret = rhs
	} else {
		ret = rcvr
	}
	leng = DefaultDigits
	if int32(len(ret.mant)) > leng {
		leng = int32(len(ret.mant))
	}
	return ret.OpPlus(RxSetWithDigit(leng))
}

// NetRexx: method min(rhs=Rexx) returns Rexx signals NumberFormatException
//
// Differences: argument rhs cannot be nil
func (rcvr *Rexx) Min(rhs *Rexx) (*Rexx, error) {
	if rhs == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument rhs cannot be nil")
	}
	var ret *Rexx
	if rcvr.ind == NotaNum {
		return nil, RxException(9, string(rcvr.chars))
	}
	if rhs.ind == NotaNum {
		return nil, RxException(9, string(rcvr.chars))
	}
	leng := DefaultDigits
	if int32(len(rcvr.mant)) > leng {
		leng = int32(len(rcvr.mant))
	}
	if int32(len(rhs.mant)) > leng {
		leng = int32(len(rhs.mant))
	}
	value, err := rcvr.docompare(RxSetWithDigit(leng), rhs)
	if err != nil {
		return nil, err
	}
	if value > 0 {
		ret = rhs
	} else {
		ret = rcvr
	}
	leng = DefaultDigits
	if int32(len(ret.mant)) > leng {
		leng = int32(len(ret.mant))
	}
	return ret.OpPlus(RxSetWithDigit(leng))
}

// NetRexx: method OpAdd(set=RexxSet,rhs=Rexx) returns Rexx
//
// Differences: argument rhs cannot be nil
func (rcvr *Rexx) OpAdd(set *RexxSet, rhs *Rexx) (*Rexx, error) {
	if rhs == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument rhs cannot be nil")
	}
	var resdig int32 = 0
	resform := int8(0)
	usel := []rune(nil)
	user := []rune(nil)
	var newlen int32 = 0
	var i int32 = 0
	var tlen int32 = 0
	ca := rune(0)
	cb := rune(0)
	lhs := rcvr
	if lhs.ind == NotaNum {
		return nil, RxException(9, string(lhs.chars))
	}
	if rhs.ind == NotaNum {
		return nil, RxException(9, string(rhs.chars))
	}
	if set == nil {
		resdig = DefaultDigits
		resform = DefaultForm
	} else {
		resdig = set.Digits
		resform = set.Form
	}
	if lhs.ind == 0 {
		res := rxfromclone(rhs)
		res.chars = nil
		res.dig = resdig
		res.form = resform
		value, err := res.finish(resdig, false)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
	if rhs.ind == 0 {
		res := rxfromclone(lhs)
		res.chars = nil
		res.dig = resdig
		res.form = resform
		value, err := res.finish(resdig, false)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
	if int32(len(lhs.mant)) > resdig+1 {
		lhs = rxfromclone(lhs)
		lhs.cut(resdig)
	}
	if int32(len(rhs.mant)) > resdig+1 {
		rhs = rxfromclone(rhs)
		rhs.cut(resdig)
	}
	res := rxfromempty()
	res.dig = resdig
	res.form = resform
	if lhs.exp == rhs.exp {
		usel = lhs.mant
		user = rhs.mant
		res.exp = lhs.exp
	} else if lhs.exp > rhs.exp {
		newlen = int32(len(lhs.mant)) + lhs.exp - rhs.exp
		if newlen > int32(len(rhs.mant))+resdig+1 {
			res.mant = lhs.mant
			res.exp = lhs.exp
			res.ind = lhs.ind
			if int32(len(res.mant)) < resdig {
				res.mant = make([]rune, resdig)
				copy(res.mant, lhs.mant)
				_28 := resdig - 1
				i = int32(len(lhs.mant))
				for ; i <= _28; i++ {
					res.mant[i] = '0'
				}
				res.exp = res.exp - (resdig - int32(len(lhs.mant)))
			}
			value, err := res.finish(resdig, false)
			if err != nil {
				return nil, err
			}
			return value, nil
		}
		res.exp = rhs.exp
		if newlen > resdig+1 {
			tlen = newlen - resdig - 1
			user = make([]rune, int32(len(rhs.mant))-tlen)
			copy(user, rhs.mant[:len(user)])
			res.exp = res.exp + tlen
			newlen = resdig + 1
		} else {
			user = rhs.mant
		}
		if newlen > int32(len(lhs.mant)) {
			usel = make([]rune, newlen)
			copy(usel, lhs.mant)
			_29 := newlen - 1
			i = int32(len(lhs.mant))
			for ; i <= _29; i++ {
				usel[i] = '0'
			}
		} else {
			usel = lhs.mant
		}
	} else {
		newlen = int32(len(rhs.mant)) + rhs.exp - lhs.exp
		if newlen > int32(len(lhs.mant))+resdig+1 {
			res.mant = rhs.mant
			res.exp = rhs.exp
			res.ind = rhs.ind
			if int32(len(res.mant)) < resdig {
				res.mant = make([]rune, resdig)
				copy(res.mant, rhs.mant)
				_30 := resdig - 1
				i = int32(len(rhs.mant))
				for ; i <= _30; i++ {
					res.mant[i] = '0'
				}
				res.exp = res.exp - (resdig - int32(len(rhs.mant)))
			}
			value, err := res.finish(resdig, false)
			if err != nil {
				return nil, err
			}
			return value, nil
		}
		res.exp = lhs.exp
		if newlen > resdig+1 {
			tlen = newlen - resdig - 1
			usel = make([]rune, int32(len(lhs.mant))-tlen)
			copy(usel, lhs.mant)
			res.exp = res.exp + tlen
			newlen = resdig + 1
		} else {
			usel = lhs.mant
		}
		if newlen > int32(len(rhs.mant)) {
			user = make([]rune, newlen)
			copy(user, rhs.mant)
			_31 := newlen - 1
			i = int32(len(rhs.mant))
			for ; i <= _31; i++ {
				user[i] = '0'
			}
		} else {
			user = rhs.mant
		}
	}
	if lhs.ind == rhs.ind {
		temp := charaddsub(usel, user, 1)
		res.mant = temp
		res.ind = lhs.ind
		value, err := res.finish(resdig, false)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
	res.ind = lhs.ind
	if len(usel) > len(user) {
	} else if len(usel) < len(user) {
		t := usel
		usel = user
		user = t
		res.ind = int8(-res.ind)
	} else {
		ia := 0
		ib := 0
		ea := len(usel) - 1
		eb := len(user) - 1
		for {
			if ia <= ea {
				ca = usel[ia]
			} else {
				if ib > eb {
					res.mant = make([]rune, 1)
					res.mant[0] = '0'
					res.ind = int8(0)
					res.exp = 0
					return res, nil
				}
				ca = '0'
			}
			if ib <= eb {
				cb = user[ib]
			} else {
				cb = '0'
			}
			if ca != cb {
				if ca < cb {
					t := usel
					usel = user
					user = t
					res.ind = int8(-res.ind)
				}
				break
			}
			ia++
			ib++
		}
	}
	tempa := charaddsub(usel, user, -1)
	res.mant = tempa
	if int32(len(res.mant)) > resdig {
		res.round(resdig)
	}
	if res.mant[0] == '0' {
		_32 := int32(len(res.mant))
		i = 1
		for ; i <= _32; i++ {
			if i == int32(len(res.mant)) {
				res.mant = make([]rune, 1)
				res.mant[0] = '0'
				res.ind = int8(0)
				res.exp = 0
				return res, nil
			}
			if res.mant[i] != '0' {
				_new := make([]rune, int32(len(res.mant))-i)
				copy(_new, res.mant[i:])
				res.mant = _new
				break
			}
		}
	}
	return res, nil
}

// NetRexx: method OpAnd(set=RexxSet,rhs=Rexx) returns boolean
//
// Differences: argument rhs cannot be nil
func (rcvr *Rexx) OpAnd(set *RexxSet, rhs *Rexx) (bool, error) {
	if rhs == nil {
		// Prevents Go Panic
		return false, RxException(0, "argument rhs cannot be nil")
	}
	arg1, err := rcvr.ToBool()
	if err != nil {
		return false, err
	}
	arg2, err := rhs.ToBool()
	if err != nil {
		return false, err
	}
	value := arg1 && arg2
	return value, nil
}

// NetRexx: method OpCc(set=RexxSet,rhs=Rexx) returns Rexx
//
// Differences: if rhs == nil, returns rcvr
func (rcvr *Rexx) OpCc(set *RexxSet, rhs *Rexx) *Rexx {
	if rhs == nil {
		// Prevents Go Panic
		return rcvr
	}
	return rcvr.concat(set, rhs, 0)
}

// NetRexx: method OpCcblank(set=RexxSet,rhs=Rexx) returns Rexx
//
// Differences: if rhs == nil, returns rcvr
func (rcvr *Rexx) OpCcBlank(set *RexxSet, rhs *Rexx) *Rexx {
	if rhs == nil {
		// Prevents Go Panic
		return rcvr
	}
	return rcvr.concat(set, rhs, 1)
}

// NetRexx: method OpDiv(set=RexxSet,rhs=Rexx) returns Rexx
//
// Differences: argument rhs cannot be nil
func (rcvr *Rexx) OpDiv(set *RexxSet, rhs *Rexx) (*Rexx, error) {
	if rhs == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument rhs cannot be nil")
	}
	return rcvr.dodivide('D', set, rhs)
}

// NetRexx: method OpDivI(set=RexxSet,rhs=Rexx) returns Rexx
//
// Differences: argument rhs cannot be nil
func (rcvr *Rexx) OpDivI(set *RexxSet, rhs *Rexx) (*Rexx, error) {
	if rhs == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument rhs cannot be nil")
	}
	return rcvr.dodivide('I', set, rhs)
}

// NetRexx: method OpEq(set=RexxSet,rhs=Rexx) returns boolean
//
// Differences: None
func (rcvr *Rexx) OpEq(set *RexxSet, rhs *Rexx) (bool, error) {
	if rhs == nil {
		return false, nil
	}
	value, err := rcvr.docompare(set, rhs)
	if err != nil {
		return false, err
	}
	return value == 0, nil
}

// NetRexx: method OpEqS(set=RexxSet,rhs=Rexx) returns boolean
//
// Differences: None
func (rcvr *Rexx) OpEqS(set *RexxSet, rhs *Rexx) bool {
	if rhs == nil {
		return false
	}
	return rcvr.docomparestrict(set, rhs) == 0
}

// NetRexx: method OpGt(set=RexxSet,rhs=Rexx) returns boolean
//
// Differences: None
func (rcvr *Rexx) OpGt(set *RexxSet, rhs *Rexx) (bool, error) {
	if rhs == nil {
		return true, nil
	}
	value, err := rcvr.docompare(set, rhs)
	if err != nil {
		return false, err
	}
	return value > 0, nil
}

// NetRexx: method OpGtEq(set=RexxSet,rhs=Rexx) returns boolean
//
// Differences: None
func (rcvr *Rexx) OpGtEq(set *RexxSet, rhs *Rexx) (bool, error) {
	if rhs == nil {
		return true, nil
	}
	value, err := rcvr.docompare(set, rhs)
	if err != nil {
		return false, err
	}
	return value >= 0, nil
}

// NetRexx: method OpGtEqS(set=RexxSet,rhs=Rexx) returns boolean
//
// Differences: None
func (rcvr *Rexx) OpGtEqS(set *RexxSet, rhs *Rexx) bool {
	if rhs == nil {
		return true
	}
	return rcvr.docomparestrict(set, rhs) >= 0
}

// NetRexx: method OpGtS(set=RexxSet,rhs=Rexx) returns boolean
//
// Differences: None
func (rcvr *Rexx) OpGtS(set *RexxSet, rhs *Rexx) bool {
	if rhs == nil {
		return true
	}
	return rcvr.docomparestrict(set, rhs) > 0
}

// NetRexx: method OpLt(set=RexxSet,rhs=Rexx) returns boolean
//
// Differences: None
func (rcvr *Rexx) OpLt(set *RexxSet, rhs *Rexx) (bool, error) {
	if rhs == nil {
		return false, nil
	}
	value, err := rcvr.docompare(set, rhs)
	if err != nil {
		return false, err
	}
	return value < 0, nil
}

// NetRexx: method OpLtEq(set=RexxSet,rhs=Rexx) returns boolean
//
// Differences: None
func (rcvr *Rexx) OpLtEq(set *RexxSet, rhs *Rexx) (bool, error) {
	if rhs == nil {
		return false, nil
	}
	value, err := rcvr.docompare(set, rhs)
	if err != nil {
		return false, err
	}
	return value <= 0, nil
}

// NetRexx: method OpLtEqS(set=RexxSet,rhs=Rexx) returns boolean
//
// Differences: None
func (rcvr *Rexx) OpLtEqS(set *RexxSet, rhs *Rexx) bool {
	if rhs == nil {
		return false
	}
	return rcvr.docomparestrict(set, rhs) <= 0
}

// NetRexx: method OpLtS(set=RexxSet,rhs=Rexx) returns boolean
//
// Differences: None
func (rcvr *Rexx) OpLtS(set *RexxSet, rhs *Rexx) bool {
	if rhs == nil {
		return false
	}
	return rcvr.docomparestrict(set, rhs) < 0
}

// NetRexx: method OpMinus(set=RexxSet) returns Rexx
//
// Differences: None
func (rcvr *Rexx) OpMinus(set *RexxSet) (*Rexx, error) {
	if rcvr.ind == NotaNum {
		return nil, RxException(9, string(rcvr.chars))
	}
	res := rxfromclone(rcvr)
	res.ind = int8(-res.ind)
	res.chars = []rune(nil)
	if set == nil {
		res.dig = DefaultDigits
		res.form = DefaultForm
	} else {
		res.dig = set.Digits
		res.form = set.Form
	}
	value, err := res.finish(res.dig, false)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// NetRexx: method OpMult(set=RexxSet,rhs=Rexx) returns Rexx
//
// Differences: argument rhs cannot be nil
func (rcvr *Rexx) OpMult(set *RexxSet, rhs *Rexx) (*Rexx, error) {
	if rhs == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument rhs cannot be nil")
	}
	var resdig int32 = 0
	resform := int8(0)
	multer := []rune(nil)
	mand := []rune(nil)
	var mult int32 = 0
	newmand := []rune(nil)
	lhs := rcvr
	if lhs.ind == NotaNum {
		return nil, RxException(9, string(lhs.chars))
	}
	if rhs.ind == NotaNum {
		return nil, RxException(9, string(rhs.chars))
	}
	if set == nil {
		resdig = DefaultDigits
		resform = DefaultForm
	} else {
		resdig = set.Digits
		resform = set.Form
	}
	if int32(len(lhs.mant)) > resdig+1 {
		lhs = rxfromclone(lhs)
		lhs.cut(resdig)
	}
	if int32(len(rhs.mant)) > resdig+1 {
		rhs = rxfromclone(rhs)
		rhs.cut(resdig)
	}
	res := rxfromempty()
	res.dig = resdig
	res.form = resform
	res.ind = int8(lhs.ind * rhs.ind)
	res.exp = lhs.exp + rhs.exp
	if len(lhs.mant) < len(rhs.mant) {
		multer = lhs.mant
		mand = rhs.mant
	} else {
		multer = rhs.mant
		mand = lhs.mant
	}
	acc := make([]rune, 1)
	acc[0] = '0'
	n := len(multer) - 1
	for ; n >= 0; n-- {
		mult = int32(multer[n]) - int32('0')
		if mult > 0 {
			temp := charaddsub(acc, mand, mult)
			acc = temp
		}
		if n == 0 {
			break
		}
		newmand = make([]rune, len(mand)+1)
		copy(newmand, mand)
		newmand[len(mand)] = '0'
		mand = newmand
	}
	res.mant = acc
	value, err := res.finish(resdig, false)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// NetRexx: method OpNot(set=RexxSet) returns boolean
//
// Differences: None
func (rcvr *Rexx) OpNot(set *RexxSet) (bool, error) {
	arg1, err := rcvr.ToBool()
	if err != nil {
		return false, err
	}
	return !arg1, nil
}

// NetRexx: method OpNotEq(set=RexxSet,rhs=Rexx) returns boolean
//
// Differences: argument rhs cannot be nil
func (rcvr *Rexx) OpNotEq(set *RexxSet, rhs *Rexx) (bool, error) {
	if rhs == nil {
		// Prevents Go Panic
		return false, RxException(0, "argument rhs cannot be nil")
	}
	value, err := rcvr.docompare(set, rhs)
	if err != nil {
		return false, err
	}
	return value != 0, nil
}

// NetRexx: method OpNotEqS(set=RexxSet,rhs=Rexx) returns boolean
//
// Differences: If rhs is nil, return true
func (rcvr *Rexx) OpNotEqS(set *RexxSet, rhs *Rexx) bool {
	if rhs == nil {
		// Prevents Go Panic
		return true
	}
	return rcvr.docomparestrict(set, rhs) != 0
}

// NetRexx: method OpOr(set=RexxSet,rhs=Rexx) returns boolean
//
// Differences: argument rhs cannot be nil
func (rcvr *Rexx) OpOr(set *RexxSet, rhs *Rexx) (bool, error) {
	if rhs == nil {
		// Prevents Go Panic
		return false, RxException(0, "argument rhs cannot be nil")
	}
	arg1, err := rcvr.ToBool()
	if err != nil {
		return false, err
	}
	arg2, err := rhs.ToBool()
	if err != nil {
		return false, err
	}
	value := arg1 || arg2
	return value, nil
}

// NetRexx: method OpPlus(set=RexxSet) returns Rexx
//
// Differences: None
func (rcvr *Rexx) OpPlus(set *RexxSet) (*Rexx, error) {
	if rcvr.ind == NotaNum {
		return nil, RxException(9, string(rcvr.chars))
	}
	res := rxfromclone(rcvr)
	res.chars = nil
	if set == nil {
		res.dig = DefaultDigits
		res.form = DefaultForm
	} else {
		res.dig = set.Digits
		res.form = set.Form
	}
	value, err := res.finish(res.dig, false)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// NetRexx: method OpPow(set=RexxSet,rhs=Rexx) returns Rexx
//
// Differences: argument rhs cannot be nil
func (rcvr *Rexx) OpPow(set *RexxSet, rhs *Rexx) (*Rexx, error) {
	if rhs == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument rhs cannot be nil")
	}
	var d int32 = 0
	var L int32 = 0
	var workset *RexxSet
	if rcvr.ind == NotaNum {
		return nil, RxException(9, string(rcvr.chars))
	}
	if set == nil {
		d = DefaultDigits
	} else {
		d = set.Digits
	}
	if int32(len(rhs.mant)) > d {
		rhs = rxfromclone(rhs)
		rhs.round(d)
	}
	if int32(len(rhs.mant))+rhs.exp > d {
		return nil, RxException(1, "Too many digits: "+rhs.ToString())
	}
	n, err := rhs.intcheck(MinArg, MaxArg)
	if err != nil {
		return nil, err
	}
	lhs := rcvr
	if int32(len(lhs.mant)) > d+1 {
		lhs = rxfromclone(lhs)
		lhs.cut(d)
	}
	one := RxFromInt32(1)
	if rhs.exp == 0 {
		L = int32(len(rhs.mant))
	} else {
		rx01 := RxFromInt32(n).Length()
		value, err := rx01.ToInt32()
		if err != nil {
			return nil, err
		}
		L = value
	}
	if set == nil {
		workset = RxSetWithDigit(DefaultDigits + L + 1)
	} else {
		workset = RxSetWithDigitandForm(set.Digits+L+1, set.Form)
	}
	res := one
	if n == 0 {
		return res, nil
	}
	if n < 0 {
		n = int32(-n)
	}
	lastbit := 31
	seenbit := false
	i := 1
	for ; ; i++ {
		n = n << 1
		if n < 0 {
			seenbit = true
			rx02, err := res.OpMult(workset, lhs)
			if err != nil {
				return nil, err
			}
			res = rx02
		}
		if i == lastbit {
			break
		}
		if !seenbit {
			continue
		}
		rx03, err := res.OpMult(workset, res)
		if err != nil {
			return nil, err
		}
		res = rx03
	}
	if rhs.ind < 0 {
		rx04, err := one.OpDiv(workset, res)
		if err != nil {
			return nil, err
		}
		res = rx04
	}
	rx05, err := res.OpDiv(set, one)
	if err != nil {
		return nil, err
	}
	return rx05, nil
}

// NetRexx: method OpRem(set=RexxSet,rhs=Rexx) returns Rexx
//
// Differences: argument rhs cannot be nil
func (rcvr *Rexx) OpRem(set *RexxSet, rhs *Rexx) (*Rexx, error) {
	if rhs == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument rhs cannot be nil")
	}
	return rcvr.dodivide('R', set, rhs)
}

// NetRexx: method OpSub(set=RexxSet,rhs=Rexx) returns Rexx
//
// Differences: argument rhs cannot be nil
func (rcvr *Rexx) OpSub(set *RexxSet, rhs *Rexx) (*Rexx, error) {
	if rhs == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument rhs cannot be nil")
	}
	if rcvr.ind == NotaNum {
		return nil, RxException(9, string(rcvr.chars))
	}
	if rhs.ind == NotaNum {
		return nil, RxException(9, string(rhs.chars))
	}
	newrhs := rxfromclone(rhs)
	newrhs.ind = int8(-newrhs.ind)
	return rcvr.OpAdd(set, newrhs)
}

// NetRexx: method OpXor(set=RexxSet,rhs=Rexx) returns boolean
//
// Differences: argument rhs cannot be nil
func (rcvr *Rexx) OpXor(set *RexxSet, rhs *Rexx) (bool, error) {
	if rhs == nil {
		// Prevents Go Panic
		return false, RxException(0, "argument rhs cannot be nil")
	}
	arg1, err := rcvr.ToBool()
	if err != nil {
		return false, err
	}
	arg2, err := rhs.ToBool()
	value := arg1 != arg2
	if err != nil {
		return false, err
	}
	return value, nil
}

// NetRexx: method overlay(new=Rexx,n=Rexx 1,length=Rexx new.intlength(),pad=Rexx " ") signals BadArgumentException returns Rexx
//
// Differences:  arguments _new_, n, length and pad cannot be nil
func (rcvr *Rexx) Overlay(_new *Rexx, n *Rexx, length *Rexx, pad *Rexx) (*Rexx, error) {
	if n == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument n cannot be nil")
	}
	num, err := n.intcheck(1, MaxArg)
	if err != nil {
		return nil, err
	}
	if length == nil {
		return nil, RxException(0, "argument length cannot be nil")
	}
	leng, err := length.intcheck(0, MaxArg)
	if err != nil {
		return nil, err
	}
	if pad == nil {
		return nil, RxException(0, "argument pad cannot be nil")
	}
	padchar, err := pad.padcheck()
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if _new == nil {
		return nil, RxException(0, "argument _new cannot be nil")
	}
	if _new.chars == nil {
		_new.chars = _new.layout()
	}
	return rx(overlay(rcvr.chars, _new.chars, num, leng, padchar), true), nil
}

// NetRexx: method pos(needle=Rexx,start=Rexx 1) returns Rexx signals BadArgumentException
//
// Differences: neither argument needle nor start can be nil
func (rcvr *Rexx) Pos(needle *Rexx, start *Rexx) (*Rexx, error) {
	if needle == nil || start == nil {
		// Prevents Go Panic
		return nil, RxException(0, "neither argument needle nor start can be nil")
	}
	var j int32 = 0
	value, err := start.intcheck(1, MaxArg)
	if err != nil {
		return nil, err
	}
	startoff := value - 1
	if needle.chars == nil {
		needle.chars = needle.layout()
	}
	if int32(len(needle.chars)) == 0 {
		return RxFromInt32(0), nil
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	_19 := int32(len(rcvr.chars)) - int32(len(needle.chars))
	i := startoff
i:
	for ; i <= _19; i++ {
		_20 := int32(len(needle.chars)) - 1
		j = 0
		for ; j <= _20; j++ {
			if needle.chars[j] != rcvr.chars[i+j] {
				continue i
			}
		}
		return RxFromInt32(i + 1), nil
	}
	return RxFromInt32(0), nil
}

// NetRexx: method reverse() returns Rexx
//
// Differences: None
func (rcvr *Rexx) Reverse() (*Rexx, error) {
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	j := len(rcvr.chars)
	if j == 0 {
		return RxFromString(""), nil
	}
	res := make([]rune, j)
	_21 := j - 1
	i := 0
	for ; i <= _21; i++ {
		j--
		res[i] = rcvr.chars[j]
	}
	return rx(res, true), nil
}

// NetRexx: method right(length=Rexx,pad=Rexx " ") returns Rexx signals NotCharacterException,BadArgumentException
//
// Differences: argument length and pad cannot be nil
func (rcvr *Rexx) Right(length *Rexx, pad *Rexx) (*Rexx, error) {
	if length == nil || pad == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument length cannot be nil")
	}
	leng, err := length.intcheck(0, MaxArg)
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	trim := int32(len(rcvr.chars)) - leng
	if trim >= 0 {
		n := RxFromInt32(trim + 1)
		check, err := n.ToInt32()
		if err != nil {
			return nil, err
		}
		x := RxFromInt32(rcvr.intlength() + 1 - check)
		length, err := x.Max(RxFromInt8(int8(0)))
		if err != nil {
			return nil, err
		}
		value, err := rcvr.SubStr(n, length, ToRxFromRunes([]rune(" ")))
		if err != nil {
			return nil, err
		}
		return value, nil
	}
	padchar, err := pad.padcheck()
	if err != nil {
		return nil, err
	}
	res := make([]rune, leng)
	_22 := int(-trim) - 1
	i := 0
	for ; i <= _22; i++ {
		res[i] = padchar
	}
	copy(res[i:], rcvr.chars[0:len(res)-i])
	return rx(res, true), nil
}

// NetRexx: method sequence(final=Rexx) returns Rexx signals BadArgumentException,NotCharacterException
//
// Differences: argument _final cannot be nil
func (rcvr *Rexx) Sequence(_final *Rexx) (*Rexx, error) {
	if _final == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument _final cannot be nil")
	}
	startchar, err := rcvr.padcheck()
	if err != nil {
		return nil, err
	}
	finalchar, err := _final.padcheck()
	if err != nil {
		return nil, err
	}
	istart := int(startchar)
	leng := int(finalchar) - istart + 1
	if leng <= 0 {
		return nil, RxException(1, "final<start")
	}
	car := make([]rune, leng)
	_23 := leng - 1
	i := 0
	for ; i <= _23; i++ {
		car[i] = rune(i + istart)
	}
	return rx(car, true), nil
}

// NetRexx: method sign() returns Rexx signals NumberFormatException
//
// Differences: None
func (rcvr *Rexx) Sign() (*Rexx, error) {
	if rcvr.ind == NotaNum {
		return nil, RxException(9, string(rcvr.chars))
	}
	return RxFromInt8(rcvr.ind), nil
}

// NetRexx: method significance returns int
//
// Differences: None
func (rcvr *Rexx) Significance() int32 {
	if rcvr.ind == NotaNum {
		return 0
	}
	return int32(len(rcvr.mant))
}

// NetRexx: method space(n=Rexx 1,pad=Rexx " ") returns Rexx
//
// Differences: neither argument n nor pad can be nil
func (rcvr *Rexx) Space(n *Rexx, pad *Rexx) (*Rexx, error) {
	if n == nil || pad == nil {
		// Prevents Go Panic
		return nil, RxException(0, "neither argument n nor pad can be nil")
	}
	gap, err := n.intcheck(0, MaxArg)
	if err != nil {
		return nil, err
	}
	padchar, err := pad.padcheck()
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	return ToRxFromRunes(space(rcvr.chars, gap, padchar)), nil
}

// NetRexx: method strip(opt=Rexx "B",pad=Rexx " ") returns Rexx signals BadArgumentException
//
// Differences: neither argument opt nor pad can be nil
func (rcvr *Rexx) Strip(opt *Rexx, pad *Rexx) (*Rexx, error) {
	if opt == nil || pad == nil {
		// Prevents Go Panic
		return nil, RxException(0, "neither argument opt nor pad can be nil")
	}
	startoff := 0
	endoff := 0
	ochar, err := opt.optioncheck("BLT")
	if err != nil {
		return nil, err
	}
	padchar, err := pad.padcheck()
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if len(rcvr.chars) == 0 {
		return RxFromString(""), nil
	}
	if ochar == 'T' {
		startoff = 0
	} else {
		_24 := len(rcvr.chars) - 1
		startoff = 0
		for ; startoff <= _24; startoff++ {
			if rcvr.chars[startoff] != padchar {
				break
			}
		}
		if startoff == len(rcvr.chars) {
			return RxFromString(""), nil
		}
	}
	if ochar == 'L' {
		endoff = len(rcvr.chars) - 1
	} else {
		endoff = len(rcvr.chars) - 1
		for ; endoff >= 0; endoff-- {
			if rcvr.chars[endoff] != padchar {
				break
			}
		}
		if endoff < 0 {
			return RxFromString(""), nil
		}
	}
	if startoff == 0 {
		if endoff == len(rcvr.chars)-1 {
			return rxfromclone(rcvr), nil
		}
	}
	leng := endoff - startoff + 1
	subchars := make([]rune, leng)
	copy(subchars, rcvr.chars[startoff:])
	return rx(subchars, true), nil
}

// NetRexx: method substr(n=Rexx,length=Rexx(this.intlength()+1-n.toint()).max(0),pad=Rexx " ") signals NotCharacterException,BadArgumentException returns Rexx
//
// Differences: arguments n, length or pad cannot be nil
func (rcvr *Rexx) SubStr(n *Rexx, length *Rexx, pad *Rexx) (*Rexx, error) {
	if n == nil || length == nil || pad == nil {
		// Prevents Go Panic
		return nil, RxException(0, "arguments n, length or pad cannot be nil")
	}
	value, err := n.intcheck(1, MaxArg)
	if err != nil {
		return nil, err
	}
	startoff := value - 1
	leng, err := length.intcheck(0, MaxArg)
	if err != nil {
		return nil, err
	}
	padchar, err := pad.padcheck()
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	thislen := int32(len(rcvr.chars))
	subchars := make([]rune, leng)
	if startoff+leng <= thislen {
		copy(subchars, rcvr.chars[startoff:])
	} else {
		if startoff < thislen {
			copy(subchars, rcvr.chars[startoff:])
		} else {
			startoff = thislen
		}
		_25 := leng - 1
		i := thislen - startoff
		for ; i <= _25; i++ {
			subchars[i] = padchar
		}
	}
	return rx(subchars, true), nil
}

// NetRexx: method subword(n=Rexx,length=this.length()) returns Rexx
//
// Differences: neither argument n nor length can be nil
func (rcvr *Rexx) SubWord(n *Rexx, length *Rexx) (*Rexx, error) {
	if n == nil || length == nil {
		// Prevents Go Panic
		return nil, RxException(0, "neither argument n nor length can be nil")
	}
	start, err := n.intcheck(1, MaxArg)
	if err != nil {
		return nil, err
	}
	leng, err := length.intcheck(0, MaxArg)
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	return ToRxFromRunes(subword(rcvr.chars, start, leng)), nil
}

// NetRexx: method toboolean() returns boolean signals NotLogicException
//
// Differences: None
func (rcvr *Rexx) ToBool() (bool, error) {
	if rcvr.ind == iszero {
		return false, nil
	}
	if rcvr.ind == ispos {
		if len(rcvr.mant) == 1 {
			if rcvr.mant[0] == '1' {
				return true, nil
			}
		}
	}
	return false, RxException(8, "Boolean must be 0 or 1.  Found: "+rcvr.ToString())
}

// NetRexx: method tofloat() returns float signals NumberFormatException
//
// Differences: None
func (rcvr *Rexx) ToFloat32() (float32, error) {
	var dub float64
	dub, err := rcvr.ToFloat64()
	if err != nil {
		return 0, err
	}
	if dub > math.MaxFloat32 || dub < math.SmallestNonzeroFloat32 {
		return 0, RxException(9, "Overflow")
	}
	return float32(dub), nil
}

// NetRexx: method todouble() returns double signals NumberFormatException
//
// Differences: None
func (rcvr *Rexx) ToFloat64() (float64, error) {
	if rcvr.ind == NotaNum {
		return 0, RxException(9, string(rcvr.chars))
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	dub, err := strconv.ParseFloat(string(rcvr.chars), 64)
	if err != nil {
		if math.IsInf(dub, 0) {
			return 0, RxException(9, "Overflow")
		} else {
			return 0, err
		}
	}
	return dub, nil
}

// NetRexx: method toshort() returns short signals NumberFormatException
//
// Differences: None
func (rcvr *Rexx) ToInt16() (int16, error) {
	num, err := rcvr.ToInt32()
	if err != nil {
		return 0, err
	}
	if num > 32767 || num < -32768 {
		return 0, RxException(9, "Conversion overflow")
	}
	return int16(num), nil
}

// NetRexx: method toint() returns int signals NumberFormatException
//
// Differences: None
func (rcvr *Rexx) ToInt32() (int32, error) {
	var cstart int32 = 0
	var useexp int32 = 0
	if rcvr.ind == NotaNum {
		return 0, RxException(9, string(rcvr.chars))
	}
	if rcvr.ind == iszero {
		return 0, nil
	}
	lodigit := int32(len(rcvr.mant) - 1)
	if rcvr.exp < 0 {
		lodigit = lodigit + rcvr.exp
		if lodigit < 0 {
			cstart = 0
		} else {
			cstart = lodigit + 1
		}
		_8 := int32(len(rcvr.mant) - 1)
		j := cstart
		for ; j <= _8; j++ {
			if rcvr.mant[j] != '0' {
				return 0, RxException(9, ToRxFromRunes([]rune("Decimal part non-zero:")).OpCcBlank(nil, rcvr).ToString())
			}
		}
		if lodigit < 0 {
			return 0, nil
		}
		useexp = 0
	} else {
		if rcvr.exp+int32(len(rcvr.mant)) > 10 {
			return 0, RxException(9, "Conversion overflow")
		}
		useexp = rcvr.exp
	}
	var result int32 = 0
	lastresult := result
	_9 := lodigit + useexp
	var i int32 = 0
	for ; i <= _9; i++ {
		result = result * 10
		if i <= lodigit {
			result = result + (int32(rcvr.mant[i]) - int32('0'))
		}
		if result < lastresult {
			if rcvr.ind == isneg {
				if result == math.MinInt32 {
					if i == lodigit+useexp {
						return result, nil
					}
				}
			}
			return 0, RxException(9, "Conversion overflow")
		}
		lastresult = result
	}
	if rcvr.ind > 0 {
		return result, nil
	}
	return int32(-result), nil
}

// NetRexx: method tolong() returns long signals NumberFormatException
//
// Differences: None
func (rcvr *Rexx) ToInt64() (int64, error) {
	var cstart int32 = 0
	var useexp int32 = 0
	if rcvr.ind == NotaNum {
		return 0, RxException(9, string(rcvr.chars))
	}
	if rcvr.ind == 0 {
		return 0, nil
	}
	lodigit := int32(len(rcvr.mant)) - 1
	if rcvr.exp < 0 {
		lodigit = lodigit + rcvr.exp
		if lodigit < 0 {
			cstart = 0
		} else {
			cstart = lodigit + 1
		}
		_10 := int32(len(rcvr.mant)) - 1
		j := cstart
		for ; j <= _10; j++ {
			if rcvr.mant[j] != '0' {
				return 0, RxException(9, ToRxFromRunes([]rune("Decimal part non-zero:")).OpCcBlank(nil, rcvr).ToString())
			}
		}
		if lodigit < 0 {
			return 0, nil
		}
		useexp = 0
	} else {
		if rcvr.exp+int32(len(rcvr.mant)) >= 20 {
			return 0, RxException(9, "Conversion overflow")
		}
		useexp = rcvr.exp
	}
	result := int64(0)
	lastresult := result
	_11 := lodigit + useexp
	var i int32 = 0
	for ; i <= _11; i++ {
		result = result * 10
		if i <= lodigit {
			result = result + (int64(rcvr.mant[i]) - int64('0'))
		}
		if result < lastresult {
			if rcvr.ind < 0 {
				if result == math.MinInt64 {
					if i == lodigit+useexp {
						return int64(result), nil
					}
				}
			}
			return 0, RxException(9, "Conversion overflow")
		}
		lastresult = result
	}
	if rcvr.ind > 0 {
		return int64(result), nil
	}
	return int64(-result), nil
}

// NetRexx: method tobyte() returns byte signals NumberFormatException
//
// Differences: None
func (rcvr *Rexx) ToInt8() (int8, error) {
	num, err := rcvr.ToInt32()
	if err != nil {
		return 0, err
	}
	if num > 127 || num < -128 {
		return 0, RxException(9, "Conversion overflow")
	}
	return int8(num), nil
}

// NetRexx: method tochar() returns char signals NotCharacterException
//
// Differences: None
func (rcvr *Rexx) ToRune() (rune, error) {
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if len(rcvr.chars) != 1 {
		return 0, RxException(7, string(rcvr.chars))
	}
	return rcvr.chars[0], nil
}

// NetRexx: method toCharArray() returns char[]
//
// Differences: None
func (rcvr *Rexx) ToRunes() []rune {
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	res := make([]rune, len(rcvr.chars))
	copy(res, rcvr.chars)
	return res
}

// NetRexx: method toString() returns String
//
// Differences: None
func (rcvr *Rexx) ToString() string {
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	value := string(rcvr.chars)
	return value
}

// NetRexx: method translate(tableo=Rexx,tablei=Rexx,pad=Rexx " ") signals BadArgumentException returns Rexx
//
// Differences: arguments tableo, tablei or pad cannot be nil
func (rcvr *Rexx) Translate(tableo *Rexx, tablei *Rexx, pad *Rexx) (*Rexx, error) {
	if tableo == nil || tablei == nil || pad == nil {
		// Prevents Go Panic
		return nil, RxException(0, "arguments tableo, tablei or pad cannot be nil")
	}
	padchar, err := pad.padcheck()
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if tableo.chars == nil {
		tableo.chars = tableo.layout()
	}
	if tablei.chars == nil {
		tablei.chars = tablei.layout()
	}
	return ToRxFromRunes(translate(rcvr.chars, tableo.chars, tablei.chars, padchar)), nil
}

// NetRexx: method trunc(n=Rexx 0) returns Rexx signals BadArgumentException,NumberFormatException
//
// Differences: argument n cannot be nil
func (rcvr *Rexx) Trunc(n *Rexx) (*Rexx, error) {
	if n == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument n cannot be nil")
	}
	if rcvr.ind == NotaNum {
		return nil, RxException(9, string(rcvr.chars))
	}
	after, err := n.intcheck(0, MaxArg)
	if err != nil {
		return nil, err
	}
	value, err := trunc(rcvr, after)
	if err != nil {
		return nil, err
	}
	return RxFromRunes(value), nil
}

// NetRexx: method upper(n=Rexx 1,length=this.length()) signals BadArgumentException returns Rexx
//
// Differences: neither argument n nor length can be nil
func (rcvr *Rexx) Upper(n *Rexx, length *Rexx) (*Rexx, error) {
	if n == nil || length == nil {
		// Prevents Go Panic
		return nil, RxException(0, "neither argument n nor length can be nil")
	}
	value, err := n.intcheck(1, MaxArg)
	if err != nil {
		return nil, err
	}
	startoff := value - 1
	leng, err := length.intcheck(0, MaxArg)
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	j := int32(len(rcvr.chars))
	if j == 0 {
		return RxFromString(""), nil
	}
	res := make([]rune, j)
	if leng < j || startoff > 0 {
		copy(res, rcvr.chars)
	}
	_26 := j - 1
	_27 := leng
	i := startoff
	for ; i <= _26 && _27 > 0; i++ {
		res[i] = unicode.ToUpper(rcvr.chars[i])
		_27--
	}
	return rx(res, true), nil
}

// NetRexx: method verify(list=Rexx,opt=Rexx "N",start=Rexx 1) returns Rexx
//
// Differences: arguments list, opt or start cannot be nil
func (rcvr *Rexx) Verify(list *Rexx, opt *Rexx, start *Rexx) (*Rexx, error) {
	if list == nil || opt == nil || start == nil {
		// Prevents Go Panic
		return nil, RxException(0, "arguments list, opt or start cannot be nil")
	}
	ochar, err := opt.optioncheck("NM")
	if err != nil {
		return nil, err
	}
	from, err := start.intcheck(1, MaxArg)
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if list.chars == nil {
		list.chars = list.layout()
	}
	if ochar == 'N' {
		return RxFromInt32(verifyn(rcvr.chars, list.chars, from)), nil
	}
	return RxFromInt32(verifym(rcvr.chars, list.chars, from)), nil
}

// NetRexx: method word(n=Rexx) returns Rexx
//
// Differences: argument n cannot be nil
func (rcvr *Rexx) Word(n *Rexx) (*Rexx, error) {
	if n == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument n cannot be nil")
	}
	value, err := rcvr.SubWord(n, RxFromInt8(int8(1)))
	if err != nil {
		return nil, err
	}
	return value, nil
}

// NetRexx: method wordindex(n=Rexx) returns Rexx
//
// Differences: argument n cannot be nil
func (rcvr *Rexx) WordIndex(n *Rexx) (*Rexx, error) {
	if n == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument n cannot be nil")
	}
	from, err := n.intcheck(1, MaxArg)
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	return RxFromInt32(wordindex(rcvr.chars, from)), nil
}

// NetRexx: method wordlength(n=Rexx) returns Rexx
//
// Differences: argument n cannot be nil
func (rcvr *Rexx) WordLength(n *Rexx) (*Rexx, error) {
	if n == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument n cannot be nil")
	}
	from, err := n.intcheck(1, MaxArg)
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	return RxFromInt32(wordlength(rcvr.chars, from)), nil
}

// NetRexx: method wordpos(needle=Rexx,num=Rexx 1) returns Rexx
//
// Differences: neither argument needle nor num can be nil
func (rcvr *Rexx) WordPos(needle *Rexx, num *Rexx) (*Rexx, error) {
	if needle == nil || num == nil {
		// Prevents Go Panic
		return nil, RxException(0, "neither argument needle nor num can be nil")
	}
	n, err := num.intcheck(1, MaxArg)
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if needle.chars == nil {
		needle.chars = needle.layout()
	}
	return RxFromInt32(wordpos(needle.chars, rcvr.chars, n)), nil
}

// NetRexx: method words() returns Rexx
//
// Differences: None
func (rcvr *Rexx) Words() (*Rexx, error) {
	return RxFromInt32(rcvr.intwords()), nil
}

// NetRexx: method x2b returns Rexx
//
// Differences: None
func (rcvr *Rexx) X2b() (*Rexx, error) {
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if len(rcvr.chars) == 0 {
		return nil, RxException(1, "No digits")
	}
	value, err := x2b(rcvr)
	if err != nil {
		return nil, err
	}
	return RxFromRunes(value), nil
}

// NetRexx: method x2c returns Rexx
//
// Differences: None
func (rcvr *Rexx) X2c() (*Rexx, error) {
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if len(rcvr.chars) == 0 {
		return nil, RxException(1, "No digits")
	}
	value, err := x2c(rcvr)
	if err != nil {
		return nil, err
	}
	return RxFromRune(value), nil
}

// NetRexx: method x2d(n=Rexx) returns Rexx signals BadArgumentException
//
// Differences: argument n cannot be nil, allows -1 to simulate X2d()
func (rcvr *Rexx) X2d(n *Rexx) (*Rexx, error) {
	if n == nil {
		// Prevents Go Panic
		return nil, RxException(0, "argument n cannot be nil")
	}
	// allow -1
	num, err := n.intcheck(-1, MaxArg)
	if err != nil {
		return nil, err
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	value, err := x2d(rcvr, num)
	if err != nil {
		return nil, err
	}
	return ToRxFromRunes(value), nil
}

// NetRexx: method concat(set=RexxSet,rhs=Rexx,blanks=int) private returns Rexx
//
// Differences: None
func (rcvr *Rexx) concat(set *RexxSet, rhs *Rexx, blanks int32) *Rexx {
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if rhs.chars == nil {
		rhs.chars = rhs.layout()
	}
	res := make([]rune, int32(len(rcvr.chars))+int32(len(rhs.chars))+blanks)
	copy(res, rcvr.chars)
	if blanks > 0 {
		_37 := int32(len(rcvr.chars)) + blanks - 1
		i := int32(len(rcvr.chars))
		for ; i <= _37; i++ {
			res[i] = ' '
		}
	}
	copy(res[int32(len(rcvr.chars))+blanks:], rhs.chars)
	return rx(res, true)
}

// NetRexx: method cut(d=int) private
//
// Differences: None
func (rcvr *Rexx) cut(d int32) {
	adjust := int32(len(rcvr.mant)) - d - 1
	if adjust <= 0 {
		return
	}
	rcvr.exp = rcvr.exp + adjust
	use := make([]rune, d+1)
	copy(use, rcvr.mant)
	rcvr.mant = use
	rcvr.chars = nil
	return
}

// NetRexx: method docompare(set=RexxSet,rhs=Rexx) private final returns int
//
// Differences: None
func (rcvr *Rexx) docompare(set *RexxSet, rhs *Rexx) (int8, error) {
	cl := rune(0)
	cr := rune(0)
	if rcvr.ind != NotaNum {
		if rhs.ind != NotaNum {
			for {
				if rcvr.ind == rhs.ind && rcvr.exp == rhs.exp {
					thislength := int32(len(rcvr.mant))
					if thislength < int32(len(rhs.mant)) {
						return int8(-rcvr.ind), nil
					}
					if thislength > int32(len(rhs.mant)) {
						return rcvr.ind, nil
					}
					rcvr.dig = DefaultDigits
					if set != nil {
						rcvr.dig = set.Digits
					}
					if thislength <= rcvr.dig {
						_40 := thislength - 1
						var i int32 = 0
						for ; i <= _40; i++ {
							if rcvr.mant[i] < rhs.mant[i] {
								return int8(-rcvr.ind), nil
							}
							if rcvr.mant[i] > rhs.mant[i] {
								return rcvr.ind, nil
							}
						}
						return 0, nil
					}
				} else {
					if rcvr.ind < rhs.ind {
						return -1, nil
					}
					if rcvr.ind > rhs.ind {
						return 1, nil
					}
				}
				newrhs := rxfromclone(rhs)
				newrhs.ind = int8(-newrhs.ind)
				res, err := rcvr.OpAdd(set, newrhs)
				if err != nil {
					return 0, err
				}
				return res.ind, nil
			}
		}
	}
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	sl := rcvr.chars
	if rhs.chars == nil {
		rhs.chars = rhs.layout()
	}
	sr := rhs.chars
	il := 0
	ir := 0
	el := len(sl) - 1
	er := len(sr) - 1
	for ; el >= 0; el-- {
		if !(sl[el] == ' ') {
			break
		}
	}
	for ; er >= 0; er-- {
		if !(sr[er] == ' ') {
			break
		}
	}
	_41 := el
	il = 0
	for ; il <= _41; il++ {
		if !(sl[il] == ' ') {
			break
		}
	}
	_42 := er
	ir = 0
	for ; ir <= _42; ir++ {
		if !(sr[ir] == ' ') {
			break
		}
	}
_43:
	for {
		if il <= el {
			cl = sl[il]
		} else {
			cl = ' '
		}
		if ir <= er {
			cr = sr[ir]
		} else {
			cr = ' '
		}
		if cr == cl {
			if cr == ' ' {
				if il > el {
					if ir > er {
						break _43
					}
				}
			}
			il++
			ir++
			continue _43
		}
		cl = unicode.ToLower(cl)
		cr = unicode.ToLower(cr)
		if cr != cl {
			if cl < cr {
				return -1, nil
			}
			return 1, nil
		}
		il++
		ir++
	}
	return 0, nil
}

// NetRexx: method docomparestrict(set=RexxSet,rhs=Rexx) private final returns int
//
// Differences: None
func (rcvr *Rexx) docomparestrict(set *RexxSet, rhs *Rexx) int32 {
	cl := rune(0)
	cr := rune(0)
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	sl := rcvr.chars
	if rhs.chars == nil {
		rhs.chars = rhs.layout()
	}
	sr := rhs.chars
	il := 0
	ir := 0
	el := len(sl) - 1
	er := len(sr) - 1
_44:
	for {
		if il <= el {
			cl = sl[il]
		} else {
			if ir > er {
				break _44
			}
			cl = '\000'
		}
		if ir <= er {
			cr = sr[ir]
		} else {
			cr = '\000'
		}
		if cr != cl {
			if cl < cr {
				return -1
			}
			return 1
		}
		il++
		ir++
	}
	return 0
}

// NetRexx: method dodivide(code=char,set=RexxSet,rhs=Rexx) private final returns Rexx signals NumberFormatException,DivideException
//
// Differences: None
func (rcvr *Rexx) dodivide(code rune, set *RexxSet, rhs *Rexx) (*Rexx, error) {
	var resdig int32 = 0
	resform := int8(0)
	var res *Rexx
	var ca int32 = 0
	var d int32 = 0
	if rcvr.ind == NotaNum {
		return nil, RxException(9, string(rcvr.chars))
	}
	if rhs.ind == NotaNum {
		return nil, RxException(9, string(rhs.chars))
	}
	if rhs.ind == 0 {
		return nil, RxException(4, "Divide by 0")
	}
	if set == nil {
		resdig = DefaultDigits
		resform = DefaultForm
	} else {
		resdig = set.Digits
		resform = set.Form
	}
	lhs := rcvr
	if lhs.ind == 0 {
		return RxFromInt32(0), nil
	}
	if int32(len(lhs.mant)) > resdig+1 {
		lhs = rxfromclone(lhs)
		lhs.cut(resdig)
	}
	if int32(len(rhs.mant)) > resdig+1 {
		rhs = rxfromclone(rhs)
		rhs.cut(resdig)
	}
	newexp := lhs.exp - rhs.exp + int32(len(lhs.mant)) - int32(len(rhs.mant))
	if newexp < 0 {
		if code != 'D' {
			if code == 'I' {
				return RxFromInt32(0), nil
			}
			res = rxfromclone(lhs)
			res.dig = resdig
			res.form = resform
			value, err := res.finish(resdig, false)
			if err != nil {
				return nil, err
			}
			return value, nil
		}
	}
	res = rxfromempty()
	res.ind = int8(lhs.ind * rhs.ind)
	res.exp = newexp
	res.dig = resdig
	res.form = resform
	res.mant = make([]rune, resdig+1)
	newlen := resdig + resdig + 1
	var1 := make([]rune, newlen)
	copy(var1, lhs.mant)
	_33 := newlen - 1
	i := int32(len(lhs.mant))
	for ; i <= _33; i++ {
		var1[i] = '0'
	}
	var2 := make([]rune, newlen)
	copy(var2, rhs.mant)
	_34 := newlen - 1
	i = int32(len(rhs.mant))
	for ; i <= _34; i++ {
		var2[i] = '0'
	}
	c2b := (int32(var2[0])-int32('0'))*10 + int32(var2[1]) - int32('0') + 1
	var have int32 = 0
outer:
	for {
		var thisdigit int32 = 0
	inner:
		for {
			if len(var1) < len(var2) {
				break inner
			}
			if len(var1) == len(var2) {
			compare:
				for {
					_35 := int32(len(var1)) - 1
					var i int32 = 0
					for ; i <= _35; i++ {
						if var1[i] < var2[i] {
							break inner
						}
						if var1[i] > var2[i] {
							break compare
						}
					}
					if code != 'R' {
						thisdigit++
						res.mant[have] = rune(thisdigit + int32('0'))
						have++
						break outer
					}
					if !false {
						break
					}
				}
				ca = int32(var1[0]) - int32('0')
			} else {
				ca = (int32(var1[0]) - int32('0')) * 10
				if len(var1) > 1 {
					ca = ca + int32(var1[1]) - int32('0')
				}
			}
			mult := ca * 10 / c2b
			if mult == 0 {
				mult = 1
			}
			thisdigit = thisdigit + mult
			var1 = charaddsub(var1, var2, int32(-mult))
			if var1[0] != '0' {
				continue inner
			}
			d = int32(len(var1))
			_36 := d - 2
			var start int32 = 0
		start:
			for ; start <= _36; start++ {
				if var1[start] != '0' {
					break start
				}
				d--
			}
			reduced := make([]rune, d)
			copy(reduced, var1[start:])
			var1 = reduced
		}
		if have != 0 || thisdigit != 0 {
			res.mant[have] = rune(thisdigit + int32('0'))
			have++
			if have == resdig+1 {
				break outer
			}
			if var1[0] == '0' {
				break outer
			}
		}
		if code != 'D' {
			if res.exp <= 0 {
				break outer
			}
		}
		res.exp = res.exp - 1
		newvar2 := make([]rune, len(var2)-1)
		copy(newvar2, var2)
		var2 = newvar2
	}
	if code == 'I' || code == 'R' {
		if have+res.exp > resdig {
			return nil, RxException(4, "Integer overflow")
		}
	}
	if code != 'R' {
		if have == 0 {
			return RxFromInt32(0), nil
		}
		if have == int32(len(res.mant)) {
			res.round(resdig)
			have = resdig
		}
		newmant := make([]rune, have)
		copy(newmant, res.mant)
		res.mant = newmant
		value, err := res.finish(resdig, true)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
	if have == 0 {
		res = rxfromclone(lhs)
		res.chars = nil
		res.dig = resdig
		res.form = resform
		value, err := res.finish(resdig, false)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
	if var1[0] == '0' {
		return RxFromInt32(0), nil
	}
	res.ind = lhs.ind
	padding := resdig + resdig + 1 - int32(len(lhs.mant))
	res.exp = res.exp - padding + lhs.exp
	d = int32(len(var1))
	i = d - 1
i:
	for ; i >= 1; i-- {
		if !(res.exp < lhs.exp && res.exp < rhs.exp) {
			break
		}
		if var1[i] != '0' {
			break i
		}
		d--
		res.exp = res.exp + 1
	}
	if d < int32(len(var1)) {
		newvar1 := make([]rune, d)
		copy(newvar1, var1)
		var1 = newvar1
	}
	res.mant = var1
	value, err := res.finish(resdig, false)
	if err != nil {
		return nil, err
	}
	return value, nil
}

// NetRexx: method finish(roundlen=int,strip=boolean) private final returns Rexx
//
// Differences: None
func (rcvr *Rexx) finish(roundlen int32, strip bool) (*Rexx, error) {
	if int32(len(rcvr.mant)) > roundlen {
		rcvr.round(roundlen)
	}
	if strip {
		d := len(rcvr.mant)
		i := d - 1
		for ; i >= 1; i-- {
			if rcvr.mant[i] != '0' {
				break
			}
			d--
			rcvr.exp++
		}
		if d < len(rcvr.mant) {
			newmant := make([]rune, d)
			copy(newmant, rcvr.mant)
			rcvr.mant = newmant
		}
	}
	_45 := len(rcvr.mant) - 1
	i := 0
	for ; i <= _45; i++ {
		if rcvr.mant[i] != '0' {
			mag := rcvr.exp + int32(len(rcvr.mant)) - 1
			if mag < MinExp || mag > MaxExp {
			overflow:
				for {
					if rcvr.form == ENGINEERING {
						sig := mag % 3
						if sig < 0 {
							sig = 3 + sig
						}
						mag = mag - sig
						if mag >= MinExp {
							if mag <= MinExp {
								break overflow
							}
						}
					}
					return nil, RxException(5, strconv.Itoa(int(mag)))
				}
			}
			return rcvr, nil
		}
	}
	if len(rcvr.mant) == 1 {
		if rcvr.ind == 0 {
			if rcvr.exp == 0 {
				return rcvr, nil
			}
		}
	}
	return RxFromInt32(0), nil
}

// NetRexx: method intcheck(min=int,max=int) private returns int signals BadArgumentException
//
// Differences: None
func (rcvr *Rexx) intcheck(min int32, max int32) (int32, error) {
	var i int32 = 0
	if rcvr.ind == NotaNum {
		return 0, RxException(9, "Not a number")
	}
	lodigit := len(rcvr.mant) - 1
	if rcvr.exp < 0 {
		i = rcvr.exp + 1
		for ; i <= 0; i++ {
			if rcvr.mant[lodigit] == '0' {
				lodigit--
			} else {
				return 0, RxException(9, ToRxFromRunes([]rune("Non-zero decimal part in")).OpCcBlank(nil, rcvr).ToString())
			}
		}
	}
	i, err := rcvr.ToInt32()
	if err != nil {
		return 0, err
	}
	if i < min {
		return 0, RxException(1, fmt.Sprintf("%v%v%v%v%v%v", "Argument ", i, " ", "<", " ", min))
	}
	if i > max {
		return 0, RxException(1, fmt.Sprintf("%v%v%v%v%v%v", "Argument ", i, " ", ">", " ", max))
	}
	return i, nil
}

// NetRexx: method intlength() private final returns int
//
// Differences: None
func (rcvr *Rexx) intlength() int32 {
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	return int32(len(rcvr.chars))
}

// NetRexx: method intwords() private final returns int
//
// Differences: None
func (rcvr *Rexx) intwords() int32 {
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	return words(rcvr.chars)
}

// NetRexx: method layout() private returns char[]
//
// Differences: None
func (rcvr *Rexx) layout() []rune {
	var s rune
	var sb strings.Builder
	var euse int32 = 0
	var sig int32 = 0
	if rcvr.dig < 0 {
		SayString(fmt.Sprintf("%v%v", "Internal error: Rexx: bad dig ", rcvr.dig))
		rcvr.dig = DefaultDigits
	}
	mag := rcvr.exp + int32(len(rcvr.mant))
	if mag > rcvr.dig || mag < -5 {
		if rcvr.form != PLAIN {
			if rcvr.ind == isneg {
				sb.WriteRune('-')
			}
			euse = rcvr.exp + int32(len(rcvr.mant)) - 1
			if rcvr.form == DefaultForm {
				sb.WriteRune(rcvr.mant[0])
				if len(rcvr.mant) > 1 {
					sb.WriteRune('.')
					for i := 1; i <= len(rcvr.mant)-1; i++ {
						sb.WriteRune(rcvr.mant[i])
					}
				}
			} else {
				for {
					sig = euse % 3
					if sig < 0 {
						sig = 3 + sig
					}
					euse = euse - sig
					sig++
					if sig >= int32(len(rcvr.mant)) {
						for i := 0; i < len(rcvr.mant); i++ {
							sb.WriteRune(rcvr.mant[i])
						}
						_12 := sig - int32(len(rcvr.mant))
						for ; _12 > 0; _12-- {
							sb.WriteRune('0')
						}
					} else {
						var i int32
						for i = 0; i < sig; i++ {
							sb.WriteRune(rcvr.mant[i])
						}
						sb.WriteRune('.')
						for i := sig; i < int32(len(rcvr.mant)); i++ {
							sb.WriteRune(rcvr.mant[i])
						}
					}
					if !false {
						break
					}
				}
			}
			if euse != 0 {
				if euse < 0 {
					s = '-'
					euse = int32(-euse)
				} else {
					s = '+'
				}
				sb.WriteRune('E')
				sb.WriteRune(s)
				chrs := []rune(strconv.Itoa(int(euse)))
				for i := 0; i < len(chrs); i++ {
					sb.WriteRune(chrs[i])
				}
			}
			return []rune(sb.String())
		}
	}
	if rcvr.exp == 0 {
		if rcvr.ind >= 0 {
			return rcvr.mant
		}
		sb.WriteRune('-')
		for i := 0; i < len(rcvr.mant); i++ {
			sb.WriteRune(rcvr.mant[i])
		}
		return []rune(sb.String())
	}
	if rcvr.ind == isneg {
		sb.WriteRune('-')
	}
	if mag < 1 {
		sb.WriteRune('0')
		sb.WriteRune('.')
		_13 := int(-mag)
		for ; _13 > 0; _13-- {
			sb.WriteRune('0')
		}
		for i := 0; i < len(rcvr.mant); i++ {
			sb.WriteRune(rcvr.mant[i])
		}
		return []rune(sb.String())
	}
	if mag > int32(len(rcvr.mant)) {
		for i := 0; i < len(rcvr.mant); i++ {
			sb.WriteRune(rcvr.mant[i])
		}
		_14 := mag - int32(len(rcvr.mant))
		for ; _14 > 0; _14-- {
			sb.WriteRune('0')
		}
		return []rune(sb.String())
	}
	var i int32
	for i = 0; i < mag; i++ {
		sb.WriteRune(rcvr.mant[i])
	}
	sb.WriteRune('.')
	for i = mag; i < int32(len(rcvr.mant)); i++ {
		sb.WriteRune(rcvr.mant[i])
	}
	return []rune(sb.String())
}

// NetRexx: method layoutnum() shared returns char[]
//
// Differences: None
func (rcvr *Rexx) layoutnum() []rune {
	return rcvr.layout()
}

// NetRexx: method optioncheck(oklist=String) private returns char signals BadArgumentException
//
// Differences: None
func (rcvr *Rexx) optioncheck(oklist string) (rune, error) {
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if len(rcvr.chars) == 0 {
		return 0, RxException(1, "Null option string")
	}
	ochar := rcvr.chars[0]
	uchar := unicode.ToUpper(ochar)
	if !strings.ContainsRune(oklist, uchar) {
		return 0, RxException(1, fmt.Sprintf("%v%v%v%v%v%v", "Bad Option character ", string(ochar), " ", "[", string(uchar), "]"))
	}
	return uchar, nil
}

// NetRexx: method padcheck() private returns char signals NotCharacterException
//
// Differences: None
func (rcvr *Rexx) padcheck() (rune, error) {
	if rcvr.chars == nil {
		rcvr.chars = rcvr.layout()
	}
	if len(rcvr.chars) != 1 {
		return 0, RxException(7, string(rcvr.chars))
	}
	return rcvr.chars[0], nil
}

// NetRexx: method round(d=int) private
//
// Differences: None
func (rcvr *Rexx) round(d int32) {
	var use []rune
	adjust := int32(len(rcvr.mant)) - d
	if adjust <= 0 {
		return
	}
	rcvr.exp = rcvr.exp + adjust
	if adjust > 1 {
		use = make([]rune, d+1)
		copy(use, rcvr.mant)
	} else {
		use = rcvr.mant
	}
	add := make([]rune, 1)
	add[0] = '5'
	_new := charaddsub(use, add, 1)
	if int32(len(_new)) > d+1 {
		rcvr.exp++
	}
	res := make([]rune, d)
	copy(res, _new)
	rcvr.mant = res
	rcvr.chars = nil
	return
}

// NetRexx: method charaddsub(a=char[],b=char[],m=int) static private returns char[]
//
// Differences: None
func charaddsub(a []rune, b []rune, m int32) []rune {
	var da int32 = 0
	ap := int32(len(a)) - 1
	bp := int32(len(b)) - 1
	maxarr := ap
	if bp > maxarr {
		maxarr = bp
	}
	res := make([]rune, maxarr+1)
	var carry int32 = 0
	op := maxarr
	for ; op >= 0; op-- {
		da = carry
		if ap >= 0 {
			da = da + int32(a[ap]) - int32('0')
			ap--
		}
		if bp >= 0 {
			da = da + (int32(b[bp])-int32('0'))*m
			bp--
		}
		if da < 0 {
			da = da + 100
			carry = da/10 - 10
			res[op] = rune(da%10 + int32('0'))
			continue
		}
		if da > 9 {
			carry = da / 10
			res[op] = rune(da%10 + int32('0'))
			continue
		}
		carry = 0
		res[op] = rune(da + int32('0'))
	}
	if carry == 0 {
		return res
	}
	_new := make([]rune, maxarr+2)
	_new[0] = rune(carry + int32('0'))
	copy(_new[1:], res)
	return _new
}

// NetRexx: This method does not exist
//
// Returns the digit value for char
// This is an internal helper function.
func getdigit(char rune) int32 {
	if unicode.IsDigit(char) {
		if int(char) >= 66720 {
			match := uint32(char)
			for _, seek := range unicode.Nd.R32 {
				if seek.Lo <= match && match <= seek.Hi {
					var number int32 = 0
					for scan := seek.Lo; scan <= seek.Hi; scan += seek.Stride {
						if number > 9 {
							number = 0
						}
						if scan == match {
							return number
						}
						number++
					}
				}
			}
		} else {
			match := uint16(char)
			for _, seek := range unicode.Nd.R16 {
				if seek.Lo <= match && match <= seek.Hi {
					var number int32 = 0
					for scan := seek.Lo; scan <= seek.Hi; scan += seek.Stride {
						if number > 9 {
							number = 0
						}
						if scan == match {
							return number
						}
						number++
					}
				}
			}
		}
	}
	return -1
}

// NetRexx: method sa2ca(s=String[]) private static returns char[]
//
// Differences: None
func sa2ca(s []string) []rune {
	seglen := 0
	blanks := len(s) - 1
	if blanks < 0 {
		return make([]rune, 0)
	}
	leng := blanks
	_38 := blanks
	i := 0
	for ; i <= _38; i++ {
		leng = leng + len([]rune(s[i]))
	}
	res := make([]rune, leng)
	out := 0
	_39 := blanks
	i = 0
	for ; i <= _39; i++ {
		seglen = len([]rune(s[i]))
		copy(res[out:], []rune(s[i]))
		if i == blanks {
			break
		}
		out = out + seglen + 1
		res[out-1] = ' '
	}
	return res
}
