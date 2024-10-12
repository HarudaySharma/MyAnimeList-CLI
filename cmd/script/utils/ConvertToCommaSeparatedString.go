package utils

import (
	"fmt"
	"strings"
)

type ConvertToCommaSeperatedStringParams[T any] struct {
	Data            []T
	MaxLineLen      int
	SpaceAfterComma bool
}

func ConvertToCommaSeperatedString[T any](p ConvertToCommaSeperatedStringParams[T]) string {
	fieldsStr := strings.Builder{}
	lineLen := 0
	for i, d := range p.Data {
		if i != 0 {
			fieldsStr.WriteString(",")
			if p.SpaceAfterComma {
				fieldsStr.WriteString(" ")
			}
		}

		if p.MaxLineLen != 0 && lineLen > p.MaxLineLen {
			fieldsStr.WriteString("\n")
			lineLen = 0
		}

		str, ok := any(d).(fmt.Stringer)
		appendStr := ""
		if ok {
			appendStr = str.String()
		} else {
			appendStr = fmt.Sprintf("%v", d)
		}
		fieldsStr.WriteString(appendStr)

		lineLen += len(appendStr)
	}

	return fieldsStr.String()
}
