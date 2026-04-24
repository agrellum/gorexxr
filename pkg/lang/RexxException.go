package lang

import (
	"errors"
)

var ExceptionType = []string{"ArgumentNullException", "BadArgumentException", "BadColumnException", "BadNumericException", "DivideException", "ExponentOverflowException", "NoOtherwiseException", "NotCharacterException", "NotLogicException", "NumberFormatException"}

// Differences: All Rexx Exceptions are combined in this one function.
//
//	By number and name.
//	0 - ArgumentNullException
//	1 - BadArgumentException
//	2 - BadColumnException
//	3 - BadNumericException
//	4 - DivideException
//	5 - ExponentOverflowException
//	6 - NoOtherwiseException
//	7 - NotCharacterException
//	8 - NotLogicException
//	9 - NumberFormatException"
func RxException(kind int, detail string) error {
	if kind > 9 {
		// Prevents Go Panic
		kind = 1
	}
	message := ExceptionType[kind] + " : " + detail
	return errors.New(message)
}
