package lang

const SCIENTIFIC int8 = 0
const ENGINEERING int8 = 1
const PLAIN int8 = 2
const DEFAULT_DIGITS int32 = 9
const DEFAULT_FORM = SCIENTIFIC

type RexxSet struct {
	Digits int32
	Form   int8
}

// NetRexx: method RexxSet() returns RexxSet
//
// Default: digits = 9, form = scientific
//
// Differences: None
func RxSet() (rcvr *RexxSet) {
	rcvr = &RexxSet{}
	rcvr.Digits = DEFAULT_DIGITS
	rcvr.Form = DEFAULT_FORM
	return
}

// NetRexx: method RexxSet(old=RexxSet)
//
// Differences: If old is Nil, call RxSet()
func RxSetFromOld(old *RexxSet) (rcvr *RexxSet) {
	if old != nil {
		rcvr = RxSet()
		rcvr.Digits = old.Digits
		rcvr.Form = old.Form
	} else {
		// Prevents Go Panic
		rcvr = RxSet()
	}
	return
}

// NetRexx: method RexxSet(newdigits=int) returns RexxSet
//
// Differences: None
func RxSetWithDigit(newdigits int32) (rcvr *RexxSet) {
	rcvr = RxSet()
	rcvr.Digits = newdigits
	return
}

// NetRexx: method RexxSet(newdigits=int,newform=byte) returns RexxSet
//
// Differences: None
func RxSetWithDigitandForm(newdigits int32, newform int8) (rcvr *RexxSet) {
	rcvr = RxSet()
	rcvr.Digits = newdigits
	rcvr.Form = newform
	return
}

// NetRexx: method formword returns Rexx
//
// Differences: None
func (rcvr *RexxSet) FormWord() *Rexx {
	if rcvr.Form == SCIENTIFIC {
		return ToRxFromRunes([]rune("scientific"))
	}
	return ToRxFromRunes([]rune("engineering"))
}

// NetRexx: method setDigits(d=Rexx)
//
// Differences: If d is nil, throw an Exception
func (rcvr *RexxSet) SetDigits(d *Rexx) error {
	if d == nil {
		// Prevents Go Panic
		return RxException(0, "cannot use nil as type Rexx in argument")
	}
	r, err := d.OpPlus(rcvr)
	if err != nil {
		return err
	}
	if r.ind == ispos {
		rx01, err := r.DataType(RxFromRune('w'))
		if err != nil {
			return err
		}
		test := rx01.OpEqS(nil, RxFromRune('1'))
		if test {
			if int32(len(r.mant))+r.exp <= 9 {
				num, err := r.ToInt32()
				if err != nil {
					return err
				}
				rcvr.Digits = num
				return nil
			}
		}
	}
	return RxException(3, d.ToString())
}

// NetRexx: method setForm(f=Rexx)
//
// Differences:  If f is nil, throw an Exception
func (rcvr *RexxSet) SetForm(f *Rexx) error {
	if f == nil {
		// Prevents Go Panic
		return RxException(0, "cannot use nil as type Rexx in argument")
	}
	eng, err := f.OpEq(nil, ToRxFromRunes([]rune("engineering")))
	if err != nil {
		return err
	}
	sci, err := f.OpEq(nil, ToRxFromRunes([]rune("scientific")))
	if err != nil {
		return err
	}
	if eng {
		rcvr.Form = ENGINEERING
	} else if sci {
		rcvr.Form = SCIENTIFIC
	} else {
		return RxException(3, f.ToString())
	}
	return nil
}
