package lang

import (
	"bytes"
	//"fmt"
	"io"
	"math"
	"os"
	"testing"
)

// Note: Rexx("") may substituted for rxfromempty() or null in some places.
// Some examples use different Netrexx calls but return the same values.
// On some tests, the returned values do not match NetRexx.
// 	This port has added checks to prevent panics.

// func rx(s []rune, trynum bool) (rcvr *Rexx)
// method Rexx(s=char[],trynum=boolean) not accessible in NetRexx
// Test_1000 to Test_1060 just use Rexx("some value")

func Test_1000(t *testing.T) {

	// -- say Rexx("")

	got := rx(nil, true).ToString()
	want := ""

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1001(t *testing.T) {

	// -- say Rexx("")

	got := rx(nil, false).ToString()
	want := ""

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1002(t *testing.T) {

	// -- say Rexx('a')

	got := rx([]rune{'a'}, true).ToString()
	want := "a"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1003(t *testing.T) {

	// -- say Rexx("")

	got := rx([]rune{}, false).ToString()
	want := ""

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1004(t *testing.T) {

	// -- say Rexx('a').tochar()

	value, err := rx([]rune{'a'}, false).ToRune()
	want := 'a'

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1005(t *testing.T) {

	// -- say Rexx('Ā').tochar()

	value, err := rx([]rune{'Ā'}, true).ToRune()
	want := 'Ā'

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1006(t *testing.T) {

	// -- say Rexx("½.½")

	got := rx([]rune("½.½"), true).ToString()
	want := "½.½"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1007(t *testing.T) {

	// -- say Rexx("-a.bc.G-f")

	got := rx([]rune("-a.bc.G-f"), false).ToString()
	want := "-a.bc.G-f"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1008(t *testing.T) {

	// -- say Rexx("0A")

	got := rx([]rune{'0', 'A'}, true).ToString()
	want := "0A"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1009(t *testing.T) {

	// -- say Rexx(" ")

	got := rx([]rune{32}, true).ToString()
	want := " "

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1010(t *testing.T) {

	// -- say Rexx(" 1 ")

	got := rx([]rune{32, 49, 32}, true).ToString()
	want := " 1 "

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1011(t *testing.T) {

	// -- say Rexx("-----.12345")

	got := rx([]rune("-----.12345"), true).ToString()
	want := "-----.12345"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1012(t *testing.T) {

	// -- say Rexx("--ABC---")

	got := rx([]rune{45, 45, 65, 66, 67, 45, 45, 45}, true).ToString()
	want := "--ABC---"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1013(t *testing.T) {

	// -- say Rexx("-000001.12345")

	got := rx([]rune("-000001.12345"), true).ToString()
	want := "-000001.12345"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1014(t *testing.T) {

	// -- say Rexx("++++.12345")

	got := rx([]rune("++++.12345"), true).ToString()
	want := "++++.12345"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1015(t *testing.T) {

	// -- say Rexx("+.12345")

	got := rx([]rune("+.12345"), true).ToString()
	want := "+.12345"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1016(t *testing.T) {

	// -- say Rexx("+++")

	got := rx([]rune("+++"), true).ToString()
	want := "+++"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1017(t *testing.T) {

	// -- say Rexx("-")

	got := rx([]rune("-"), true).ToString()
	want := "-"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1018(t *testing.T) {

	// -- say Rexx(" 2 ")

	got := rx([]rune{32, 50, 32}, true).ToString()
	want := " 2 "

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1019(t *testing.T) {

	// -- say Rexx(" 0:")

	got := rx([]rune{32, 48, 58}, true).ToString()
	want := " 0:"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1020(t *testing.T) {

	// -- say Rexx("......")

	got := rx([]rune("......"), true).ToString()
	want := "......"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1021(t *testing.T) {

	// -- say Rexx(" 0?")

	got := rx([]rune{32, 48, 63}, true).ToString()
	want := " 0?"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1022(t *testing.T) {

	// -- say Rexx(" 0۱")

	got := rx([]rune{32, 48, 0x06F1}, true).ToString()
	want := " 0۱"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1023(t *testing.T) {

	// -- say Rexx(" 0E")

	got := rx([]rune{32, 48, 69}, true).ToString()
	want := " 0E"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1024(t *testing.T) {

	// -- say Rexx(" 0")

	got := rx([]rune{32, 48}, true).ToString()
	want := " 0"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1025(t *testing.T) {

	// -- say Rexx("0E")

	got := rx([]rune("0E"), true).ToString()
	want := "0E"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1026(t *testing.T) {

	// -- say Rexx(" 0E00000000000c")

	got := rx([]rune{32, 48, 69, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 99}, true).ToString()
	want := " 0E00000000000c"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1027(t *testing.T) {

	// -- say Rexx("1E+99999999999")

	got := rx([]rune("1E+99999999999"), true).ToString()
	want := "1E+99999999999"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1028(t *testing.T) {

	// -- say Rexx("-.011111111111111111111111111E-۱011111111111111111111111111")

	got := rx([]rune("-.011111111111111111111111111E-۱011111111111111111111111111"), true).ToString()
	want := "-.011111111111111111111111111E-۱011111111111111111111111111"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1029(t *testing.T) {

	// -- say Rexx("0.1E-510")

	got := rx([]rune("0.1E-510"), true).ToString()
	want := "0.1E-510"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1030(t *testing.T) {

	// -- say Rexx("0.1E+510")

	got := rx([]rune("0.1E+510"), true).ToString()
	want := "0.1E+510"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1031(t *testing.T) {

	// -- say Rexx("+1.0E/۱")

	got := rx([]rune("+1.0E/۱"), true).ToString()
	want := "+1.0E/۱"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1032(t *testing.T) {

	// -- say Rexx("+1.0E++1")

	got := rx([]rune("+1.0E++1"), true).ToString()
	want := "+1.0E++1"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1033(t *testing.T) {

	// -- say Rexx("1.0E+=")

	got := rx([]rune("1.0E+="), true).ToString()
	want := "1.0E+="

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1034(t *testing.T) {

	// -- say Rexx("+1.0E+Ꭾ")

	got := rx([]rune("+1.0E+Ꭾ"), true).ToString()
	want := "+1.0E+Ꭾ"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1035(t *testing.T) {

	// -- say Rexx("1.0E+½")

	got := rx([]rune("1.0E+½"), true).ToString()
	want := "1.0E+½"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1036(t *testing.T) {

	// -- say Rexx("۱۱.۱۱E+۱")

	got := rx([]rune("۱۱.۱۱E+۱"), true).ToString()
	want := "۱۱.۱۱E+۱"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1037(t *testing.T) {

	// -- say Rexx("4.4E+444")

	got := rx([]rune("4.4E+444"), true).ToString()
	want := "4.4E+444"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1038(t *testing.T) {

	// -- say Rexx("0.1E-510")

	got := rx([]rune("0.1E-510"), true).ToString()
	want := "0.1E-510"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1039(t *testing.T) {

	// -- say Rexx("۱۱۱۱.0E+6")

	got := rx([]rune("۱۱۱۱.0E+6"), true).ToString()
	want := "۱۱۱۱.0E+6"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1040(t *testing.T) {

	// -- say Rexx('.').tochar()

	value, err := rx([]rune{'.'}, true).ToRune()
	want := '.'

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1041(t *testing.T) {

	// -- say Rexx("1.0")

	got := rx([]rune("1.0"), true).ToString()
	want := "1.0"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1042(t *testing.T) {

	// -- say Rexx("+.1038")

	got := rx([]rune("+.1038"), true).ToString()
	want := "+.1038"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1043(t *testing.T) {

	// -- say Rexx("            +0۱1.0E-9۱9۱99       ")

	got := rx([]rune("            +0۱1.0E-9۱9۱99       "), true).ToString()
	want := "            +0۱1.0E-9۱9۱99       "

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1044(t *testing.T) {

	// -- say Rexx("100")

	got := rx([]rune("100"), true).ToString()
	want := "100"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1045(t *testing.T) {

	// -- say Rexx("¼.¼")

	got := rx([]rune("¼.¼"), true).ToString()
	want := "¼.¼"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1046(t *testing.T) {

	// -- say Rexx("٠.٠٠٠٠٠")

	got := rx([]rune("٠.٠٠٠٠٠"), true).ToString()
	want := "٠.٠٠٠٠٠"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1047(t *testing.T) {

	// -- say Rexx("            +0۱2.0E-8۱7۱33       ")

	got := rx([]rune("            +0۱2.0E-8۱7۱33       "), true).ToString()
	want := "            +0۱2.0E-8۱7۱33       "

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1048(t *testing.T) {

	// -- say Rexx("1.¼")

	got := rx([]rune("1.¼"), true).ToString()
	want := "1.¼"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1049(t *testing.T) {

	// -- say Rexx("+7۱3.0E-5۱4۱22")

	got := rx([]rune("+7۱3.0E-5۱4۱22"), true).ToString()
	want := "+7۱3.0E-5۱4۱22"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1050(t *testing.T) {

	// -- say Rexx("1.¼123")

	got := rx([]rune("1.¼123"), true).ToString()
	want := "1.¼123"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1051(t *testing.T) {

	// -- say Rexx("+0۱3.0E-5۱4۱22")

	got := rx([]rune("+0۱3.0E-5۱4۱22"), true).ToString()
	want := "+0۱3.0E-5۱4۱22"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1052(t *testing.T) {

	// -- say Rexx("0")

	got := rx([]rune("0"), true).ToString()
	want := "0"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1053(t *testing.T) {

	// -- say Rexx("-000001.12345")

	got := rx([]rune("-000001.12345"), true).ToString()
	want := "-000001.12345"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1054(t *testing.T) {

	// -- say Rexx("3.0")

	got := rx([]rune("3.0"), true).ToString()
	want := "3.0"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1055(t *testing.T) {

	// -- say Rexx("0.2E-053")

	got := rx([]rune("0.2E-053"), true).ToString()
	want := "0.2E-053"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1056(t *testing.T) {

	// -- say Rexx("۱.۱E+۱")

	got := rx([]rune("۱.۱E+۱"), true).ToString()
	want := "۱.۱E+۱"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1057(t *testing.T) {

	// -- say Rexx("٠.٠٠")

	got := rx([]rune("٠.٠٠"), true).ToString()
	want := "٠.٠٠"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1058(t *testing.T) {

	// -- say Rexx("-1.7")

	got := rx([]rune("-1.7"), true).ToString()
	want := "-1.7"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1059(t *testing.T) {

	// -- say Rexx("200")

	got := rx([]rune("200"), true).ToString()
	want := "200"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1060(t *testing.T) {

	// -- say Rexx("+.17")

	got := rx([]rune("+.17"), true).ToString()
	want := "+.17"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func RxFromBool(flag bool) (rcvr *Rexx)

func Test_1061(t *testing.T) {

	// -- say Rexx("1").toboolean() -- will return "1"

	value, err := RxFromBool(true).ToBool()
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1062(t *testing.T) {

	// -- say Rexx("0").toboolean() -- will return "0"

	value, err := RxFromBool(false).ToBool()
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func rxfromclone(in *Rexx) (rcvr *Rexx)

func Test_1063(t *testing.T) {

	// -- say Rexx(Rexx null).toString() -- cannot use null

	got := rxfromclone(nil).ToString()
	want := ""

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1064(t *testing.T) {

	// -- say Rexx(Rexx "1")

	got := rxfromclone(RxFromBool(true)).ToString()
	want := "1"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func rxfromempty() (rcvr *Rexx)

func Test_1065(t *testing.T) {

	// -- say Rexx("").toString()

	got := rxfromempty().ToString()
	want := ""

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func RxFromFloat32(num float32) (rcvr *Rexx)

func Test_1066(t *testing.T) {

	// -- say Rexx(float 1.123).toString()

	got := RxFromFloat32(1.123).ToString()
	want := "1.123"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1067(t *testing.T) {

	// -- say Rexx(float 1.0).tofloat()

	value, err := RxFromFloat32(1.0).ToFloat32()
	want := float32(1.0)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func RxFromFloat64(num float64) (rcvr *Rexx)

func Test_1068(t *testing.T) {

	// -- say Rexx(double 1.123).toString()

	got := RxFromFloat64(1.123).ToString()
	want := "1.123"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1069(t *testing.T) {

	// -- say Rexx(double 1.0).todouble()

	value, err := RxFromFloat64(1.0).ToFloat64()
	want := float64(1.0)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func RxFromInt16(num int16) (rcvr *Rexx)

func Test_1070(t *testing.T) {

	// -- say Rexx(short 1).toString()

	got := RxFromInt16(1).ToString()
	want := "1"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1071(t *testing.T) {

	// -- say Rexx(short 1).toshort()

	value, err := RxFromInt16(1).ToInt16()
	want := int16(1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func RxFromInt32(num int32) (rcvr *Rexx)

func Test_1072(t *testing.T) {

	// -- say Rexx(int 1).toString()

	got := RxFromInt32(1).ToString()
	want := "1"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1073(t *testing.T) {

	// -- say Rexx(int 1).toint()

	value, err := RxFromInt32(1).ToInt32()
	want := int32(1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1074(t *testing.T) {

	// -- say Rexx(int 0).toString()

	got := RxFromInt32(0).ToString()
	want := "0"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1075(t *testing.T) {

	// -- say Rexx(int 0).toint()

	value, err := RxFromInt32(0).ToInt32()
	want := int32(0)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1076(t *testing.T) {

	// -- say Rexx(int -8).toString()

	got := RxFromInt32(-8).ToString()
	want := "-8"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1077(t *testing.T) {

	// -- say Rexx(int -8).toint()

	value, err := RxFromInt32(-8).ToInt32()
	want := int32(-8)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1078(t *testing.T) {

	// -- say Rexx(int 2147483647).toString()

	got := RxFromInt32(2147483647).ToString()
	want := "2147483647"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1079(t *testing.T) {

	// -- say Rexx(int 2147483647).toint()

	value, err := RxFromInt32(2147483647).ToInt32()
	want := int32(2147483647)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1080(t *testing.T) {

	// -- say Rexx(int -2147483648).toString() -- java.lang.NumberFormatException

	got := RxFromInt32(-2147483648).ToString()
	want := "-2147483648"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1081(t *testing.T) {

	// -- say Rexx(int -2147483648).toint() -- java.lang.NumberFormatException

	value, err := RxFromInt32(-2147483648).ToInt32()
	want := int32(-2147483648)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1082(t *testing.T) {

	// -- say Rexx(int -12).toString()

	got := RxFromInt32(-12).ToString()
	want := "-12"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1083(t *testing.T) {

	// -- say Rexx(int -12).toint()

	value, err := RxFromInt32(-12).ToInt32()
	want := int32(-12)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func RxFromInt64(num int64) (rcvr *Rexx)

func Test_1084(t *testing.T) {

	// -- say Rexx(long 9223372036854775807).toString()

	got := RxFromInt64(9223372036854775807).ToString()
	want := "9223372036854775807"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1085(t *testing.T) {

	// -- say Rexx(long 9223372036854775807).tolong()

	value, err := RxFromInt64(9223372036854775807).ToInt64()
	want := int64(9223372036854775807)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func RxFromInt8(num int8) (rcvr *Rexx)

func Test_1086(t *testing.T) {

	// -- say Rexx(byte -127).toString()

	got := RxFromInt8(-127).ToString()
	want := "-127"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1087(t *testing.T) {

	// -- say Rexx(byte -127).tobyte()

	value, err := RxFromInt8(-127).ToInt8()
	want := int8(-127)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func RxFromRune(inchar rune) (rcvr *Rexx)

func Test_1088(t *testing.T) {

	// -- say Rexx(char(32)).toString()

	got := RxFromRune(32).ToString()
	want := " "

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1089(t *testing.T) {

	// -- say Rexx(char(32)).tochar()

	value, err := RxFromRune(32).ToRune()
	want := ' '

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func RxFromRunes(in []rune) (rcvr *Rexx)

func Test_1090(t *testing.T) {

	// -- test = char[3]; test[0] = char(32); test[1] = char(65); test[2] = char(32); say Rexx(test).toString()

	got := RxFromRunes([]rune{32, 65, 32}).ToString()
	want := " A "

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1091(t *testing.T) {

	// -- test = char[3]; test[0] = char(32); test[1] = char(65); test[2] = char(32); say Rexx(test).tochararray()

	got := RxFromRunes([]rune{32, 65, 32}).ToRunes()
	want := " A "

	if string(got) != want {
		t.Errorf("got %v, wanted %v", string(got), want)
	}

}

// func RxFromString(str string) (rcvr *Rexx)

func Test_1092(t *testing.T) {

	// -- say Rexx("ABCDefgh").toString()

	got := RxFromString("ABCDefgh").ToString()
	want := "ABCDefgh"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1093(t *testing.T) {

	// -- say Rexx("7.62939453629394536293945362939453E+28")

	got := RxFromString("7.62939453629394536293945362939453E+28")
	want := "7.62939453629394536293945362939453E+28"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

// func RxFromStrings(strings []string) (rcvr *Rexx)

func Test_1094(t *testing.T) {

	// -- test = string[3]; test[0] = "1"; test[1] = "2"; test[2] = "S"; say Rexx(test).toString()

	got := RxFromStrings([]string{"1", "2", "S"}).ToString()
	want := "1 2 S"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func ToRuneFromRunes(s []rune) (rune, error)

func Test_1095(t *testing.T) {

	// -- say Rexx("").tochar(null) -- java.lang.NullPointerException

	_, err := ToRuneFromRunes(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1096(t *testing.T) {

	// -- say test = char[3]; test[0] = '1'; test[1] = '2'; test[2] = '3'; say Rexx(test).tochar()

	_, err := ToRuneFromRunes([]rune("123"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1097(t *testing.T) {

	// -- test = char[1]; test[0] = '1'; say Rexx(test).tochar()

	value, err := ToRuneFromRunes([]rune{'1'})
	want := '1'

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func ToRuneFromString(s string) (rune, error)

func Test_1098(t *testing.T) {

	// -- say Rexx("").tochar("123")

	_, err := ToRuneFromString("123")

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1099(t *testing.T) {

	// -- say Rexx("").tochar(String 'B' )

	value, err := ToRuneFromString(string('B'))
	want := 'B'

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1100(t *testing.T) {

	// -- say Rexx("").tochar(String 'Ꭳ' )

	value, err := ToRuneFromString(string('Ꭳ'))
	want := 'Ꭳ'

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func ToRunesFromRexx(r *Rexx) []rune

func Test_1101(t *testing.T) {

	// -- say Rexx("").tochararray(null)

	got := ToRunesFromRexx(nil)

	if got != nil {
		t.Errorf("got %v, wanted %v", got, nil)
	}

}

func Test_1102(t *testing.T) {

	// -- say Rexx("").tochararray()

	got := ToRunesFromRexx(rxfromempty())
	want := ""

	if string(got) != want {
		t.Errorf("got %v, wanted %v", string(got), want)
	}

}

func Test_1103(t *testing.T) {

	// -- say Rexx("").tochararray(Rexx("123"))

	got := ToRunesFromRexx(RxFromString("123"))
	want := "123"

	if string(got) != want {
		t.Errorf("got %v, wanted %v", string(got), want)
	}

}

func Test_1104(t *testing.T) {

	// -- say Rexx("").tochararray("A B C")

	got := ToRunesFromRexx(RxFromString("A B C"))
	want := "A B C"

	if string(got) != want {
		t.Errorf("got %v, wanted %v", string(got), want)
	}

}

// func ToRunesFromRune(c rune) []rune

func Test_1105(t *testing.T) {

	// -- say Rexx("").tochararray('Y')

	got := ToRunesFromRune('Y')
	want := []rune{'Y'}

	if string(got) != string(want) {
		t.Errorf("got %v, wanted %v", string(got), string(want))
	}

}

// func ToRxFromRunes(ca []rune) *Rexx

func Test_1106(t *testing.T) {

	// -- say Rexx("").toRexx(char[] "LONGshort")

	got := ToRxFromRunes([]rune("LONGshort"))
	want := "LONGshort"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1107(t *testing.T) {

	// -- say Rexx.toRexx(char[] null)
	// Netrexx will return null but this Go port returns ""

	got := ToRxFromRunes(nil).ToString()
	want := ""

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func ToRxFromString(s string) *Rexx

func Test_1108(t *testing.T) {

	// -- say Rexx.toRexx("")
	// Notes: Go does not allow a string to be nil.

	got := ToRxFromString("").ToString()
	want := ""

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1109(t *testing.T) {

	// -- say Rexx.toRexx(String "ᎠᎡᎢᎣᎤᎥᎦᎧᎨᎩᎪᎫᎬᎭᎮᎯᎰᎱᎲᎳᎴᎵᎶᎷᎸᎹᎺᎻᎼᎽᎾᎿᏀᏁᏂᏃᏄᏅᏆᏇᏈᏉᏊᏋᏌᏍᏎᏏᏐᏑᏒᏓᏔᏕᏖᏗᏘᏙᏚᏛᏜᏝᏞᏟᏠᏡᏢᏣᏤᏥᏦᏧᏨᏩᏪᏫᏬᏭᏮᏯᏰᏱᏲᏳᏴ ")

	got := ToRxFromString("ᎠᎡᎢᎣᎤᎥᎦᎧᎨᎩᎪᎫᎬᎭᎮᎯᎰᎱᎲᎳᎴᎵᎶᎷᎸᎹᎺᎻᎼᎽᎾᎿᏀᏁᏂᏃᏄᏅᏆᏇᏈᏉᏊᏋᏌᏍᏎᏏᏐᏑᏒᏓᏔᏕᏖᏗᏘᏙᏚᏛᏜᏝᏞᏟᏠᏡᏢᏣᏤᏥᏦᏧᏨᏩᏪᏫᏬᏭᏮᏯᏰᏱᏲᏳᏴ ")
	want := "ᎠᎡᎢᎣᎤᎥᎦᎧᎨᎩᎪᎫᎬᎭᎮᎯᎰᎱᎲᎳᎴᎵᎶᎷᎸᎹᎺᎻᎼᎽᎾᎿᏀᏁᏂᏃᏄᏅᏆᏇᏈᏉᏊᏋᏌᏍᏎᏏᏐᏑᏒᏓᏔᏕᏖᏗᏘᏙᏚᏛᏜᏝᏞᏟᏠᏡᏢᏣᏤᏥᏦᏧᏨᏩᏪᏫᏬᏭᏮᏯᏰᏱᏲᏳᏴ "

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

// func ToString(r *Rexx) string

func Test_1110(t *testing.T) {

	// -- say Rexx.toString(Rexx null)
	// Netrexx will return null but this Go port returns ""

	got := ToString(nil)
	want := ""

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1111(t *testing.T) {

	// -- say Rexx.toString(Rexx("")) -- cannot use Rexx() here

	got := ToString(rxfromempty())
	want := ""

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1112(t *testing.T) {

	// -- say Rexx.toString(Rexx("-0+"))

	got := ToString(RxFromString("-0+"))
	want := "-0+"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) Abbrev(b *Rexx, leng *Rexx) (*Rexx, error)

func Test_1113(t *testing.T) {

	// -- say Rexx("Print").Abbrev(Rexx("Pri"), null)

	_, err := RxFromString("Print").Abbrev(RxFromString("Pri"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1114(t *testing.T) {

	// -- say Rexx("Print").Abbrev(null, null)

	_, err := RxFromString("Print").Abbrev(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1115(t *testing.T) {

	// -- say Rexx("Print").Abbrev(null, Rexx("Pri").length())

	_, err := RxFromString("Print").Abbrev(nil, RxFromString("Pri").Length())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1116(t *testing.T) {

	// -- say Rexx("a").abbrev(Rexx("b"), Rexx(int 999999999 + 1))

	_, err := RxFromString("a").Abbrev(RxFromString("b"), RxFromInt32(MaxArg+1))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1117(t *testing.T) {

	// -- say Rexx("").abbrev(Rexx(""), Rexx(int 1))

	value, err := rxfromempty().Abbrev(RxFromString(""), RxFromInt32(1))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1118(t *testing.T) {

	// -- say Rexx("").abbrev(Rexx(""), Rexx(int 0))
	value, err := rxfromempty().Abbrev(RxFromString(""), RxFromInt32(0))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1119(t *testing.T) {

	// -- say Rexx("").abbrev(Rexx(""), Rexx(int 1))

	value, err := rxfromempty().Abbrev(rxfromempty(), RxFromInt32(1))

	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1120(t *testing.T) {

	// -- say Rexx("").abbrev(Rexx(""), Rexx(int 0))

	value, err := rxfromempty().Abbrev(rxfromempty(), RxFromInt32(0))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1121(t *testing.T) {

	// -- say Rexx("Print").Abbrev(Rexx("Pri"), Rexx("Pri").length())

	value, err := RxFromString("Print").Abbrev(RxFromString("Pri"), RxFromString("Pri").Length())
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Abs() (*Rexx, error)

func Test_1122(t *testing.T) {

	// -- say Rexx("xyz").abs() -- java.lang.NumberFormatException

	_, err := RxFromString("xyz").Abs()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1123(t *testing.T) {

	// -- say Rexx("- 1234567.7654321").abs()

	value, err := RxFromString("- 1234567.7654321").Abs()
	want := "1234567.7654321"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1124(t *testing.T) {

	// -- say Rexx(" -0.307").abs()

	value, err := RxFromString(" -0.307").Abs()
	want := "0.307"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1125(t *testing.T) {

	// -- say Rexx("123.45E+16").abs()

	value, err := RxFromString("123.45E+16").Abs()
	want := "1.2345E+18"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1126(t *testing.T) {

	// -- say Rexx("-1").abs()

	value, err := RxFromString("-1").Abs()
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) B2d(bil *Rexx) (*Rexx, error)

func Test_1127(t *testing.T) {

	// -- say Rexx("1").b2d(null) -- java.lang.NullPointerException

	_, err := RxFromString("1").B2d(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1128(t *testing.T) {

	// -- say Rexx("1").b2d(Rexx("-2147483649")) -- java.lang.NumberFormatException

	_, err := RxFromString("1").B2d(RxFromString("-2147483649"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1129(t *testing.T) {

	// -- say Rexx("1").b2d(Rexx("-9")) -- returns 1 but MaxExp and digits not set

	new_set := RxSet()
	new_set.SetForm(RxFromString("scientific"))
	new_set.SetDigits(RxFromString("16"))
	MaxExp = -1

	_, err := RxFromString("1").B2d(RxFromString("-9"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999

}

func Test_1130(t *testing.T) {

	// -- say Rexx("1").b2d(Rexx("0"))

	value, err := RxFromString("1").B2d(RxFromString("0"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1131(t *testing.T) {

	// -- say Rexx("1").b2d(Rexx("-1"))

	value, err := RxFromString("1").B2d(RxFromString("-1"))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1132(t *testing.T) {

	// -- say Rexx("10000001").b2d(Rexx("8"))

	value, err := RxFromString("10000001").B2d(RxFromString("8"))
	want := "-127"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1133(t *testing.T) {

	// -- say Rexx("11111111").b2d(Rexx("9999999999"))  -- netrexx.lang.BadArgumentException: Argument 1410065407 > 999999999

	_, err := RxFromString("11111111").B2d(RxFromString("9999999999"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1134(t *testing.T) {

	// -- say Rexx("1").b2d(Rexx("9999"))

	value, err := RxFromString("1").B2d(RxFromString("9999"))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1135(t *testing.T) {

	// -- say Rexx("10000001").b2d(Rexx(int 16))

	value, err := RxFromString("10000001").B2d(RxFromInt32(16))
	want := "129"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1136(t *testing.T) {

	// -- say Rexx("1").b2d(Rexx("99"))
	// MaxArg was found, but is not accessible in NetRexx

	MaxArg = 99

	_, err := RxFromString("1").B2d(RxFromString("99"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	// Reset
	MaxArg = 999999999
}

func Test_1137(t *testing.T) {

	// -- say Rexx("123").b2d(Rexx(int 128)) -- netrexx.lang.BadArgumentException: Bad binary 00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000123

	_, err := RxFromString("123").B2d(RxFromInt32(128))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1138(t *testing.T) {

	// -- say Rexx("0").b2d(Rexx("1"))

	value, err := RxFromString("0").B2d(RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1139(t *testing.T) {

	// -- say Rexx("10000000000000000000000000000000").b2d(Rexx(int 32))

	value, err := RxFromString("10000000000000000000000000000000").B2d(RxFromInt32(32))
	want := "-2147483648"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1140(t *testing.T) {

	// -- say Rexx("11111111110000000000000000000000000000000").b2d(Rexx(int 32))

	value, err := RxFromString("11111111110000000000000000000000000000000").B2d(RxFromInt32(32))
	want := "-2147483648"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1141(t *testing.T) {

	// -- say Rexx("11111111110000000000000000000000000000000").b2d(Rexx(int -2147483648)) -- java.lang.NumberFormatException: Conversion overflow

	value, err := RxFromString("11111111110000000000000000000000000000000").B2d(RxFromInt32(-2147483648))

	want := "2196875771904"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1142(t *testing.T) {

	// -- say Rexx("11").b2d(Rexx(int 9999999))

	value, err := RxFromString("11").B2d(RxFromInt32(9999999))
	want := "3"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1143(t *testing.T) {

	// -- say Rexx("1").b2d(Rexx(int -999999999))

	value, err := RxFromString("1").B2d(RxFromInt32(-999999999))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1144(t *testing.T) {

	// -- say Rexx("").b2d(Rexx(int -999999999))

	value, err := RxFromString("").B2d(RxFromInt32(-999999999))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1145(t *testing.T) {

	// -- say Rexx("0101011").B2d(Rexx("2"))	-- test uses variables not accessible in NetRexx

	MaxArg = 0
	MinArg = 0
	MaxExp = 0
	MinExp = 0

	new_set := RxSet()
	new_set.SetDigits(RxFromString("3"))

	_, err := RxFromString("0101011").B2d(RxFromString("2"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxArg = 999999999
	MinArg = -999999999
	MaxExp = 999999999
	MinExp = -999999999

}

// func (rcvr *Rexx) B2x() (*Rexx, error)

func Test_1146(t *testing.T) {

	// -- say Rexx().b2x() -- java.lang.NullPointerException

	_, err := rxfromempty().B2x()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1147(t *testing.T) {

	// -- say Rexx("11000011").b2x()

	value, err := RxFromString("11000011").B2x()
	want := "C3"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1148(t *testing.T) {

	// -- say Rexx("1.01E+999999999999999999999999999").b2x() -- netrexx.lang.BadArgumentException: Bad binary 1.01E+999999999999999999999999999

	_, err := RxFromString("1.01E+999999999999999999999999999").B2x()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1149(t *testing.T) {

	// -- say Rexx("-1").b2x() -- netrexx.lang.BadArgumentException: Bad binary -1

	_, err := RxFromString("-1").B2x()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1150(t *testing.T) {

	// -- say Rexx("1").b2x()

	value, err := RxFromString("1").B2x()
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1151(t *testing.T) {

	// -- say Rexx("0").b2x()

	value, err := RxFromString("0").B2x()
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) C2d() (*Rexx, error)

func Test_1152(t *testing.T) {

	// -- say Rexx().C2d() -- java.lang.NullPointerException

	_, err := rxfromempty().C2d()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1153(t *testing.T) {

	// -- say Rexx("1").C2d()

	value, err := RxFromString("1").C2d()
	want := "49"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1154(t *testing.T) {

	// -- say Rexx(int 6).C2d()

	value, err := RxFromInt32(6).C2d()
	want := "54"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) C2x() (*Rexx, error)

func Test_1155(t *testing.T) {

	// -- say Rexx().C2x() -- java.lang.NullPointerException

	_, err := rxfromempty().C2x()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1156(t *testing.T) {

	// -- say Rexx(char '\r').C2x()

	value, err := RxFromRune('\r').C2x()
	want := "D"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1157(t *testing.T) {

	// -- say Rexx(char 'M').C2x()

	value, err := RxFromRune('M').C2x()
	want := "4D"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1158(t *testing.T) {

	// -- say Rexx(char(4095)).C2x()

	value, err := RxFromRune(4095).C2x()
	want := "FFF"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1159(t *testing.T) {

	// -- say Rexx(char(0x104A0)).C2x() -- incorrect return
	// Osmanya Digit Zero
	// C2x returns only 4 chars in range of char 0 to FFFF in Java

	_, err := RxFromRune(0x104A0).C2x()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1160(t *testing.T) {

	// -- say Rexx(char(0xFFFF)).C2x()

	value, err := RxFromRune(0xFFFF).C2x()
	want := "FFFF"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1161(t *testing.T) {

	// -- say Rexx('۱').C2x() -- test uses variables not accessible in NetRexx

	MaxArg = 0
	MinArg = 0
	MaxExp = 0
	MinExp = 0

	new_set := RxSet()
	new_set.SetForm(RxFromString("engineering"))
	new_set.SetDigits(RxFromString("3"))

	RxFromRune('۱').C2x()

	MaxArg = 999999999
	MinArg = -999999999
	MaxExp = 999999999
	MinExp = -999999999

}

func Test_1162(t *testing.T) {

	// -- numeric digits 3; say Rexx('A').C2x() -- test uses variables not accessible in NetRexx

	MaxArg = 0

	new_set := RxSet()
	new_set.SetDigits(RxFromString("3"))

	RxFromRune('A').C2x()

	MaxArg = 999999999

}

func Test_1163(t *testing.T) {

	// -- say Rexx(char "\x00").C2x() -- test uses variables not accessible in NetRexx

	MaxArg = 0
	MinArg = 0
	MaxExp = 0
	MinExp = 0

	new_set := RxSet()
	new_set.SetForm(RxFromString("engineering"))
	new_set.SetDigits(RxFromString("3"))

	_, err := RxFromRune(0).C2x()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxArg = 999999999
	MinArg = -999999999
	MaxExp = 999999999
	MinExp = -999999999

}

// func (rcvr *Rexx) Center(wid *Rexx, pad *Rexx) (*Rexx, error)

func Test_1164(t *testing.T) {

	// -- say Rexx("111").center(null, Rexx(" ")) -- java.lang.NullPointerException

	_, err := RxFromString("111").Center(nil, RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1165(t *testing.T) {

	// -- say Rexx("111").center(null, null)

	_, err := RxFromString("111").Center(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1166(t *testing.T) {

	// -- say Rexx("111").center(Rexx("1"), null) -- java.lang.NullPointerException

	_, err := RxFromString("111").Center(RxFromString("1"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1167(t *testing.T) {

	// -- say Rexx("111").center(Rexx("1"), Rexx()) -- java.lang.NullPointerException

	_, err := RxFromString("111").Center(RxFromString("1"), rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1168(t *testing.T) {

	// -- say Rexx("111").center(Rexx("1"), Rexx("  ")) -- netrexx.lang.NotCharacterException

	_, err := RxFromString("111").Center(RxFromString("1"), RxFromString("  "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1169(t *testing.T) {

	// -- say Rexx("111").center(Rexx("5"), Rexx(" "))

	value, err := RxFromString("111").Center(RxFromString("5"), RxFromString(" "))
	want := " 111 "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Centre(wid *Rexx, pad *Rexx) (*Rexx, error)

func Test_1170(t *testing.T) {

	// -- say Rexx("22").centre(null, Rexx(" ")) -- java.lang.NullPointerException

	_, err := RxFromString("22").Centre(nil, RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1171(t *testing.T) {

	// -- say Rexx("22").centre(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("22").Centre(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1172(t *testing.T) {

	// -- say Rexx("0000").centre(Rexx("1"), null) -- java.lang.NullPointerException

	_, err := RxFromString("0000").Centre(RxFromString("1"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1173(t *testing.T) {

	// -- say Rexx("333").centre(Rexx(int 999999999+1), Rexx(" "))

	_, err := RxFromString("333").Centre(RxFromInt32(MaxArg+1), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}
}

func Test_1174(t *testing.T) {

	// -- say Rexx("333").centre(Rexx("3"), Rexx()) -- java.lang.NullPointerException

	_, err := RxFromString("333").Centre(RxFromString("3"), rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1175(t *testing.T) {

	// -- say Rexx("").centre(Rexx("1"), Rexx(" "))

	value, err := rxfromempty().Centre(RxFromString("1"), RxFromString(" "))
	want := " "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1176(t *testing.T) {

	// -- say Rexx("2").centre(Rexx("3"), Rexx(" "))

	value, err := RxFromString("2").Centre(RxFromString("3"), RxFromString(" "))
	want := " 2 "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) ChangeStr(old *Rexx, _new *Rexx) *Rexx

func Test_1177(t *testing.T) {

	// -- say Rexx("elephant").changestr(Rexx(""), Rexx("X"))

	got := RxFromString("elephant").ChangeStr(nil, RxFromString("X"))
	want := "elephant"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1178(t *testing.T) {

	// -- say Rexx("elephant").changestr(Rexx("e"), null) -- java.lang.NullPointerException

	got := RxFromString("elephant").ChangeStr(RxFromString("e"), nil)
	want := "elephant"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1179(t *testing.T) {

	// -- say Rexx("elephant").changestr(null, null) -- java.lang.NullPointerException

	got := RxFromString("elephant").ChangeStr(nil, nil)
	want := "elephant"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1180(t *testing.T) {

	// -- say Rexx("").changestr(Rexx("e"), Rexx("X"))

	got := rxfromempty().ChangeStr(RxFromString("e"), RxFromString("X"))
	want := ""

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1181(t *testing.T) {

	// -- say Rexx("elephant").changestr(Rexx(""), Rexx("X"))

	got := RxFromString("elephant").ChangeStr(rxfromempty(), RxFromString("X"))
	want := "elephant"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1182(t *testing.T) {

	// -- say Rexx("elephant").changestr(Rexx("e"), Rexx(""))

	got := RxFromString("elephant").ChangeStr(RxFromString("e"), rxfromempty())
	want := "lphant"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1183(t *testing.T) {

	// -- say Rexx("elephant").changestr(Rexx(""), Rexx("X"))

	got := RxFromString("elephant").ChangeStr(nil, RxFromString("X"))
	want := "elephant"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1184(t *testing.T) {

	// -- say Rexx("elephant").changestr(Rexx("ph"), Rexx("X"))

	got := RxFromString("elephant").ChangeStr(RxFromString("ph"), RxFromString("X"))
	want := "eleXant"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1185(t *testing.T) {

	// -- say Rexx("elephant").changestr(Rexx("ph"), Rexx("hph"))

	got := RxFromString("elephant").ChangeStr(RxFromString("ph"), RxFromString("hph"))
	want := "elehphant"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1186(t *testing.T) {

	// -- say Rexx("elephant").changestr(Rexx("e"), Rexx(""))

	got := RxFromString("elephant").ChangeStr(RxFromString("e"), RxFromString(""))
	want := "lphant"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1187(t *testing.T) {

	// -- say Rexx("elephant").changestr(Rexx(""), Rexx("!!"))

	got := RxFromString("elephant").ChangeStr(RxFromString(""), RxFromString("!!"))
	want := "elephant"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

// func (rcvr *Rexx) CharAt(index int32) (rune, error)

func Test_1188(t *testing.T) {

	// -- say Rexx("").CharAt(1)

	_, err := rxfromempty().CharAt(1)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1189(t *testing.T) {

	// -- say Rexx("Go").CharAt(0)

	value, err := RxFromString("Go").CharAt(0)
	want := 'G'

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Compare(target *Rexx, pad *Rexx) (*Rexx, error)

func Test_1190(t *testing.T) {

	// -- say Rexx("ab ").compare(null, Rexx("x")) -- java.lang.NullPointerException

	_, err := RxFromString("ab ").Compare(nil, RxFromString("x"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1191(t *testing.T) {

	// -- say Rexx("ab ").compare(Rexx("ab"), null) -- java.lang.NullPointerException

	_, err := RxFromString("ab ").Compare(RxFromString("ab"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1192(t *testing.T) {

	// -- say Rexx("abc").compare(Rexx(""), Rexx(""))) -- netrexx.lang.NotCharacterException

	_, err := RxFromString("abc").Compare(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1193(t *testing.T) {

	// -- say Rexx("ab ").compare(Rexx("ab"), Rexx()) -- java.lang.NullPointerException

	_, err := RxFromString("ab ").Compare(RxFromString("ab"), rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1194(t *testing.T) {

	// -- say Rexx("").compare(Rexx("ab"), Rexx(" "))

	value, err := rxfromempty().Compare(RxFromString("ab"), RxFromString(" "))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1195(t *testing.T) {

	// -- say Rexx("ab ").compare(Rexx(""), Rexx(" "))

	value, err := RxFromString("ab ").Compare(rxfromempty(), RxFromString(" "))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1196(t *testing.T) {

	// -- say Rexx("ab ").compare(Rexx("ab"), Rexx("x"))

	value, err := RxFromString("ab ").Compare(RxFromString("ab"), RxFromString("x"))
	want := "3"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1197(t *testing.T) {

	// -- say Rexx("ab-- ").compare(Rexx("ab"), Rexx("-"))

	value, err := RxFromString("ab-- ").Compare(RxFromString("ab"), RxFromString("-"))
	want := "5"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Copies(n *Rexx) (*Rexx, error)

func Test_1198(t *testing.T) {

	// -- say Rexx("abc").copies(null) -- java.lang.NullPointerException

	_, err := RxFromString("abc").Copies(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1199(t *testing.T) {

	// -- say Rexx("abc").copies(Rexx(int 999999999+1))

	_, err := RxFromString("abc").Copies(RxFromInt32(MaxArg + 1))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1200(t *testing.T) {

	// -- say Rexx("").copies(Rexx(int 0))

	value, err := rxfromempty().Copies(RxFromInt32(0))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1201(t *testing.T) {

	// -- say Rexx("").copies(Rexx(int 1))

	value, err := rxfromempty().Copies(RxFromInt32(1))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1202(t *testing.T) {

	// -- say Rexx("abc").copies(Rexx(byte 0))

	value, err := RxFromString("abc").Copies(RxFromInt8(0))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1203(t *testing.T) {

	// -- say Rexx("abc").copies(Rexx(byte 1))

	value, err := RxFromString("abc").Copies(RxFromInt8(1))
	want := "abc"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) CountStr(b *Rexx) *Rexx

func Test_1204(t *testing.T) {

	// -- say Rexx("elephant").countstr(Rexx(""))

	got := RxFromString("elephant").CountStr(nil)
	want := "0"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1205(t *testing.T) {

	// -- say Rexx("").countstr(Rexx("e"))

	got := rxfromempty().CountStr(RxFromString("e"))
	want := "0"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1206(t *testing.T) {

	// -- say Rexx("elephant").countstr(Rexx(""))

	got := RxFromString("elephant").CountStr(rxfromempty())
	want := "0"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1207(t *testing.T) {

	// -- say Rexx("elephant").countstr(Rexx("e"))

	got := RxFromString("elephant").CountStr(RxFromString("e"))
	want := "2"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1208(t *testing.T) {

	// -- say Rexx("elephant").countstr(Rexx("ph"))

	got := RxFromString("elephant").CountStr(RxFromString("ph"))
	want := "1"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

// func (rcvr *Rexx) D2b(dil *Rexx) (*Rexx, error)

func Test_1209(t *testing.T) {

	// -- say Rexx("-1").d2b(null) -- java.lang.NullPointerException

	_, err := RxFromString("-1").D2b(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1210(t *testing.T) {

	// -- say Rexx("-1").d2b(Rexx("0"))

	value, err := RxFromString("-1").D2b(RxFromString("0"))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1211(t *testing.T) {

	// -- say Rexx("259.9").d2b(Rexx("zip"))

	_, err := RxFromString("259.9").D2b(ToRxFromRunes([]rune("zip")))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1212(t *testing.T) {

	// -- say Rexx(int -999999999 - 1).d2b(Rexx("zip"))

	_, err := RxFromInt32(MinArg - 1).D2b(ToRxFromRunes([]rune("zip")))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1213(t *testing.T) {

	// -- say Rexx("009").d2b(Rexx("zip"))

	value, err := RxFromString("009").D2b(ToRxFromRunes([]rune("zip")))
	want := "1001"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1214(t *testing.T) {

	// -- say Rexx("-1").d2b(Rexx("ax")) -- java.lang.NumberFormatException: ax

	_, err := RxFromString("-1").D2b(RxFromString("ax"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1215(t *testing.T) {

	// -- say Rexx("-77e-9999999").d2b(Rexx("1"))

	value, err := RxFromString("-77e-9999999").D2b(RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1216(t *testing.T) {

	// -- say Rexx("-127").d2b(Rexx("8"))

	value, err := RxFromString("-127").D2b(RxFromString("8"))
	want := "10000001"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1217(t *testing.T) {

	// -- say Rexx("-127").d2b(Rexx(int 16))

	value, err := RxFromString("-127").D2b(RxFromInt32(16))
	want := "1111111110000001"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1218(t *testing.T) {

	// -- say Rexx("259.9").d2b(Rexx(long 8)) -- java.lang.NumberFormatException: Decimal part non-zero: 3.9

	_, err := RxFromString("259.9").D2b(RxFromInt64(8))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1219(t *testing.T) {

	// -- say Rexx("").d2b(Rexx("zip"))

	_, err := rxfromempty().D2b(ToRxFromRunes([]rune("zip")))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1220(t *testing.T) {

	// -- say Rexx("-127").d2b(Rexx()) -- java.lang.NullPointerException

	_, err := RxFromString("-127").D2b(rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1221(t *testing.T) {

	// -- say Rexx("0").d2b(Rexx("zip"))
	// Notes: "zip" is the default argument in NetRexx

	value, err := RxFromString("0").D2b(ToRxFromRunes([]rune("zip")))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1222(t *testing.T) {

	// -- say Rexx("000000000000000000000000000000000000001").d2b(Rexx("zip"))

	value, err := RxFromString("000000000000000000000000000000000000001").D2b(ToRxFromRunes([]rune("zip")))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1223(t *testing.T) {

	// -- say Rexx("000000").d2b(Rexx("-1")) -- netrexx.lang.BadArgumentException: Argument -1 < 0

	_, err := RxFromString("000000").D2b(RxFromString("-1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1224(t *testing.T) {

	// -- say Rexx("129").d2b(Rexx(byte 8))

	value, err := RxFromString("129").D2b(RxFromInt8(8))
	want := "10000001"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1225(t *testing.T) {

	// -- Rexx('A').D2b(Rexx("2")) -- java.lang.NumberFormatException: A

	_, err := RxFromRune('A').D2b(RxFromString("2"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

// func (rcvr *Rexx) D2c() (*Rexx, error)

func Test_1226(t *testing.T) {

	// -- say Rexx("AX").d2c() -- java.lang.NumberFormatException: AX

	_, err := RxFromString("AX").D2c()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1227(t *testing.T) {

	// -- say Rexx("-1").d2c() -- java.lang.NumberFormatException: Encoding bad -1

	_, err := RxFromString("-1").D2c()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1228(t *testing.T) {

	// -- say Rexx("77").d2c()

	value, err := RxFromString("77").D2c()
	want := "M"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1229(t *testing.T) {

	// -- say Rexx("+77").d2c()

	value, err := RxFromString("+77").D2c()
	want := "M"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) D2x(n *Rexx) (*Rexx, error)

func Test_1230(t *testing.T) {

	// -- say Rexx("9").d2x(null) -- java.lang.NullPointerException

	_, err := RxFromString("9").D2x(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1231(t *testing.T) {

	// -- say Rexx("9").d2x(Rexx("-2")) -- netrexx.lang.BadArgumentException: Argument -2 < 0

	_, err := RxFromString("9").D2x(RxFromString("-2"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1232(t *testing.T) {

	// -- say Rexx().d2x(Rexx("1")) -- java.lang.NullPointerException

	value, err := rxfromempty().D2x(RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1233(t *testing.T) {

	// -- say Rexx("").d2x(Rexx("1")) -- java.lang.NumberFormatException

	_, err := RxFromString("").D2x(RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1234(t *testing.T) {

	// -- say Rexx("19").d2x(Rexx("2"))

	value, err := RxFromString("19").D2x(RxFromString("2"))
	want := "13"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1235(t *testing.T) {

	// -- say Rexx("19").d2x()
	// Notes: This port allows "-1" for no argument passed

	value, err := RxFromString("19").D2x(RxFromString("-1"))
	want := "13"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1236(t *testing.T) {

	// -- say Rexx("1.2345E+18").d2x() -- netrexx.lang.DivideException: Integer overflow

	_, err := RxFromString("1.2345E+18").D2x(RxFromString("-1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

// func (rcvr *Rexx) DataType(opt *Rexx) (*Rexx, error)

func Test_1237(t *testing.T) {

	// -- say Rexx("ABC").datatype(null) -- java.lang.NullPointerException

	_, err := RxFromString("ABC").DataType(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1238(t *testing.T) {

	// -- say Rexx().datatype(Rexx()) -- java.lang.NullPointerException

	_, err := rxfromempty().DataType(rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1239(t *testing.T) {

	// -- say Rexx("").datatype(Rexx(char 'M'))

	value, err := RxFromString("").DataType(RxFromRune('M'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1240(t *testing.T) {

	// -- say Rexx("").datatype(Rexx("A"))

	value, err := rxfromempty().DataType(RxFromString("A"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1241(t *testing.T) {

	// -- say Rexx(char 'Z').datatype(Rexx(char 'A'))

	value, err := RxFromRune('Z').DataType(RxFromRune('A'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1242(t *testing.T) {

	// -- say Rexx(".........").datatype(Rexx(char 'A'))

	value, err := RxFromString(".........").DataType(RxFromRune('A'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1243(t *testing.T) {

	// -- say Rexx("0").datatype(Rexx(char 'B'))

	value, err := RxFromString("0").DataType(RxFromRune('B'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1244(t *testing.T) {

	// -- say Rexx("101").datatype(Rexx(char 'B'))

	value, err := RxFromString("101").DataType(RxFromRune('B'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1245(t *testing.T) {

	// -- say Rexx("a").datatype(Rexx(char 'B'))

	value, err := RxFromString("a").DataType(RxFromRune('B'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1246(t *testing.T) {

	// -- say Rexx("19").datatype(Rexx(char 'D'))

	value, err := RxFromString("19").DataType(RxFromRune('D'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1247(t *testing.T) {

	// -- say Rexx("12.3").datatype(Rexx(char 'D'))

	value, err := RxFromString("12.3").DataType(RxFromRune('D'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1248(t *testing.T) {

	// -- say Rexx(char 'a').datatype(Rexx(char 'L'))

	value, err := RxFromRune('a').DataType(RxFromRune('L'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1249(t *testing.T) {

	// -- say Rexx(char 'A').datatype(Rexx(char 'L'))

	value, err := RxFromRune('A').DataType(RxFromRune('L'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1250(t *testing.T) {

	// -- say Rexx("LaArca").datatype(Rexx(char 'M'))

	value, err := RxFromString("LaArca").DataType(RxFromRune('M'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1251(t *testing.T) {

	// -- say Rexx(char '1').datatype(Rexx(char 'M'))

	value, err := RxFromRune('1').DataType(RxFromRune('M'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1252(t *testing.T) {

	// -- say Rexx("UPPERS@lowers").datatype(Rexx(char 'M'))

	value, err := RxFromString("UPPERS@lowers").DataType(RxFromRune('M'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1253(t *testing.T) {

	// -- say Rexx("12.3").datatype(Rexx(char 'N'))

	value, err := RxFromString("12.3").DataType(RxFromRune('N'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1254(t *testing.T) {

	// -- say Rexx(char '\u06F1').datatype(Rexx(char 'N'))

	value, err := RxFromRune('\u06F1').DataType(RxFromRune('N'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1255(t *testing.T) {

	// -- say Rexx("1A").datatype(Rexx(char 'N'))

	value, err := RxFromString("1A").DataType(RxFromRune('N'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1256(t *testing.T) {

	// -- say Rexx("Alpha_is_start").datatype(Rexx(char 'S'))

	value, err := RxFromString("Alpha_is_start").DataType(RxFromRune('S'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1257(t *testing.T) {

	// -- say Rexx("_is_start").datatype(Rexx(char 'S'))

	value, err := RxFromString("_is_start").DataType(RxFromRune('S'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1258(t *testing.T) {

	// -- say Rexx("A8_is_part").datatype(Rexx(char 'S'))

	value, err := RxFromString("A8_is_part").DataType(RxFromRune('S'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1259(t *testing.T) {

	// -- say Rexx("b_is_part").datatype(Rexx(char 'S'))

	value, err := RxFromString("b_is_part").DataType(RxFromRune('S'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1260(t *testing.T) {

	// -- say Rexx("8_is_start").datatype(Rexx(char 'S'))

	value, err := RxFromString("8_is_start").DataType(RxFromRune('S'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1261(t *testing.T) {

	// -- say Rexx("@_is_start").datatype(Rexx(char 'S'))

	value, err := RxFromString("@_is_start").DataType(RxFromRune('S'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1262(t *testing.T) {

	// -- say Rexx("A+_is_part").datatype(Rexx(char 'S'))

	value, err := RxFromString("A+_is_part").DataType(RxFromRune('S'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1263(t *testing.T) {

	// -- say Rexx("A@_is_part").datatype(Rexx(char 'S'))

	value, err := RxFromString("A@_is_part").DataType(RxFromRune('S'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1264(t *testing.T) {

	// -- say Rexx("A").datatype(Rexx(char 'U'))

	value, err := RxFromString("A").DataType(RxFromRune('U'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1265(t *testing.T) {

	// -- say Rexx("1A").datatype(Rexx(char 'U'))

	value, err := RxFromString("1A").DataType(RxFromRune('U'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1266(t *testing.T) {

	// -- say Rexx("NOTANUMB").datatype(Rexx(char 'W'))

	value, err := RxFromString("NOTANUMB").DataType(RxFromRune('W'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1267(t *testing.T) {

	// -- say Rexx("1.99999999999999999").datatype(Rexx(char 'W'))

	value, err := RxFromString("1.99999999999999999").DataType(RxFromRune('W'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1268(t *testing.T) {

	// -- say Rexx("1.0").datatype(Rexx(char 'W'))

	value, err := RxFromString("1.0").DataType(RxFromRune('W'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1269(t *testing.T) {

	// -- say Rexx("12.3").datatype(Rexx(char 'W'))

	value, err := RxFromString("12.3").DataType(RxFromRune('W'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1270(t *testing.T) {

	// -- say Rexx(".1E-999999999").datatype(Rexx("W")) -- netrexx.lang.ExponentOverflowException: -1000000000

	_, err := RxFromString(".1E-999999999").DataType(RxFromString("W"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1271(t *testing.T) {

	// -- say Rexx("0").datatype(Rexx(char 'W'))

	value, err := RxFromString("0").DataType(RxFromRune('W'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1272(t *testing.T) {

	// -- say Rexx("123.3").datatype(Rexx(char 'W'))

	value, err := RxFromString("123.3").DataType(RxFromRune('W'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1273(t *testing.T) {

	// -- say Rexx("9.999E+999999999").datatype(Rexx("W"))

	value, err := RxFromString("9.999E+999999999").DataType(RxFromString("W"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1274(t *testing.T) {

	// -- say Rexx("BCd3").datatype(Rexx(char 'X'))

	value, err := RxFromString("BCd3").DataType(RxFromRune('X'))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1275(t *testing.T) {

	// -- say Rexx("BCgd3").datatype(Rexx(char 'X'))

	value, err := RxFromString("BCgd3").DataType(RxFromRune('X'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1276(t *testing.T) {

	// -- say Rexx("").datatype(Rexx(char 'X'))

	value, err := rxfromempty().DataType(RxFromRune('X'))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1277(t *testing.T) {

	// --

	MaxArg = 0
	MinArg = 0
	MaxExp = 0
	MinExp = 0

	new_set := RxSet()
	new_set.SetDigits(RxFromString("3"))

	_, err := RxFromString("1234").DataType(RxFromString("A"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxArg = 999999999
	MinArg = -999999999
	MaxExp = 999999999
	MinExp = -999999999

}

func Test_1278(t *testing.T) {

	// --

	MaxArg = 0
	MinArg = 0
	MaxExp = 0
	MinExp = 0

	new_set := RxSet()
	new_set.SetDigits(RxFromString("3"))

	_, err := RxFromString("1234").DataType(RxFromString("B"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxArg = 999999999
	MinArg = -999999999
	MaxExp = 999999999
	MinExp = -999999999

}

func Test_1279(t *testing.T) {

	// --

	MaxArg = 0
	MinArg = 0
	MaxExp = 0
	MinExp = 0

	new_set := RxSet()
	new_set.SetDigits(RxFromString("3"))

	_, err := RxFromString("1234").DataType(RxFromString("D"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxArg = 999999999
	MinArg = -999999999
	MaxExp = 999999999
	MinExp = -999999999

}

func Test_1280(t *testing.T) {

	// --

	MaxArg = 0
	MinArg = 0
	MaxExp = 0
	MinExp = 0

	new_set := RxSet()
	new_set.SetDigits(RxFromString("3"))

	_, err := RxFromString("1234").DataType(RxFromString("L"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxArg = 999999999
	MinArg = -999999999
	MaxExp = 999999999
	MinExp = -999999999

}

func Test_1281(t *testing.T) {

	// --

	MaxArg = 0
	MinArg = 0
	MaxExp = 0
	MinExp = 0

	new_set := RxSet()
	new_set.SetDigits(RxFromString("3"))

	_, err := RxFromString("1234").DataType(RxFromString("M"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxArg = 999999999
	MinArg = -999999999
	MaxExp = 999999999
	MinExp = -999999999

}

func Test_1282(t *testing.T) {

	// --

	MaxArg = 0
	MinArg = 0
	MaxExp = 0
	MinExp = 0

	new_set := RxSet()
	new_set.SetDigits(RxFromString("3"))

	_, err := RxFromString("1234").DataType(RxFromString("S"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxArg = 999999999
	MinArg = -999999999
	MaxExp = 999999999
	MinExp = -999999999

}

func Test_1283(t *testing.T) {

	// --

	MaxArg = 0
	MinArg = 0
	MaxExp = 0
	MinExp = 0

	new_set := RxSet()
	new_set.SetDigits(RxFromString("3"))

	_, err := RxFromString("1234").DataType(RxFromString("U"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxArg = 999999999
	MinArg = -999999999
	MaxExp = 999999999
	MinExp = -999999999

}

func Test_1284(t *testing.T) {

	// --

	MaxArg = 0
	MinArg = 0
	MaxExp = 0
	MinExp = 0

	new_set := RxSet()
	new_set.SetDigits(RxFromString("3"))

	_, err := RxFromString("1234").DataType(RxFromString("X"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxArg = 999999999
	MinArg = -999999999
	MaxExp = 999999999
	MinExp = -999999999

}

func Test_1285(t *testing.T) {

	// --

	MaxArg = 0
	MinArg = 0
	MaxExp = 0
	MinExp = 0

	new_set := RxSet()
	new_set.SetDigits(RxFromString("3"))

	_, err := RxFromString("1234").DataType(RxFromString("W"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxArg = 999999999
	MinArg = -999999999
	MaxExp = 999999999
	MinExp = -999999999

}

// func (rcvr *Rexx) DelStr(n *Rexx, length *Rexx) (*Rexx, error)

func Test_1286(t *testing.T) {

	// -- say Rexx("abcd").delstr(null, Rexx("-1")) -- java.lang.NullPointerException

	_, err := RxFromString("abcd").DelStr(nil, RxFromString("-1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1287(t *testing.T) {

	// -- say Rexx("abcd").delstr(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("abcd").DelStr(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1288(t *testing.T) {

	// -- say Rexx("abcd").delstr(Rexx("3"), null) -- java.lang.NullPointerException

	_, err := RxFromString("abcd").DelStr(RxFromString("3"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1289(t *testing.T) {

	// -- say Rexx("abcd").delstr(Rexx(""), Rexx("1")) -- java.lang.NumberFormatException: Not a number

	_, err := RxFromString("abcd").DelStr(RxFromString(""), RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1290(t *testing.T) {

	// -- say Rexx("abcd").delstr(Rexx("1"), Rexx("-1"))  -- netrexx.lang.BadArgumentException: Argument -1 < 0

	_, err := RxFromString("abcd").DelStr(RxFromString("1"), RxFromString("-1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1291(t *testing.T) {

	// -- say Rexx("abcd").delstr(Rexx("1"), Rexx("-50"))  -- netrexx.lang.BadArgumentException: Argument -50 < 0

	_, err := RxFromString("abcd").DelStr(RxFromString("1"), RxFromString("-50"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1292(t *testing.T) {

	// -- say Rexx("").delstr(Rexx("1"), Rexx("1"))

	value, err := rxfromempty().DelStr(RxFromString("1"), RxFromString("1"))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1293(t *testing.T) {

	// -- say Rexx("abcd").delstr(Rexx("1"), Rexx("ab").length())

	value, err := RxFromString("abcd").DelStr(RxFromString("1"), RxFromString("ab").Length())
	want := "cd"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1294(t *testing.T) {

	// -- say Rexx("abcd").delstr(Rexx("1"), Rexx("abcd").length())

	value, err := RxFromString("abcd").DelStr(RxFromString("1"), RxFromString("abcd").Length())
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1295(t *testing.T) {

	// -- say Rexx("abcd").delstr(Rexx("3"), Rexx("abcd").length())

	value, err := RxFromString("abcd").DelStr(RxFromString("3"), RxFromString("abcd").Length())
	want := "ab"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1296(t *testing.T) {

	// -- say Rexx("abcde").delstr(Rexx("3"), Rexx("2"))

	value, err := RxFromString("abcde").DelStr(RxFromString("3"), RxFromString("2"))
	want := "abe"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1297(t *testing.T) {

	// -- say Rexx("abcde").delstr(Rexx("6"), Rexx("5"))

	value, err := RxFromString("abcde").DelStr(RxFromString("6"), RxFromString("5"))
	want := "abcde"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) DelWord(n *Rexx, length *Rexx) (*Rexx, error)

func Test_1298(t *testing.T) {

	// -- say Rexx("Now is the  time").delword(null, Rexx("2")) -- java.lang.NullPointerException

	_, err := RxFromString("Now is the  time").DelWord(nil, RxFromString("2"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1299(t *testing.T) {

	// -- say Rexx("Now is the  time").delword(Rexx("2"), null) -- java.lang.NullPointerException

	_, err := RxFromString("Now is the  time").DelWord(RxFromString("2"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1300(t *testing.T) {

	// -- say Rexx("Now is the  time").delword(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("Now is the  time").DelWord(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1301(t *testing.T) {

	// -- say Rexx("").delword(Rexx(""), Rexx("")) -- java.lang.NumberFormatException: Not a number

	_, err := rxfromempty().DelWord(rxfromempty(), rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1302(t *testing.T) {

	// -- say Rexx("ab").delword(Rexx("1"), Rexx(int 999999999+1))

	_, err := RxFromString("ab").DelWord(RxFromString("1"), RxFromInt32(MaxArg+1))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1303(t *testing.T) {

	// -- say Rexx("").delword(Rexx("3"), Rexx("2")).Length() -- both test test length

	value, err := rxfromempty().DelWord(RxFromString("3"), RxFromString("2"))

	want := 0

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := len(value.ToString())
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1304(t *testing.T) {

	// -- say Rexx("Now is the  time").delword(Rexx("2"), Rexx("2"))

	value, err := RxFromString("Now is the  time").DelWord(RxFromString("2"), RxFromString("2"))
	want := "Now time"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1305(t *testing.T) {

	// -- say Rexx("Now is the  time ").delword(Rexx("3"), Rexx("4"))

	value, err := RxFromString("Now is the  time ").DelWord(RxFromString("3"), RxFromString("4"))
	want := "Now is "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1306(t *testing.T) {

	// -- say Rexx("Now  time").delword(Rexx("5"), Rexx("2"))

	value, err := RxFromString("Now  time").DelWord(RxFromString("5"), RxFromString("2"))
	want := "Now  time"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Format(before *Rexx, after *Rexx, explaces *Rexx, exdigits *Rexx, exform *Rexx) (*Rexx, error)

func Test_1307(t *testing.T) {

	// -- say Rexx("-a.bc.G-f").format(null, Rexx("3"), null, Rexx("0"), null) -- java.lang.NumberFormatException: -a.bc.G-f

	_, err := rx([]rune("-a.bc.G-f"), false).Format(nil, RxFromString("3"), nil, RxFromString("0"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1308(t *testing.T) {

	// -- say Rexx("-12345.73").format(null, Rexx("3"), null, Rexx("0"), null)

	value, err := rx([]rune("-12345.73"), true).Format(nil, RxFromString("3"), nil, RxFromString("0"), nil)
	want := "-1.235E+4"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1309(t *testing.T) {

	// -- say Rexx("1").format(Rexx("-1"), null, null, null, null)  -- netrexx.lang.BadArgumentException: Argument -1 < 1

	_, err := RxFromString("1").Format(RxFromString("-1"), nil, nil, nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1310(t *testing.T) {

	// -- say Rexx("1.73").format(Rexx("4"), Rexx("3"), null, null, null)

	value, err := RxFromString("1.73").Format(RxFromString("4"), RxFromString("3"), nil, nil, nil)
	want := "   1.730"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1311(t *testing.T) {

	// -- say Rexx("1.73").format(null, null, null, null, null)

	value, err := RxFromString("1.73").Format(nil, nil, nil, nil, nil)
	want := "1.73"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1312(t *testing.T) {

	// -- say Rexx("3.03").format(Rexx("4"), null, null, null, null)

	value, err := RxFromString("3.03").Format(RxFromString("4"), nil, nil, nil, nil)
	want := "   3.03"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1313(t *testing.T) {

	// -- say Rexx("1").format(null, Rexx("-1"), null, null, null) -- netrexx.lang.BadArgumentException: Argument -1 < 0

	_, err := RxFromString("1").Format(nil, RxFromString("-1"), nil, nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1314(t *testing.T) {

	// -- say Rexx(" - 12.73").format(null, Rexx("4"), null, null, null)

	value, err := RxFromString(" - 12.73").Format(nil, RxFromString("4"), nil, nil, nil)
	want := "-12.7300"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1315(t *testing.T) {

	// -- say Rexx("-.76").format(Rexx("4"), Rexx("1"), null, null, null)

	value, err := RxFromString("-.76").Format(RxFromString("4"), RxFromString("1"), nil, nil, nil)
	want := "  -0.8"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1316(t *testing.T) {

	// -- say Rexx("1").format(null, null, Rexx("-1"), null, null) -- netrexx.lang.BadArgumentException: Argument -1 < 1

	_, err := RxFromString("1").Format(nil, nil, RxFromString("-1"), nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1317(t *testing.T) {

	// -- say Rexx("12345.73").format(null, null, Rexx("3"), Rexx("6"), null)

	value, err := RxFromString("12345.73").Format(nil, nil, RxFromString("3"), RxFromString("6"), nil)
	want := "12345.73     "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1318(t *testing.T) {

	// -- say Rexx("12345e+5").format(null, Rexx("3"), null, null, null)

	value, err := RxFromString("12345e+5").Format(nil, RxFromString("3"), nil, nil, nil)
	want := "1234500000.000"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1319(t *testing.T) {

	// -- say Rexx("12345e+5").format(null, Rexx("3"), null, Rexx("-1"), null) -- netrexx.lang.BadArgumentException: Argument -1 < 0

	_, err := RxFromString("12345e+5").Format(nil, RxFromString("3"), nil, RxFromString("-1"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1320(t *testing.T) {

	// -- say Rexx("123.45").format(null, Rexx("3"), Rexx("2"), Rexx("0"), null)

	value, err := RxFromString("123.45").Format(nil, RxFromString("3"), RxFromString("2"), RxFromString("0"), nil)
	want := "1.235E+02"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1321(t *testing.T) {

	// -- say Rexx("1.2345").format(null, Rexx("3"), Rexx("2"), Rexx("0"), null)

	value, err := RxFromString("1.2345").Format(nil, RxFromString("3"), RxFromString("2"), RxFromString("0"), nil)
	want := "1.235    "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1322(t *testing.T) {

	// -- say Rexx(char '3').format(Rexx("4"), null, null, null, Rexx("")) -- netrexx.lang.BadArgumentException: Null option string

	_, err := RxFromRune('3').Format(RxFromString("4"), nil, nil, nil, rx(nil, false))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1323(t *testing.T) {

	// -- say Rexx("+0000000000000000009.00000000100000000").format(null, null, Rexx("2"), Rexx("0"), Rexx(char 'e'))

	value, err := RxFromString("+0000000000000000009.00000000100000000").Format(nil, nil, RxFromString("2"), RxFromString("0"), RxFromRune('e'))
	want := "9.00000000100000000    "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1324(t *testing.T) {

	// -- say Rexx("2.73").format(Rexx(int 999999999), Rexx("5"), null, null, null)

	_, err := RxFromString("2.73").Format(RxFromInt32(MaxExp), RxFromString("5"), nil, nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1325(t *testing.T) {

	// -- say Rexx("1.73").format(Rexx("4"), Rexx("0"), null, null, null)

	value, err := RxFromString("1.73").Format(RxFromString("4"), RxFromString("0"), nil, nil, nil)
	want := "   2"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1326(t *testing.T) {

	// -- say Rexx("0.1E-510").format(null, null, null, Rexx(0), Rexx(char 'e'))

	value, err := RxFromString("0.1E-510").Format(nil, nil, nil, zero, RxFromRune('e'))
	want := "100E-513"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) HashCode() int32

func Test_1327(t *testing.T) {

	// -- say Rexx("").hashCode()

	got := rxfromempty().HashCode()
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1328(t *testing.T) {

	// -- say Rexx("").hashCode()

	got := RxFromString("").HashCode()
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1329(t *testing.T) {

	// -- say Rexx("123456789012345").hashCode()

	got := RxFromString("123456789012345").HashCode()
	want := int32(20748692)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1330(t *testing.T) {

	// -- say Rexx(char 'A').hashCode()

	got := RxFromRune('A').HashCode()
	want := int32(195)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1331(t *testing.T) {

	// -- say Rexx(int 999999999 + 1).hashCode()

	got := RxFromInt32(MaxArg + 1).HashCode()
	want := int32(408146)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) Insert(_new *Rexx, n *Rexx, length *Rexx, pad *Rexx) (*Rexx, error)

func Test_1332(t *testing.T) {

	// -- say Rexx("abc").insert(Rexx("123"), null, Rexx("123").length(), Rexx(" ")) -- java.lang.NullPointerException

	_, err := RxFromString("abc").Insert(RxFromString("123"), nil, RxFromString("123").Length(), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1333(t *testing.T) {

	// -- say Rexx("").insert(null, null, null, null) -- java.lang.NullPointerException

	_, err := RxFromString("").Insert(nil, nil, nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1334(t *testing.T) {

	// -- say Rexx("abcdef").insert(Rexx(char ' '), Rexx(int 999999999+1), Rexx(char ' ').length(), Rexx(" "))

	_, err := RxFromString("abcdef").Insert(RxFromRune(' '), RxFromInt32(MaxArg+1), RxFromRune(' ').Length(), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1335(t *testing.T) {

	// -- say Rexx("abc").insert(Rexx("123"), Rexx("0"), null, Rexx(" ")) -- java.lang.NullPointerException

	_, err := RxFromString("abc").Insert(RxFromString("123"), RxFromString("0"), nil, RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1336(t *testing.T) {

	// -- say Rexx("abcdef").insert(Rexx(char ' '), Rexx("1"), Rexx(int 999999999+1), Rexx(" "))

	_, err := RxFromString("abcdef").Insert(RxFromRune(' '), RxFromString("1"), RxFromInt32(MaxArg+1), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1337(t *testing.T) {

	// -- say Rexx("abc").insert(Rexx("123"), Rexx("0"), Rexx("123").length(), null) -- java.lang.NullPointerException

	_, err := RxFromString("abc").Insert(RxFromString("123"), RxFromString("0"), RxFromString("123").Length(), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1338(t *testing.T) {

	// -- say Rexx("abcdef").insert(Rexx(char ' '), Rexx("1"), Rexx("1"), Rexx("ABC")) -- netrexx.lang.NotCharacterException: ABC

	_, err := RxFromString("abcdef").Insert(RxFromRune(' '), RxFromString("1"), RxFromString("1"), RxFromString("ABC"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1339(t *testing.T) {

	// -- say Rexx("").insert(Rexx("a"), Rexx("1"), Rexx("1"), Rexx("A"))

	value, err := rxfromempty().Insert(RxFromString("a"), RxFromString("1"), RxFromString("1"), RxFromString("A"))
	want := "Aa"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1340(t *testing.T) {

	// -- say Rexx("abc").insert(null, Rexx("0"), Rexx("123").length(), Rexx(" ")) -- java.lang.NullPointerException

	_, err := RxFromString("abc").Insert(nil, RxFromString("0"), RxFromString("123").Length(), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1341(t *testing.T) {

	// -- say Rexx("abcdef").insert(Rexx(""), Rexx("1"), Rexx("1"), Rexx("A"))

	value, err := RxFromString("abcdef").Insert(rxfromempty(), RxFromString("1"), RxFromString("1"), RxFromString("A"))
	want := "aAbcdef"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1342(t *testing.T) {

	// -- say Rexx("abc").insert(Rexx("123"), Rexx("0"), Rexx("123").length(), Rexx(" "))

	value, err := RxFromString("abc").Insert(RxFromString("123"), RxFromString("0"), RxFromString("123").Length(), RxFromString(" "))
	want := "123abc"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1343(t *testing.T) {

	// -- say Rexx("abcdef").insert(Rexx(char ' '), Rexx("3"), Rexx(char ' ').length(), Rexx(" "))

	value, err := RxFromString("abcdef").Insert(RxFromRune(' '), RxFromString("3"), RxFromRune(' ').Length(), RxFromString(" "))
	want := "abc def"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1344(t *testing.T) {

	// -- say Rexx("abc").insert(Rexx("123"), Rexx("5"), Rexx("6"), Rexx(" "))

	value, err := RxFromString("abc").Insert(RxFromString("123"), RxFromString("5"), RxFromString("6"), RxFromString(" "))
	want := "abc  123   "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1345(t *testing.T) {

	// -- say Rexx("abc").insert(Rexx("123"), Rexx("5"), Rexx("6"), Rexx(char '+'))

	value, err := RxFromString("abc").Insert(RxFromString("123"), RxFromString("5"), RxFromString("6"), RxFromRune('+'))
	want := "abc++123+++"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1346(t *testing.T) {

	// -- say Rexx("abc").insert(Rexx("123"), Rexx("0"), Rexx("5"), Rexx(char '-'))

	value, err := RxFromString("abc").Insert(RxFromString("123"), RxFromString("0"), RxFromString("5"), RxFromRune('-'))
	want := "123--abc"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) LastPos(needle *Rexx, start *Rexx) (*Rexx, error)

func Test_1347(t *testing.T) {

	// -- say Rexx("abc def ghi").lastpos(Rexx("2"), null) -- java.lang.NullPointerException

	_, err := RxFromString("abc def ghi").LastPos(RxFromString("2"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1348(t *testing.T) {

	// -- say Rexx("abc def ghi").lastpos(null, Rexx("1")) -- java.lang.NullPointerException

	_, err := RxFromString("abc def ghi").LastPos(nil, RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1349(t *testing.T) {

	// -- say Rexx("abc def ghi").lastpos(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("abc def ghi").LastPos(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1350(t *testing.T) {

	// -- say Rexx("a").lastpos(Rexx("a"), Rexx(int 999999999+1))

	_, err := RxFromString("a").LastPos(RxFromString("a"), RxFromInt32(MaxArg+1))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1351(t *testing.T) {

	// -- say Rexx("").lastpos(Rexx("a"), Rexx("1"))

	value, err := rxfromempty().LastPos(RxFromString("a"), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1352(t *testing.T) {

	// -- say Rexx("abc def ghi").lastpos(Rexx(char ' '), Rexx("77777"))

	value, err := RxFromString("abc def ghi").LastPos(RxFromRune(' '), RxFromString("77777"))
	want := "8"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1353(t *testing.T) {

	// -- say Rexx("a").lastpos(Rexx(""), Rexx("1"))

	value, err := RxFromString("a").LastPos(rxfromempty(), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1354(t *testing.T) {

	// -- say Rexx("abc def ghi").lastpos(Rexx(char ' '), Rexx("2"))

	value, err := RxFromString("abc def ghi").LastPos(RxFromRune(' '), RxFromString("2"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1355(t *testing.T) {

	// -- say Rexx("abc def ghi").lastpos(Rexx(char ' '), Rexx("4"))

	value, err := RxFromString("abc def ghi").LastPos(RxFromRune(' '), RxFromString("4"))
	want := "4"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1356(t *testing.T) {

	// -- say Rexx("abc def ghi").lastpos(Rexx(char ' '), Rexx("7"))

	value, err := RxFromString("abc def ghi").LastPos(RxFromRune(' '), RxFromString("7"))
	want := "4"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1357(t *testing.T) {

	// -- say Rexx("abc def ghi").lastpos(Rexx("2"), Rexx("1"))

	value, err := RxFromString("abc def ghi").LastPos(RxFromString("2"), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Left(length *Rexx, pad *Rexx) (*Rexx, error)

func Test_1358(t *testing.T) {

	// -- say Rexx("abc d").left(Rexx("2"), null) -- java.lang.NullPointerException

	_, err := RxFromString("abc d").Left(RxFromString("2"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1359(t *testing.T) {

	// -- say Rexx("abc d").left(null, Rexx(" ")) -- java.lang.NullPointerException

	_, err := RxFromString("abc d").Left(nil, RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1360(t *testing.T) {

	// -- say Rexx("abc d").left(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("abc d").Left(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1361(t *testing.T) {

	// -- say Rexx("a").left(Rexx(int 999999999+1), Rexx("1"))

	_, err := RxFromString("a").Left(RxFromInt32(MaxArg+1), RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1362(t *testing.T) {

	// -- say Rexx("abc d").left(Rexx("2"), Rexx(" "))

	value, err := RxFromString("abc d").Left(RxFromString("2"), RxFromString(" "))
	want := "ab"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1363(t *testing.T) {

	// -- say Rexx("abc d").left(Rexx("8"), Rexx(" "))

	value, err := RxFromString("abc d").Left(RxFromString("8"), RxFromString(" "))
	want := "abc d   "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1364(t *testing.T) {

	// -- say Rexx("abc d").left(Rexx("8"), Rexx(char '.'))

	value, err := RxFromString("abc d").Left(RxFromString("8"), RxFromRune('.'))
	want := "abc d..."

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1365(t *testing.T) {

	// -- say Rexx("abc defg").left(Rexx("6"), Rexx(" "))

	value, err := RxFromString("abc defg").Left(RxFromString("6"), RxFromString(" "))
	want := "abc de"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Length() *Rexx

func Test_1366(t *testing.T) {

	// -- say Rexx("abc d").length()

	got := RxFromString("abc d").Length()
	want := "5"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

// func (rcvr *Rexx) Lower(n *Rexx, length *Rexx) (*Rexx, error)

func Test_1367(t *testing.T) {

	// -- say Rexx("SumA").lower(Rexx("2"), null) -- java.lang.NullPointerException

	_, err := RxFromString("SumA").Lower(RxFromString("2"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1368(t *testing.T) {

	// -- say Rexx("SumA").lower(null, Rexx("1")) -- java.lang.NullPointerException

	_, err := RxFromString("SumA").Lower(nil, RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1369(t *testing.T) {

	// -- say Rexx("SumA").lower(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("SumA").Lower(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1370(t *testing.T) {

	// -- say Rexx("").lower(Rexx("-11"), Rexx("").length()) -- netrexx.lang.BadArgumentException: Argument -11 < 1

	_, err := RxFromString("").Lower(RxFromString("-11"), RxFromString("").Length())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1371(t *testing.T) {

	// -- say Rexx("SumA").lower(Rexx(int 999999999+1), Rexx("SumA").length())

	_, err := RxFromString("SumA").Lower(RxFromInt32(MaxArg+1), RxFromString("SumA").Length())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1372(t *testing.T) {

	// -- say Rexx("SUMB").lower(Rexx("2"), Rexx("-2")) -- netrexx.lang.BadArgumentException: Argument -2 < 0

	_, err := RxFromString("SUMB").Lower(RxFromString("2"), RxFromString("-2"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1373(t *testing.T) {

	// -- say Rexx("SumA").lower(Rexx("1"), Rexx(int 999999999+1))

	_, err := RxFromString("SumA").Lower(RxFromString("1"), RxFromInt32(MaxArg+1))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1374(t *testing.T) {

	// -- say Rexx("").lower(Rexx("1"), Rexx("1"))

	value, err := rxfromempty().Lower(RxFromString("1"), RxFromString("1"))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1375(t *testing.T) {

	// -- say Rexx("").lower(Rexx("1"), Rexx("1"))

	value, err := RxFromString("").Lower(RxFromString("1"), RxFromString("1"))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1376(t *testing.T) {

	// -- say Rexx("").lower(Rexx("1"), Rexx("").length())

	value, err := RxFromString("").Lower(RxFromString("1"), RxFromString("").Length())
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1377(t *testing.T) {

	// -- say Rexx("SUMB").lower(Rexx("2"), Rexx("2"))

	value, err := RxFromString("SUMB").Lower(RxFromString("2"), RxFromString("2"))
	want := "SumB"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1378(t *testing.T) {

	// -- say Rexx("SumA").lower(Rexx("1"), Rexx("SumA").length())

	value, err := RxFromString("SumA").Lower(RxFromString("1"), RxFromString("SumA").Length())
	want := "suma"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1379(t *testing.T) {

	// -- say Rexx("SuMB").lower(Rexx("1"), Rexx("1"))

	value, err := RxFromString("SuMB").Lower(RxFromString("1"), RxFromString("1"))
	want := "suMB"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1380(t *testing.T) {

	// -- say Rexx("SumA").lower(Rexx("2"), Rexx("1"))

	value, err := RxFromString("SumA").Lower(RxFromString("2"), RxFromString("1"))
	want := "SumA"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1381(t *testing.T) {

	// -- say Rexx("SumA").lower(Rexx("2"), Rexx("SumA").length())

	value, err := RxFromString("SumA").Lower(RxFromString("2"), RxFromString("SumA").Length())
	want := "Suma"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Max(rhs *Rexx) (*Rexx, error)

func Test_1382(t *testing.T) {

	// -- say Rexx("5.00").max(null) -- java.lang.NullPointerException

	_, err := RxFromString("5.00").Max(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1383(t *testing.T) {

	// -- say Rexx("A").max(Rexx("1")) -- java.lang.NumberFormatException: A

	_, err := RxFromString("A").Max(RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1384(t *testing.T) {

	// -- say Rexx("1").max(Rexx("A")) -- java.lang.NumberFormatException: 1

	_, err := RxFromString("1").Max(RxFromString("A"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1385(t *testing.T) {

	// -- say Rexx("123456700000").max(Rexx("1234567E+5"))

	value, err := RxFromString("123456700000").Max(RxFromString("1234567E+5"))
	want := "123456700000"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1386(t *testing.T) {

	// -- say Rexx("1234567E+5").max(Rexx("123456700000"))

	value, err := RxFromString("1234567E+5").Max(RxFromString("123456700000"))
	want := "1.234567E+11"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1387(t *testing.T) {

	// -- say Rexx("9.47109959E+230565093").max(Rexx("73354723.2"))

	value, err := RxFromString("9.47109959E+230565093").Max(RxFromString("73354723.2"))
	want := "9.47109959E+230565093"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1388(t *testing.T) {

	// -- say Rexx("9.47109959E+230565093").max(Rexx("73354723.2"))

	MaxExp = 99 // set to trigger Exception

	_, err := RxFromString("9.47109959E+230565093").Max(RxFromString("73354723.2"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999 // then reset

}

func Test_1389(t *testing.T) {

	// -- say Rexx("-1").max(Rexx("1"))

	value, err := RxFromString("-1").Max(RxFromString("1"))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1390(t *testing.T) {

	// -- say Rexx("0").max(Rexx("1"))

	value, err := RxFromString("0").Max(RxFromString("1"))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1391(t *testing.T) {

	// -- say Rexx("1234").max(Rexx("1234567E+5"))

	value, err := RxFromString("1234").Max(RxFromString("1234567E+5"))
	want := "1.234567E+11"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1392(t *testing.T) {

	// -- say Rexx("5.00").max(Rexx("6"))

	value, err := RxFromString("5.00").Max(RxFromString("6"))
	want := "6"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1393(t *testing.T) {

	// -- say Rexx("+1").max(Rexx("-1"))

	value, err := RxFromString("+1").Max(RxFromString("-1"))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1394(t *testing.T) {

	// -- say Rexx("1.0").max(Rexx("1.00"))

	value, err := RxFromString("1.0").Max(RxFromString("1.00"))
	want := "1.0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1395(t *testing.T) {

	// -- say Rexx("1.00").max(Rexx("1.0"))

	value, err := RxFromString("1.00").Max(RxFromString("1.0"))
	want := "1.00"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1396(t *testing.T) {

	// -- say Rexx("100").max(Rexx(byte 0))

	value, err := RxFromString("100").Max(RxFromInt8(int8(0)))
	want := "100"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1397(t *testing.T) {

	// -- say Rexx("101010101").max(Rexx("010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011"))

	value, err := RxFromString("101010101").Max(RxFromString("010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011"))
	want := "10101011010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1398(t *testing.T) {

	// -- say Rexx("010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011").max(Rexx("101010101"))

	value, err := RxFromString("010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011").Max(RxFromString("101010101"))
	want := "10101011010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1399(t *testing.T) {

	// -- say Rexx(int -999999999 - 1).max(Rexx(int -999999999 - 1)) -- NetRexx returns number in a different form

	value, err := RxFromInt32(MinArg - 1).Max(RxFromInt32(MinArg - 1))
	want := "-1000000000"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1400(t *testing.T) {

	// -- say Rexx(int 999999999 + 1).max(Rexx(int 999999999 + 1))

	value, err := RxFromInt32(MaxArg + 1).Max(RxFromInt32(MaxArg + 1))
	want := "1000000000"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1401(t *testing.T) {

	// -- say Rexx("9.47109959E+230565093").max(Rexx("73354723.2"))

	value, err := RxFromString("9.47109959E+230565093").Max(RxFromString("73354723.2"))
	want := "9.47109959E+230565093"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Min(rhs *Rexx) (*Rexx, error)

func Test_1402(t *testing.T) {

	// -- say Rexx("5.00").min(null) -- java.lang.NullPointerException

	_, err := RxFromString("5.00").Min(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1403(t *testing.T) {

	// -- say Rexx("A").min(Rexx("1")) -- java.lang.NumberFormatException: A

	_, err := RxFromString("A").Min(RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1404(t *testing.T) {

	// -- say Rexx("1").min(Rexx("A")) -- java.lang.NumberFormatException: 1

	_, err := RxFromString("1").Min(RxFromString("A"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1405(t *testing.T) {

	// -- say Rexx("123456700000").min(Rexx("1234567E+5"))

	value, err := RxFromString("123456700000").Min(RxFromString("1234567E+5"))
	want := "123456700000"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1406(t *testing.T) {

	// -- say Rexx("1234567E+5").min(Rexx("123456700000"))

	value, err := RxFromString("1234567E+5").Min(RxFromString("123456700000"))
	want := "1.234567E+11"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1407(t *testing.T) {

	// -- say Rexx("9.47109959E+230565093").min(Rexx("73354723.2")) -- test uses variables not accessible in NetRexx

	MaxExp = 99

	_, err := RxFromString("9.47109959E+230565093").Min(RxFromString("73354723.2"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999

}

func Test_1408(t *testing.T) {

	// -- say Rexx("+1").min(Rexx("-1"))

	value, err := RxFromString("+1").Min(RxFromString("-1"))
	want := "-1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1409(t *testing.T) {

	// -- say Rexx("100").min(Rexx(byte 0))

	value, err := RxFromString("100").Min(RxFromInt8(int8(0)))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1410(t *testing.T) {

	// -- say Rexx("010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011").min(Rexx("101010101"))

	value, err := RxFromString("010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011").Min(RxFromString("101010101"))
	want := "101010101"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1411(t *testing.T) {

	// -- say Rexx("-1").min(Rexx("1"))

	value, err := RxFromString("-1").Min(RxFromString("1"))
	want := "-1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1412(t *testing.T) {

	// -- say Rexx("0").min(Rexx("1"))

	value, err := RxFromString("0").Min(RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1413(t *testing.T) {

	// -- say Rexx("1234").min(Rexx("1234567E+5"))

	value, err := RxFromString("1234").Min(RxFromString("1234567E+5"))
	want := "1234"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1414(t *testing.T) {

	// -- say Rexx("5.00").min(Rexx("6"))

	value, err := RxFromString("5.00").Min(RxFromString("6"))
	want := "5.00"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1415(t *testing.T) {

	// -- say Rexx("1.0").min(Rexx("1.00"))

	value, err := RxFromString("1.0").Min(RxFromString("1.00"))
	want := "1.0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1416(t *testing.T) {

	// -- say Rexx("1.00").min(Rexx("1.0"))

	value, err := RxFromString("1.00").Min(RxFromString("1.0"))
	want := "1.00"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1417(t *testing.T) {

	// -- say Rexx("101010101").min(Rexx("010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011"))

	value, err := RxFromString("101010101").Min(RxFromString("010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011010101011"))
	want := "101010101"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1418(t *testing.T) {

	// -- say Rexx(int -999999999 - 1).min(Rexx(int -999999999 - 1))  -- NetRexx returns number in a different form

	value, err := RxFromInt32(MinArg - 1).Min(RxFromInt32(MinArg - 1))
	want := "-1000000000"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1419(t *testing.T) {

	// -- say Rexx(int 999999999 + 1).min(Rexx(int 999999999 + 1))

	value, err := RxFromInt32(MaxArg + 1).Min(RxFromInt32(MaxArg + 1))
	want := "1000000000"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1420(t *testing.T) {

	// -- say Rexx("9").min(Rexx("7"))

	value, err := RxFromString("9").Min(RxFromString("7"))
	want := "7"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpAdd(set *RexxSet, rhs *Rexx) (*Rexx, error)

func Test_1421(t *testing.T) {

	// -- say Rexx("-123").OpAdd(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("-123").OpAdd(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1422(t *testing.T) {

	// -- say Rexx("abc").OpAdd(null, Rexx("-123")) -- java.lang.NumberFormatException: abc

	_, err := RxFromString("abc").OpAdd(nil, RxFromString("-123"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1423(t *testing.T) {

	// -- say Rexx("-123").OpAdd(null, Rexx("abc")) -- java.lang.NumberFormatException: abc

	_, err := RxFromString("-123").OpAdd(nil, RxFromString("abc"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1424(t *testing.T) {

	// -- say Rexx().OpAdd(null, Rexx("abc")) -- java.lang.NullPointerException

	_, err := rxfromempty().OpAdd(nil, RxFromString("abc"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1425(t *testing.T) {

	// -- say Rexx("-123").OpAdd(null, Rexx("-1"))

	value, err := RxFromString("-123").OpAdd(nil, RxFromString("-1"))
	want := "-124"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1426(t *testing.T) {

	// -- say Rexx("-123").OpAdd(RexxSet(8, 1), Rexx("-1")) work around

	value, err := RxFromString("-123").OpAdd(RxSetWithDigitandForm(8, int8(1)), RxFromString("-1"))
	want := "-124"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1427(t *testing.T) {

	// -- say Rexx("0").OpAdd(null, Rexx("-9.999E+999999999"))

	value, err := RxFromString("0").OpAdd(nil, RxFromString("-9.999E+999999999"))
	want := "-9.999E+999999999"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1428(t *testing.T) {

	// -- say Rexx("0").OpAdd(null, Rexx("-9.999E+999999999"))

	MaxExp = 99

	_, err := RxFromString("0").OpAdd(nil, RxFromString("-9.999E+999999999"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999 // reset

}

func Test_1429(t *testing.T) {

	// -- say Rexx("0").OpAdd(null, Rexx("-9.999E+999999999"))

	value, err := RxFromString("0").OpAdd(nil, RxFromString("-9.999E+999999999"))
	want := "-9.999E+999999999"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1430(t *testing.T) {

	// -- say Rexx.toRexx("-0.9e-999999999").OpAdd(null, Rexx.toRexx("-00")) -- netrexx.lang.ExponentOverflowException: -1000000000

	_, err := ToRxFromString("-0.9e-999999999").OpAdd(nil, ToRxFromString("-00"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1431(t *testing.T) {

	// -- say Rexx("5").OpAdd(null, Rexx("0"))

	value, err := RxFromString("5").OpAdd(nil, RxFromString("0"))
	want := "5"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1432(t *testing.T) {

	// -- say Rexx("-34359738368").OpAdd(null, Rexx(byte -128))

	value, err := RxFromString("-34359738368").OpAdd(nil, RxFromInt8(int8(-128)))
	want := "-3.43597385E+10"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1433(t *testing.T) {

	// -- say Rexx(int 65).OpAdd(RexxSet(0, byte 0), Rexx("1.0"))

	value, err := RxFromInt32(65).OpAdd(RxSetWithDigitandForm(0, int8(0)), RxFromString("1.0"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1434(t *testing.T) {

	// -- say Rexx(int 120000000).OpAdd(null, Rexx("-999999999999"))

	value, err := RxFromInt32(120000000).OpAdd(nil, RxFromString("-999999999999"))
	want := "-9.99880000E+11"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1435(t *testing.T) {

	// -- say Rexx("-123").OpAdd(null, Rexx("2"))

	value, err := RxFromString("-123").OpAdd(nil, RxFromString("2"))
	want := "-121"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1436(t *testing.T) {

	// -- say Rexx("-56746.8689E+934981942").OpAdd(RexxSet(17, 1), Rexx("471002521."))

	value, err := RxFromString("-56746.8689E+934981942").OpAdd(RxSetWithDigitandForm(17, int8(1)), RxFromString("471002521."))
	want := "-567.46868900000000E+934981944"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1437(t *testing.T) {

	// -- say Rexx("-56746.8689E+934981942").OpAdd(RexxSet(17, 1), Rexx("471002521."))

	value, err := RxFromString("-56746.8689E+934981942").OpAdd(RxSetWithDigitandForm(17, int8(1)), RxFromString("471002521."))
	want := "-567.46868900000000E+934981944"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1438(t *testing.T) {

	// -- say Rexx("-56746.8689E+934981942").OpAdd(RexxSet(17, 1), Rexx("471002521."))

	MaxExp = 99

	_, err := RxFromString("-56746.8689E+934981942").OpAdd(RxSetWithDigitandForm(17, int8(1)), RxFromString("471002521."))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999 // reset

}

func Test_1439(t *testing.T) {

	// -- say Rexx("-34359738368").OpAdd(null, Rexx(byte -128))

	value, err := RxFromString("-34359738368").OpAdd(nil, RxFromInt8(int8(-128)))
	want := "-3.43597385E+10"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1440(t *testing.T) {

	// -- say Rexx("  1E+5").OpAdd(null, Rexx(byte -128))

	value, err := RxFromString("  1E+5").OpAdd(nil, RxFromInt8(int8(-128)))
	want := "99872"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1441(t *testing.T) {

	// -- say Rexx("  1E+5").OpAdd(null, Rexx(byte -128))

	value, err := RxFromString("  1E+5").OpAdd(nil, RxFromInt8(int8(-128)))
	want := "99872"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1442(t *testing.T) {

	// -- say Rexx("-34359738368").OpAdd(null, Rexx(byte -128))

	value, err := RxFromString("-34359738368").OpAdd(nil, RxFromInt8(int8(-128)))
	want := "-3.43597385E+10"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1443(t *testing.T) {

	// -- say Rexx("-000001.12345").OpAdd(RexxSet(17, 1), Rexx("7.62939453E+28"))

	value, err := RxFromString("-000001.12345").OpAdd(RxSetWithDigitandForm(17, int8(1)), RxFromString("7.62939453E+28"))
	want := "76.293945300000000E+27"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1444(t *testing.T) {

	// -- say Rexx("+1.23456789012345E-0").OpAdd(RexxSet(17, 1), Rexx("9E+999999999"))

	value, err := RxFromString("+1.23456789012345E-0").OpAdd(RxSetWithDigitandForm(17, int8(1)), RxFromString("9E+999999999"))
	want := "9.0000000000000000E+999999999"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1445(t *testing.T) {

	// -- say Rexx("+1.23456789012345E-0").OpAdd(RexxSet(17, 1), Rexx("9E+999999999"))

	MaxExp = 99

	_, err := RxFromString("+1.23456789012345E-0").OpAdd(RxSetWithDigitandForm(17, int8(1)), RxFromString("9E+999999999"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999 // reset

}

func Test_1446(t *testing.T) {

	// -- say Rexx(int 120000000).OpAdd(null, Rexx("-999999999999"))

	value, err := RxFromInt32(120000000).OpAdd(nil, RxFromString("-999999999999"))
	want := "-9.99880000E+11"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1447(t *testing.T) {

	// -- say Rexx("1.235E+02").OpAdd(null, Rexx(byte -128))

	value, err := RxFromString("1.235E+02").OpAdd(nil, RxFromInt8(int8(-128)))
	want := "-4.5"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1448(t *testing.T) {

	// -- say Rexx("1.369E+02").OpAdd(null, Rexx(byte -128))

	value, err := RxFromString("1.369E+02").OpAdd(nil, RxFromInt8(int8(-128)))
	want := "8.9"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1449(t *testing.T) {

	// -- say Rexx(int 120000000).OpAdd(null, Rexx("-999999999999"))

	value, err := RxFromInt32(120000000).OpAdd(nil, RxFromString("-999999999999"))
	want := "-9.99880000E+11"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1450(t *testing.T) {

	// -- say Rexx("1E+999999999").OpAdd(null, Rexx("9E+999999999")) -- netrexx.lang.ExponentOverflowException: 1000000000

	_, err := RxFromString("1E+999999999").OpAdd(nil, RxFromString("9E+999999999"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1451(t *testing.T) {

	// -- say Rexx("0.4444444444").OpAdd(null, Rexx("0.5555555550"))

	value, err := RxFromString("0.4444444444").OpAdd(nil, RxFromString("0.5555555550"))
	want := "0.999999999"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1452(t *testing.T) {

	// -- say Rexx("1").OpAdd(null, Rexx("۱"))

	value, err := RxFromString("1").OpAdd(nil, RxFromString("۱"))
	want := "2"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1453(t *testing.T) {

	// -- say Rexx("-123").OpAdd(null, Rexx("2"))

	value, err := RxFromString("-123").OpAdd(nil, RxFromString("2"))
	want := "-121"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1454(t *testing.T) {

	// -- say Rexx(int 120000000).OpAdd(null, Rexx("-999999999999"))

	value, err := RxFromInt32(120000000).OpAdd(nil, RxFromString("-999999999999"))
	want := "-9.99880000E+11"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1455(t *testing.T) {

	// -- say Rexx("1.235E+02").OpAdd(null, Rexx(byte -128))

	value, err := RxFromString("1.235E+02").OpAdd(nil, RxFromInt8(int8(-128)))
	want := "-4.5"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1456(t *testing.T) {

	// -- say Rexx("1").OpAdd(null, Rexx("-1.0"))

	value, err := RxFromString("1").OpAdd(nil, RxFromString("-1.0"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1457(t *testing.T) {

	// -- say Rexx("1.235E+02").OpAdd(RexxSet(17, 1), Rexx(byte -128))

	value, err := RxFromString("1.235E+02").OpAdd(RxSetWithDigitandForm(17, int8(1)), RxFromInt8(int8(-128)))
	want := "-4.5"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1458(t *testing.T) {

	// -- say Rexx("1.371E+02").OpAdd(null, Rexx(byte -128))

	value, err := RxFromString("1.371E+02").OpAdd(nil, RxFromInt8(int8(-128)))
	want := "9.1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1459(t *testing.T) {

	// -- say Rexx("10.23456785").OpAdd(RexxSet(9), Rexx("-10.23456789"))

	value, err := RxFromString("10.23456785").OpAdd(RxSetWithDigit(9), RxFromString("-10.23456789"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1460(t *testing.T) {

	// -- say Rexx(int 120000000).OpAdd(null, Rexx("-999999999999"))

	value, err := RxFromInt32(120000000).OpAdd(nil, RxFromString("-999999999999"))
	want := "-9.99880000E+11"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1461(t *testing.T) {

	// -- say Rexx("  1E+5").OpAdd(null, Rexx(byte -128))

	value, err := RxFromString("  1E+5").OpAdd(nil, RxFromInt8(int8(-128)))
	want := "99872"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1462(t *testing.T) {

	// -- say Rexx("  1E+5").OpAdd(null, Rexx(byte -128))

	value, err := RxFromString("  1E+5").OpAdd(nil, RxFromInt8(int8(-128)))
	want := "99872"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1463(t *testing.T) {

	// -- say Rexx("1.235E+02").OpAdd(RexxSet(17, 1), Rexx(byte -128))

	value, err := RxFromString("1.235E+02").OpAdd(RxSetWithDigitandForm(17, int8(1)), RxFromInt8(int8(-128)))
	want := "-4.5"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1464(t *testing.T) {

	// -- say Rexx("1.235E+02").OpAdd(RexxSet(17, 1), Rexx(byte -128))

	MaxExp = 99

	value, err := RxFromString("1.235E+02").OpAdd(RxSetWithDigitandForm(17, int8(1)), RxFromInt8(int8(-128)))
	want := "-4.5"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	MaxExp = 999999999 // reset

}

func Test_1465(t *testing.T) {

	// -- say Rexx(int 120000000).OpAdd(null, Rexx("-999999999999"))

	value, err := RxFromInt32(120000000).OpAdd(nil, RxFromString("-999999999999"))
	want := "-9.99880000E+11"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1466(t *testing.T) {

	// -- say Rexx("-56267E-10").OpAdd(RexxSet(9), Rexx("0"))
	// ToString() triggers code block in layout

	value, err := RxFromString("-56267E-10").OpAdd(RxSetWithDigit(9), RxFromString("0"))
	want := "-0.0000056267"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1467(t *testing.T) {

	// -- say Rexx("1E+2").OpAdd(RexxSet(15), Rexx("1E+4"))
	// ToString() triggers code blocks in layout

	value, err := RxFromString("1E+2").OpAdd(RxSetWithDigit(15), RxFromString("1E+4"))
	want := "10100"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1468(t *testing.T) {

	// --

	MaxArg = 0
	MinArg = 0
	MaxExp = 0
	MinExp = 0

	new_set := RxSet()
	new_set.SetForm(RxFromString("engineering"))
	new_set.SetDigits(RxFromString("3"))

	RxFromString("+00678.5432").OpAdd(new_set, RxFromString("9"))

	MaxArg = 999999999
	MinArg = -999999999
	MaxExp = 999999999
	MinExp = -999999999

}

func Test_1469(t *testing.T) {

	// -- say Rexx("0.4444444444").OpAdd(null, Rexx("0.5555555550"))

	value, err := RxFromString("0.4444444444").OpAdd(nil, RxFromString("0.5555555550"))
	want := "0.999999999"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpAnd(set *RexxSet, rhs *Rexx) (bool, error)

func Test_1470(t *testing.T) {

	// -- say Rexx("1").OpAnd(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("1").OpAnd(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1471(t *testing.T) {

	// -- say Rexx("2").OpAnd(null, Rexx()) -- java.lang.NullPointerException

	_, err := RxFromString("2").OpAnd(nil, rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1472(t *testing.T) {

	// -- say Rexx("").OpAnd(null, Rexx()) -- java.lang.NullPointerException

	_, err := RxFromString("").OpAnd(nil, rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1473(t *testing.T) {

	// -- say Rexx("0").OpAnd(null, Rexx("2")) -- netrexx.lang.NotLogicException: Boolean must be 0 or 1.  Found: 2

	_, err := RxFromString("0").OpAnd(nil, RxFromString("2"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1474(t *testing.T) {

	// -- say Rexx().OpAnd(null, Rexx())

	value, err := rxfromempty().OpAnd(nil, rxfromempty())
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1475(t *testing.T) {

	// -- say Rexx("0").OpAnd(null, Rexx())

	value, err := RxFromString("0").OpAnd(nil, rxfromempty())
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1476(t *testing.T) {

	// -- say Rexx("1").OpAnd(null, Rexx())

	value, err := RxFromString("1").OpAnd(nil, rxfromempty())
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1477(t *testing.T) {

	// -- say Rexx().OpAnd(null, Rexx("0"))

	value, err := rxfromempty().OpAnd(nil, RxFromString("0"))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1478(t *testing.T) {

	// -- say Rexx().OpAnd(null, Rexx("1"))

	value, err := rxfromempty().OpAnd(nil, RxFromString("1"))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1479(t *testing.T) {

	// -- say Rexx("0").OpAnd(null, Rexx("0"))

	value, err := RxFromString("0").OpAnd(nil, RxFromString("0"))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1480(t *testing.T) {

	// -- say Rexx("0").OpAnd(null, Rexx("1"))

	value, err := RxFromString("0").OpAnd(nil, RxFromString("1"))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1481(t *testing.T) {

	// -- say Rexx("1").OpAnd(null, Rexx("0"))

	value, err := RxFromString("1").OpAnd(nil, RxFromString("0"))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1482(t *testing.T) {

	// -- say Rexx("1").OpAnd(null, Rexx("1"))

	value, err := RxFromString("1").OpAnd(nil, RxFromString("1"))
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpCc(set *RexxSet, rhs *Rexx) *Rexx

func Test_1483(t *testing.T) {

	// -- say Rexx("1").OpCc(null, null) -- java.lang.NullPointerException
	// This port returns rcvr

	got := RxFromString("1").OpCc(nil, nil)
	want := "1"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1484(t *testing.T) {

	// -- say Rexx("1").OpCc(null, Rexx("1"))

	got := RxFromString("1").OpCc(nil, RxFromString("1"))
	want := "11"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1485(t *testing.T) {

	// -- say Rexx("").OpCc(null, Rexx("1"))

	got := RxFromString("").OpCc(nil, RxFromString("1"))
	want := "1"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1486(t *testing.T) {

	// -- say Rexx().OpCc(null, Rexx()) -- java.lang.NullPointerException
	// This port returns rcvr

	got := rxfromempty().OpCc(nil, rxfromempty())
	want := ""

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

// func (rcvr *Rexx) OpCcBlank(set *RexxSet, rhs *Rexx) *Rexx

func Test_1487(t *testing.T) {

	// -- say Rexx("1").OpCcBlank(null, null) -- java.lang.NullPointerException
	// -- This port returns rcvr

	got := RxFromString("1").OpCcBlank(nil, nil)
	want := "1"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_1488(t *testing.T) {

	// -- say Rexx("1").OpCcBlank(null, Rexx("1"))

	got := RxFromString("1").OpCcBlank(nil, RxFromString("1"))
	want := "1 1"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

// func (rcvr *Rexx) OpDiv(set *RexxSet, rhs *Rexx) (*Rexx, error)

func Test_1489(t *testing.T) {

	// -- say Rexx("1").OpDiv(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("1").OpDiv(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1490(t *testing.T) {

	// -- say Rexx("1").OpDiv(null, Rexx("1"))

	value, err := RxFromString("1").OpDiv(nil, RxFromString("1"))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1491(t *testing.T) {

	// -- say Rexx("a").OpDiv(RexxSet(), Rexx("2")) -- java.lang.NumberFormatException: a

	_, err := RxFromString("a").OpDiv(RxSet(), RxFromString("2"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1492(t *testing.T) {

	// -- say Rexx(char '\ubEEf').OpDiv(RexxSet(), Rexx("-1")) -- java.lang.NumberFormatException: 뻯

	_, err := RxFromRune('\ubEEf').OpDiv(RxSet(), RxFromString("-1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1493(t *testing.T) {

	// -- say Rexx("2").OpDiv(RexxSet(), Rexx("a")) -- java.lang.NumberFormatException: a

	_, err := RxFromString("2").OpDiv(RxSet(), RxFromString("a"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1494(t *testing.T) {

	// -- say Rexx("-0.9E-999999999").OpDiv(big_set, Rexx("-000001.12345")) -- netrexx.lang.ExponentOverflowException: -1000000002

	test_set := RxSet()

	test_set.SetForm(RxFromString("engineering"))

	_, err := RxFromString("-0.9E-999999999").OpDiv(test_set, RxFromString("-000001.12345"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	test_set.SetForm(RxFromInt8(DEFAULT_FORM)) // reset

}

func Test_1495(t *testing.T) {

	// -- say Rexx("1").OpDivI(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("1").OpDivI(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

// func (rcvr *Rexx) OpDivI(set *RexxSet, rhs *Rexx) (*Rexx, error)

func Test_1496(t *testing.T) {

	// -- say Rexx("1").OpDivI(null, Rexx("1"))

	value, err := RxFromString("1").OpDivI(nil, RxFromString("1"))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1497(t *testing.T) {

	// -- say Rexx("-296590035").OpDivI(RexxSet(), Rexx("-481734529"))

	value, err := RxFromString("-296590035").OpDivI(RxSet(), RxFromString("-481734529"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpEq(set *RexxSet, rhs *Rexx) (bool, error)

func Test_1498(t *testing.T) {

	// -- say Rexx("1").OpEq(null, Rexx())

	value, err := RxFromString("1").OpEq(nil, nil)
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1499(t *testing.T) {

	// -- say Rexx("9E+999999999").OpEq(RexxSet(17, byte 1), Rexx("+1.23456789012345E-0"))

	MaxExp = 99

	_, err := RxFromString("9E+999999999").OpEq(RxSetWithDigitandForm(17, int8(1)), RxFromString("+1.23456789012345E-0"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999 // reset

}

func Test_1500(t *testing.T) {

	// -- say Rexx("1").OpEq(null, Rexx("1"))

	value, err := RxFromString("1").OpEq(nil, RxFromString("1"))
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1501(t *testing.T) {

	// -- say Rexx("0.444444444444444").OpEq(RexxSet(15), Rexx("0.555555555555553"))

	value, err := RxFromString("0.444444444444444").OpEq(RxSetWithDigit(15), RxFromString("0.555555555555553"))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1502(t *testing.T) {

	// -- say Rexx("ab").OpEq(RexxSet(15), Rexx(char 'A'))

	value, err := RxFromString("ab").OpEq(RxSetWithDigit(15), RxFromRune('A'))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpEqS(set *RexxSet, rhs *Rexx) bool

func Test_1503(t *testing.T) {

	// -- say Rexx("1").OpEqS(null, Rexx(""))

	got := RxFromString("1").OpEqS(nil, nil)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1504(t *testing.T) {

	// -- say Rexx("1").OpEqS(null, Rexx("1"))

	got := RxFromString("1").OpEqS(nil, RxFromString("1"))
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) OpGt(set *RexxSet, rhs *Rexx) (bool, error)

func Test_1505(t *testing.T) {

	// -- say Rexx("1").OpGt(null, Rexx(""))

	value, err := RxFromString("1").OpGt(nil, nil)
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1506(t *testing.T) {

	// -- say Rexx("9E+999999999").OpGt(RexxSet(17, 1), Rexx("+1.23456789012345E-0"))

	value, err := RxFromString("9E+999999999").OpGt(RxSetWithDigitandForm(17, int8(1)), RxFromString("+1.23456789012345E-0"))
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1507(t *testing.T) {

	// -- say Rexx("9E+999999999").OpGt(RexxSet(17, 1), Rexx("+1.23456789012345E-0"))

	MaxExp = 99

	_, err := RxFromString("9E+999999999").OpGt(RxSetWithDigitandForm(17, int8(1)), RxFromString("+1.23456789012345E-0"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999 // reset

}

func Test_1508(t *testing.T) {

	// -- say Rexx("1").OpGt(null, Rexx("1"))

	value, err := RxFromString("1").OpGt(nil, RxFromString("1"))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpGtEq(set *RexxSet, rhs *Rexx) (bool, error)

func Test_1509(t *testing.T) {

	// -- say Rexx("1").OpGtEq(null, Rexx())

	value, err := RxFromString("1").OpGtEq(nil, nil)
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1510(t *testing.T) {

	// -- say Rexx("9E+999999999").OpGtEq(RexxSet(17, 1), Rexx("+1.23456789012345E-0"))

	value, err := RxFromString("9E+999999999").OpGtEq(RxSetWithDigitandForm(17, int8(1)), RxFromString("+1.23456789012345E-0"))
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1511(t *testing.T) {

	// -- say Rexx("9E+999999999").OpGtEq(RexxSet(17, 1), Rexx("+1.23456789012345E-0"))

	MaxExp = 99

	_, err := RxFromString("9E+999999999").OpGtEq(RxSetWithDigitandForm(17, int8(1)), RxFromString("+1.23456789012345E-0"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999 // reset

}

func Test_1512(t *testing.T) {

	// -- say Rexx("1").OpGtEq(null, Rexx("1"))

	value, err := RxFromString("1").OpGtEq(nil, RxFromString("1"))
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpGtEqS(set *RexxSet, rhs *Rexx) bool

func Test_1513(t *testing.T) {

	// -- say Rexx("1").OpGtEqS(null, Rexx(""))

	got := RxFromString("1").OpGtEqS(nil, nil)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1514(t *testing.T) {

	// -- say Rexx("1").OpGtEqS(null, Rexx("1"))

	got := RxFromString("1").OpGtEqS(nil, RxFromString("1"))
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) OpGtS(set *RexxSet, rhs *Rexx) bool

func Test_1515(t *testing.T) {

	// -- say Rexx("1").OpGtS(null, Rexx(""))

	got := RxFromString("1").OpGtS(nil, nil)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1516(t *testing.T) {

	// -- say Rexx("1").OpGtS(null, Rexx("1"))

	got := RxFromString("1").OpGtS(nil, RxFromString("1"))
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) OpLt(set *RexxSet, rhs *Rexx) (bool, error)

func Test_1517(t *testing.T) {

	// -- say Rexx("1").OpLt(null, Rexx(""))

	value, err := RxFromString("1").OpLt(nil, nil)
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1518(t *testing.T) {

	// -- say Rexx("9E+999999999").OpLt(RexxSet(17, 1), Rexx("+1.23456789012345E-0"))

	value, err := RxFromString("9E+999999999").OpLt(RxSetWithDigitandForm(17, int8(1)), RxFromString("+1.23456789012345E-0"))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1519(t *testing.T) {

	// -- say Rexx("9E+999999999").OpLt(RexxSet(17, 1), Rexx("+1.23456789012345E-0"))

	MaxExp = 99

	_, err := RxFromString("9E+999999999").OpLt(RxSetWithDigitandForm(17, int8(1)), RxFromString("+1.23456789012345E-0"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999 // reset

}

func Test_1520(t *testing.T) {

	// -- say Rexx("1").OpLt(null, Rexx("1"))

	value, err := RxFromString("1").OpLt(nil, RxFromString("1"))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpLtEq(set *RexxSet, rhs *Rexx) (bool, error)

func Test_1521(t *testing.T) {

	// -- say Rexx("1").OpLt(null, Rexx(""))

	value, err := RxFromString("1").OpLtEq(nil, nil)
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1522(t *testing.T) {

	// -- say Rexx("9E+999999999").OpLtEq(RexxSet(17, 1), Rexx("+1.23456789012345E-0"))

	value, err := RxFromString("9E+999999999").OpLtEq(RxSetWithDigitandForm(17, int8(1)), RxFromString("+1.23456789012345E-0"))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1523(t *testing.T) {

	// -- say Rexx("9E+999999999").OpLtEq(RexxSet(17, 1), Rexx("+1.23456789012345E-0"))

	MaxExp = 99

	_, err := RxFromString("9E+999999999").OpLtEq(RxSetWithDigitandForm(17, int8(1)), RxFromString("+1.23456789012345E-0"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999 // reset

}

func Test_1524(t *testing.T) {

	// -- say Rexx("1").OpLtEq(null, Rexx("1"))

	value, err := RxFromString("1").OpLtEq(nil, RxFromString("1"))
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpLtEqS(set *RexxSet, rhs *Rexx) bool

func Test_1525(t *testing.T) {

	// -- say Rexx("1").OpLtEqS(null, Rexx(""))

	got := RxFromString("1").OpLtEqS(nil, nil)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1526(t *testing.T) {

	// -- say Rexx("1").OpLtEqS(null, Rexx("1"))

	got := RxFromString("1").OpLtEqS(nil, RxFromString("1"))
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) OpLtS(set *RexxSet, rhs *Rexx) bool

func Test_1527(t *testing.T) {

	// -- say Rexx("1").OpLtS(null, Rexx(""))

	got := RxFromString("1").OpLtS(nil, nil)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1528(t *testing.T) {

	// -- say Rexx("1").OpLtS(null, Rexx("1"))

	got := RxFromString("1").OpLtS(nil, RxFromString("1"))
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) OpMinus(set *RexxSet) (*Rexx, error)

func Test_1529(t *testing.T) {

	// -- say Rexx("A").OpMinus(null) -- java.lang.NumberFormatException: A

	_, err := RxFromString("A").OpMinus(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1530(t *testing.T) {

	// -- say Rexx("1").OpMinus(null)

	value, err := RxFromString("1").OpMinus(nil)
	want := "-1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1531(t *testing.T) {

	// -- say Rexx("1").OpMinus(RexxSet(9, 1))

	value, err := RxFromString("1").OpMinus(RxSetWithDigitandForm(9, int8(1)))
	want := "-1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1532(t *testing.T) {

	// -- say Rexx.toRexx("-0.9e-999999999").OpMinus(null) -- netrexx.lang.ExponentOverflowException: -1000000000

	_, err := ToRxFromString("-0.9e-999999999").OpMinus(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1533(t *testing.T) {

	// -- say Rexx("-1").OpMinus(null)

	value, err := RxFromString("-1").OpMinus(nil)
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpMult(set *RexxSet, rhs *Rexx) (*Rexx, error)

func Test_1534(t *testing.T) {

	// -- say Rexx(int 50).OpMult(null, null) -- java.lang.NullPointerException

	_, err := RxFromInt32(50).OpMult(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1535(t *testing.T) {

	// -- say Rexx(char ' ').OpMult(null, Rexx(byte -128)) -- java.lang.NumberFormatException:

	_, err := RxFromRune(' ').OpMult(nil, RxFromInt8(int8(-128)))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1536(t *testing.T) {

	// -- say Rexx(int 50).OpMult(null, Rexx("a")) -- java.lang.NumberFormatException: a

	_, err := RxFromInt32(50).OpMult(nil, RxFromString("a"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1537(t *testing.T) {

	// -- say Rexx(byte -128).OpMult(null, Rexx("-34359738368"))

	value, err := RxFromInt8(int8(-128)).OpMult(nil, RxFromString("-34359738368"))
	want := "4.39804651E+12"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1538(t *testing.T) {

	// -- say Rexx(int 50).OpMult(RexxSet(9), Rexx(int 0))

	value, err := RxFromInt32(50).OpMult(RxSetWithDigit(9), RxFromInt32(0))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1539(t *testing.T) {

	// -- say Rexx("-34359738368").OpMult(null, Rexx(byte -128))

	value, err := RxFromString("-34359738368").OpMult(nil, RxFromInt8(int8(-128)))
	want := "4.39804651E+12"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1540(t *testing.T) {

	// -- say Rexx(byte -128).OpMult(null, Rexx("-34359738368"))

	value, err := RxFromInt8(int8(-128)).OpMult(nil, RxFromString("-34359738368"))
	want := "4.39804651E+12"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1541(t *testing.T) {

	// -- say Rexx("7E+777777777").OpMult(RexxSet(9), Rexx("+1.23456789012345E-0"))

	value, err := RxFromString("7E+777777777").OpMult(RxSetWithDigit(9), RxFromString("+1.23456789012345E-0"))
	want := "8.64197523E+777777777"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1542(t *testing.T) {

	// -- say Rexx("50.000000").OpMult(null, Rexx("17.000"))

	value, err := RxFromString("50.000000").OpMult(nil, RxFromString("17.000"))
	want := "850.000000"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1543(t *testing.T) {

	// -- say Rexx(int 50).OpMult(null, Rexx(int 17))

	value, err := RxFromInt32(50).OpMult(nil, RxFromInt32(17))
	want := "850"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1544(t *testing.T) {

	// -- say Rexx(int 50).OpMult(null, Rexx(int 2))

	value, err := RxFromInt32(50).OpMult(nil, RxFromInt32(2))
	want := "100"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1545(t *testing.T) {

	// -- say Rexx(int 17).OpMult(null, Rexx(int 2))

	value, err := RxFromInt32(17).OpMult(nil, RxFromInt32(2))
	want := "34"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1546(t *testing.T) {

	// -- say Rexx("9E+999999999").OpMult(RexxSet(9), Rexx("+1.23456789012345E-0")) -- netrexx.lang.ExponentOverflowException: 1000000000

	_, err := RxFromString("9E+999999999").OpMult(RxSetWithDigit(9), RxFromString("+1.23456789012345E-0"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1547(t *testing.T) {

	// -- say Rexx(int 1).OpMult(null, Rexx(int 1))

	value, err := RxFromInt32(1).OpMult(nil, RxFromInt32(1))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpNot(set *RexxSet) (bool, error)

func Test_1548(t *testing.T) {

	// -- say Rexx("").OpNot(null) -- netrexx.lang.NotLogicException: Boolean must be 0 or 1.  Found:

	_, err := RxFromString("").OpNot(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1549(t *testing.T) {

	// -- say Rexx("2").OpNot(null) -- netrexx.lang.NotLogicException: Boolean must be 0 or 1.  Found: 2

	_, err := RxFromString("2").OpNot(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1550(t *testing.T) {

	// -- say Rexx().OpNot(null)

	value, err := rxfromempty().OpNot(nil)

	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1551(t *testing.T) {

	// -- say Rexx("0").OpNot(null)

	value, err := RxFromString("0").OpNot(nil)
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1552(t *testing.T) {

	// -- say Rexx("1").OpNot(null)

	value, err := RxFromString("1").OpNot(nil)
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpNotEq(set *RexxSet, rhs *Rexx) (bool, error)

func Test_1553(t *testing.T) {

	// -- say Rexx("1").OpNotEq(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("1").OpNotEq(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1554(t *testing.T) {

	// -- say Rexx("9.47109959E+230565093").OpNotEq(null, Rexx("73354723.2"))

	value, err := RxFromString("9.47109959E+230565093").OpNotEq(nil, RxFromString("73354723.2"))
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1555(t *testing.T) {

	// -- say Rexx("9.47109959E+230565093").OpNotEq(null, Rexx("73354723.2"))

	MaxExp = 99

	_, err := RxFromString("9.47109959E+230565093").OpNotEq(nil, RxFromString("73354723.2"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999

}

func Test_1556(t *testing.T) {

	// -- say Rexx("0").OpNotEq(null, Rexx("1"))

	value, err := RxFromString("0").OpNotEq(nil, RxFromString("1"))
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1557(t *testing.T) {

	// -- say Rexx("").OpNotEq(null, Rexx(""))

	value, err := RxFromString("").OpNotEq(nil, RxFromString(""))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1558(t *testing.T) {

	// -- say Rexx(" ").OpNotEq(null, Rexx(" "))

	value, err := RxFromString(" ").OpNotEq(nil, RxFromString(" "))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1559(t *testing.T) {

	// -- say Rexx("").OpNotEq(null, Rexx(""))

	value, err := rxfromempty().OpNotEq(nil, RxFromString(""))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1560(t *testing.T) {

	// -- say Rexx("").OpNotEq(null, Rexx(""))

	value, err := RxFromString("").OpNotEq(nil, rxfromempty())
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1561(t *testing.T) {

	// -- say Rexx(" ").OpNotEq(null, Rexx(""))

	value, err := RxFromString(" ").OpNotEq(nil, rxfromempty())
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1562(t *testing.T) {

	// -- say Rexx("").OpNotEq(null, Rexx(""))

	value, err := rxfromempty().OpNotEq(nil, rxfromempty())
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpNotEqS(set *RexxSet, rhs *Rexx) bool

func Test_1563(t *testing.T) {

	// -- say Rexx("9.47109959E+230565093").OpNotEqS(null, Rexx("73354723.2"))

	got := RxFromString("9.47109959E+230565093").OpNotEqS(nil, RxFromString("73354723.2"))
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1564(t *testing.T) {

	// -- say Rexx("9.47109959E+230565093").OpNotEqS(null, null) -- java.lang.NullPointerException

	got := RxFromString("9.47109959E+230565093").OpNotEqS(nil, nil)
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) OpOr(set *RexxSet, rhs *Rexx) (bool, error)

func Test_1565(t *testing.T) {

	// -- say Rexx("1").OpOr(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("1").OpOr(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1566(t *testing.T) {

	// -- say Rexx("A").OpOr(null, Rexx("1")) -- netrexx.lang.NotLogicException: Boolean must be 0 or 1.  Found: A

	_, err := RxFromString("A").OpOr(nil, RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1567(t *testing.T) {

	// -- say Rexx("1").OpOr(null, Rexx("A")) -- netrexx.lang.NotLogicException: Boolean must be 0 or 1.  Found: A

	_, err := RxFromString("1").OpOr(nil, RxFromString("A"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1568(t *testing.T) {

	// -- say Rexx("1").OpOr(null, Rexx("1"))

	value, err := RxFromString("1").OpOr(nil, RxFromString("1"))
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpPlus(set *RexxSet) (*Rexx, error)

func Test_1569(t *testing.T) {

	// -- say Rexx("A").OpPlus(null) -- java.lang.NumberFormatException: A

	_, err := RxFromString("A").OpPlus(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1570(t *testing.T) {

	// -- say Rexx("1").OpPlus(null)

	value, err := RxFromString("1").OpPlus(nil)
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1571(t *testing.T) {

	// -- say Rexx("1").OpPlus(RexxSet(9, 1))

	value, err := RxFromString("1").OpPlus(RxSetWithDigitandForm(9, int8(1)))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1572(t *testing.T) {

	// -- say Rexx.toRexx("-0.9e-999999999").OpPlus(null) -- netrexx.lang.ExponentOverflowException: -1000000000

	_, err := ToRxFromString("-0.9e-999999999").OpPlus(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1573(t *testing.T) {

	// -- say Rexx("-1").OpPlus(null)

	value, err := RxFromString("-1").OpPlus(nil)
	want := "-1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpPow(set *RexxSet, rhs *Rexx) (*Rexx, error)

func Test_1574(t *testing.T) {

	// -- say Rexx("-123").OpPow(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("-123").OpPow(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1575(t *testing.T) {

	// -- say Rexx(char '/').OpPow(null, Rexx(byte -128)) -- java.lang.NumberFormatException: /

	_, err := RxFromRune('/').OpPow(nil, RxFromInt8(int8(-128)))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1576(t *testing.T) {

	// -- say Rexx(int 1).OpPow(null, Rexx(int 1))

	value, err := RxFromInt32(1).OpPow(nil, RxFromInt32(1))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1577(t *testing.T) {

	// -- say Rexx(int 1).OpPow(RexxSet(16), Rexx(int 1))

	value, err := RxFromInt32(1).OpPow(RxSetWithDigit(16), RxFromInt32(1))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1578(t *testing.T) {

	// -- say Rexx.toRexx("2").OpPow(RexxSet(9, byte 1), Rexx.toRexx("2.147483647E+8"))

	value, err := ToRxFromString("2").OpPow(RxSetWithDigitandForm(9, int8(1)), ToRxFromString("2.147483647E+8"))
	want := "24.2547639E+64645698"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1579(t *testing.T) {

	// -- say Rexx(byte -128).OpPow(null, Rexx("-34359738368")) -- netrexx.lang.BadArgumentException: Too many digits: -3.43597384E+10

	_, err := RxFromInt8(int8(-128)).OpPow(nil, RxFromString("-34359738368"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1580(t *testing.T) {

	// -- say Rexx("100").OpPow(RexxSet(16), Rexx("1E+9"))  -- netrexx.lang.BadArgumentException: Argument 1000000000 > 999999999

	_, err := RxFromString("100").OpPow(RxSetWithDigit(16), RxFromString("1E+9"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1581(t *testing.T) {

	// -- say Rexx("-34359738368").OpPow(null, Rexx(byte -128))

	value, err := RxFromString("-34359738368").OpPow(nil, RxFromInt8(int8(-128)))
	want := "2.43007366E-1349"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1582(t *testing.T) {

	// -- say Rexx.toRexx("3454").OpPow(RexxSet(17, byte 0), Rexx.toRexx("-12E+3"))

	value, err := ToRxFromString("3454").OpPow(RxSetWithDigitandForm(17, int8(0)), ToRxFromString("-12E+3"))
	want := "1.3551932277805279E-42460"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1583(t *testing.T) {

	// -- say Rexx.toRexx("+1.1E-999999999").OpPow(null, Rexx.toRexx("+10")) -- netrexx.lang.ExponentOverflowException: -1999999998

	_, err := ToRxFromString("+1.1E-999999999").OpPow(nil, ToRxFromString("+10"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1584(t *testing.T) {

	// -- say Rexx.toRexx("+1.1E-999999999").OpPow(RexxSet(17, byte 0), Rexx.toRexx("+10")) -- netrexx.lang.ExponentOverflowException: -1999999998

	_, err := ToRxFromString("+1.1E-999999999").OpPow(RxSetWithDigitandForm(17, int8(0)), ToRxFromString("+10"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1585(t *testing.T) {

	// -- say Rexx(int 50).OpPow(null, Rexx("0"))

	value, err := RxFromInt32(50).OpPow(nil, RxFromString("0"))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1586(t *testing.T) {

	// -- say Rexx(char '0').OpPow(null, Rexx(byte -128)) -- netrexx.lang.DivideException: Divide by 0

	_, err := RxFromRune('0').OpPow(nil, RxFromInt8(int8(-128)))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1587(t *testing.T) {

	// -- say Rexx(int 50).OpPow(null, Rexx("2"))

	value, err := RxFromInt32(50).OpPow(nil, RxFromString("2"))
	want := "2500"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1588(t *testing.T) {

	// -- say Rexx.toRexx("7198.0493E+436250299").OpPow(null, Rexx.toRexx("71E+2")) -- netrexx.lang.ExponentOverflowException: 1308750908

	_, err := ToRxFromString("7198.0493E+436250299").OpPow(nil, ToRxFromString("71E+2"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1589(t *testing.T) {

	// -- say Rexx(int 70).OpPow(null, Rexx(int 17))

	value, err := RxFromInt32(70).OpPow(nil, RxFromInt32(17))
	want := "2.32630514E+31"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1590(t *testing.T) {

	// -- say Rexx("50.000000").OpPow(null, Rexx("17.000"))

	value, err := RxFromString("50.000000").OpPow(nil, RxFromString("17.000"))
	want := "7.62939453E+28"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1591(t *testing.T) {

	// -- say Rexx.toRexx("+1.1E-914679999").OpPow(null, Rexx.toRexx("+10")) -- netrexx.lang.ExponentOverflowException: -1829359998

	_, err := ToRxFromString("+1.1E-914679999").OpPow(nil, ToRxFromString("+10"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1592(t *testing.T) {

	// -- say Rexx(int -7).OpPow(null, Rexx("3"))

	value, err := RxFromInt32(-7).OpPow(nil, RxFromString("3"))
	want := "-343"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1593(t *testing.T) {

	// -- say Rexx(char '0').OpPow(null, Rexx(byte -68)) -- netrexx.lang.DivideException: Divide by 0

	_, err := RxFromRune('0').OpPow(nil, RxFromInt8(int8(-68)))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1594(t *testing.T) {

	// -- say Rexx(char '9').OpPow(null, Rexx(byte -6))

	value, err := RxFromRune('9').OpPow(nil, RxFromInt8(int8(-6)))
	want := "0.00000188167642"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1595(t *testing.T) {

	// -- say Rexx("-123").OpPow(null, Rexx("2"))

	value, err := RxFromString("-123").OpPow(nil, RxFromString("2"))
	want := "15129"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1596(t *testing.T) {

	// --

	_, err := RxFromString("-9.999E+999999999").OpPow(RxSetWithDigit(3), RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

// func (rcvr *Rexx) OpRem(set *RexxSet, rhs *Rexx) (*Rexx, error)

func Test_1597(t *testing.T) {

	// -- say Rexx("50").OpRem(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("50").OpRem(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1598(t *testing.T) {

	// -- say Rexx("-1.1E-999999999").OpRem(null, Rexx("1E-999999999")) -- netrexx.lang.ExponentOverflowException: -1000000000

	_, err := RxFromString("-1.1E-999999999").OpRem(nil, RxFromString("1E-999999999"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1599(t *testing.T) {

	// -- say Rexx("0").OpRem(null, Rexx("50"))

	value, err := RxFromString("0").OpRem(nil, RxFromString("50"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1600(t *testing.T) {

	// -- say Rexx("1").OpRem(null, Rexx("3"))

	value, err := RxFromString("1").OpRem(nil, RxFromString("3"))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1601(t *testing.T) {

	// -- say Rexx("1E+999999999").OpRem(null, Rexx("9E+999999999"))

	value, err := RxFromString("1E+999999999").OpRem(nil, RxFromString("9E+999999999"))
	want := "1E+999999999"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1602(t *testing.T) {

	// -- say Rexx("3.93591888E-236595626").OpRem(RexxSet(17, 1), Rexx("7242375.00"))

	value, err := RxFromString("3.93591888E-236595626").OpRem(RxSetWithDigitandForm(17, int8(1)), RxFromString("7242375.00"))
	want := "3.93591888E-236595626"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1603(t *testing.T) {

	// -- say Rexx("50").OpRem(null, Rexx("2"))

	value, err := RxFromString("50").OpRem(nil, RxFromString("2"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1604(t *testing.T) {

	// -- say Rexx("50").OpRem(null, Rexx("50"))

	value, err := RxFromString("50").OpRem(nil, RxFromString("50"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1605(t *testing.T) {

	// -- say Rexx("1E+999999999").OpRem(RexxSet(70, 1), Rexx("9E+999999999")) -- test uses variables not accessible in NetRexx

	MaxExp = 9999

	_, err := RxFromString("1E+999999999").OpRem(RxSetWithDigitandForm(70, int8(1)), RxFromString("9E+999999999"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999 //reset

}

func Test_1606(t *testing.T) {

	// -- say Rexx("3.93591888E-236595626").OpRem(RexxSet(17, 1), Rexx("7242375.00")) -- test uses variables not accessible in NetRexx

	MinExp = -9999

	_, err := RxFromString("3.93591888E-236595626").OpRem(RxSetWithDigitandForm(17, int8(1)), RxFromString("7242375.00"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MinExp = -999999999 // reset

}

// func (rcvr *Rexx) OpSub(set *RexxSet, rhs *Rexx) (*Rexx, error)

func Test_1607(t *testing.T) {

	// -- say Rexx("50").OpSub(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("50").OpSub(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1608(t *testing.T) {

	// -- say Rexx("abc").OpSub(null, Rexx()) -- java.lang.NullPointerException

	_, err := RxFromString("abc").OpSub(nil, rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1609(t *testing.T) {

	// -- say Rexx().OpSub(null, Rexx("abc")) -- java.lang.NullPointerException

	_, err := rxfromempty().OpSub(nil, RxFromString("abc"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1610(t *testing.T) {

	// -- say Rexx("-50").OpSub(null, Rexx("2"))

	value, err := RxFromString("-50").OpSub(nil, RxFromString("2"))
	want := "-52"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1611(t *testing.T) {

	// -- say Rexx("0").OpSub(null, Rexx("50"))

	value, err := RxFromString("0").OpSub(nil, RxFromString("50"))
	want := "-50"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1612(t *testing.T) {

	// -- say Rexx("10.23456785").OpSub(RexxSet(9), Rexx("10.23456789"))

	value, err := RxFromString("10.23456785").OpSub(RxSetWithDigit(9), RxFromString("10.23456789"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1613(t *testing.T) {

	// -- say Rexx("50").OpSub(null, Rexx("17"))

	value, err := RxFromString("50").OpSub(nil, RxFromString("17"))
	want := "33"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) OpXor(set *RexxSet, rhs *Rexx) (bool, error)

func Test_1614(t *testing.T) {

	// -- say Rexx("1").OpXor(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("1").OpXor(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1615(t *testing.T) {

	// -- say Rexx("2").OpXor(null, Rexx()) -- java.lang.NullPointerException

	_, err := RxFromString("2").OpXor(nil, rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1616(t *testing.T) {

	// -- say Rexx("0").OpXor(null, Rexx("2")) -- netrexx.lang.NotLogicException: Boolean must be 0 or 1.  Found: 2

	_, err := RxFromString("0").OpXor(nil, RxFromString("2"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1617(t *testing.T) {

	// -- say Rexx("0").OpXor(null, Rexx("0"))

	value, err := RxFromString("0").OpXor(nil, RxFromString("0"))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1618(t *testing.T) {

	// -- say Rexx("0").OpXor(null, Rexx("1"))

	value, err := RxFromString("0").OpXor(nil, RxFromString("1"))
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1619(t *testing.T) {

	// -- say Rexx("1").OpXor(null, Rexx("0"))

	value, err := RxFromString("1").OpXor(nil, RxFromString("0"))
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1620(t *testing.T) {

	// -- say Rexx("1").OpXor(null, Rexx("1"))

	value, err := RxFromString("1").OpXor(nil, RxFromString("1"))
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Overlay(_new *Rexx, n *Rexx, length *Rexx, pad *Rexx) (*Rexx, error)

func Test_1621(t *testing.T) {

	// -- say Rexx("abc").overlay(null, null, null, null) -- java.lang.NullPointerException

	_, err := RxFromString("abc").Overlay(nil, nil, nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1622(t *testing.T) {

	// -- say Rexx("abc").overlay(Rexx("123"), null, Rexx("6"), Rexx(char '+')) -- java.lang.NullPointerException

	_, err := RxFromString("abc").Overlay(RxFromString("123"), nil, RxFromString("6"), RxFromRune('+'))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1623(t *testing.T) {

	// -- say Rexx("abcdef").overlay(Rexx(char '.'), Rexx(int 999999999+1), Rexx("2"), Rexx(" "))

	_, err := RxFromString("abcdef").Overlay(RxFromRune('.'), RxFromInt32(MaxArg+1), RxFromString("2"), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1624(t *testing.T) {

	// -- say Rexx("abc").overlay(Rexx("123"), Rexx("5"), null, Rexx(char '+')) -- java.lang.NullPointerException

	_, err := RxFromString("abc").Overlay(RxFromString("123"), RxFromString("5"), nil, RxFromRune('+'))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1625(t *testing.T) {

	// -- say Rexx("abcdef").overlay(Rexx(char '.'), Rexx("3"), Rexx(int 999999999+1), Rexx(" "))

	_, err := RxFromString("abcdef").Overlay(RxFromRune('.'), RxFromString("3"), RxFromInt32(MaxArg+1), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1626(t *testing.T) {

	// -- say Rexx("abc").overlay(Rexx("123"), Rexx("5"), Rexx("6"), null) -- java.lang.NullPointerException

	_, err := RxFromString("abc").Overlay(RxFromString("123"), RxFromString("5"), RxFromString("6"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1627(t *testing.T) {

	// -- say Rexx("abcdef").overlay(Rexx(char '.'), Rexx("3"), Rexx("2"), Rexx("AX")) -- netrexx.lang.NotCharacterException: AX

	_, err := RxFromString("abcdef").Overlay(RxFromRune('.'), RxFromString("3"), RxFromString("2"), RxFromString("AX"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1628(t *testing.T) {

	// -- say Rexx("").overlay(Rexx(char '.'), Rexx("3"), Rexx("2"), Rexx(" "))

	value, err := rxfromempty().Overlay(RxFromRune('.'), RxFromString("3"), RxFromString("2"), RxFromString(" "))
	want := "  . "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1629(t *testing.T) {

	// -- say Rexx("abc").overlay(null, Rexx("5"), Rexx("6"), Rexx(char '+')) -- java.lang.NullPointerException

	_, err := RxFromString("abc").Overlay(nil, RxFromString("5"), RxFromString("6"), RxFromRune('+'))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1630(t *testing.T) {

	// -- say Rexx("abcdef").overlay(Rexx(""), Rexx("3"), Rexx("2"), Rexx(" "))

	value, err := RxFromString("abcdef").Overlay(rxfromempty(), RxFromString("3"), RxFromString("2"), RxFromString(" "))
	want := "ab  ef"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1631(t *testing.T) {

	// -- say Rexx("abc").overlay(Rexx("123"), Rexx("5"), Rexx("6"), Rexx(char '+'))

	value, err := RxFromString("abc").Overlay(RxFromString("123"), RxFromString("5"), RxFromString("6"), RxFromRune('+'))
	want := "abc+123+++"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1632(t *testing.T) {

	// -- say Rexx("abcd").overlay(Rexx("qq"), Rexx("4"), Rexx("qq").length(), Rexx(" "))

	value, err := RxFromString("abcd").Overlay(RxFromString("qq"), RxFromString("4"), RxFromString("qq").Length(), RxFromString(" "))
	want := "abcqq"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1633(t *testing.T) {

	// -- say Rexx("abcd").overlay(Rexx("qq"), Rexx("1"), Rexx("qq").length(), Rexx(" "))

	value, err := RxFromString("abcd").Overlay(RxFromString("qq"), RxFromString("1"), RxFromString("qq").Length(), RxFromString(" "))
	want := "qqcd"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1634(t *testing.T) {

	// -- say Rexx("abcdef").overlay(Rexx(char ' '), Rexx("3"), Rexx(char ' ').length(), Rexx(" "))

	value, err := RxFromString("abcdef").Overlay(RxFromRune(' '), RxFromString("3"), RxFromRune(' ').Length(), RxFromString(" "))
	want := "ab def"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1635(t *testing.T) {

	// -- say Rexx("abcdef").overlay(Rexx(char '.'), Rexx("3"), Rexx("2"), Rexx(" "))

	value, err := RxFromString("abcdef").Overlay(RxFromRune('.'), RxFromString("3"), RxFromString("2"), RxFromString(" "))
	want := "ab. ef"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Pos(needle *Rexx, start *Rexx) (*Rexx, error)

func Test_1636(t *testing.T) {

	// -- say Rexx("50").pos(null, Rexx("1")) -- java.lang.NullPointerException

	_, err := RxFromString("50").Pos(nil, RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1637(t *testing.T) {

	// -- say Rexx("50").pos(Rexx("1"), null) -- java.lang.NullPointerException

	_, err := RxFromString("50").Pos(RxFromString("1"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1638(t *testing.T) {

	// -- say Rexx("50").pos(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("50").Pos(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1639(t *testing.T) {

	// -- say Rexx("a").pos(Rexx("1"), Rexx())  -- java.lang.NullPointerException

	_, err := RxFromString("a").Pos(RxFromString("1"), rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1640(t *testing.T) {

	// -- say Rexx().pos(Rexx(), Rexx("1")) -- java.lang.NullPointerException
	// -- say Rexx().pos(Rexx(""), Rexx("1"))

	value, err := rxfromempty().Pos(rxfromempty(), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1641(t *testing.T) {

	// -- say Rexx("Saturday").pos(Rexx(""), Rexx("1"))

	value, err := RxFromString("Saturday").Pos(RxFromString(""), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1642(t *testing.T) {

	// -- say Rexx("a").pos(Rexx(""), Rexx("1"))

	value, err := RxFromString("a").Pos(rxfromempty(), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1643(t *testing.T) {

	// -- say Rexx("").pos(Rexx("1"), Rexx("1"))

	value, err := rxfromempty().Pos(RxFromString("1"), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1644(t *testing.T) {

	// -- say Rexx("Saturday").pos(Rexx("day"), Rexx("1"))

	value, err := RxFromString("Saturday").Pos(RxFromString("day"), RxFromString("1"))
	want := "6"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1645(t *testing.T) {

	// -- say Rexx("abc def ghi").pos(Rexx(char ' '), Rexx("1"))

	value, err := RxFromString("abc def ghi").Pos(RxFromRune(' '), RxFromString("1"))
	want := "4"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1646(t *testing.T) {

	// -- say Rexx("abc def ghi").pos(Rexx(char ' '), Rexx("5"))

	value, err := RxFromString("abc def ghi").Pos(RxFromRune(' '), RxFromString("5"))
	want := "8"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1647(t *testing.T) {

	// -- say Rexx("50").pos(Rexx("2"), Rexx("1"))

	value, err := RxFromString("50").Pos(RxFromString("2"), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1648(t *testing.T) {

	// -- say Rexx("abc def ghi").pos(Rexx(char 'x'), Rexx("1"))

	value, err := RxFromString("abc def ghi").Pos(RxFromRune('x'), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1649(t *testing.T) {

	// -- say Rexx("1").pos(Rexx("1"), Rexx("2"))

	value, err := RxFromString("1").Pos(RxFromString("1"), RxFromString("2"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Reverse() (*Rexx, error)

func Test_1650(t *testing.T) {

	// -- say Rexx("").reverse()

	value, err := rxfromempty().Reverse()
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1651(t *testing.T) {

	// -- say Rexx("").reverse()

	value, err := RxFromString("").Reverse()
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1652(t *testing.T) {

	// -- say Rexx("ABc.").reverse()

	value, err := RxFromString("ABc.").Reverse()
	want := ".cBA"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1653(t *testing.T) {

	// -- say Rexx("Tranquility").reverse()

	value, err := RxFromString("Tranquility").Reverse()
	want := "ytiliuqnarT"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1654(t *testing.T) {

	// -- say Rexx("XYZ ").reverse()

	value, err := RxFromString("XYZ ").Reverse()
	want := " ZYX"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1655(t *testing.T) {

	// -- say Rexx("0").reverse()

	value, err := RxFromString("0").Reverse()
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Right(length *Rexx, pad *Rexx) (*Rexx, error)

func Test_1656(t *testing.T) {

	// -- say Rexx("-123").right(null, Rexx("1")) -- java.lang.NullPointerException

	_, err := RxFromString("-123").Right(nil, RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1657(t *testing.T) {

	// -- say Rexx("-123").right(Rexx("2"), null)
	// This port does not allow pad to be nil.

	_, err := RxFromString("-123").Right(RxFromString("2"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1658(t *testing.T) {

	// -- say Rexx("-123").right(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("-123").Right(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1659(t *testing.T) {

	// -- say Rexx("abc def").right(Rexx(int 999999999+1), Rexx(" "))

	_, err := RxFromString("abc def").Right(RxFromInt32(MaxArg+1), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1660(t *testing.T) {

	// -- say Rexx("abc def").right(Rexx("-6.7"), Rexx(" ")) -- java.lang.NumberFormatException: Non-zero decimal part in -6.7

	_, err := RxFromString("abc def").Right(RxFromString("-6.7"), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1661(t *testing.T) {

	// -- say Rexx("").right(Rexx("5"), Rexx(" "))

	value, err := rxfromempty().Right(RxFromString("5"), RxFromString(" "))
	want := "     "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1662(t *testing.T) {

	// -- say Rexx("-123").right(Rexx("2"), Rexx("1"))

	value, err := RxFromString("-123").Right(RxFromString("2"), RxFromString("1"))
	want := "23"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1663(t *testing.T) {

	// -- say Rexx("abc def").right(Rexx("0"), Rexx(" "))

	value, err := RxFromString("abc def").Right(RxFromString("0"), RxFromString(" "))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1664(t *testing.T) {

	// -- say Rexx("abc def").right(Rexx("5"), Rexx(" "))

	value, err := RxFromString("abc def").Right(RxFromString("5"), RxFromString(" "))
	want := "c def"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1665(t *testing.T) {

	// -- say Rexx("abc def").right(Rexx("8"), Rexx("jjj")) -- netrexx.lang.NotCharacterException: jjj

	_, err := RxFromString("abc def").Right(RxFromString("8"), RxFromString("jjj"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1666(t *testing.T) {

	// -- say Rexx("12").right(Rexx("5"), Rexx(char '0'))

	value, err := RxFromString("12").Right(RxFromString("5"), RxFromRune('0'))
	want := "00012"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1667(t *testing.T) {

	// -- say Rexx("abc  d").right(Rexx("8"), Rexx(" "))

	value, err := RxFromString("abc  d").Right(RxFromString("8"), RxFromString(" "))
	want := "  abc  d"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1668(t *testing.T) {

	// -- say Rexx("7635111111111111111666666666666666666666666666666666666666666222222222222219875220.444444444444444").right(Rexx("76351111111111111116666666666666666666666666666666666666666662222222222222198752215"), Rexx("7635111111111111111666666666666666666666666666666666666666666222222222222219875220.555555555555553")) -- java.lang.NumberFormatException: Conversion overflow

	_, err := RxFromString("7635111111111111111666666666666666666666666666666666666666666222222222222219875220.444444444444444").Right(RxFromString("76351111111111111116666666666666666666666666666666666666666662222222222222198752215"), RxFromString("7635111111111111111666666666666666666666666666666666666666666222222222222219875220.555555555555553"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1669(t *testing.T) {

	// --

	MaxArg = 0
	MinArg = 0
	MaxExp = 0
	MinExp = 0

	new_set := RxSet()
	new_set.SetForm(RxFromString("engineering"))
	new_set.SetDigits(RxFromString("3"))

	RxFromString("-56267E-5").Right(RxFromInt32(0), RxFromRune('o'))

	MaxArg = 999999999
	MinArg = -999999999
	MaxExp = 999999999
	MinExp = -999999999

}

// func (rcvr *Rexx) Sequence(_final *Rexx) (*Rexx, error)

func Test_1670(t *testing.T) {

	// -- say Rexx(char 'a').sequence(null) -- java.lang.NullPointerException

	_, err := RxFromRune('a').Sequence(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1671(t *testing.T) {

	// -- say Rexx().sequence(Rexx(char 'f')) -- java.lang.NullPointerException

	_, err := rxfromempty().Sequence(RxFromRune('f'))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1672(t *testing.T) {

	// -- say Rexx("").sequence(Rexx(char 'f')) -- netrexx.lang.NotCharacterException:

	_, err := RxFromString("").Sequence(RxFromRune('f'))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1673(t *testing.T) {

	// -- say Rexx(char 'a').sequence(Rexx()) -- java.lang.NullPointerException

	_, err := RxFromRune('a').Sequence(rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1674(t *testing.T) {

	// -- say Rexx(char 'q').sequence(Rexx(char 'k')) -- netrexx.lang.BadArgumentException: final<start

	_, err := RxFromRune('q').Sequence(RxFromRune('k'))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1675(t *testing.T) {

	// -- say Rexx(char 'a').sequence(Rexx(char 'f'))

	value, err := RxFromRune('a').Sequence(RxFromRune('f'))
	want := "abcdef"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1676(t *testing.T) {

	// -- say Rexx(char(40)).sequence(Rexx(char '\x29'))

	value, err := RxFromRune(rune(40)).Sequence(RxFromRune('\x29'))
	want := "()"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1677(t *testing.T) {

	// -- say Rexx(char(0xfffe)).sequence(Rexx(char(0xffff)))

	value, err := RxFromRune(rune(0xfffe)).Sequence(RxFromRune(0xffff))
	want := "￾￿"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Sign() (*Rexx, error)

func Test_1678(t *testing.T) {

	// -- say Rexx(char 'q').sign() -- java.lang.NumberFormatException: q

	_, err := RxFromRune('q').Sign()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1679(t *testing.T) {

	// -- say Rexx(" -0.307").sign()

	value, err := RxFromString(" -0.307").Sign()
	want := "-1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1680(t *testing.T) {

	// -- say Rexx("0.0").sign()

	value, err := RxFromString("0.0").Sign()
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1681(t *testing.T) {

	// -- say Rexx("12.3").sign()

	value, err := RxFromString("12.3").Sign()
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Significance() int32

func Test_1682(t *testing.T) {

	// -- say Rexx(char 'a').significance()

	got := RxFromRune('a').Significance()
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1683(t *testing.T) {

	// -- say Rexx("12.3").significance()

	got := RxFromString("12.3").Significance()
	want := int32(3)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) Space(n *Rexx, pad *Rexx) (*Rexx, error)

func Test_1684(t *testing.T) {

	// -- say Rexx("abc  def  ").space(null, Rexx("+")) -- java.lang.NullPointerException

	_, err := RxFromString("abc  def  ").Space(nil, RxFromString("+"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1685(t *testing.T) {

	// -- say Rexx("abc  def  ").space(Rexx("2"), null) -- java.lang.NullPointerException

	_, err := RxFromString("abc  def  ").Space(RxFromString("2"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1686(t *testing.T) {

	// -- say Rexx("abc  def  ").space(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("abc  def  ").Space(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1687(t *testing.T) {

	// -- say Rexx(char 'a').space(Rexx(int 999999999+1), Rexx())

	_, err := RxFromRune('a').Space(RxFromInt32(MaxArg+1), rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1688(t *testing.T) {

	// -- say Rexx(char 'a').space(Rexx(), Rexx()) -- java.lang.NullPointerException

	_, err := RxFromRune('a').Space(rxfromempty(), rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1689(t *testing.T) {

	// -- say Rexx("").space(Rexx("1"), Rexx(" "))

	value, err := rxfromempty().Space(RxFromString("1"), RxFromString(" "))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1690(t *testing.T) {

	// -- say Rexx("  abc def ").space(Rexx("3"), Rexx(" "))

	value, err := RxFromString("  abc def ").Space(RxFromString("3"), RxFromString(" "))
	want := "abc   def"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1691(t *testing.T) {

	// -- say Rexx("abc  def  ").space(Rexx("0"), Rexx(" "))

	value, err := RxFromString("abc  def  ").Space(RxFromString("0"), RxFromString(" "))
	want := "abcdef"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1692(t *testing.T) {

	// -- say Rexx("abc  def  ").space(Rexx("1"), Rexx(" "))

	value, err := RxFromString("abc  def  ").Space(RxFromString("1"), RxFromString(" "))
	want := "abc def"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1693(t *testing.T) {

	// -- say Rexx("abc  def  ").space(Rexx("1"), Rexx(" "))

	value, err := RxFromString("abc  def  ").Space(RxFromString("1"), RxFromString(" "))
	want := "abc def"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1694(t *testing.T) {

	// -- say Rexx("abc  def  ").space(Rexx("2"), Rexx(char '+'))

	value, err := RxFromString("abc  def  ").Space(RxFromString("2"), RxFromRune('+'))
	want := "abc++def"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1695(t *testing.T) {

	// -- say Rexx("abc  def  ").space(Rexx("2"), Rexx("+"))

	value, err := RxFromString("abc  def  ").Space(RxFromString("2"), RxFromString("+"))
	want := "abc++def"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Strip(opt *Rexx, pad *Rexx) (*Rexx, error)

func Test_1696(t *testing.T) {

	// -- say Rexx("-123").strip(null, Rexx("1")) -- java.lang.NullPointerException

	_, err := RxFromString("-123").Strip(nil, RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1697(t *testing.T) {

	// -- say Rexx("-123").strip(Rexx("2"), null) -- netrexx.lang.BadArgumentException: Bad Option character 2 [2]

	_, err := RxFromString("-123").Strip(RxFromString("2"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1698(t *testing.T) {

	// -- say Rexx("-123").strip(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("-123").Strip(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1699(t *testing.T) {

	// -- say Rexx().strip(Rexx("1"), Rexx(" ")) -- java.lang.NullPointerException

	_, err := rxfromempty().Strip(RxFromString("1"), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1700(t *testing.T) {

	// -- say Rexx("-123").strip(Rexx("2"), Rexx("1")) -- netrexx.lang.BadArgumentException: Bad Option character 2 [2]

	_, err := RxFromString("-123").Strip(RxFromString("2"), RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1701(t *testing.T) {

	// -- say Rexx(char 'a').strip(Rexx("B"), Rexx("AXE")) -- netrexx.lang.NotCharacterException: AXE

	_, err := RxFromRune('a').Strip(RxFromString("B"), RxFromString("AXE"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1702(t *testing.T) {

	// -- say Rexx("").strip(Rexx("B"), Rexx(" "))

	value, err := rxfromempty().Strip(RxFromString("B"), RxFromString(" "))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1703(t *testing.T) {

	// -- say Rexx("").strip(Rexx("B"), Rexx(" "))

	value, err := RxFromString("").Strip(RxFromString("B"), RxFromString(" "))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1704(t *testing.T) {

	// -- say Rexx("0000").strip(Rexx(char 'T'), Rexx("0"))

	value, err := RxFromString("0000").Strip(RxFromRune('T'), RxFromString("0"))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1705(t *testing.T) {

	// -- say Rexx("  Ab c  ").strip(Rexx(char 'L'), Rexx(" "))

	value, err := RxFromString("  Ab c  ").Strip(RxFromRune('L'), RxFromString(" "))
	want := "Ab c  "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1706(t *testing.T) {

	// -- say Rexx("  aB c  ").strip(Rexx(char 'L'), Rexx(" "))

	value, err := RxFromString("  aB c  ").Strip(RxFromRune('L'), RxFromString(" "))
	want := "aB c  "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1707(t *testing.T) {

	// -- say Rexx("0000").strip(Rexx("l"), Rexx("0"))

	value, err := RxFromString("0000").Strip(RxFromString("l"), RxFromString("0"))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1708(t *testing.T) {

	// -- say Rexx("  ab C  ").strip(Rexx(char 'L'), Rexx(" "))

	value, err := RxFromString("  ab C  ").Strip(RxFromRune('L'), RxFromString(" "))
	want := "ab C  "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1709(t *testing.T) {

	// -- say Rexx("  ab c  ").strip(Rexx(char 'L'), Rexx(" "))

	value, err := RxFromString("  ab c  ").Strip(RxFromRune('L'), RxFromString(" "))
	want := "ab c  "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1710(t *testing.T) {

	// -- say Rexx("  ab c  ").strip(Rexx(char 't'), Rexx(" "))

	value, err := RxFromString("  ab c  ").Strip(RxFromRune('t'), RxFromString(" "))
	want := "  ab c"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1711(t *testing.T) {

	// -- say Rexx("0000").strip(Rexx(char 'T'), Rexx("0"))

	value, err := RxFromString("0000").Strip(RxFromRune('T'), RxFromString("0"))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1712(t *testing.T) {

	// -- say Rexx("  ab c  ").strip(Rexx("B"), Rexx(" "))

	value, err := RxFromString("  ab c  ").Strip(RxFromString("B"), RxFromString(" "))
	want := "ab c"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1713(t *testing.T) {

	// -- say Rexx("0012.700").strip(Rexx(char 'b'), Rexx("0"))

	value, err := RxFromString("0012.700").Strip(RxFromRune('b'), RxFromString("0"))
	want := "12.7"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1714(t *testing.T) {

	// -- say Rexx("1001").strip(Rexx("l"), Rexx("0"))

	value, err := RxFromString("1001").Strip(RxFromString("l"), RxFromString("0"))
	want := "1001"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1715(t *testing.T) {

	// -- say Rexx("12.70000").strip(Rexx(char 't'), Rexx("0"))

	value, err := RxFromString("12.70000").Strip(RxFromRune('t'), RxFromString("0"))
	want := "12.7"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1716(t *testing.T) {

	// -- say Rexx("0012.700").strip(Rexx(char 'b'), Rexx("0"))

	value, err := RxFromString("0012.700").Strip(RxFromRune('b'), RxFromString("0"))
	want := "12.7"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) SubStr(n *Rexx, length *Rexx, pad *Rexx) (*Rexx, error)

func Test_1717(t *testing.T) {

	// -- say Rexx("-123").substr(null, Rexx("1"), Rexx(" ")) -- java.lang.NullPointerException

	_, err := RxFromString("-123").SubStr(nil, RxFromString("1"), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1718(t *testing.T) {

	// -- say Rexx("-123").substr(Rexx("2"), null, Rexx(" ")) -- java.lang.NullPointerException

	_, err := RxFromString("-123").SubStr(RxFromString("2"), nil, RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1719(t *testing.T) {

	// -- say Rexx("-123").substr(Rexx("2"), Rexx("1"), null) -- java.lang.NullPointerException

	_, err := RxFromString("-123").SubStr(RxFromString("2"), RxFromString("1"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1720(t *testing.T) {

	// -- say Rexx("-123").substr(null, null, null) -- java.lang.NullPointerException

	_, err := RxFromString("-123").SubStr(nil, nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1721(t *testing.T) {

	// -- say Rexx("abc").substr(Rexx(int 999999999+1), Rexx("4"), Rexx(" "))

	_, err := RxFromString("abc").SubStr(RxFromInt32(MaxArg+1), RxFromString("4"), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1722(t *testing.T) {

	// -- say Rexx("abc").substr(Rexx("2"), Rexx(int 999999999+1), Rexx(" "))

	_, err := RxFromString("abc").SubStr(RxFromString("2"), RxFromInt32(MaxArg+1), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1723(t *testing.T) {

	// -- say Rexx("abc").substr(Rexx("2"), Rexx("4"), Rexx("AXE")) -- netrexx.lang.NotCharacterException: AXE

	_, err := RxFromString("abc").SubStr(RxFromString("2"), RxFromString("4"), RxFromString("AXE"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1724(t *testing.T) {

	// -- say Rexx("").substr(Rexx("1"), Rexx("4"), Rexx(" "))

	value, err := rxfromempty().SubStr(RxFromString("1"), RxFromString("4"), RxFromString(" "))
	want := "    "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1725(t *testing.T) {

	// -- say Rexx("-123").substr(Rexx("2"), Rexx("1"), Rexx(" "))

	value, err := RxFromString("-123").SubStr(RxFromString("2"), RxFromString("1"), RxFromString(" "))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1726(t *testing.T) {

	// -- say Rexx("abc").substr(Rexx(int 2), Rexx("2"), Rexx(" "))

	value, err := RxFromString("abc").SubStr(RxFromInt32(2), RxFromString("2"), RxFromString(" "))
	want := "bc"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1727(t *testing.T) {

	// -- say Rexx("abc").substr(Rexx("2"), Rexx("4"), Rexx(" "))

	value, err := RxFromString("abc").SubStr(RxFromString("2"), RxFromString("4"), RxFromString(" "))
	want := "bc  "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1728(t *testing.T) {

	// -- say Rexx("abc").substr(Rexx("2"), Rexx("6"), Rexx(char '.'))

	value, err := RxFromString("abc").SubStr(RxFromString("2"), RxFromString("6"), RxFromRune('.'))
	want := "bc...."

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1729(t *testing.T) {

	// -- say Rexx("abc").substr(Rexx("5"), Rexx("6"), Rexx(char '.'))

	value, err := RxFromString("abc").SubStr(RxFromString("5"), RxFromString("6"), RxFromRune('.'))
	want := "......"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1730(t *testing.T) {

	// -- say Rexx("abc").substr(Rexx("5"), Rexx("4"), Rexx(" "))

	value, err := RxFromString("abc").SubStr(RxFromString("5"), RxFromString("4"), RxFromString(" "))
	want := "    "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1731(t *testing.T) {

	// -- say Rexx("541").substr(Rexx("2"), Rexx("4"), Rexx(" "))

	value, err := RxFromString("541").SubStr(RxFromString("2"), RxFromString("4"), RxFromString(" "))
	want := "41  "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) SubWord(n *Rexx, length *Rexx) (*Rexx, error)

func Test_1732(t *testing.T) {

	// -- say Rexx("-123").subword(Rexx("2"), Rexx("0"))

	value, err := RxFromString("-123").SubWord(RxFromString("2"), RxFromString("0"))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1733(t *testing.T) {

	// -- say Rexx("-123").subword(Rexx("2"), null) -- java.lang.NullPointerException

	_, err := RxFromString("-123").SubWord(RxFromString("2"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1734(t *testing.T) {

	// -- say Rexx("-123").subword(null, Rexx("1")) -- java.lang.NullPointerException

	_, err := RxFromString("-123").SubWord(nil, RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1735(t *testing.T) {

	// -- say Rexx("-123").subword(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("-123").SubWord(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1736(t *testing.T) {

	// -- say Rexx("abc").subword(Rexx(int 999999999+1), Rexx(" "))

	_, err := RxFromString("abc").SubWord(RxFromInt32(MaxArg+1), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1737(t *testing.T) {

	// -- say Rexx("abc").subword(Rexx("2"), Rexx("AXE")) -- java.lang.NumberFormatException: Not a number

	_, err := RxFromString("abc").SubWord(RxFromString("2"), RxFromString("AXE"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1738(t *testing.T) {

	// -- say Rexx("").subword(Rexx("2"), Rexx("4"))

	value, err := rxfromempty().SubWord(RxFromString("2"), RxFromString("4"))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1739(t *testing.T) {

	// -- say Rexx("-123").subword(Rexx("2"), Rexx("1"))

	value, err := RxFromString("-123").SubWord(RxFromString("2"), RxFromString("1"))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1740(t *testing.T) {

	// -- say Rexx("Now is the  time").subword(Rexx("2"), Rexx("2"))

	value, err := RxFromString("Now is the  time").SubWord(RxFromString("2"), RxFromString("2"))
	want := "is the"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1741(t *testing.T) {

	// -- say Rexx("Now is the  time").subword(Rexx("3"), Rexx("Now is the  time").length())

	value, err := RxFromString("Now is the  time").SubWord(RxFromString("3"), RxFromString("Now is the  time").Length())
	want := "the  time"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1742(t *testing.T) {

	// -- say Rexx("Now is the  time").subword(Rexx("5"), Rexx("Now is the  time").length())

	value, err := RxFromString("Now is the  time").SubWord(RxFromString("5"), RxFromString("Now is the  time").Length())
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1743(t *testing.T) {

	// -- say Rexx("Now is the  time").subword(Rexx("5"), Rexx("Now is the  time").length())

	value, err := RxFromString("Now is the  time").SubWord(RxFromString("5"), RxFromString("Now is the  time").Length())
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) ToBool() (bool, error)

func Test_1744(t *testing.T) {

	// -- say Rexx(boolean 0).toboolean()

	value, err := RxFromBool(false).ToBool()
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1745(t *testing.T) {

	// -- say Rexx("0").toboolean()

	value, err := RxFromString("0").ToBool()
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1746(t *testing.T) {

	// -- say Rexx().toboolean()

	value, err := rxfromempty().ToBool()
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1747(t *testing.T) {

	// -- say Rexx("-0").toboolean()

	value, err := RxFromString("-0").ToBool()
	want := false

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1748(t *testing.T) {

	// -- say Rexx(boolean 1).toboolean()

	value, err := RxFromBool(true).ToBool()
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1749(t *testing.T) {

	// -- say Rexx("1").toboolean()

	value, err := RxFromString("1").ToBool()
	want := true

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1750(t *testing.T) {

	// -- say Rexx("").toboolean() -- netrexx.lang.NotLogicException: Boolean must be 0 or 1.  Found:

	_, err := RxFromString("").ToBool()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1751(t *testing.T) {

	// -- say Rexx("a").toboolean() -- netrexx.lang.NotLogicException: Boolean must be 0 or 1.  Found: a

	_, err := RxFromString("a").ToBool()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

// func (rcvr *Rexx) ToFloat32() (float32, error)

func Test_1752(t *testing.T) {

	// -- say Rexx(char ' ').tofloat() -- java.lang.NumberFormatException:

	_, err := RxFromRune(' ').ToFloat32()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1753(t *testing.T) {

	// -- say Rexx(float 1.401298464324817070923729583289916131280e-45).tofloat()
	// Notes: Use strings for very large numbers.
	// Calling RxFromFloat32(float32(1.401298464324817070923729583289916131280e-45)).ToFloat32() causes a NumberFormatException : Overflow

	value, err := RxFromString("1.401298464324817070923729583289916131280e-45").ToFloat32()
	want := float32(1.4e-45)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1754(t *testing.T) {

	// -- say Rexx(float 1.401298464324817070923729583289916131280e-45).tofloat()

	_, err := RxFromFloat32(1.401298464324817070923729583289916131280e-45).ToFloat32()
	// Same as above Test but causes a NumberFormatException : Overflow when called this way.

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1755(t *testing.T) {

	// -- say Rexx("1.401298464324817e-45").tofloat()

	value, err := RxFromString("1.401298464324817e-45").ToFloat32()
	want := float32(1.4e-45)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) ToFloat64() (float64, error)

func Test_1756(t *testing.T) {

	// -- say Rexx(char ' ').todouble() -- java.lang.NumberFormatException:

	_, err := RxFromRune(' ').ToFloat64()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1757(t *testing.T) {

	// -- say Rexx("").todouble() -- java.lang.NumberFormatException:

	_, err := RxFromString("").ToFloat64()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1758(t *testing.T) {

	// -- say Rexx().todouble() -- java.lang.NullPointerException

	_, err := rxfromempty().ToFloat64()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1759(t *testing.T) {

	// -- say Rexx("-1.797693134862315708145274237317043567981e+309").todouble() -- java.lang.NumberFormatException: Overflow

	_, err := RxFromString("-1.797693134862315708145274237317043567981e+309").ToFloat64()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1760(t *testing.T) {

	// -- say Rexx("1.797693134862315708145274237317043567981e+309").todouble() -- java.lang.NumberFormatException: Overflow

	_, err := RxFromString("1.797693134862315708145274237317043567981e+309").ToFloat64()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1761(t *testing.T) {

	// -- say Rexx(char(0x11452)).todouble() -- java.lang.NumberFormatException: ᑒ

	_, err := RxFromRune(0x11452).ToFloat64()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1762(t *testing.T) {

	// -- say Rexx("1.0").todouble()

	value, err := RxFromString("1.0").ToFloat64()
	want := float64(1.0)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1763(t *testing.T) {

	// -- say Rexx("1.7976931348623157e+308").todouble()

	value, err := RxFromString("1.7976931348623157e+308").ToFloat64()
	want := float64(1.7976931348623157e308)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) ToInt16() (int16, error)

func Test_1764(t *testing.T) {

	// -- say Rexx("abc").toshort() -- java.lang.NumberFormatException: abc

	_, err := RxFromString("abc").ToInt16()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1765(t *testing.T) {

	// -- say Rexx(int 123456).toshort() -- java.lang.NumberFormatException: Conversion overflow

	_, err := RxFromInt32(123456).ToInt16()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1766(t *testing.T) {

	// -- say Rexx(int 1).toshort()

	value, err := RxFromInt32(1).ToInt16()
	want := int16(1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) ToInt32() (int32, error)

func Test_1767(t *testing.T) {

	// -- say Rexx(char[] 'a', boolean 1).toint()

	_, err := rx([]rune{'a'}, true).ToInt32()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1768(t *testing.T) {

	// -- say Rexx(char[] '0', boolean 0).toint()

	_, err := rx([]rune{'0'}, false).ToInt32()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1769(t *testing.T) {

	// -- say Rexx(char[] '0', boolean 1).toint()

	value, err := rx([]rune{'0'}, true).ToInt32()
	want := int32(0)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1770(t *testing.T) {

	// -- say Rexx("0-2").toint() -- java.lang.NumberFormatException: 0-2

	_, err := RxFromString("0-2").ToInt32()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1771(t *testing.T) {

	// -- say Rexx(char[] '0', boolean 1).toint()

	value, err := rx([]rune{'0'}, true).ToInt32()
	want := int32(0)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1772(t *testing.T) {

	// -- say Rexx().toint() -- java.lang.NullPointerException

	value, err := rxfromempty().ToInt32()
	want := int32(0)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1773(t *testing.T) {

	// -- say Rexx(char[] "123.4568888", boolean 1).toint()

	_, err := rx([]rune("123.4568888"), true).ToInt32()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1774(t *testing.T) {

	// -- say Rexx(char[] "-.1", boolean 1).toint()

	_, err := rx([]rune("-.1"), true).ToInt32()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1775(t *testing.T) {

	// -- say Rexx("-.01234568").toint() -- java.lang.NumberFormatException: Decimal part non-zero: -.01234568

	_, err := RxFromString("-.01234568").ToInt32()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1776(t *testing.T) {

	// -- say Rexx("-0.00000000009").toint() -- java.lang.NumberFormatException: Decimal part non-zero: -0.00000000009

	_, err := RxFromString("-0.00000000009").ToInt32()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1777(t *testing.T) {

	// -- say Rexx(char[] "7453.4", boolean 1).toint()

	_, err := rx([]rune("7453.4"), true).ToInt32()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1778(t *testing.T) {

	// -- say Rexx("-100.10101").toint() -- java.lang.NumberFormatException: Decimal part non-zero: -100.10101

	_, err := RxFromString("-100.10101").ToInt32()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1779(t *testing.T) {

	// -- say Rexx(char[] "999.999", boolean 1).toint()

	_, err := rx([]rune("999.999"), true).ToInt32()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1780(t *testing.T) {

	// -- say Rexx(char[] ".4E+121", boolean 1).toint()

	_, err := rx([]rune(".4E+121"), true).ToInt32()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1781(t *testing.T) {

	// -- say Rexx(float -1.54e+20).toint() -- java.lang.NumberFormatException: Conversion overflow

	_, err := RxFromFloat32(-1.54e+20).ToInt32()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1782(t *testing.T) {

	// -- say Rexx("000001001101").toint()

	value, err := RxFromString("000001001101").ToInt32()
	want := int32(1001101)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1783(t *testing.T) {

	// -- say Rexx("10010101").toint()

	value, err := RxFromString("10010101").ToInt32()
	want := int32(10010101)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1784(t *testing.T) {

	// -- say Rexx(int 2147483647).toint()

	value, err := RxFromInt32(2147483647).ToInt32()
	want := int32(2147483647)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1785(t *testing.T) {

	// -- say Rexx(char[] "3", boolean 1).toint()

	value, err := rx([]rune("3"), true).ToInt32()
	want := int32(3)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1786(t *testing.T) {

	// -- say Rexx("1").toint()

	value, err := RxFromString("1").ToInt32()
	want := int32(1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1787(t *testing.T) {

	// -- say Rexx(int -2147483648).toint() -- java.lang.NumberFormatException: Conversion overflow
	// This port runs command

	value, err := RxFromInt32(-2147483648).ToInt32()
	want := int32(-2147483648)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1788(t *testing.T) {

	// -- say Rexx("4294967294").toint() -- java.lang.NumberFormatException: Conversion overflow

	_, err := RxFromString("4294967294").ToInt32()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1789(t *testing.T) {

	// -- say Rexx(char[] "3", boolean 1).toint()

	value, err := rx([]rune("3"), true).ToInt32()
	want := int32(3)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1790(t *testing.T) {

	// -- say Rexx(char[] "-1000E-003", boolean 1).toint()

	value, err := rx([]rune("-1000E-003"), true).ToInt32()
	want := int32(-1000e-003)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) ToInt64() (int64, error)

func Test_1791(t *testing.T) {

	// -- say Rexx("abc").tolong() -- java.lang.NumberFormatException: abc

	_, err := RxFromString("abc").ToInt64()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1792(t *testing.T) {

	// -- say Rexx("0").tolong()

	value, err := RxFromString("0").ToInt64()
	want := int64(0)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1793(t *testing.T) {

	// -- say Rexx().tolong()

	value, err := rxfromempty().ToInt64()
	want := int64(0)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1794(t *testing.T) {

	// -- say Rexx("-.01234568").tolong() -- java.lang.NumberFormatException: Decimal part non-zero: -.01234568

	_, err := RxFromString("-.01234568").ToInt64()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1795(t *testing.T) {

	// -- say Rexx("-0.1234567890").tolong() -- java.lang.NumberFormatException: Decimal part non-zero: -0.1234567890

	_, err := RxFromString("-0.1234567890").ToInt64()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1796(t *testing.T) {

	// -- say Rexx(".1").tolong() -- java.lang.NumberFormatException: Decimal part non-zero: .1

	_, err := RxFromString(".1").ToInt64()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1797(t *testing.T) {

	// -- say Rexx("90.0E-1").tolong()

	value, err := RxFromString("90.0E-1").ToInt64()
	want := int64(9)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1798(t *testing.T) {

	// -- say Rexx.toRexx("+00678.5432").tolong() -- java.lang.NumberFormatException: Decimal part non-zero: +00678.5432

	_, err := ToRxFromString("+00678.5432").ToInt64()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1799(t *testing.T) {

	// -- say Rexx("56260E-1").tolong()

	value, err := RxFromString("56260E-1").ToInt64()
	want := int64(5626)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1800(t *testing.T) {

	// -- say Rexx("20.11E-1").tolong() -- java.lang.NumberFormatException: Decimal part non-zero: 20.11E-1

	_, err := RxFromString("20.11E-1").ToInt64()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1801(t *testing.T) {

	// -- say Rexx("56260E-1").tolong()

	value, err := RxFromString("56260E-1").ToInt64()
	want := int64(5626)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1802(t *testing.T) {

	// -- say Rexx(double 9.1e+40).tolong() -- java.lang.NumberFormatException: Conversion overflow

	_, err := RxFromFloat64(9.1e+40).ToInt64()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1803(t *testing.T) {

	// -- say Rexx("18446744073709551615").tolong() -- java.lang.NumberFormatException: Conversion overflow

	_, err := RxFromString("18446744073709551615").ToInt64()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1804(t *testing.T) {

	// -- say Rexx(char[] "۱۱۱۱.0E+6", boolean 1).tolong()

	value, err := rx([]rune("۱۱۱۱.0E+6"), true).ToInt64()
	want := int64(1111000000)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1805(t *testing.T) {

	// -- say Rexx(int 123456).tolong()

	value, err := RxFromInt32(123456).ToInt64()
	want := int64(123456)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1806(t *testing.T) {

	// -- say Rexx("1.").tolong()

	value, err := RxFromString("1.").ToInt64()
	want := int64(1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1807(t *testing.T) {

	// -- say Rexx("1").tolong()

	value, err := RxFromString("1").ToInt64()
	want := int64(1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1808(t *testing.T) {

	// -- say Rexx("998890068.").tolong()

	value, err := RxFromString("998890068.").ToInt64()
	want := int64(998890068)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1809(t *testing.T) {

	// -- say Rexx("-9223372036854775808").tolong()

	value, err := RxFromString("-9223372036854775808").ToInt64()
	want := int64(-9223372036854775808)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1810(t *testing.T) {

	// -- say Rexx("-9223372036854775809").tolong() -- java.lang.NumberFormatException: Conversion overflow

	_, err := RxFromString("-9223372036854775809").ToInt64()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1811(t *testing.T) {

	// -- say Rexx("-9223372036854775808").tolong()

	value, err := RxFromString("-9223372036854775808").ToInt64()
	want := int64(-9223372036854775808)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1812(t *testing.T) {

	// -- say Rexx("9223372036854775807").tolong()

	value, err := RxFromString("9223372036854775807").ToInt64()
	want := int64(9223372036854775807)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1813(t *testing.T) {

	// -- say Rexx(char[] "-1000E-003", boolean 1).tolong()

	value, err := rx([]rune("-1000E-003"), true).ToInt64()
	want := int64(-1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) ToInt8() (int8, error)

func Test_1814(t *testing.T) {

	// -- say Rexx("A").tobyte() -- java.lang.NumberFormatException: A

	_, err := RxFromString("A").ToInt8()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1815(t *testing.T) {

	// -- say Rexx("-129").tobyte() -- java.lang.NumberFormatException: Conversion overflow

	_, err := RxFromString("-129").ToInt8()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1816(t *testing.T) {

	// -- say Rexx("6").tobyte()

	value, err := RxFromString("6").ToInt8()
	want := int8(6)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) ToRune() (rune, error)

func Test_1817(t *testing.T) {

	// -- say Rexx().tochar() -- java.lang.NullPointerException

	_, err := rxfromempty().ToRune()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1818(t *testing.T) {

	// -- say Rexx("ABC").tochar()  -- netrexx.lang.NotCharacterException: ABC

	_, err := RxFromString("ABC").ToRune()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1819(t *testing.T) {

	// -- say Rexx("A").tochar()

	value, err := RxFromString("A").ToRune()
	want := 'A'

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) ToRunes() []rune

func Test_1820(t *testing.T) {

	// -- say Rexx("").tochararray()

	got := rxfromempty().ToRunes()

	want := ""

	if string(got) != want {
		t.Errorf("got %v, wanted %v", string(got), want)
	}

}

func Test_1821(t *testing.T) {

	// -- say Rexx("ABCDEFGH").tochararray()

	got := RxFromString("ABCDEFGH").ToRunes()
	want := "ABCDEFGH"

	if string(got) != want {
		t.Errorf("got %v, wanted %v", string(got), want)
	}

}

// func (rcvr *Rexx) ToString() string

func Test_1822(t *testing.T) {

	// -- say Rexx("").ToString()

	got := rxfromempty().ToString()
	want := ""

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1823(t *testing.T) {

	// -- say Rexx("-1234+5678.").ToString()

	got := RxFromString("-1234+5678.").ToString()
	want := "-1234+5678."

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1824(t *testing.T) {

	// -- say RxFromStrings([]string{}).ToString()

	got := RxFromStrings([]string{}).ToString()
	want := ""

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) Translate(tableo *Rexx, tablei *Rexx, pad *Rexx) (*Rexx, error)

func Test_1825(t *testing.T) {

	// -- say Rexx("-1234").translate(null, Rexx("1"), Rexx(" ")) -- java.lang.NullPointerException

	_, err := RxFromString("-1234").Translate(nil, RxFromString("1"), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1826(t *testing.T) {

	// -- say Rexx("-1234").translate(Rexx("2"), null, Rexx(" ")) -- java.lang.NullPointerException

	_, err := RxFromString("-1234").Translate(RxFromString("2"), nil, RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1827(t *testing.T) {

	// -- say Rexx("-1234").translate(Rexx("2"), Rexx("1"), null) -- java.lang.NullPointerException

	_, err := RxFromString("-1234").Translate(RxFromString("2"), RxFromString("1"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1828(t *testing.T) {

	// -- say Rexx("-1234").translate(null, null, null) -- java.lang.NullPointerException

	_, err := RxFromString("-1234").Translate(nil, nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1829(t *testing.T) {

	// -- say Rexx("abcdef").translate(Rexx("12"), Rexx("ec"), Rexx("AXE")) -- netrexx.lang.NotCharacterException: AXE

	_, err := RxFromString("abcdef").Translate(RxFromString("12"), RxFromString("ec"), RxFromString("AXE"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1830(t *testing.T) {

	// -- say Rexx("").translate(Rexx("12"), Rexx("ec"), Rexx(" "))

	value, err := rxfromempty().Translate(RxFromString("12"), RxFromString("ec"), RxFromString(" "))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1831(t *testing.T) {

	// -- say Rexx("abcdef").translate(Rexx(""), Rexx("ec"), Rexx(" "))

	value, err := RxFromString("abcdef").Translate(rxfromempty(), RxFromString("ec"), RxFromString(" "))
	want := "ab d f"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1832(t *testing.T) {

	// -- say Rexx("abcdef").translate(Rexx("12"), Rexx(""), Rexx(" "))

	value, err := RxFromString("abcdef").Translate(RxFromString("12"), rxfromempty(), RxFromString(" "))
	want := "abcdef"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1833(t *testing.T) {

	// -- say Rexx("-1234").translate(Rexx("2"), Rexx("1"), Rexx(" "))

	value, err := RxFromString("-1234").Translate(RxFromString("2"), RxFromString("1"), RxFromString(" "))
	want := "-2234"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1834(t *testing.T) {

	// -- say Rexx("4123").translate(Rexx("abcd"), Rexx("1234"), Rexx(" "))

	value, err := RxFromString("4123").Translate(RxFromString("abcd"), RxFromString("1234"), RxFromString(" "))
	want := "dabc"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1835(t *testing.T) {

	// -- say Rexx("4123").translate(Rexx("abcd"), Rexx("1234"), Rexx(" "))

	value, err := RxFromString("4123").Translate(RxFromString("abcd"), RxFromString("1234"), RxFromString(" "))
	want := "dabc"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1836(t *testing.T) {

	// -- say Rexx("4123").translate(Rexx("hods"), Rexx("1234"), Rexx(" "))

	value, err := RxFromString("4123").Translate(RxFromString("hods"), RxFromString("1234"), RxFromString(" "))
	want := "shod"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1837(t *testing.T) {

	// -- say Rexx("4312").translate(Rexx("hods"), Rexx("1234"), Rexx(" "))

	value, err := RxFromString("4312").Translate(RxFromString("hods"), RxFromString("1234"), RxFromString(" "))
	want := "sdho"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1838(t *testing.T) {

	// -- say Rexx("abbc").translate(Rexx(char '&'), Rexx(char 'b'), Rexx(" "))

	value, err := RxFromString("abbc").Translate(RxFromRune('&'), RxFromRune('b'), RxFromString(" "))
	want := "a&&c"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1839(t *testing.T) {

	// -- say Rexx("abcdef").translate(Rexx("12"), Rexx("abcd"), Rexx(char '.'))

	value, err := RxFromString("abcdef").Translate(RxFromString("12"), RxFromString("abcd"), RxFromRune('.'))
	want := "12..ef"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1840(t *testing.T) {

	// -- say Rexx("abcdef").translate(Rexx("12"), Rexx("ec"), Rexx(" "))

	value, err := RxFromString("abcdef").Translate(RxFromString("12"), RxFromString("ec"), RxFromString(" "))
	want := "ab2d1f"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Trunc(n *Rexx) (*Rexx, error)

func Test_1841(t *testing.T) {

	// -- say Rexx("-1234").trunc(null) -- java.lang.NullPointerException

	_, err := RxFromString("-1234").Trunc(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1842(t *testing.T) {

	// -- say Rexx("A.3").trunc(Rexx(int -11)) -- java.lang.NumberFormatException: A.3

	_, err := RxFromString("A.3").Trunc(RxFromInt32(-11))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1843(t *testing.T) {

	// -- say Rexx("12.3").trunc(Rexx(float 1.1)) -- java.lang.NumberFormatException: Non-zero decimal part in 1.1

	_, err := RxFromString("12.3").Trunc(RxFromFloat32(1.1))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1844(t *testing.T) {

	// -- say Rexx.toRexx("-0.9e-999999999").trunc(Rexx.toRexx("-00")) -- netrexx.lang.ExponentOverflowException: -1000000000

	_, err := ToRxFromString("-0.9e-999999999").Trunc(ToRxFromString("-00"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1845(t *testing.T) {

	// -- say Rexx("0").trunc(Rexx("1")) -- uses different call

	value, err := rxfromempty().Trunc(RxFromString("1"))
	want := "0.0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1846(t *testing.T) {

	// -- say Rexx("-1234").trunc(Rexx("2"))

	value, err := RxFromString("-1234").Trunc(RxFromString("2"))
	want := "-1234.00"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1847(t *testing.T) {

	// -- say Rexx("0").trunc(Rexx("2"))

	value, err := RxFromString("0").Trunc(RxFromString("2"))
	want := "0.00"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1848(t *testing.T) {

	// -- say Rexx("12.3").trunc(Rexx("0"))

	value, err := RxFromString("12.3").Trunc(RxFromString("0"))
	want := "12"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1849(t *testing.T) {

	// -- say Rexx("127").trunc(Rexx("2"))

	value, err := RxFromString("127").Trunc(RxFromString("2"))
	want := "127.00"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1850(t *testing.T) {

	// -- say Rexx("127.09782").trunc(Rexx("3"))

	value, err := RxFromString("127.09782").Trunc(RxFromString("3"))
	want := "127.097"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1851(t *testing.T) {

	// -- say Rexx("127.1").trunc(Rexx("3"))

	value, err := RxFromString("127.1").Trunc(RxFromString("3"))
	want := "127.100"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1852(t *testing.T) {

	// -- say Rexx("-014.014").trunc(Rexx(int 1929))

	value, err := RxFromString("-014.014").Trunc(RxFromInt32(1929))
	want := "-14.014000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1853(t *testing.T) {

	// -- say Rexx("-014.014").trunc(Rexx(int 1929))

	new_set := RxSet()
	new_set.SetDigits(RxFromString("16"))

	MaxExp = 9
	MinExp = -9

	_, err := RxFromString("-014.014").Trunc(RxFromInt32(1929))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999
	MinExp = -999999999

}

func Test_1854(t *testing.T) {

	// -- say Rexx("۱۱۱۱.0E+6").trunc(Rexx(int 1))

	value, err := RxFromString("۱۱۱۱.0E+6").Trunc(RxFromInt32(1))
	want := "1111000000.0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Upper(n *Rexx, length *Rexx) (*Rexx, error)

func Test_1855(t *testing.T) {

	// -- say Rexx("-1234").upper(null, Rexx("1")) -- java.lang.NullPointerException

	_, err := RxFromString("-1234").Upper(nil, RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1856(t *testing.T) {

	// -- say Rexx("-1234").upper(Rexx("2"), null) -- java.lang.NullPointerException

	_, err := RxFromString("-1234").Upper(RxFromString("2"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1857(t *testing.T) {

	// -- say Rexx("-1234").upper(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("-1234").Upper(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1858(t *testing.T) {

	// -- say Rexx("Mad Sheep").upper(Rexx("0"), Rexx("Mad Sheep").length()) -- netrexx.lang.BadArgumentException: Argument 0 < 1

	_, err := RxFromString("Mad Sheep").Upper(RxFromString("0"), RxFromString("Mad Sheep").Length())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1859(t *testing.T) {

	// -- say Rexx("Mad Sheep").upper(Rexx("1"), Rexx(int 999999999+1))

	_, err := RxFromString("Mad Sheep").Upper(RxFromString("1"), RxFromInt32(MaxArg+1))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1860(t *testing.T) {

	// -- say Rexx("").upper(Rexx("1"), Rexx("Mad Sheep").length())

	value, err := rxfromempty().Upper(RxFromString("1"), RxFromString("Mad Sheep").Length())
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1861(t *testing.T) {

	// -- say Rexx("").upper(Rexx("1"), Rexx("").length())

	value, err := RxFromString("").Upper(RxFromString("1"), RxFromString("").Length())
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1862(t *testing.T) {

	// -- say Rexx("-1234").upper(Rexx("2"), Rexx("1"))

	value, err := RxFromString("-1234").Upper(RxFromString("2"), RxFromString("1"))
	want := "-1234"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1863(t *testing.T) {

	// -- say Rexx("Mad sheep").upper(Rexx("5"), Rexx("1"))

	value, err := RxFromString("Mad sheep").Upper(RxFromString("5"), RxFromString("1"))
	want := "Mad Sheep"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1864(t *testing.T) {

	// -- say Rexx("Mad sheep").upper(Rexx("5"), Rexx("4"))

	value, err := RxFromString("Mad sheep").Upper(RxFromString("5"), RxFromString("4"))
	want := "Mad SHEEp"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1865(t *testing.T) {

	// -- say Rexx("Mad sheep").upper(Rexx("5"), Rexx("Mad sheep").length())

	value, err := RxFromString("Mad sheep").Upper(RxFromString("5"), RxFromString("Mad sheep").Length())
	want := "Mad SHEEP"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1866(t *testing.T) {

	// -- say Rexx("Fou-Baa").upper(Rexx("1"), Rexx("Fou-Baa").length())

	value, err := RxFromString("Fou-Baa").Upper(RxFromString("1"), RxFromString("Fou-Baa").Length())
	want := "FOU-BAA"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1867(t *testing.T) {

	// -- say Rexx("Mad Sheep").upper(Rexx("1"), Rexx("Mad Sheep").length())

	value, err := RxFromString("Mad Sheep").Upper(RxFromString("1"), RxFromString("Mad Sheep").Length())
	want := "MAD SHEEP"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1868(t *testing.T) {

	// -- say Rexx("tinganon").upper(Rexx("1"), Rexx("1"))

	value, err := RxFromString("tinganon").Upper(RxFromString("1"), RxFromString("1"))
	want := "Tinganon"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Verify(list *Rexx, opt *Rexx, start *Rexx) (*Rexx, error)

func Test_1869(t *testing.T) {

	// -- say Rexx("-123").verify(null, Rexx("1"), Rexx(" ")) -- netrexx.lang.BadArgumentException: Bad Option character 1 [1]

	_, err := RxFromString("-123").Verify(nil, RxFromString("1"), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1870(t *testing.T) {

	// -- say Rexx("-123").verify(Rexx("2"), null, Rexx(" ")) -- java.lang.NullPointerException

	_, err := RxFromString("-123").Verify(RxFromString("2"), nil, RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1871(t *testing.T) {

	// -- say Rexx("-123").verify(Rexx("2"), Rexx("1"), null) -- netrexx.lang.BadArgumentException: Bad Option character 1 [1]

	_, err := RxFromString("-123").Verify(RxFromString("2"), RxFromString("1"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1872(t *testing.T) {

	// -- say Rexx("-123").verify(null, null, null) -- java.lang.NullPointerException

	_, err := RxFromString("-123").Verify(nil, nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1873(t *testing.T) {

	// -- say Rexx("-123").verify(Rexx("2"), Rexx("1"), Rexx(" ")) -- netrexx.lang.BadArgumentException: Bad Option character 1 [1]

	_, err := RxFromString("-123").Verify(RxFromString("2"), RxFromString("1"), RxFromString(" "))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1874(t *testing.T) {

	// -- say Rexx("1Z3").verify(Rexx("1234567890"), Rexx(""), Rexx("1")) -- netrexx.lang.BadArgumentException: Null option string

	_, err := RxFromString("1Z3").Verify(RxFromString("1234567890"), RxFromString(""), RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1875(t *testing.T) {

	// -- say Rexx("1Z3").verify(Rexx("1234567890"), Rexx("N"), Rexx("-1")) -- netrexx.lang.BadArgumentException: Argument -1 < 1

	_, err := RxFromString("1Z3").Verify(RxFromString("1234567890"), RxFromString("N"), RxFromString("-1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1876(t *testing.T) {

	// -- say Rexx("").verify(Rexx("1234567890"), Rexx("N"), Rexx("1"))

	value, err := rxfromempty().Verify(RxFromString("1234567890"), RxFromString("N"), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1877(t *testing.T) {

	// -- say Rexx("1Z3").verify(Rexx(""), Rexx("N"), Rexx("1"))

	value, err := RxFromString("1Z3").Verify(rxfromempty(), RxFromString("N"), RxFromString("1"))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1878(t *testing.T) {

	// -- say Rexx("ABCDE").verify(Rexx(""), Rexx(char 'n'), Rexx("3"))

	value, err := RxFromString("ABCDE").Verify(RxFromString(""), RxFromRune('n'), RxFromString("3"))
	want := "3"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1879(t *testing.T) {

	// -- say Rexx("123").verify(Rexx("1234567890"), Rexx("N"), Rexx("1"))

	value, err := RxFromString("123").Verify(RxFromString("1234567890"), RxFromString("N"), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1880(t *testing.T) {

	// -- say Rexx("1P3Q4").verify(Rexx("1234567890"), Rexx(char 'N'), Rexx("3"))

	value, err := RxFromString("1P3Q4").Verify(RxFromString("1234567890"), RxFromRune('N'), RxFromString("3"))
	want := "4"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1881(t *testing.T) {

	// -- say Rexx("1Z3").verify(Rexx("1234567890"), Rexx("N"), Rexx("1"))

	value, err := RxFromString("1Z3").Verify(RxFromString("1234567890"), RxFromString("N"), RxFromString("1"))
	want := "2"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1882(t *testing.T) {

	// -- say Rexx("AB3CD5").verify(Rexx("1234567890"), Rexx(char 'm'), Rexx("4"))

	value, err := RxFromString("AB3CD5").Verify(RxFromString("1234567890"), RxFromRune('m'), RxFromString("4"))
	want := "6"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1883(t *testing.T) {

	// -- say Rexx("AB4T").verify(Rexx("1234567890"), Rexx(char 'M'), Rexx("1"))

	value, err := RxFromString("AB4T").Verify(RxFromString("1234567890"), RxFromRune('M'), RxFromString("1"))
	want := "3"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1884(t *testing.T) {

	// -- say Rexx("AB4T").verify(Rexx("1234567890"), Rexx(char 'M'), Rexx("0")) -- netrexx.lang.BadArgumentException: Argument 0 < 1

	_, err := RxFromString("AB4T").Verify(RxFromString("1234567890"), RxFromRune('M'), RxFromString("0"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

// func (rcvr *Rexx) Word(n *Rexx) (*Rexx, error)

func Test_1885(t *testing.T) {

	// -- say Rexx("-123").word(null) -- java.lang.NullPointerException

	_, err := RxFromString("-123").Word(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1886(t *testing.T) {

	// -- say Rexx("1Z3").word(Rexx()) -- java.lang.NullPointerException

	_, err := RxFromString("1Z3").Word(rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1887(t *testing.T) {

	// -- say Rexx("-123").word(Rexx("2"))

	value, err := RxFromString("-123").Word(RxFromString("2"))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1888(t *testing.T) {

	// -- say Rexx("Now is the time").word(Rexx("3"))

	value, err := RxFromString("Now is the time").Word(RxFromString("3"))
	want := "the"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1889(t *testing.T) {

	// -- say Rexx("Now is the time").word(Rexx("5"))

	value, err := RxFromString("Now is the time").Word(RxFromString("5"))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1890(t *testing.T) {

	// -- say Rexx("Now is the time").word(Rexx("5"))

	value, err := RxFromString("Now is the time").Word(RxFromString("5"))
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) WordIndex(n *Rexx) (*Rexx, error)

func Test_1891(t *testing.T) {

	// -- say Rexx("-123").wordindex(null) -- java.lang.NullPointerException

	_, err := RxFromString("-123").WordIndex(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1892(t *testing.T) {

	// -- say Rexx("1Z3").wordindex(Rexx()) -- java.lang.NullPointerException

	_, err := RxFromString("1Z3").WordIndex(rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1893(t *testing.T) {

	// -- say Rexx("").wordindex(Rexx("1"))

	value, err := rxfromempty().WordIndex(RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1894(t *testing.T) {

	// -- say Rexx("-123").wordindex(Rexx("2"))

	value, err := RxFromString("-123").WordIndex(RxFromString("2"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1895(t *testing.T) {

	// -- say Rexx("Now is the time").wordindex(Rexx("3"))

	value, err := RxFromString("Now is the time").WordIndex(RxFromString("3"))
	want := "8"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1896(t *testing.T) {

	// -- say Rexx("Now is the time").wordindex(Rexx("6"))

	value, err := RxFromString("Now is the time").WordIndex(RxFromString("6"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) WordLength(n *Rexx) (*Rexx, error)

func Test_1897(t *testing.T) {

	// -- say Rexx("-123").wordlength(null) -- java.lang.NullPointerException

	_, err := RxFromString("-123").WordLength(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1898(t *testing.T) {

	// -- say Rexx("1Z3").wordlength(Rexx()) -- java.lang.NullPointerException

	_, err := RxFromString("1Z3").WordLength(rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1899(t *testing.T) {

	// -- say Rexx("").wordlength(Rexx("1"))

	value, err := rxfromempty().WordLength(RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1900(t *testing.T) {

	// -- say Rexx("-123").wordlength(Rexx("2"))

	value, err := RxFromString("-123").WordLength(RxFromString("2"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1901(t *testing.T) {

	// -- say Rexx("Now comes the time").wordlength(Rexx("2"))

	value, err := RxFromString("Now comes the time").WordLength(RxFromString("2"))
	want := "5"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1902(t *testing.T) {

	// -- say Rexx("Now is the time").wordlength(Rexx("2"))

	value, err := RxFromString("Now is the time").WordLength(RxFromString("2"))
	want := "2"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1903(t *testing.T) {

	// -- say Rexx("Now is the time").wordlength(Rexx("6"))

	value, err := RxFromString("Now is the time").WordLength(RxFromString("6"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) WordPos(needle *Rexx, num *Rexx) (*Rexx, error)

func Test_1904(t *testing.T) {

	// -- say Rexx("-123").wordpos(null, Rexx("1")) -- java.lang.NullPointerException

	_, err := RxFromString("-123").WordPos(nil, RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1905(t *testing.T) {

	// -- say Rexx("-123").wordpos(Rexx("2"), null) -- java.lang.NullPointerException

	_, err := RxFromString("-123").WordPos(RxFromString("2"), nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1906(t *testing.T) {

	// -- say Rexx("-123").wordpos(null, null) -- java.lang.NullPointerException

	_, err := RxFromString("-123").WordPos(nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1907(t *testing.T) {

	// -- say Rexx("Now is the time").wordpos(Rexx("The"), Rexx()) -- java.lang.NullPointerException

	_, err := RxFromString("Now is the time").WordPos(RxFromString("The"), rxfromempty())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1908(t *testing.T) {

	// -- say Rexx("").wordpos(Rexx("The"), Rexx("1"))

	value, err := rxfromempty().WordPos(RxFromString("The"), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1909(t *testing.T) {

	// -- say Rexx("Now is the time").wordpos(Rexx(""), Rexx("1"))

	value, err := RxFromString("Now is the time").WordPos(rxfromempty(), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1910(t *testing.T) {

	// -- say Rexx("-123").wordpos(Rexx("2"), Rexx("1"))

	value, err := RxFromString("-123").WordPos(RxFromString("2"), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1911(t *testing.T) {

	// -- say Rexx("Now is the time").wordpos(Rexx("The"), Rexx("1"))

	value, err := RxFromString("Now is the time").WordPos(RxFromString("The"), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1912(t *testing.T) {

	// -- say Rexx("Now is the time").wordpos(Rexx("is   the"), Rexx("1"))

	value, err := RxFromString("Now is the time").WordPos(RxFromString("is   the"), RxFromString("1"))
	want := "2"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1913(t *testing.T) {

	// -- say Rexx("Now is the time").wordpos(Rexx("is the"), Rexx("1"))

	value, err := RxFromString("Now is the time").WordPos(RxFromString("is the"), RxFromString("1"))
	want := "2"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1914(t *testing.T) {

	// -- say Rexx("Now is the time").wordpos(Rexx("is time"), Rexx("1"))

	value, err := RxFromString("Now is the time").WordPos(RxFromString("is time"), RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1915(t *testing.T) {

	// -- say Rexx("Now is the time").wordpos(Rexx("the"), Rexx("1"))

	value, err := RxFromString("Now is the time").WordPos(RxFromString("the"), RxFromString("1"))
	want := "3"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1916(t *testing.T) {

	// -- say Rexx("To be or not to be").wordpos(Rexx("be"), Rexx("1"))

	value, err := RxFromString("To be or not to be").WordPos(RxFromString("be"), RxFromString("1"))
	want := "2"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1917(t *testing.T) {

	// -- say Rexx("To be or not to be").wordpos(Rexx("be"), Rexx("3"))

	value, err := RxFromString("To be or not to be").WordPos(RxFromString("be"), RxFromString("3"))
	want := "6"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) Words() (*Rexx, error)

func Test_1918(t *testing.T) {

	// -- say Rexx("").words()

	value, err := rxfromempty().Words()
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1919(t *testing.T) {

	// -- say Rexx(" ").words()

	value, err := RxFromString(" ").Words()
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1920(t *testing.T) {

	// -- say Rexx("").words()

	value, err := RxFromString("").Words()
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1921(t *testing.T) {

	// -- say Rexx("Now  time").words()

	value, err := RxFromString("Now  time").Words()
	want := "2"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1922(t *testing.T) {

	// -- say Rexx("Now is the  time ").words()

	value, err := RxFromString("Now is the  time ").Words()
	want := "4"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1923(t *testing.T) {

	// -- say Rexx("Now is the time").words()

	value, err := RxFromString("Now is the time").Words()
	want := "4"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) X2b() (*Rexx, error)

func Test_1924(t *testing.T) {

	// -- say Rexx("").x2b() -- netrexx.lang.BadArgumentException: No digits

	_, err := rxfromempty().X2b()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1925(t *testing.T) {

	// -- say Rexx(" ").x2b() -- netrexx.lang.BadArgumentException: Bad hexadecimal

	_, err := RxFromString(" ").X2b()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1926(t *testing.T) {

	// -- say Rexx("J").x2b() -- netrexx.lang.BadArgumentException: Bad hexadecimal J

	_, err := RxFromString("J").X2b()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1927(t *testing.T) {

	// -- say Rexx("1C1").x2b()

	value, err := RxFromString("1C1").X2b()
	want := "000111000001"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1928(t *testing.T) {

	// -- say Rexx(char '7').x2b()

	value, err := RxFromRune('7').X2b()
	want := "0111"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1929(t *testing.T) {

	// -- say Rexx("C3").x2b()

	value, err := RxFromString("C3").X2b()
	want := "11000011"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) X2c() (*Rexx, error)

func Test_1930(t *testing.T) {

	// -- say Rexx(").x2c() -- netrexx.lang.BadArgumentException: No digits

	_, err := rxfromempty().X2c()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1931(t *testing.T) {

	// -- say Rexx(" ").x2c() -- netrexx.lang.BadArgumentException: Bad hexadecimal

	_, err := RxFromString(" ").X2c()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1932(t *testing.T) {

	// -- say Rexx("J").x2c() -- netrexx.lang.BadArgumentException: Bad hexadecimal J

	_, err := RxFromString("J").X2c()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1933(t *testing.T) {

	// -- say Rexx("FFFFFFFFFFFFFFFFFFFFFFFF").x2c() -- netrexx.lang.BadArgumentException: Too big FFFFFFFFFFFFFFFFFFFFFFFF

	_, err := RxFromString("FFFFFFFFFFFFFFFFFFFFFFFF").X2c()

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1934(t *testing.T) {

	// -- say '0'.x2c()

	value, err := RxFromRune('0').X2c()
	want := rune(0)

	result, _ := value.ToRune()

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := result
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1935(t *testing.T) {

	// -- say Rexx("004D").x2c()

	value, err := RxFromString("004D").X2c()
	want := "M"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1936(t *testing.T) {

	// -- say Rexx("4d").x2c()

	value, err := RxFromString("4d").X2c()
	want := "M"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) X2d(n *Rexx) (*Rexx, error)

func Test_1937(t *testing.T) {

	// -- say Rexx("-123").x2d(null) -- java.lang.NullPointerException

	_, err := RxFromString("-123").X2d(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1938(t *testing.T) {

	// -- say Rexx(" ").x2d(Rexx(int 999999999 + 1))

	_, err := RxFromString(" ").X2d(RxFromInt32(MaxArg + 1))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1939(t *testing.T) {

	// -- say Rexx("").x2d(Rexx("1"))

	value, err := rxfromempty().X2d(RxFromString("1"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1940(t *testing.T) {

	// -- say Rexx("J").x2d(Rexx("1")) -- netrexx.lang.BadArgumentException: Bad hexadecimal J

	_, err := RxFromString("J").X2d(RxFromString("1"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_1941(t *testing.T) {

	// -- say Rexx("-123").x2d(Rexx("2"))

	value, err := RxFromString("-123").X2d(RxFromString("2"))
	want := "35"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1942(t *testing.T) {

	// -- say Rexx("0031").x2d(Rexx("0"))

	value, err := RxFromString("0031").X2d(RxFromString("0"))
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1943(t *testing.T) {

	// -- say Rexx("0E").x2d()
	// -- This port allows -1 as an argument
	// -- 	to bypass golang not enough arguments in call

	value, err := RxFromString("0E").X2d(RxFromString("-1"))
	want := "14"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1944(t *testing.T) {

	// -- say Rexx("81").x2d()
	// -- This port allows -1 as an argument
	// -- 	to bypass golang not enough arguments in call

	value, err := RxFromString("81").X2d(RxFromString("-1"))
	want := "129"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1945(t *testing.T) {

	// -- say Rexx("81").x2d(Rexx("2"))

	value, err := RxFromString("81").X2d(RxFromString("2"))
	want := "-127"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1946(t *testing.T) {

	// -- say Rexx("81").x2d(Rexx("4"))

	value, err := RxFromString("81").X2d(RxFromString("4"))
	want := "129"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1947(t *testing.T) {

	// -- say Rexx("F081").x2d(Rexx("1"))

	value, err := RxFromString("F081").X2d(RxFromString("1"))
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1948(t *testing.T) {

	// -- say Rexx("F081").x2d(Rexx("2"))

	value, err := RxFromString("F081").X2d(RxFromString("2"))
	want := "-127"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1949(t *testing.T) {

	// -- say Rexx("F081").x2d(Rexx("3"))

	value, err := RxFromString("F081").X2d(RxFromString("3"))
	want := "129"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1950(t *testing.T) {

	// -- say Rexx("F081").x2d(Rexx("4"))

	value, err := RxFromString("F081").X2d(RxFromString("4"))
	want := "-3967"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1951(t *testing.T) {

	// -- say Rexx("F81").x2d()
	// -- This port allows -1 as an argument
	// -- 	to bypass golang not enough arguments in call

	value, err := RxFromString("F81").X2d(RxFromString("-1"))
	want := "3969"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1952(t *testing.T) {

	// -- say Rexx("FF81").x2d()
	// -- This port allows -1 as an argument
	// -- 	to bypass golang not enough arguments in call

	value, err := RxFromString("FF81").X2d(RxFromString("-1"))
	want := "65409"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1953(t *testing.T) {

	// -- say Rexx("c6f0").x2d()
	// -- This port allows -1 as an argument
	// -- 	to bypass golang not enough arguments in call

	value, err := RxFromString("c6f0").X2d(RxFromString("-1"))
	want := "50928"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func (rcvr *Rexx) concat(set *RexxSet, rhs *Rexx, blanks int32) *Rexx

// func (rcvr *Rexx) cut(d int32)

func Test_1954(t *testing.T) {

	// -- say "Shortcut used just to test two block codes"
	// -- say "cut and round need proper tests"

	got := RxFromString("0")

	got.cut(5)   // Shortcut
	got.round(5) //Shortcut

	RxException(12, "test") //Shortcut

	want := "0"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

// func (rcvr *Rexx) docompare(set *RexxSet, rhs *Rexx) (int8, error)

// func (rcvr *Rexx) docomparestrict(set *RexxSet, rhs *Rexx) int32

// func (rcvr *Rexx) dodivide(code rune, set *RexxSet, rhs *Rexx) (*Rexx, error)

// func (rcvr *Rexx) finish(roundlen int32, strip bool) (*Rexx, error)

// func (rcvr *Rexx) intcheck(min int32, max int32) (int32, error)

// func (rcvr *Rexx) intlength() int32

// func (rcvr *Rexx) intwords() int32

// func (rcvr *Rexx) layout() []rune

// func (rcvr *Rexx) layoutnum() []rune

// func (rcvr *Rexx) optioncheck(oklist string) (rune, error)

// func (rcvr *Rexx) padcheck() (rune, error)

// func (rcvr *Rexx) round(d int32)

// func charaddsub(a []rune, b []rune, m int32) []rune

// func getdigit(char rune) int32

func Test_1955(t *testing.T) {

	// --

	unigroup := [370]rune{
		'\u0030', '\u0031', '\u0032', '\u0033', '\u0034', '\u0035', '\u0036', '\u0037', '\u0038', '\u0039',
		'\u0660', '\u0661', '\u0662', '\u0663', '\u0664', '\u0665', '\u0666', '\u0667', '\u0668', '\u0669',
		'\u06F0', '\u06F1', '\u06F2', '\u06F3', '\u06F4', '\u06F5', '\u06F6', '\u06F7', '\u06F8', '\u06F9',
		'\u07C0', '\u07C1', '\u07C2', '\u07C3', '\u07C4', '\u07C5', '\u07C6', '\u07C7', '\u07C8', '\u07C9',
		'\u0966', '\u0967', '\u0968', '\u0969', '\u096A', '\u096B', '\u096C', '\u096D', '\u096E', '\u096F',
		'\u09E6', '\u09E7', '\u09E8', '\u09E9', '\u09EA', '\u09EB', '\u09EC', '\u09ED', '\u09EE', '\u09EF',
		'\u0A66', '\u0A67', '\u0A68', '\u0A69', '\u0A6A', '\u0A6B', '\u0A6C', '\u0A6D', '\u0A6E', '\u0A6F',
		'\u0AE6', '\u0AE7', '\u0AE8', '\u0AE9', '\u0AEA', '\u0AEB', '\u0AEC', '\u0AED', '\u0AEE', '\u0AEF',
		'\u0B66', '\u0B67', '\u0B68', '\u0B69', '\u0B6A', '\u0B6B', '\u0B6C', '\u0B6D', '\u0B6E', '\u0B6F',
		'\u0BE6', '\u0BE7', '\u0BE8', '\u0BE9', '\u0BEA', '\u0BEB', '\u0BEC', '\u0BED', '\u0BEE', '\u0BEF',
		'\u0C66', '\u0C67', '\u0C68', '\u0C69', '\u0C6A', '\u0C6B', '\u0C6C', '\u0C6D', '\u0C6E', '\u0C6F',
		'\u0CE6', '\u0CE7', '\u0CE8', '\u0CE9', '\u0CEA', '\u0CEB', '\u0CEC', '\u0CED', '\u0CEE', '\u0CEF',
		'\u0D66', '\u0D67', '\u0D68', '\u0D69', '\u0D6A', '\u0D6B', '\u0D6C', '\u0D6D', '\u0D6E', '\u0D6F',
		'\u0DE6', '\u0DE7', '\u0DE8', '\u0DE9', '\u0DEA', '\u0DEB', '\u0DEC', '\u0DED', '\u0DEE', '\u0DEF',
		'\u0E50', '\u0E51', '\u0E52', '\u0E53', '\u0E54', '\u0E55', '\u0E56', '\u0E57', '\u0E58', '\u0E59',
		'\u0ED0', '\u0ED1', '\u0ED2', '\u0ED3', '\u0ED4', '\u0ED5', '\u0ED6', '\u0ED7', '\u0ED8', '\u0ED9',
		'\u0F20', '\u0F21', '\u0F22', '\u0F23', '\u0F24', '\u0F25', '\u0F26', '\u0F27', '\u0F28', '\u0F29',
		'\u1040', '\u1041', '\u1042', '\u1043', '\u1044', '\u1045', '\u1046', '\u1047', '\u1048', '\u1049',
		'\u1090', '\u1091', '\u1092', '\u1093', '\u1094', '\u1095', '\u1096', '\u1097', '\u1098', '\u1099',
		'\u17E0', '\u17E1', '\u17E2', '\u17E3', '\u17E4', '\u17E5', '\u17E6', '\u17E7', '\u17E8', '\u17E9',
		'\u1810', '\u1811', '\u1812', '\u1813', '\u1814', '\u1815', '\u1816', '\u1817', '\u1818', '\u1819',
		'\u1946', '\u1947', '\u1948', '\u1949', '\u194A', '\u194B', '\u194C', '\u194D', '\u194E', '\u194F',
		'\u19D0', '\u19D1', '\u19D2', '\u19D3', '\u19D4', '\u19D5', '\u19D6', '\u19D7', '\u19D8', '\u19D9',
		'\u1A80', '\u1A81', '\u1A82', '\u1A83', '\u1A84', '\u1A85', '\u1A86', '\u1A87', '\u1A88', '\u1A89',
		'\u1A90', '\u1A91', '\u1A92', '\u1A93', '\u1A94', '\u1A95', '\u1A96', '\u1A97', '\u1A98', '\u1A99',
		'\u1B50', '\u1B51', '\u1B52', '\u1B53', '\u1B54', '\u1B55', '\u1B56', '\u1B57', '\u1B58', '\u1B59',
		'\u1BB0', '\u1BB1', '\u1BB2', '\u1BB3', '\u1BB4', '\u1BB5', '\u1BB6', '\u1BB7', '\u1BB8', '\u1BB9',
		'\u1C40', '\u1C41', '\u1C42', '\u1C43', '\u1C44', '\u1C45', '\u1C46', '\u1C47', '\u1C48', '\u1C49',
		'\u1C50', '\u1C51', '\u1C52', '\u1C53', '\u1C54', '\u1C55', '\u1C56', '\u1C57', '\u1C58', '\u1C59',
		'\uA620', '\uA621', '\uA622', '\uA623', '\uA624', '\uA625', '\uA626', '\uA627', '\uA628', '\uA629',
		'\uA8D0', '\uA8D1', '\uA8D2', '\uA8D3', '\uA8D4', '\uA8D5', '\uA8D6', '\uA8D7', '\uA8D8', '\uA8D9',
		'\uA900', '\uA901', '\uA902', '\uA903', '\uA904', '\uA905', '\uA906', '\uA907', '\uA908', '\uA909',
		'\uA9D0', '\uA9D1', '\uA9D2', '\uA9D3', '\uA9D4', '\uA9D5', '\uA9D6', '\uA9D7', '\uA9D8', '\uA9D9',
		'\uA9F0', '\uA9F1', '\uA9F2', '\uA9F3', '\uA9F4', '\uA9F5', '\uA9F6', '\uA9F7', '\uA9F8', '\uA9F9',
		'\uAA50', '\uAA51', '\uAA52', '\uAA53', '\uAA54', '\uAA55', '\uAA56', '\uAA57', '\uAA58', '\uAA59',
		'\uABF0', '\uABF1', '\uABF2', '\uABF3', '\uABF4', '\uABF5', '\uABF6', '\uABF7', '\uABF8', '\uABF9',
		'\uFF10', '\uFF11', '\uFF12', '\uFF13', '\uFF14', '\uFF15', '\uFF16', '\uFF17', '\uFF18', '\uFF19',
	}

	hexgroup := [260]rune{
		0x104A0, 0x104A1, 0x104A2, 0x104A3, 0x104A4, 0x104A5, 0x104A6, 0x104A7, 0x104A8, 0x104A9,
		0x10D30, 0x10D31, 0x10D32, 0x10D33, 0x10D34, 0x10D35, 0x10D36, 0x10D37, 0x10D38, 0x10D39,
		0x11066, 0x11067, 0x11068, 0x11069, 0x1106A, 0x1106B, 0x1106C, 0x1106D, 0x1106E, 0x1106F,
		0x110F0, 0x110F1, 0x110F2, 0x110F3, 0x110F4, 0x110F5, 0x110F6, 0x110F7, 0x110F8, 0x110F9,
		0x11136, 0x11137, 0x11138, 0x11139, 0x1113A, 0x1113B, 0x1113C, 0x1113D, 0x1113E, 0x1113F,
		0x111D0, 0x111D1, 0x111D2, 0x111D3, 0x111D4, 0x111D5, 0x111D6, 0x111D7, 0x111D8, 0x111D9,
		0x112F0, 0x112F1, 0x112F2, 0x112F3, 0x112F4, 0x112F5, 0x112F6, 0x112F7, 0x112F8, 0x112F9,
		0x11450, 0x11451, 0x11452, 0x11453, 0x11454, 0x11455, 0x11456, 0x11457, 0x11458, 0x11459,
		0x114D0, 0x114D1, 0x114D2, 0x114D3, 0x114D4, 0x114D5, 0x114D6, 0x114D7, 0x114D8, 0x114D9,
		0x11650, 0x11651, 0x11652, 0x11653, 0x11654, 0x11655, 0x11656, 0x11657, 0x11658, 0x11659,
		0x116C0, 0x116C1, 0x116C2, 0x116C3, 0x116C4, 0x116C5, 0x116C6, 0x116C7, 0x116C8, 0x116C9,
		0x11730, 0x11731, 0x11732, 0x11733, 0x11734, 0x11735, 0x11736, 0x11737, 0x11738, 0x11739,
		0x118E0, 0x118E1, 0x118E2, 0x118E3, 0x118E4, 0x118E5, 0x118E6, 0x118E7, 0x118E8, 0x118E9,
		0x11C50, 0x11C51, 0x11C52, 0x11C53, 0x11C54, 0x11C55, 0x11C56, 0x11C57, 0x11C58, 0x11C59,
		0x11D50, 0x11D51, 0x11D52, 0x11D53, 0x11D54, 0x11D55, 0x11D56, 0x11D57, 0x11D58, 0x11D59,
		0x11DA0, 0x11DA1, 0x11DA2, 0x11DA3, 0x11DA4, 0x11DA5, 0x11DA6, 0x11DA7, 0x11DA8, 0x11DA9,
		0x16A60, 0x16A61, 0x16A62, 0x16A63, 0x16A64, 0x16A65, 0x16A66, 0x16A67, 0x16A68, 0x16A69,
		0x16B50, 0x16B51, 0x16B52, 0x16B53, 0x16B54, 0x16B55, 0x16B56, 0x16B57, 0x16B58, 0x16B59,
		0x1D7CE, 0x1D7CF, 0x1D7D0, 0x1D7D1, 0x1D7D2, 0x1D7D3, 0x1D7D4, 0x1D7D5, 0x1D7D6, 0x1D7D7,
		0x1D7D8, 0x1D7D9, 0x1D7DA, 0x1D7DB, 0x1D7DC, 0x1D7DD, 0x1D7DE, 0x1D7DF, 0x1D7E0, 0x1D7E1,
		0x1D7E2, 0x1D7E3, 0x1D7E4, 0x1D7E5, 0x1D7E6, 0x1D7E7, 0x1D7E8, 0x1D7E9, 0x1D7EA, 0x1D7EB,
		0x1D7EC, 0x1D7ED, 0x1D7EE, 0x1D7EF, 0x1D7F0, 0x1D7F1, 0x1D7F2, 0x1D7F3, 0x1D7F4, 0x1D7F5,
		0x1D7F6, 0x1D7F7, 0x1D7F8, 0x1D7F9, 0x1D7FA, 0x1D7FB, 0x1D7FC, 0x1D7FD, 0x1D7FE, 0x1D7FF,
		//next two rows : go1.13.1 openbsd[6.6]/amd64 unicode.IsDigit not working for me
		0x1E140, 0x1E141, 0x1E142, 0x1E143, 0x1E144, 0x1E145, 0x1E146, 0x1E147, 0x1E148, 0x1E149,
		0x1E2F0, 0x1E2F1, 0x1E2F2, 0x1E2F3, 0x1E2F4, 0x1E2F5, 0x1E2F6, 0x1E2F7, 0x1E2F8, 0x1E2F9,
		0x1E950, 0x1E951, 0x1E952, 0x1E953, 0x1E954, 0x1E955, 0x1E956, 0x1E957, 0x1E958, 0x1E959,
	}

	count := 0
	var col int32 = 0
	for row := 1; row <= 37; row++ {
		for col = 0; col <= 9; col++ {
			number, err := RxFromRune(unigroup[count]).ToInt32()
			if err != nil {
				t.Errorf("Expecting a value.")
				//fmt.Printf("err: row %d col %d\n", row, col)
			} else {
				if number != col {
					t.Errorf("got %v, wanted %v", number, col)
					//fmt.Printf("number: %d row %d col %d\n", number, row, col)
				}
			}
			count++
		}
	}

	count = 0
	for row := 1; row <= 24; row++ {
		for col = 0; col <= 9; col++ {
			number, err := RxFromRune(hexgroup[count]).ToInt32()
			if err != nil {
				t.Errorf("Expecting a value.")
				//fmt.Printf("err: row %d col %d\n", row, col)
			} else {
				if number != col {
					t.Errorf("got %v, wanted %v", number, col)
					//fmt.Printf("number: %d row %d col %d\n", number, row, col)
				}
			}
			count++
		}
	}

}

// func sa2ca(s []string) []rune

// func RxException(kind int, detail string) error

// func rxio() (rcvr *RexxIO)

func Test_1956(t *testing.T) {

	// --

	rxio := rxio()

	if rxio == nil {
		t.Errorf("Call should not be nil")
	}

}

// func Ask() (*Rexx, error)

func Test_1957(t *testing.T) {

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	input := "input\n"
	r, w, _ := os.Pipe()
	os.Stdin = r

	go func() {
		w.Write([]byte(input))
		w.Close()
	}()

	result, _ := Ask()

	if result.ToString() != "input" {
		t.Errorf("Expected input, got %s", result.ToString())
	}
}

func Test_1958(t *testing.T) {

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r

	go func() {
		w.Write([]byte(nil))
		w.Close()
	}()

	_, err := Ask()

	if err.Error() != "NotCharacterException : input error." {
		t.Errorf("Expected an error")
	}

}

// func AskOne() (*Rexx, error)

func Test_1959(t *testing.T) {

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	input := "1\n"
	r, w, _ := os.Pipe()
	os.Stdin = r

	go func() {
		w.Write([]byte(input))
		w.Close()
	}()

	result, _ := AskOne()

	if result.ToString() != "1" {
		t.Errorf("Expected character \"1\", got %s", result.ToString())
	}

}

func Test_1960(t *testing.T) {

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	r, w, _ := os.Pipe()
	os.Stdin = r

	go func() {
		w.Write([]byte(nil))
		w.Close()
	}()

	_, err := AskOne()

	if err.Error() != "NotCharacterException : input error." {
		t.Errorf("Expected an error")
	}

}

// captureStdout replaces os.Stdout with a buffer and returns it.
// https://rednafi.com/go/capture-console-output/
func captureStdout(f func()) string {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	f() // run the function that writes to stdout

	_ = w.Close()
	var buf bytes.Buffer
	_, _ = io.Copy(&buf, r)
	os.Stdout = old

	return buf.String()
}

// func Say(aline []rune) bool

func Test_1961(t *testing.T) {

	output := captureStdout(func() {
		Say(make([]rune, 0))
	})

	expected := "\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func Test_1962(t *testing.T) {

	vas := make([]rune, 1)
	vas[0] = '\000'

	output := captureStdout(func() {
		Say(vas)
	})

	expected := ""
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// func SayBool(b bool) bool

func Test_1963(t *testing.T) {

	output := captureStdout(func() {
		SayBool(true)
	})

	expected := "1\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func Test_1964(t *testing.T) {

	output := captureStdout(func() {
		SayBool(false)
	})

	expected := "0\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// func SayFloat32(f float32) bool

func Test_1965(t *testing.T) {

	output := captureStdout(func() {
		SayFloat32(float32(0.0))
	})

	expected := "0\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// func SayFloat64(d float64) bool

func Test_1966(t *testing.T) {

	output := captureStdout(func() {
		SayFloat64(float64(0.0))
	})

	expected := "0\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// func SayInt64(n int64) bool

func Test_1967(t *testing.T) {

	output := captureStdout(func() {
		SayInt64(int64(0))
	})

	expected := "0\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// func SayRexx(line *Rexx) bool

func Test_1968(t *testing.T) {

	// -- say null -- different from NetRexx?

	output := captureStdout(func() {
		SayRexx(nil)
	})

	expected := "\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func Test_1969(t *testing.T) {

	output := captureStdout(func() {
		SayRexx(RxFromRune('a'))
	})

	expected := "a\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func Test_1970(t *testing.T) {

	// --

	output := captureStdout(func() {
		SayRexx(rxfromclone(rxfromempty()))
	})

	expected := "\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func Test_1971(t *testing.T) {

	// --

	output := captureStdout(func() {
		SayRexx(RxFromRune('a'))
	})

	expected := "a\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func Test_1972(t *testing.T) {

	// --

	output := captureStdout(func() {
		SayRexx(rxfromclone(rxfromempty()))
	})

	expected := "\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}

}

// func SayRune(c rune) bool

func Test_1973(t *testing.T) {

	// --

	output := captureStdout(func() {
		SayRune('a')
	})

	expected := "a\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// func SayString(str string) bool

func Test_1974(t *testing.T) {

	// --

	output := captureStdout(func() {
		SayString("a")
	})

	expected := "a\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

func Test_1975(t *testing.T) {

	// --

	output := captureStdout(func() {
		Say(nil)
	})

	expected := "\n"
	if output != expected {
		t.Errorf("Expected %q, got %q", expected, output)
	}
}

// func RxNode(argleaf *Rexx) (rcvr *RexxNode)

func Test_1976(t *testing.T) {

	MyNode := RxNode(RxFromString("MyNode"))

	if MyNode.Leaf.ToString() != MyNode.InitLeaf.ToString() {
		t.Errorf("got %v, wanted %v", MyNode.Leaf.ToString(), MyNode.InitLeaf.ToString())
	}

}

// func (rcvr *Rexx) Clear()

func Test_1977(t *testing.T) {

	// New stem named drawer
	drawer := RxFromString("no drawer with that label in this file cabinet")

	//  add node labeled drawer A
	value, err := drawer.GetNode(RxFromRune('A'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer A Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('A'))
		want := "Drawer A Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer B
	value, err = drawer.GetNode(RxFromRune('B'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer B Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('B'))
		want := "Drawer B Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer C
	value, err = drawer.GetNode(RxFromRune('C'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer C Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('C'))
		want := "Drawer C Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer D
	value, err = drawer.GetNode(RxFromRune('D'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Start placing our files in this one.")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('D'))
		want := "Start placing our files in this one."

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}
	}

	// keep reference for Drawer D - we will put something here ...
	folder := value.Leaf

	// add folder X to drawer D
	value, err = folder.GetNode(RxFromRune('X'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("a folder labeled X")

		// call again with the new key to test that value was set
		value, err = folder.GetNode(RxFromRune('X'))
		want := "a folder labeled X"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	// add folder Y to drawer D
	value, err = folder.GetNode(RxFromRune('Y'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("a folder labeled Y")

		// call again with the new key to test that value was set
		value, err = folder.GetNode(RxFromRune('Y'))
		want := "a folder labeled Y"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	// add folder Z to drawer D
	value, err = folder.GetNode(RxFromRune('Z'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("a folder labeled Z")

		// call again with the new key to test that value was set
		value, err = folder.GetNode(RxFromRune('Z'))
		want := "a folder labeled Z"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	// keep reference for folder Z in Drawer D for the Clear() test below
	folder_Z_in_Drawer_D := value.Leaf

	// func (rcvr *Rexx) Clear()

	folder_Z_in_Drawer_D.Clear()

	if !folder_Z_in_Drawer_D.IsEmpty() {
		t.Errorf("got %v, wanted %v", !folder_Z_in_Drawer_D.IsEmpty(), true)
	}

	got := drawer.ContainsKey(RxFromString("E"))
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = drawer.ContainsKey(RxFromString("C"))
	want = true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = folder.ContainsValue(RxFromString("a folder labeled X"))
	want = true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = drawer.ContainsValue(RxFromString("No Drawer Here"))
	want = false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = rxfromempty().ContainsValue(RxFromInt32(1))
	want = false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) ContainsKey(key *Rexx) bool

// func (rcvr *Rexx) ContainsValue(value *Rexx) bool

// func (rcvr *Rexx) CompareTo(i2 string) (int32, error)

func Test_1978(t *testing.T) {

	// --

	value, err := rxfromclone(RxFromBool(true)).CompareTo("")
	want := int32(1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1979(t *testing.T) {

	// -- say Rexx("").compareto(Object "")

	value, err := rxfromempty().CompareTo("")
	want := int32(0)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1980(t *testing.T) {

	// -- say "1".compareto(Object "0")

	value, err := RxFromString("1").CompareTo("0")
	want := int32(1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1981(t *testing.T) {

	// -- say "1".compareto(Object "1")

	value, err := RxFromString("1").CompareTo("1")
	want := int32(0)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1982(t *testing.T) {

	// --

	value, err := RxFromString("1").CompareTo("2")
	want := int32(-1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1983(t *testing.T) {

	// --

	value, err := RxFromString("1").CompareTo("a.+E")
	want := int32(-1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1984(t *testing.T) {

	// --

	value, err := RxFromString("a").CompareTo("0")
	want := int32(1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1985(t *testing.T) {

	// --

	value, err := RxFromString("a").CompareTo("")
	want := int32(1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1986(t *testing.T) {

	// --

	value, err := rxfromempty().CompareTo("0")
	want := int32(0)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1987(t *testing.T) {

	// --

	value, err := rxfromclone(nil).CompareTo("0")
	want := int32(-1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1988(t *testing.T) {

	// --

	value, err := RxFromString("9.47109959E+230565093").CompareTo(RxFromString("73354723.2").ToString())
	want := int32(1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1989(t *testing.T) {

	// --

	MinExp = 9

	value, err := RxFromString("9E+999999999").CompareTo("+1.23456789012345E-0")
	want := int32(1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	MinExp = -999999999 // reset

}

func Test_1990(t *testing.T) {

	// --

	value, err := RxFromString("\x00").CompareTo("\x00")
	want := int32(0)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_1991(t *testing.T) {

	// --

	MaxExp = 9

	_, err := RxFromString("9E+999999999").CompareTo(RxFromString("+1.23456789012345E-0").ToString())

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999 // reset

}

// func (rcvr *Rexx) CopyIndexed(r *Rexx) *Rexx

func Test_1992(t *testing.T) {

	// another stem named foo
	foo := RxFromString("def")

	//  add node with key a
	value, err := foo.GetNode(RxFromRune('a'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("1")

		// call again with the new key to test that value was set
		value, err = foo.GetNode(RxFromRune('a'))
		want := "1"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node with key b
	value, err = foo.GetNode(RxFromRune('b'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("2")

		// call again with the new key to test that value was set
		value, err = foo.GetNode(RxFromRune('b'))
		want := "2"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	// another stem named bar
	bar := RxFromString("ghi")

	//  add node with key b - same key as in used in foo - will be combined in merge
	value, err = bar.GetNode(RxFromRune('b'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("b")

		// call again with the new key to test that value was set
		value, err = bar.GetNode(RxFromRune('b'))
		want := "b"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node with key c
	value, err = bar.GetNode(RxFromRune('c'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("c")

		// call again with the new key to test that value was set
		value, err = bar.GetNode(RxFromRune('c'))
		want := "c"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	// merge bar into foo
	merged := foo.CopyIndexed(bar)

	// get the keys after the merge of bar with foo
	all_keys := merged.Keys()

	// loop through keys
	for _, key := range all_keys {

		value, err = merged.GetNode(key)

		if err == nil {

			if key.ToString() == "a" {

				got := value.Leaf.ToString()
				want := "1"

				if got != want {
					t.Errorf("got %v, wanted %v", got, want)
				}

			}

			// The value of b.Leaf is now 'b' replacing the 2 that was set in foo
			if key.ToString() == "b" {

				got := value.Leaf.ToString()
				want := "b"

				if got != want {
					t.Errorf("got %v, wanted %v", got, want)
				}

			}

			if key.ToString() == "c" {

				got := value.Leaf.ToString()
				want := "c"

				if got != want {
					t.Errorf("got %v, wanted %v", got, want)
				}
			}

		} else {
			t.Errorf("Should not have been an error with merge of bar with foo")
		}

	}

	//  add another node with key d to merged but do not set Leaf value
	value, err = merged.GetNode(RxFromRune('d'))

	// Notes: no Leaf value for 'd' will trigger code block below
	// code block -- if rxnode.Leaf == rxnode.InitLeaf {
	// code block -- 		continue
	// code block -- }
	// when rxfromempty().CopyIndexed(foo) is called
	rxfromempty().CopyIndexed(foo)

}

func Test_1993(t *testing.T) {

	// ---

	empty_stem := ToRxFromString("stem")
	empty_merge := ToRxFromString("merge")

	merged := empty_stem.CopyIndexed(empty_merge)

	got := merged.Size(0)
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1994(t *testing.T) {

	// ---

	too := RxFromRune('E')

	value, err := too.GetNode(RxFromRune('+'))

	if err == nil {
		value.Leaf = RxFromRune('+')
	} else {
		t.Errorf("Did not expect an error")
	}

	value, err = too.GetNode(RxFromRune('-'))

	if err == nil {
		value.Leaf = nil
	} else {
		t.Errorf("Did not expect an error")
	}

	merged := RxFromRune('0').CopyIndexed(too)

	got := merged.Size(0)
	// nil Leaf not counted
	want := int32(1)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) Equals(rhs *Rexx) bool

func Test_1995(t *testing.T) {

	// --

	// nil as argument will return false
	got := RxFromBool(false).Equals(nil)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1996(t *testing.T) {

	// --

	got := RxFromBool(true).Equals(RxFromBool(false))
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1997(t *testing.T) {

	// --

	got := RxFromBool(false).Equals(RxFromBool(true))
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1998(t *testing.T) {

	// --

	got := RxFromBool(true).Equals(RxFromBool(true))
	want := true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_1999(t *testing.T) {

	// -- say Rexx("A").Equals(null)

	got := RxFromString("A").Equals(rxfromempty())
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) Exists(key *Rexx) *Rexx

func Test_2000(t *testing.T) {

	// --

	// New stem named drawer
	drawer := RxFromString("no drawer with that label in this file cabinet")

	//  add node labeled drawer A
	value, err := drawer.GetNode(RxFromRune('A'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer A Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('A'))
		want := "Drawer A Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer B
	value, err = drawer.GetNode(RxFromRune('B'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer B Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('B'))
		want := "Drawer B Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer C
	value, err = drawer.GetNode(RxFromRune('C'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer C Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('C'))
		want := "Drawer C Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer D
	value, err = drawer.GetNode(RxFromRune('D'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Start placing our files in this one.")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('D'))
		want := "Start placing our files in this one."

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}
	}

	got := drawer.Exists(RxFromRune('C')).ToString()
	want := "1"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = drawer.Exists(RxFromRune('Q')).ToString()
	want = "0"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	// Notes: case senstive, there is a drawer with 'C' but not 'c'
	got = drawer.Exists(RxFromRune('c')).ToString()
	want = "0"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) Get(key *Rexx) *Rexx

func Test_2001(t *testing.T) {

	// --

	// New stem named drawer
	drawer := RxFromString("no drawer with that label in this file cabinet")

	//  add node labeled drawer A
	value, err := drawer.GetNode(RxFromRune('A'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer A Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('A'))
		want := "Drawer A Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer B
	value, err = drawer.GetNode(RxFromRune('B'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer B Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('B'))
		want := "Drawer B Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer C
	value, err = drawer.GetNode(RxFromRune('C'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer C Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('C'))
		want := "Drawer C Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer D
	value, err = drawer.GetNode(RxFromRune('D'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Start placing our files in this one.")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('D'))
		want := "Start placing our files in this one."

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}
	}

	result, err := drawer.Get(RxFromString("B"))
	want := "Drawer B Empty"

	if err != nil {
		t.Errorf("Expecting a result.")
	} else {
		got := result.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	// Notes: this func can return nil
	result, err = drawer.Get(RxFromString("Bad Key"))

	if result != nil {
		t.Errorf("Expecting a nil result.")
	}

	// Notes: this func can return nil
	result, err = drawer.Get(nil)

	if result != nil {
		t.Errorf("Expecting a nil result.")
	}

}

// func (rcvr *Rexx) GetNode(key *Rexx) *RexxNode

func Test_2002(t *testing.T) {

	// --

	// New stem named drawer
	drawer := RxFromString("no drawer with that label in this file cabinet")

	//  add node labeled drawer A
	value, err := drawer.GetNode(RxFromRune('A'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer A Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('A'))
		want := "Drawer A Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer B
	value, err = drawer.GetNode(RxFromRune('B'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer B Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('B'))
		want := "Drawer B Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer C
	value, err = drawer.GetNode(RxFromRune('C'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer C Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('C'))
		want := "Drawer C Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer D
	value, err = drawer.GetNode(RxFromRune('D'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Start placing our files in this one.")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('D'))
		want := "Start placing our files in this one."

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}
	}

	// Error: ArgumentNullException : cannot use nil as type Rexx in argument
	value, err = drawer.GetNode(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	//  add node labeled drawer E
	value, err = drawer.GetNode(RxFromRune('E'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("New Drawer")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('E'))
		want := "New Drawer"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}
	}

	// Size should be 5 the nil call not added
	got := drawer.Size(0)
	want := int32(5)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) IsEmpty() bool

func Test_2003(t *testing.T) {

	// --

	// New stem named drawer
	drawer := RxFromString("no drawer with that label in this file cabinet")

	//  add node labeled drawer A
	value, err := drawer.GetNode(RxFromRune('A'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer A Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('A'))
		want := "Drawer A Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer B
	value, err = drawer.GetNode(RxFromRune('B'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer B Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('B'))
		want := "Drawer B Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer C
	value, err = drawer.GetNode(RxFromRune('C'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer C Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('C'))
		want := "Drawer C Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	got := drawer.IsEmpty()
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) IsIndexed() *Rexx

func Test_2004(t *testing.T) {

	// --

	// New stem named drawer
	drawer := RxFromString("no drawer with that label in this file cabinet")

	//  add node labeled drawer A
	value, err := drawer.GetNode(RxFromRune('A'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer A Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('A'))
		want := "Drawer A Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer B
	value, err = drawer.GetNode(RxFromRune('B'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer B Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('B'))
		want := "Drawer B Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer C
	value, err = drawer.GetNode(RxFromRune('C'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer C Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('C'))
		want := "Drawer C Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	got := drawer.IsIndexed().ToString()
	want := "1"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2005(t *testing.T) {

	got := rxfromempty().IsIndexed().ToString()
	want := "0"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) Keys() []*Rexx

func Test_2006(t *testing.T) {

	got := len(rxfromempty().Keys())
	want := 0

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) Put(key *Rexx, value *Rexx) *Rexx

func Test_2007(t *testing.T) {

	// --

	// New stem named drawer
	drawer := RxFromString("no drawer with that label in this file cabinet")

	//  add node labeled drawer A
	value, err := drawer.GetNode(RxFromRune('A'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer A Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('A'))
		want := "Drawer A Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer B
	value, err = drawer.GetNode(RxFromRune('B'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer B Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('B'))
		want := "Drawer B Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer C
	value, err = drawer.GetNode(RxFromRune('C'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer C Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('C'))
		want := "Drawer C Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	// Error: ArgumentNullException : cannot use nil as type Rexx in argument
	value, err = drawer.GetNode(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	//  try to add Leaf to drawer Q
	_, err = drawer.Put(RxFromString("Q"), RxFromString("Not a Drawer"))

	// Error: BadArgumentException : key not found
	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	//  try to use a nil key
	_, err = drawer.Put(nil, RxFromString("Not a Drawer"))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	// Change a key's value
	_, err = drawer.Put(RxFromString("A"), RxFromString("Change A \"Drawer A Empty\" to New Leaf Value"))

	if err == nil {

		value, err = drawer.GetNode(RxFromRune('A'))
		want := "Change A \"Drawer A Empty\" to New Leaf Value"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	} else {
		t.Errorf("Should not have been an error with merge of bar with foo")
	}

}

// func (rcvr *Rexx) PutAll(t *Rexx) (*Rexx, error)

func Test_2008(t *testing.T) {

	// --

	// New stem named drawer
	drawer := RxFromString("no drawer with that label in this file cabinet")

	//  add node labeled drawer A
	value, err := drawer.GetNode(RxFromRune('A'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer A Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('A'))
		want := "Drawer A Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer B
	value, err = drawer.GetNode(RxFromRune('B'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer B Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('B'))
		want := "Drawer B Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer C
	value, err = drawer.GetNode(RxFromRune('C'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer C Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('C'))
		want := "Drawer C Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	// Error: ArgumentNullException : cannot use nil as type Rexx in argument
	_, err = drawer.PutAll(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	// Our new empty stem with default value of Cleaned
	our_new_empty_stem := RxFromString("Cleaned")

	result, err := our_new_empty_stem.PutAll(drawer)

	if err == nil {

		value, err = result.GetNode(RxFromRune('A'))
		want := "Drawer A Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

		value, err = result.GetNode(RxFromString("return stem's default value with bad key"))
		want = "Cleaned"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

}

// func (rcvr *Rexx) Remove(key *Rexx) (*Rexx, error)

func Test_2009(t *testing.T) {

	// --

	// New stem named drawer
	drawer := RxFromString("no drawer with that label in this file cabinet")

	//  add node labeled drawer A
	value, err := drawer.GetNode(RxFromRune('A'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer A Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('A'))
		want := "Drawer A Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer B
	value, err = drawer.GetNode(RxFromRune('B'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer B Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('B'))
		want := "Drawer B Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer C
	value, err = drawer.GetNode(RxFromRune('C'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer C Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('C'))
		want := "Drawer C Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	// key cannot be nil
	_, err = drawer.Remove(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	// BadArgumentException : key not found
	_, err = drawer.Remove(RxFromString("It's not here."))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	_, err = drawer.Remove(RxFromRune('B'))

	if err == nil {

		// check if key is gone
		value, err = drawer.GetNode(RxFromRune('B'))
		// Should return stem default
		want := "no drawer with that label in this file cabinet"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	} else {
		t.Errorf("Should not have been an error.")
	}

}

// func (rcvr *Rexx) Size(check int32) int32

func Test_2010(t *testing.T) {

	// --

	// stem named drawer
	drawer := RxFromString("Nothing here.")

	// create a node with key F
	// do not set  value.Leaf
	drawer.GetNode(RxFromRune('F'))

	got := drawer.Size(0)
	want := int32(1)


	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	// shortcut to test code blocks after some bug fixes
	value, _ := drawer.GetNode(RxFromRune('F'))
	value.Leaf = value.InitLeaf //Shortcut
	drawer.Size(0)
	rxfromempty().CopyIndexed(drawer) //Shortcut

}

func Test_2011(t *testing.T) {

	// --

	// New stem named drawer
	drawer := RxFromString("no drawer with that label in this file cabinet")

	//  add node labeled drawer A
	value, err := drawer.GetNode(RxFromRune('A'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer A Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('A'))
		want := "Drawer A Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer B
	value, err = drawer.GetNode(RxFromRune('B'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer B Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('B'))
		want := "Drawer B Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer C
	value, err = drawer.GetNode(RxFromRune('C'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer C Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('C'))
		want := "Drawer C Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer D
	value, err = drawer.GetNode(RxFromRune('D'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Start placing our files in this one.")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('D'))
		want := "Start placing our files in this one."

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}
	}

	// keep reference for Drawer D - we will put something here ...
	folder := value.Leaf

	// add folder X to drawer D
	value, err = folder.GetNode(RxFromRune('X'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("a folder labeled X")

		// call again with the new key to test that value was set
		value, err = folder.GetNode(RxFromRune('X'))
		want := "a folder labeled X"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	// add folder Y to drawer D
	value, err = folder.GetNode(RxFromRune('Y'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("a folder labeled Y")

		// call again with the new key to test that value was set
		value, err = folder.GetNode(RxFromRune('Y'))
		want := "a folder labeled Y"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	// add folder Z to drawer D
	value, err = folder.GetNode(RxFromRune('Z'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("a folder labeled Z")

		// call again with the new key to test that value was set
		value, err = folder.GetNode(RxFromRune('Z'))
		want := "a folder labeled Z"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	got := rxfromempty().Size(0)
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = rxfromclone(RxFromBool(false)).Size(0)
	want = int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = folder.Size(0)
	want = int32(3)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = drawer.Size(0)
	want = int32(4)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = folder.Size(1)
	want = int32(1)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func (rcvr *Rexx) TestNode(key *Rexx) bool

func Test_2012(t *testing.T) {

	// --

	// New stem named drawer
	drawer := RxFromString("no drawer with that label in this file cabinet")

	//  add node labeled drawer A
	value, err := drawer.GetNode(RxFromRune('A'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer A Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('A'))
		want := "Drawer A Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer B
	value, err = drawer.GetNode(RxFromRune('B'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer B Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('B'))
		want := "Drawer B Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer C
	value, err = drawer.GetNode(RxFromRune('C'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer C Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('C'))
		want := "Drawer C Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer D
	value, err = drawer.GetNode(RxFromRune('D'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Start placing our files in this one.")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('D'))
		want := "Start placing our files in this one."

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}
	}

	got := drawer.TestNode(RxFromString("IT_IS_NOT_HERE"))
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = drawer.TestNode(RxFromRune('A'))
	want = true

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = drawer.TestNode(RxFromRune('a'))
	want = false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = drawer.TestNode(RxFromRune('Q'))
	want = false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	got = rxfromempty().TestNode(nil)
	want = false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	// Set InitLeaf and Leaf to nil
	value, err = drawer.GetNode(RxFromRune('E'))

	if err == nil {
		value.InitLeaf = nil
		value.Leaf = nil
	}

	got = drawer.TestNode(RxFromRune('Q'))
	want = false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2013(t *testing.T) {

	// New stem named drawer
	drawer := RxFromString("no drawer with that label in this file cabinet")

	//  add node labeled drawer A
	value, err := drawer.GetNode(RxFromRune('A'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer A Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('A'))
		want := "Drawer A Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer B
	value, err = drawer.GetNode(RxFromRune('B'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer B Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('B'))
		want := "Drawer B Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer C
	value, err = drawer.GetNode(RxFromRune('C'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Drawer C Empty")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('C'))
		want := "Drawer C Empty"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	}

	//  add node labeled drawer D
	value, err = drawer.GetNode(RxFromRune('D'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Start placing our files in this one.")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('D'))
		want := "Start placing our files in this one."

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}
	}

	//  add node labeled drawer E
	value, err = drawer.GetNode(RxFromRune('E'))

	if err == nil {
		// set Leaf value
		value.Leaf = RxFromString("Junk Drawer")

		// call again with the new key to test that value was set
		value, err = drawer.GetNode(RxFromRune('E'))
		want := "Junk Drawer"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}
	}

	_, err = drawer.Remove(RxFromRune('B'))

	if err == nil {

		// check if key is gone
		value, err = drawer.GetNode(RxFromRune('B'))
		// Should return stem default
		want := "no drawer with that label in this file cabinet"

		if err != nil {
			t.Errorf("Expecting a value.")
		} else {
			got := value.Leaf.ToString()
			if got != want {
				t.Errorf("got %v, wanted %v", got, want)
			}
		}

	} else {
		t.Errorf("Should not have been an error.")
	}

	// size should be 5 after removing B - remove only resets the Leaf
	got := drawer.Size(0)
	want := int32(5)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

	//  Change E's InitLeaf to match Leaf's
	value, err = drawer.GetNode(RxFromRune('E'))

	if err == nil {
		// set InitLeaf to match Leaf's
		value.InitLeaf = value.Leaf
	}

	gota := drawer.TestNode(RxFromRune('E'))
	wanta := false

	if gota != wanta {
		t.Errorf("got %v, wanted %v", gota, wanta)
	}

	//  Change E's InitLeaf and Leaf to nil
	value, err = drawer.GetNode(RxFromRune('E'))

	if err == nil {
		// set InitLeaf and Leaf to nil
		value.InitLeaf = nil
		value.Leaf = nil
	}

	gota = drawer.TestNode(RxFromRune('E'))
	wanta = false

	if gota != wanta {
		t.Errorf("got %v, wanted %v", gota, wanta)
	}

}

// func rxparse() (rcvr *RexxParse)

func Test_2014(t *testing.T) {

	// --

	rxparse := rxparse()

	if rxparse == nil {
		t.Errorf("Call should not be nil")
	}

}

// func Parse(obj *Rexx, list []rune, vars []*Rexx) error

func Test_2015(t *testing.T) {

	// -- say parse 'This is a sentence.' v1 v2 v3
	// -- spilt subject by spaces

	subject := RxFromString("This is a sentence.")
	template := []rune{1, 10, 3, 0, 1, 2, 0}
	results := make([]*Rexx, 3)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "This"
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
		want = "is"
		if results[1].ToString() != want {
			t.Errorf("got %v, wanted %v", results[1].ToString(), want)
		}
		want = "a sentence."
		if results[2].ToString() != want {
			t.Errorf("got %v, wanted %v", results[2].ToString(), want)
		}
	}

}

func Test_2016(t *testing.T) {

	// -- say parse "jan,feb,mar,apr,may,jun" v1 ',' v2 ',' v3 ',' v4 ','v5
	// -- spilt subject by commas

	subject := RxFromString("jan,feb,mar,apr,may,jun")
	template := []rune{2, 1, 44, 10, 1, 0, 2, 1, 44, 10, 1, 1, 2, 1, 44, 10, 1, 2, 2, 1, 44, 10, 1, 3, 1, 10, 1, 4, 0}
	results := make([]*Rexx, 5)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "jan"
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
		want = "feb"
		if results[1].ToString() != want {
			t.Errorf("got %v, wanted %v", results[1].ToString(), want)
		}
		want = "mar"
		if results[2].ToString() != want {
			t.Errorf("got %v, wanted %v", results[2].ToString(), want)
		}
		want = "apr"
		if results[3].ToString() != want {
			t.Errorf("got %v, wanted %v", results[3].ToString(), want)
		}
		want = "may,jun"
		if results[4].ToString() != want {
			t.Errorf("got %v, wanted %v", results[4].ToString(), want)
		}
	}

}

func Test_2017(t *testing.T) {

	// -- say parse 'Flying pigs have wings' x1 5 x2
	// -- spilt subject at position number

	subject := RxFromString("Flying pigs have wings")
	template := []rune{3, 1, 5, 10, 1, 0, 1, 10, 1, 1, 0}
	results := make([]*Rexx, 2)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "Flyi"
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
		want = "ng pigs have wings"
		if results[1].ToString() != want {
			t.Errorf("got %v, wanted %v", results[1].ToString(), want)
		}
	}

}

func Test_2018(t *testing.T) {

	// -- say parse "This is  the text which, I think, is scanned." ... word4 .
	// -- spilt subject using periods as place holders

	subject := RxFromString("This is  the text which, I think, is scanned.")
	template := []rune{1, 10, 5, 0, 0, 0, 1, 0, 0}
	results := make([]*Rexx, 2)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "which, I think, is scanned."
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
		want = "text"
		if results[1].ToString() != want {
			t.Errorf("got %v, wanted %v", results[1].ToString(), want)
		}
	}

}

func Test_2019(t *testing.T) {

	// -- say parse "" w1 ' ' w2 ' ' w3 ' ' w4 ','
	// -- spilt empty subject by pattern ' ' and ','

	subject := RxFromString("")
	template := []rune{2, 1, 32, 10, 1, 0, 2, 1, 32, 10, 1, 1, 2, 1, 32, 10, 1, 2, 2, 1, 44, 10, 1, 3, 0}
	results := make([]*Rexx, 4)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := ""
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
		want = ""
		if results[1].ToString() != want {
			t.Errorf("got %v, wanted %v", results[1].ToString(), want)
		}
		want = ""
		if results[2].ToString() != want {
			t.Errorf("got %v, wanted %v", results[2].ToString(), want)
		}
		want = ""
		if results[3].ToString() != want {
			t.Errorf("got %v, wanted %v", results[3].ToString(), want)
		}
	}

}

func Test_2020(t *testing.T) {

	// -- say parse "L/look for/1 10" verb 2 delim +1 string (delim) rest
	// -- spilt subject using variable patterns

	subject := RxFromString("L/look for/1 10")
	template := []rune{3, 1, 2, 10, 1, 0, 4, 1, 1, 10, 1, 1, 6, 1, 10, 1, 2, 1, 10, 1, 3, 0}
	results := make([]*Rexx, 4)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "L"
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
		want = "/"
		if results[1].ToString() != want {
			t.Errorf("got %v, wanted %v", results[1].ToString(), want)
		}
		want = "look for"
		if results[2].ToString() != want {
			t.Errorf("got %v, wanted %v", results[2].ToString(), want)
		}
		want = "1 10"
		if results[3].ToString() != want {
			t.Errorf("got %v, wanted %v", results[3].ToString(), want)
		}
	}

}

func Test_2021(t *testing.T) {

	// -- say parse "This is  the text which, I think, is scanned." "," -1 x +1

	subject := RxFromString("This is  the text which, I think, is scanned.")
	template := []rune{5, 1, 1, 4, 1, 1, 10, 1, 0, 0}
	results := make([]*Rexx, 1)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "T"
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
	}

}

func Test_2022(t *testing.T) {

	// -- say parse "This is  the text which, I think, is scanned." "," 46 x +15

	subject := RxFromString("This is  the text which, I think, is scanned.")
	template := []rune{2, 1, 44, 3, 1, 46, 4, 1, 15, 10, 1, 0, 0}
	results := make([]*Rexx, 1)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := ""
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
	}

}

func Test_2023(t *testing.T) {

	// -- say parse ask a b c d e

	oldStdin := os.Stdin
	defer func() { os.Stdin = oldStdin }()

	input := "space seperated command line arguments\n"
	r, w, _ := os.Pipe()
	os.Stdin = r

	go func() {
		w.Write([]byte(input))
		w.Close()
	}()

	arg, _ := Ask()

	if arg.ToString() != "space seperated command line arguments" {
		t.Errorf("Expected input, got %s", arg.ToString())
	}

	template := []rune{1, 10, 5, 0, 1, 2, 3, 4, 0}
	results := make([]*Rexx, 5)

	err := Parse(arg, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "space"
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
		want = "seperated"
		if results[1].ToString() != want {
			t.Errorf("got %v, wanted %v", results[1].ToString(), want)
		}
		want = "command"
		if results[2].ToString() != want {
			t.Errorf("got %v, wanted %v", results[2].ToString(), want)
		}
		want = "line"
		if results[3].ToString() != want {
			t.Errorf("got %v, wanted %v", results[3].ToString(), want)
		}
		want = "arguments"
		if results[4].ToString() != want {
			t.Errorf("got %v, wanted %v", results[4].ToString(), want)
		}
	}

}

func Test_2024(t *testing.T) {

	// -- say parse "abcde" 0 a +0 b -0 "0" c "۱" 0 d -0  e

	subject := RxFromString("abcde")
	template := []rune{3, 0, 4, 0, 10, 1, 0, 5, 0, 10, 1, 1, 2, 1, 48, 2, 1, 1777, 10, 1, 2, 3, 0, 5, 0, 10, 1, 3, 1, 10, 1, 4, 0}
	results := make([]*Rexx, 5)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "abcde"
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
		want = "abcde"
		if results[1].ToString() != want {
			t.Errorf("got %v, wanted %v", results[1].ToString(), want)
		}
		want = ""
		if results[2].ToString() != want {
			t.Errorf("got %v, wanted %v", results[2].ToString(), want)
		}
		want = "abcde"
		if results[3].ToString() != want {
			t.Errorf("got %v, wanted %v", results[3].ToString(), want)
		}
		want = "abcde"
		if results[4].ToString() != want {
			t.Errorf("got %v, wanted %v", results[4].ToString(), want)
		}
	}

}

func Test_2025(t *testing.T) {

	// -- say parse "\\A" "\\" a

	subject := RxFromString("\\\" \"\\")
	template := []rune{2, 1, 92, 1, 10, 1, 0, 0}
	results := make([]*Rexx, 1)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "\" \"\\"
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
	}

}

func Test_2026(t *testing.T) {

	// -- say parse "12345 6789 0123 0555" -5 V1 =33 V2 +500 V3 -989 V4 .

	subject := RxFromString("12345 6789 0123 0555")
	template := []rune{5, 1, 5, 3, 1, 33, 10, 1, 0, 4, 2, 1, 244, 10, 1, 1, 5, 2, 3, 221, 10, 1, 2, 1, 10, 2, 3, 4, 0}
	results := make([]*Rexx, 5)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "12345 6789 0123 0555"
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
		want = ""
		if results[1].ToString() != want {
			t.Errorf("got %v, wanted %v", results[1].ToString(), want)
		}
		want = ""
		if results[2].ToString() != want {
			t.Errorf("got %v, wanted %v", results[2].ToString(), want)
		}
		want = "12345"
		if results[3].ToString() != want {
			t.Errorf("got %v, wanted %v", results[3].ToString(), want)
		}
	}

}

func Test_2027(t *testing.T) {

	// -- say parse "12345 6789 0123 0555" -5 V1 y1 V2 +500 V3 y2 V4 .

	subject := RxFromString("12345 6789 0123 0555")
	template := []rune{5, 1, 5, 4, 2, 1, 244, 10, 3, 0, 1, 2, 1, 10, 4, 3, 4, 5, 6, 0}
	results := make([]*Rexx, 7)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "12345"
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
		want = "6789"
		if results[1].ToString() != want {
			t.Errorf("got %v, wanted %v", results[1].ToString(), want)
		}
		want = "0123 0555"
		if results[2].ToString() != want {
			t.Errorf("got %v, wanted %v", results[2].ToString(), want)
		}
		want = ""
		if results[3].ToString() != want {
			t.Errorf("got %v, wanted %v", results[3].ToString(), want)
		}
		want = ""
		if results[4].ToString() != want {
			t.Errorf("got %v, wanted %v", results[4].ToString(), want)
		}
		want = ""
		if results[5].ToString() != want {
			t.Errorf("got %v, wanted %v", results[5].ToString(), want)
		}
	}

}

func Test_2028(t *testing.T) {

	// -- say parse "12345 6789 0123 0555" 5 10 15 V1

	subject := RxFromString("12345 6789 0123 0555")
	template := []rune{3, 1, 5, 3, 1, 10, 3, 1, 15, 1, 10, 1, 0, 0}
	results := make([]*Rexx, 1)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "3 0555"
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
	}

}

func Test_2029(t *testing.T) {

	// -- say parse "12345 6789 0123 0555" +1 -1 1+1 1-1 0 -0 =0 V1

	subject := RxFromString("12345 6789 0123 0555")
	template := []rune{4, 1, 1, 5, 1, 1, 3, 1, 1, 4, 1, 1, 3, 1, 1, 5, 1, 1, 3, 0, 5, 0, 3, 0, 1, 10, 1, 0, 0}
	results := make([]*Rexx, 1)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "12345 6789 0123 0555"
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
	}

}

func Test_2030(t *testing.T) {

	// -- say c = ','; parse 'To be, or not to be?' V1 (c) V2 V3 V4

	subject := RxFromString("To be, or not to be?")
	template := []rune{6, 1, 10, 1, 0, 1, 10, 3, 2, 3, 4, 0}
	results := make([]*Rexx, 5)

	// pattern may be specified as a variable
	results[1] = RxFromRune(',')

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "To be"
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
		// results[1] injected above and used in parse
		want = "or"
		if results[2].ToString() != want {
			t.Errorf("got %v, wanted %v", results[2].ToString(), want)
		}
		want = "not"
		if results[3].ToString() != want {
			t.Errorf("got %v, wanted %v", results[3].ToString(), want)
		}
		want = "to be?"
		if results[4].ToString() != want {
			t.Errorf("got %v, wanted %v", results[4].ToString(), want)
		}
	}

}

func Test_2031(t *testing.T) {

	// -- say num1 = '\x05'; start= 5; length=5; data="Fly" || num1 || "ng pigs have wings" ; parse data x1 =(start) x2 +(length) x3

	num1, _ := RxFromInt8(5).D2c()
	subject := (RxFromString("Fly").OpCc(nil, num1)).OpCc(nil, RxFromString("ng pigs have wings"))
	template := []rune{7, 1, 10, 1, 0, 8, 3, 10, 1, 2, 1, 10, 1, 4, 0}
	results := make([]*Rexx, 5)

	// patterns may be specified as a variable
	results[1] = RxFromInt8(5)
	results[3] = RxFromInt8(5)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "Fly"
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
		// results[1] injected above and used in parse
		want = "ng pi"
		if results[2].ToString() != want {
			t.Errorf("got %v, wanted %v", results[2].ToString(), want)
		}
		// results[3] injected above and used in parse
		want = "gs have wings"
		if results[4].ToString() != want {
			t.Errorf("got %v, wanted %v", results[4].ToString(), want)
		}
	}

}

func Test_2032(t *testing.T) {

	// -- say start=-5; length=-5; data='Flying pigs have wings' ; parse data x1 =(start) x2 +(length) x3

	subject := RxFromString("Flying pigs have wings")
	template := []rune{7, 1, 10, 1, 0, 8, 3, 10, 1, 2, 1, 10, 1, 4, 0}
	results := make([]*Rexx, 5)

	// patterns may be specified as a variable
	results[1] = RxFromInt8(-5)
	results[3] = RxFromInt8(-5)

	err := Parse(subject, template, results)

	// triggers "if nextcol < MinCol || nextcol > MaxCol {"

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2033(t *testing.T) {

	// -- say opts = "print"; parse (' 'opts).upper ' PR' +1 prword ' '

	// subject is the same as above NetRexx lines
	subject := RxFromString(" PRINT")
	template := []rune{2, 3, 32, 80, 82, 4, 1, 1, 2, 1, 32, 10, 1, 0, 0}
	results := make([]*Rexx, 1)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "PRINT"
		if results[0].ToString() != want {
			t.Errorf("got %v, wanted %v", results[0].ToString(), want)
		}
	}

}

func Test_2034(t *testing.T) {

	// -- say num1 = '\x05'; start= "7.62939453E+28"; length="-000001.12345"; data="Fly" || num1 || "ng pigs have wings" ; parse data x1 =(start) x2 +(length) x3

	num1, _ := RxFromInt8(5).D2c()
	subject := (RxFromString("Fly").OpCc(nil, num1)).OpCc(nil, RxFromString("ng pigs have wings"))
	template := []rune{7, 1, 10, 1, 0, 8, 3, 10, 1, 2, 1, 10, 1, 4, 0}
	results := make([]*Rexx, 5)

	// patterns may be specified as a variable
	// triggers conversion overflow
	results[1] = RxFromString("7.62939453E+28")
	results[3] = RxFromString("-000001.12345")

	err := Parse(subject, template, results)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2035(t *testing.T) {

	// --

	err := Parse(nil, nil, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2036(t *testing.T) {

	// -- parse "template is null" a

	subject := RxFromString("template is null")
	results := make([]*Rexx, 1)

	err := Parse(subject, nil, results)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2037(t *testing.T) {

	// -- parse "template is null" a

	subject := RxFromString("template is null")
	template := []rune{1, 10, 1, 0, 0}
	// pass nil results

	err := Parse(subject, template, nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2038(t *testing.T) {

	// -- say test = 9; start=5; length=18; data='Flying pigs have wings' ; parse data -(test) x1 =(start)  x2 +(length)  -48 x3

	subject := RxFromString("Flying pigs have wings")
	template := []rune{9, 0, 7, 2, 10, 1, 1, 8, 4, 10, 1, 3, 5, 1, 48, 1, 10, 1, 5, 0}
	results := make([]*Rexx, 6)

	// patterns may be specified as a variable
	results[0] = RxFromInt32(9)
	results[2] = RxFromInt32(5)
	results[4] = RxFromInt32(18)

	err := Parse(subject, template, results)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		want := "Flyi"
		if results[1].ToString() != want {
			t.Errorf("got %v, wanted %v", results[1].ToString(), want)
		}
		// results[1] injected above and used in parse
		want = "ng pigs have wings"
		if results[3].ToString() != want {
			t.Errorf("got %v, wanted %v", results[3].ToString(), want)
		}
		// results[3] injected above and used in parse
		want = "Flying pigs have wings"
		if results[5].ToString() != want {
			t.Errorf("got %v, wanted %v", results[5].ToString(), want)
		}
	}

}

// func RxSet() (rcvr *RexxSet)

func Test_2039(t *testing.T) {

	new_set := RxSet()

	if new_set.Digits != DEFAULT_DIGITS {
		t.Errorf("got %v, wanted %v", new_set.Digits, DEFAULT_DIGITS)
	}

}

// func RxSetFromOld(old *RexxSet) (rcvr *RexxSet)

func Test_2040(t *testing.T) {

	new_set := RxSetFromOld(nil)

	if new_set.Digits != DEFAULT_DIGITS {
		t.Errorf("got %v, wanted %v", new_set.Digits, DEFAULT_DIGITS)
	}

}

func Test_2041(t *testing.T) {

	old_set := RxSet()

	new_set := RxSetFromOld(old_set)

	if new_set.Digits != old_set.Digits {
		t.Errorf("got %v, wanted %v", new_set.Digits, old_set.Digits)
	}

}

// func RxSetWithDigit(newdigits int32) (rcvr *RexxSet)

func Test_2042(t *testing.T) {

	new_set := RxSetWithDigit(19)

	if new_set.Digits != 19 {
		t.Errorf("got %v, wanted %v", new_set.Digits, 19)
	}

}

// func RxSetWithDigitandForm(newdigits int32, newform int8) (rcvr *RexxSet)

func Test_2043(t *testing.T) {

	new_set := RxSetWithDigitandForm(17, int8(1))

	if new_set.Digits != 17 {
		t.Errorf("got %v, wanted %v", new_set.Digits, 17)
	}

}

// func (rcvr *RexxSet) FormWord() *Rexx

func Test_2044(t *testing.T) {

	new_set := RxSetWithDigitandForm(17, int8(1))

	if new_set.FormWord().ToString() != "engineering" {
		t.Errorf("got %v, wanted %v", new_set.FormWord().ToString(), "engineering")
	}

}

func Test_2045(t *testing.T) {

	new_set := RxSetWithDigitandForm(17, int8(0))

	if new_set.FormWord().ToString() != "scientific" {
		t.Errorf("got %v, wanted %v", new_set.FormWord().ToString(), "scientific")
	}

}

// func (rcvr *RexxSet) SetDigits(d *Rexx) error

func Test_2046(t *testing.T) {

	new_set := RxSet()

	new_set.SetDigits(nil) // RxException(0, "cannot use nil as type Rexx in argument")

	if new_set.Digits != DEFAULT_DIGITS {
		t.Errorf("got %v, wanted %v", new_set.Digits, DEFAULT_DIGITS)
	}

}

func Test_2047(t *testing.T) {

	new_set := RxSet()

	new_set.SetDigits(RxFromString("½"))

	if new_set.Digits != DEFAULT_DIGITS {
		t.Errorf("got %v, wanted %v", new_set.Digits, DEFAULT_DIGITS)
	}

}

func Test_2048(t *testing.T) {

	new_set := RxSet()

	new_set.SetDigits(RxFromString("۱۱۱۱.0E+6"))

	if new_set.Digits != DEFAULT_DIGITS {
		t.Errorf("got %v, wanted %v", new_set.Digits, DEFAULT_DIGITS)
	}

}

func Test_2049(t *testing.T) {

	new_set := RxSet()

	new_set.SetDigits(RxFromString("+000000000000000.00000000000000000000000000000001")) // NEED

	if new_set.Digits != DEFAULT_DIGITS {
		t.Errorf("got %v, wanted %v", new_set.Digits, DEFAULT_DIGITS)
	}

}

// func (rcvr *RexxSet) SetForm(f *Rexx) error

func Test_2050(t *testing.T) {

	new_set := RxSet()

	new_set.SetForm(nil)

	if new_set.Form != DEFAULT_FORM {
		t.Errorf("got %v, wanted %v", new_set.Form, DEFAULT_FORM)
	}

}

func Test_2051(t *testing.T) {

	new_set := RxSet()

	new_set.SetForm(RxFromString("engineering"))

	if new_set.Form != 1 {
		t.Errorf("got %v, wanted %v", new_set.Form, 1)
	}

}

func Test_2052(t *testing.T) {

	new_set := RxSet()

	new_set.SetForm(RxFromString("scientific"))

	if new_set.Form != 0 {
		t.Errorf("got %v, wanted %v", new_set.Form, 0)
	}

}

func Test_2053(t *testing.T) {

	new_set := RxSet()

	new_set.SetForm(RxFromString("JUST SOME JUNK"))

	if new_set.Form != 0 {
		t.Errorf("got %v, wanted %v", new_set.Form, 0)
	}

}

// func rxutil() (rcvr *RexxUtil)

func Test_2054(t *testing.T) {

	// -- say RexxUtil()

	rxutil := rxutil()

	if rxutil == nil {
		t.Errorf("Call should not be nil")
	}

}

// func d2x(d *Rexx, n int32) ([]rune, error)

func Test_2055(t *testing.T) {

	// -- say RexxUtil.d2x(Rexx("+9990000000000000001"), int 92)

	value, err := d2x(RxFromString("+9990000000000000001"), 92)
	want := "00000000000000000000000000000000000000000000000000000000000000000000000000008AA39C121A270001"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2056(t *testing.T) {

	// --

	new_set := RxSet()
	new_set.SetDigits(RxFromString("16"))

	MaxExp = 16
	MinExp = -16

	_, err := d2x(RxFromString("+9990000000000000001"), 92)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999  // reset
	MinExp = -999999999 // reset

}

func Test_2057(t *testing.T) {

	// -- say RexxUtil.d2x(Rexx("-127.24"), -0) ""

	value, err := d2x(RxFromString("-127.24"), -0)
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2058(t *testing.T) {

	// -- say RexxUtil.d2x(Rexx("-127.1"), 0)

	value, err := d2x(RxFromString("-127.1"), 0)
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2059(t *testing.T) {

	// -- say RexxUtil.d2x(Rexx("-127.00000"), 1)

	value, err := d2x(RxFromString("-127.00000"), 1)
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2060(t *testing.T) {

	// -- say RexxUtil.d2x(Rexx("-.12700000"), 1)

	_, err := d2x(RxFromString("-.12700000"), 1)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2061(t *testing.T) {

	// -- say RexxUtil.d2x(Rexx("9999999999"), -1)

	value, err := d2x(RxFromString("9999999999"), -1)
	want := "2540BE3FF"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2062(t *testing.T) {

	// -- say RexxUtil.d2x(Rexx("-0.9e-999999999"), 0) -- netrexx.lang.ExponentOverflowException: -1000000000

	_, err := d2x(RxFromString("-0.9e-999999999"), 0)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2063(t *testing.T) {

	// -- say RexxUtil.d2x(Rexx("7.62939453629394536293945362939453E+28"), -0)

	value, err := d2x(RxFromString("7.62939453629394536293945362939453E+28"), -0)
	want := ""

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2064(t *testing.T) {

	// -- say RexxUtil.d2x(Rexx("-000001.12345"), -100) -- netrexx.lang.BadArgumentException: -000001.12345

	_, err := d2x(RxFromString("-000001.12345"), -100)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2065(t *testing.T) {

	// -- say RexxUtil.d2x(null, -100)

	_, err := d2x(nil, -100)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2066(t *testing.T) {

	// --

	new_set := RxSet()
	new_set.SetForm(RxFromString("scientific"))
	new_set.SetDigits(RxFromString("16"))
	MaxExp = -1

	_, err := d2x(RxFromString("-9.9E-397"), 251)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999

}

// func Float32Pow(x float64, n int32) (float32, error)

func Test_2067(t *testing.T) {

	// -- say RexxUtil.floatPow(double 10.753567, int 6) -- NetRexx returns 1546377
	// Notes: uses RxFromFloat32(value).ToString() to match NetRexx rounding

	value, err := Float32Pow(10.753567, 6)
	// want := float32(1.5463766e+06) // also works it golang float32 to float32 tested
	want := "1546377"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := RxFromFloat32(value).ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2068(t *testing.T) {

	// --

	_, err := Float32Pow(math.Inf(-1), 2147483647)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2069(t *testing.T) {

	// --

	_, err := Float32Pow(math.Inf(1), 2147483647)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

// func Float64Pow(x float64, n int32) (float64, error)

func Test_2070(t *testing.T) {

	// -- say RexxUtil.doublePow(double 10.753567, int 6)

	sf, _ := Float64Pow(10.753567, 6)

	got := RxFromFloat64(sf)
	want := "1546376.609420583"

	if got.ToString() != want {
		t.Errorf("got %v, wanted %v", got.ToString(), want)
	}

}

func Test_2071(t *testing.T) {

	// -- say RexxUtil.doublePow(10.753567, -1) -- "0.09299239963818518" rounded up last digit

	value, err := Float64Pow(10.753567, -1)
	want := float64(0.09299239963818517)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2072(t *testing.T) {

	// -- say RexxUtil.doublePow(10.753567, 0)

	value, err := Float64Pow(10.753567, 0)
	want := float64(1)

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func format(num *Rexx, before int32, after int32, explaces int32, exdigits int32, exform rune) ([]rune, error)

func Test_2073(t *testing.T) {

	// -- say RexxUtil.format(Rexx(".2111"), 0, 2, 9, 9, 'S')

	value, err := format(RxFromString(".2111"), 0, 2, 9, 9, 'S')
	want := "0.21           "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2074(t *testing.T) {

	// -- say RexxUtil.format(Rexx(" -0.307"), 48, 48, 48, -256, 'A') -- java.lang.NegativeArraySizeException: -187
	// Notes: Numerical arguments less than -1 may produce invalid results

	value, err := format(RxFromString(" -0.307"), 48, 48, 48, -256, 'A')
	want := "                                              -0.307000000000000000000000000000000000000000000000000                                                  "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2075(t *testing.T) {

	// -- say RexxUtil.format(Rexx("1000"), 48, 48, 48, -256, 'A')  -- java.lang.NegativeArraySizeException: -189
	// Notes: Numerical arguments less than -1 may produce invalid results

	value, err := format(RxFromString("1000"), 48, 48, 48, -256, 'A')
	want := "                                            1000.000000000000000000000000000000000000000000000                                                  "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2076(t *testing.T) {

	// -- say RexxUtil.format(Rexx("-12.73"), 48, 48, 48, -256, 'A') -- java.lang.NegativeArraySizeException: -188
	// Notes: Numerical arguments less than -1 may produce invalid results

	value, err := format(RxFromString("-12.73"), 48, 48, 48, -256, 'A')
	want := "                                             -12.730000000000000000000000000000000000000000000000                                                  "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2077(t *testing.T) {

	// -- say RexxUtil.format(Rexx("!!"), 48, 48, 48, -256, 'A') -- java.lang.NullPointerException

	_, err := format(RxFromString("!!"), 48, 48, 48, -256, 'A')

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2078(t *testing.T) {

	// -- say RexxUtil.format(Rexx("0"), 57, 57, 57, -256, 'A') -- java.lang.NegativeArraySizeException: -180
	// Notes: Numerical arguments less than -1 may produce invalid results

	value, err := format(RxFromString("0"), 57, 57, 57, -256, 'A')
	want := "                                                        0.000000000000000000000000000000000000000000000000000000000                                                           "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2079(t *testing.T) {

	// -- say RexxUtil.format(Rexx("4000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"), 48, 55, 53, 88, 'A') -- netrexx.lang.BadArgumentException: 48

	_, err := format(RxFromString("4000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000"), 48, 55, 53, 88, 'A')

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2080(t *testing.T) {

	// -- say format(Rexx("-7.7E-9999998"), 48, 55, 53, -1, 'A')

	value, err := format(RxFromString("-7.7E-9999998"), 48, 55, 53, -1, 'A')
	want := "                                               0.0000000000000000000000000000000000000000000000000000000                                                       "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2081(t *testing.T) {

	// -- say RexxUtil.format(Rexx("1E-12"), 9, 11, 8, 7, char(80))

	value, err := format(RxFromString("1E-12"), 9, 11, 8, 7, rune(80))
	want := "        0.00000000000          "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2082(t *testing.T) {

	// -- say RexxUtil.format(Rexx("1.00000000000E-383"), 13, 11, 11, 12, char(106))

	new_set := RxSet()
	new_set.SetForm(RxFromString("engineering"))
	new_set.SetDigits(RxFromString("16"))

	MaxExp = 384
	MinExp = -383

	_, err := format(RxFromString("1.00000000000E-383"), 13, 11, 11, 12, rune(106))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999
	MinExp = -999999999

}

func Test_2083(t *testing.T) {

	// -- say RexxUtil.format(Rexx("1.00000000000E-383"), 13, 11, 11, 12, char(106))

	value, err := format(RxFromString("1.00000000000E-383"), 13, 11, 11, 12, rune(106))
	want := "            0.00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000000000             "

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2084(t *testing.T) {

	// -- say RexxUtil.format(Rexx("+9.999999999999999E+384"), 9, 12, 10, 14, char(8)) -- netrexx.lang.BadArgumentException: 9

	_, err := format(RxFromString("+9.999999999999999E+384"), 9, 12, 10, 14, '8')

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2085(t *testing.T) {

	// -- say RexxUtil.format(Rexx("+9.999999999999999E+384"), 11, 13, 7, 13, 'L')  -- netrexx.lang.BadArgumentException: 11

	_, err := format(RxFromString("+9.999999999999999E+384"), 11, 13, 7, 13, 'L')

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2086(t *testing.T) {

	// -- say RexxUtil.format(Rexx("-26826826826826820.555555"), 100, 21, 1, 7, 'E') -- netrexx.lang.BadArgumentException: 1

	_, err := format(RxFromString("-26826826826826820.555555"), 100, 21, 1, 7, 'E')

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2087(t *testing.T) {

	// -- say RexxUtil.format(Rexx("-000.00000000006990244916026429904498278982530170295668"), 225, 0, 38, 19, 'S')

	value, err := format(RxFromString("-000.00000000006990244916026429904498278982530170295668"), 225, 0, 38, 19, 'S')
	want := "                                                                                                                                                                                                                               -7E-00000000000000000000000000000000000011"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2088(t *testing.T) {

	// --

	new_set := RxSet()
	new_set.SetDigits(RxFromString("16"))

	MaxExp = 384
	MinExp = -383

	_, err := format(RxFromString("+9.999999999999999E+384"), 9, 12, 10, 14, '8')

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999  // reset
	MinExp = -999999999 // reset

}

func Test_2089(t *testing.T) {

	MaxArg = -1

	_, err := format(RxFromString(".2111"), 0, 2, 9, 9, 'P')

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxArg = 999999999

}

func Test_2090(t *testing.T) {

	// --

	MaxExp = -1

	_, err := format(RxFromString("-.9"), 10, 0, 10, 8, 'P')

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999 // reset

}

func Test_2091(t *testing.T) {

	// -- say RexxUtil.format(null, 10, 0, 10, 8, 'P')

	_, err := format(nil, 10, 0, 10, 8, 'P')

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2092(t *testing.T) {

	// --

	new_set := RxSet()
	new_set.SetDigits(RxFromString("16"))

	MaxExp = 99
	MinExp = -99
	MaxArg = 99
	MinArg = -99

	_, err := format(RxFromString("-4E-8"), -9, -12, 159, 15183, 'E')

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	// reset all
	MaxExp = 999999999
	MinExp = -999999999
	MaxArg = 999999999
	MinArg = -999999999

}

// func torxfromfloat64(num float64, digits int32) (*Rexx, error)

func Test_2093(t *testing.T) {

	// --

	new_set := RxSet()
	new_set.SetDigits(RxFromString("99"))

	MaxExp = 384
	MinExp = -383
	MaxArg = 384
	MinArg = -383
	value, err := torxfromfloat64(math.NaN(), 256)
	want := "NaN"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	// reset all
	MaxExp = 999999999
	MinExp = -999999999
	MaxArg = 999999999
	MinArg = -999999999

}

func Test_2094(t *testing.T) {

	// --

	new_set := RxSet()
	new_set.SetDigits(RxFromString("99"))

	MaxExp = 384
	MinExp = -383
	MaxArg = 384
	MinArg = -383

	value, err := torxfromfloat64(math.Inf(1), 256)
	want := "Infinity"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	// reset all
	MaxExp = 999999999
	MinExp = -999999999
	MaxArg = 999999999
	MinArg = -999999999

}

func Test_2095(t *testing.T) {

	// --

	new_set := RxSet()
	new_set.SetDigits(RxFromString("99"))

	MaxExp = 384
	MinExp = -383
	MaxArg = 384
	MinArg = -383

	value, err := torxfromfloat64(math.Inf(-1), 256)
	want := "Infinity"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	// reset all
	MaxExp = 999999999
	MinExp = -999999999
	MaxArg = 999999999
	MinArg = -999999999

}

func Test_2096(t *testing.T) {

	// -- say RexxUtil.doubleToRexx(double -7.7e-9999998, int 5)

	value, err := torxfromfloat64(-7.7e-9999998, 5)
	want := "0"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2097(t *testing.T) {

	// --

	_, err := torxfromfloat64(math.MaxFloat64+1, 1)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2098(t *testing.T) {

	// --

	_, err := torxfromfloat64(-math.MaxFloat64-1, 1)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2099(t *testing.T) {

	// --

	_, err := torxfromfloat64(float64(10), int32(0))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2100(t *testing.T) {

	// --

	new_set := RxSet()
	new_set.SetDigits(RxFromString("9"))

	MaxExp = 9
	MinExp = -9
	MaxArg = 9
	MinArg = -9

	value, err := torxfromfloat64(math.SmallestNonzeroFloat64, 1)
	want := "1"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

	// reset all
	MaxExp = 999999999
	MinExp = -999999999
	MaxArg = 999999999
	MinArg = -999999999

}

func Test_2101(t *testing.T) {

	// -- say RexxUtil.doubleToRexx(double -2147483649, int 5)

	value, err := torxfromfloat64(-2147483649, 5)
	want := "-2.1475E+9"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2102(t *testing.T) {

	// -- say RexxUtil.doubleToRexx(double 2147483646, int 5)

	value, err := torxfromfloat64(2147483646, 5)
	want := "2.1475E+9"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2103(t *testing.T) {

	// -- say RexxUtil.doubleToRexx(double -9223372036854775809, int 5)

	value, err := torxfromfloat64(-9223372036854775809, 5)
	want := "-9.2234E+18"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2104(t *testing.T) {

	// -- say RexxUtil.doubleToRexx(double 9223372036854775806, int 5)

	value, err := torxfromfloat64(9223372036854775806, 5)
	want := "9.2234E+18"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := value.ToString()
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

func Test_2105(t *testing.T) {

	// --

	MaxExp = 0
	MinExp = 0

	new_set := RxSet()
	new_set.SetDigits(RxFromString("7"))

	torxfromfloat64(float64(10), int32(193))

	MaxExp = 999999999
	MinExp = -999999999

}

func Test_2106(t *testing.T) {

	// --

	new_set := RxSet()
	new_set.SetDigits(RxFromString("16"))

	MaxArg = 32
	MinArg = 32
	MaxExp = 32
	MinExp = 32

	_, err := torxfromfloat64(float64(1), int32(32))

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxArg = 999999999
	MinArg = -999999999
	MaxExp = 999999999
	MinExp = -999999999

}

func Test_2107(t *testing.T) {

	// --

	MaxExp = 96
	MinExp = -95

	new_set := RxSet()
	new_set.SetDigits(RxFromString("7"))

	torxfromfloat64(float64(9.999999e+96), 2)

	MaxExp = 999999999
	MinExp = -999999999

}

// func translate(s []rune, out []rune, in []rune, pad rune) []rune

func Test_2108(t *testing.T) {

	// -- say RexxUtil.translate(null, null, null, ' ')

	got := translate(nil, nil, nil, ' ')

	if got != nil {
		t.Errorf("got %v, wanted %v", got, nil)
	}

}

// func trunc(num *Rexx, after int32) ([]rune, error)

func Test_2109(t *testing.T) {

	// -- say RexxUtil.trunc(Rexx('a'), int 5)

	_, err := trunc(RxFromRune('a'), 5)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2110(t *testing.T) {

	// -- say RexxUtil.trunc(null, int 5)

	_, err := trunc(nil, 5)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2111(t *testing.T) {

	// -- say RexxUtil.trunc(Rexx("7.62939453629394536293945362939453E+28"), 0)

	value, err := trunc(RxFromString("7.62939453629394536293945362939453E+28"), 0)
	want := "76293945362939453629394536293"

	if err != nil {
		t.Errorf("Expecting a value.")
	} else {
		got := string(value)
		if got != want {
			t.Errorf("got %v, wanted %v", got, want)
		}
	}

}

// func x2b(x *Rexx) ([]rune, error)

func Test_2112(t *testing.T) {

	// -- say RexxUtil.X2b(null)

	_, err := x2b(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

// func x2c(x *Rexx) (rune, error)

func Test_2113(t *testing.T) {

	// -- say RexxUtil.X2c(null)

	_, err := x2c(nil)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

// func x2d(x *Rexx, n int32) ([]rune, error)

func Test_2114(t *testing.T) {

	// -- say RexxUtil.x2d(null, int 5)

	_, err := x2d(nil, 5)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2115(t *testing.T) {

	// --

	new_set := RxSet()
	new_set.SetDigits(RxFromString("4"))

	MaxExp = 4
	MinExp = -4

	x2d(RxFromString("1e999999999"), 9)

	MaxExp = 999999999
	MinExp = -999999999

}

func Test_2116(t *testing.T) {

	// --

	new_set := RxSet()
	new_set.SetDigits(RxFromString("3"))

	MaxExp = 0
	MinExp = 0

	x2d(RxFromString("0.0000E-50"), 6)

	MaxExp = 999999999
	MinExp = -999999999

}

func Test_2117(t *testing.T) {

	// -- say RexxUtil.x2d(Rexx("-0.009E-99975929096475.63450425339472559646E+153"), int 21) -- netrexx.lang.BadArgumentException: Bad hexadecimal -0.009E-99975929096475.63450425339472559646E+153

	_, err := x2d(RxFromString("-0.009E-99975929096475.63450425339472559646E+153"), 21)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

}

func Test_2118(t *testing.T) {

	// --

	new_set := RxSet()
	new_set.SetForm(RxFromString("scientific"))
	new_set.SetDigits(RxFromString("16"))
	MaxExp = -1

	_, err := x2d(RxFromString("Stuff"), 1)

	if err == nil {
		t.Errorf("Expected an error, but got nil")
	}

	MaxExp = 999999999

}

// func hexint(c rune) int32

// func rxwords() (rcvr *RexxWords)

func Test_2119(t *testing.T) {

	// -- say RexxWords()

	rxwords := rxwords()

	if rxwords == nil {
		t.Errorf("Call should not be nil")
	}

}

// func abbrev(a []rune, b []rune, leng int32) bool

func Test_2120(t *testing.T) {

	// -- say RexxWords.Abbrev(char[] "\0x01", char[] "\0x01\0x02", 0)

	got := abbrev([]rune{1}, []rune{1, 2}, 0)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2121(t *testing.T) {

	// -- say RexxWords.Abbrev(char[] [32, 33, 34], char[] [35, 36, 37], 3)

	got := abbrev([]rune{32, 33, 34}, []rune{35, 36, 37}, 3)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2122(t *testing.T) {

	// -- say RexxWords.Abbrev(char[] [32], char[] [32, 49, 32], -2147483648) -- java.lang.NumberFormatException: Conversion overflow

	got := abbrev([]rune{32}, []rune{32, 49, 32}, -2147483648)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2123(t *testing.T) {

	// -- say RexxWords.Abbrev(null, null, 1)

	got := abbrev(nil, nil, 1)
	want := false

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func centre(s []rune, wid int32, pad rune) []rune

func Test_2124(t *testing.T) {

	// -- say RexxWords.Centre(char[] "A", 1, ' ')

	got := string(centre([]rune{65}, 1, ' '))
	want := "A"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2125(t *testing.T) {

	// -- say RexxWords.Centre(char[] [32, 49, 32], -2147483648, 'A') -- java.lang.NumberFormatException: Conversion overflow

	got := string(centre([]rune{32, 49, 32}, -2147483648, 'A'))
	want := ""

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2126(t *testing.T) {

	// -- say RexxWords.Centre(char[] [32, 49, 32], 2, 'A') returns "g."

	got := string(centre([]rune{32, 49, 32}, 2, 'A'))
	want := " 1"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2127(t *testing.T) {

	// -- say RexxWords.Centre(char[] [32, 49, 32], 0, 'A')

	got := string(centre([]rune{32, 49, 32}, 0, 'A'))
	want := ""

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func changestr(needle []rune, haystack []rune, _new []rune) []rune

func Test_2128(t *testing.T) {

	// -- say RexxWords.changestr(char[]"i", char[] "spring",'u')

	got := string(changestr([]rune("i"), []rune("spring"), []rune("u")))
	want := "sprung"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func compare(a []rune, b []rune, pad rune) int32

func Test_2129(t *testing.T) {

	// -- say RexxWords.Compare(char[] "1Z3", char[] "20.11E-1", 'A')

	got := compare([]rune("1Z3"), []rune("20.11E-1"), 'A')
	want := int32(1)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2130(t *testing.T) {

	// -- say RexxWords.Compare(char[] "2.147483647E+8", char[] "2.147483647E+8", 'A')

	got := compare([]rune("2.147483647E+8"), []rune("2.147483647E+8"), 'A')
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func countstr(needle []rune, haystack []rune) int32

func Test_2131(t *testing.T) {

	// -- say RexxWords.countstr(char[]"i", char[] "spring")

	got := countstr([]rune("i"), []rune("spring"))
	want := int32(1)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func delstr(s []rune, start int32, leng int32) []rune

func Test_2132(t *testing.T) {

	// -- say RexxWords.DelStr(char[] [32, 49, 32], -1, 1) -- java.lang.ArrayIndexOutOfBoundsException
	// Notes: This port adjusts zero, negative numbers or returns data unchanged to prevent Go panics.

	got := string(delstr([]rune{32, 49, 32}, -1, 1))
	want := " 1 "

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2133(t *testing.T) {

	// -- say RexxWords.DelStr(char[] [32, 49, 32], 1, -1) -- java.lang.ArrayIndexOutOfBoundsException
	// Notes: This port adjusts zero, negative numbers or returns data unchanged to prevent Go panics.

	got := string(delstr([]rune{32, 49, 32}, 1, -1))
	want := " 1 "

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func delword(s []rune, num int32, leng int32) []rune

func Test_2134(t *testing.T) {

	// -- say RexxWords.DelWord(char[] "A", 0, 0)

	got := string(delword([]rune{65}, 0, 0))
	want := "A"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2135(t *testing.T) {

	// -- say RexxWords.DelWord(char[] "A", 1, 0)
	// Notes: This port adjusts zero, negative numbers or returns data unchanged to prevent Go panics.

	got := string(delword([]rune{65}, 1, -9))
	want := "A"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func insert(chars []rune, newchars []rune, num int32, leng int32, padchar rune) []rune

func Test_2136(t *testing.T) {

	// -- say RexxWords.insert(null, char[] "xyz", -2, -2, ' ') -- java.lang.NullPointerException
	// Notes: This port adjusts zero, negative numbers or returns data unchanged to prevent Go panics.

	got := string(insert(nil, []rune("xyz"), int32(-2), int32(-2), ' '))

	want := "xyz"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2137(t *testing.T) {

	// -- say RexxWords.insert(char[] "xyz", null, 0, 0, ' ')

	got := string(insert([]rune("xyz"), nil, int32(0), int32(0), ' '))

	want := "xyz"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2138(t *testing.T) {

	// -- say RexxWords.insert(char[] "xyz", null, -2, -2, ' ') -- java.lang.ArrayIndexOutOfBoundsException: arraycopy: source index -2 out of bounds for char[3]
	// Notes: This port adjusts zero, negative numbers or returns data unchanged to prevent Go panics.

	got := string(insert([]rune("xyz"), nil, int32(-2), int32(-2), ' '))

	want := "xyz"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2139(t *testing.T) {

	// -- say RexxWords.insert(null, char[] "xyz", 0, 0, ' ') -- java.lang.NullPointerException: Cannot read the array length because "<parameter1>" is null
	// Notes: This port adjusts zero, negative numbers or returns data unchanged to prevent Go panics.

	got := string(insert(nil, []rune("xyz"), int32(0), int32(0), ' '))

	want := ""

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func nextblank(s []rune, start int32) int32

func Test_2140(t *testing.T) {

	// -- say RexxWords.nextblank(char[] "   ", 1)

	got := nextblank([]rune{32, 32, 32}, 1)
	want := int32(1)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func nextnonblank(s []rune, start int32) int32

func Test_2141(t *testing.T) {

	// -- say RexxWords.nextnonblank(char[] "   ", 1)

	got := nextnonblank([]rune{32, 32, 32}, 1)
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func overlay(chars []rune, newchars []rune, num int32, leng int32, padchar rune) []rune

func Test_2142(t *testing.T) {

	// -- say RexxWords.overlay(null, char[] "xyz", -2, -2, ' ') --  java.lang.NullPointerException
	// Notes: This port adjusts zero, negative numbers or returns data unchanged to prevent Go panics.

	got := string(overlay(nil, []rune("xyz"), int32(-2), int32(-2), ' '))
	want := "xyz"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2143(t *testing.T) {

	// -- say RexxWords.overlay(char[] "xyz", null, -2, -2, ' ') -- java.lang.ArrayIndexOutOfBoundsException
	// Notes: This port adjusts zero, negative numbers or returns data unchanged to prevent Go panics.

	got := string(overlay([]rune("xyz"), nil, int32(-2), int32(-2), ' '))
	want := "xyz"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2144(t *testing.T) {

	// -- say RexxWords.overlay(null, char[] "xyz", 0, 0, ' ') --  java.lang.NullPointerException
	// Notes: This port adjusts zero, negative numbers or returns data unchanged to prevent Go panics.

	got := string(overlay(nil, []rune("xyz"), int32(0), int32(0), ' '))
	want := ""

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2145(t *testing.T) {

	// -- say RexxWords.overlay(char[] "xyz", null, 0, 0, ' ') -- java.lang.ArrayIndexOutOfBoundsException
	// Notes: This port adjusts zero, negative numbers or returns data unchanged to prevent Go panics.

	got := string(overlay([]rune("xyz"), nil, int32(0), int32(0), ' '))
	want := "xyz"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func pos(needle []rune, haystack []rune, start int32) int32

func Test_2146(t *testing.T) {

	// --

	got := pos([]rune{0, 'ø'}, []rune{0, 'ø', 'Þ'}, -22)
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func space(s []rune, gap int32, pad rune) []rune

func Test_2147(t *testing.T) {

	// -- say RexxWords.Space(char[] " 1", 0, ' ')

	got := string(space([]rune{32, 49}, 0, ' '))
	want := "1"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2148(t *testing.T) {

	// --

	got := string(space([]rune("1 1"), int32(-9), ' '))
	want := "1 1"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func subword(s []rune, num int32, leng int32) []rune

func Test_2149(t *testing.T) {

	// -- say RexxWords.subword(char[] "Now is the  time", 1, -9)

	got := string(subword([]rune("Now is the  time"), 1, -9))
	want := "Now is the  time"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func verify(s []rune, v []rune, opt rune, start int32) int32

func Test_2150(t *testing.T) {

	// -- say RexxWords.Verify(char[] "\x01", char[] "\x01", 'N', 1)

	got := verify([]rune{1}, []rune{1}, 'N', 1)
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2151(t *testing.T) {

	// -- say RexxWords.Verify(char[] "\x01", char[] "\x01", 'M', 1)

	got := verify([]rune{1}, []rune{1}, 'M', 1)
	want := int32(1)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2152(t *testing.T) {

	// -- say RexxWords.Verify(char[] " 1 ", char[] " 1 ", 'A', 3)

	got := verify([]rune{32, 49, 32}, []rune{32, 49, 32}, 'A', 3)
	want := int32(3)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func verifym(s []rune, v []rune, start int32) int32

func Test_2153(t *testing.T) {

	// -- say RexxWords.Verifym(char[] "+0000000000000000009.00000000100000000", char[] "k", 1)

	got := verifym([]rune("+0000000000000000009.00000000100000000"), []rune("k"), 1)
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2154(t *testing.T) {

	// -- say RexxWords.verifym(char[] "\x01", char[] "\x01",  1)

	got := verifym([]rune{1}, []rune{1}, 1)
	want := int32(1)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2155(t *testing.T) {

	// -- say RexxWords.verifym(char[] "\x01", char[] "\x01",  2)

	got := verifym([]rune{1}, []rune{1}, 2)
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2156(t *testing.T) {

	// -- say RexxWords.verifym(char[] "\x01", char[] "",  1)

	got := verifym([]rune{1}, []rune{}, 1)
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2157(t *testing.T) {

	// -- say RexxWords.verifym(char[] "\x00", char[] "\x00\x00" , 1) -- use 1 instead of 0
	// Notes: This port adjusts zero, negative numbers or returns data unchanged to prevent Go panics.

	got := verifym([]rune{0}, []rune{0, 0, 32}, 0)
	want := int32(1)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func verifyn(s []rune, v []rune, start int32) int32

// func word(s []rune, num int32) []rune

func Test_2158(t *testing.T) {

	// -- say  RexxWords.Word(char[] "A", 1)

	got := string(word([]rune{65}, 1))
	want := "A"

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func wordindex(s []rune, num int32) int32

// func wordlength(s []rune, num int32) int32

func Test_2159(t *testing.T) {

	// -- say RexxWords.WordLength(char[] "\x01", 1)

	got := wordlength([]rune{1}, 1)
	want := int32(1)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2160(t *testing.T) {

	// -- say RexxWords.WordLength(char[] "\x01", 1)

	got := wordlength([]rune{1}, 1)
	want := int32(1)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func wordpos(needle []rune, haystack []rune, wpos int32) int32

func Test_2161(t *testing.T) {

	// --

	got := wordpos([]rune{0, 'ø'}, []rune{0, 'ø', 'Þ'}, 2)
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2162(t *testing.T) {

	// --

	got := wordpos([]rune{0, 0, 32}, []rune{0, 0, 32}, 1)
	want := int32(1)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2163(t *testing.T) {

	// --

	got := wordpos([]rune{0}, []rune{0, 0, 32}, 1)
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2164(t *testing.T) {

	// --

	got := wordpos([]rune{20, 32, 20}, []rune{20, 32, 20}, 2)
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2165(t *testing.T) {

	// --

	got := wordpos([]rune{32}, []rune{1}, 1)
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2166(t *testing.T) {

	// --

	got := wordpos([]rune{32}, []rune{32}, 0)
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func Test_2167(t *testing.T) {

	// --

	got := wordpos([]rune{65, 32, 32, 66, 32}, []rune{}, -3)
	want := int32(0)

	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

// func words(s []rune) int32

