package utils

import (
	"fmt"
	"strings"
)

func ConvertToCommaSeperatedString[T any](data []T) string {
	fieldsStr := strings.Builder{}
	for i, d := range data {
		if i != 0 {
			fieldsStr.WriteString(",")
		}
		str, ok := any(d).(fmt.Stringer)
		if ok {
			fieldsStr.WriteString(str.String())
		} else {
			fieldsStr.WriteString(fmt.Sprintf("%v", d))
		}
	}

	return fieldsStr.String() 
}
