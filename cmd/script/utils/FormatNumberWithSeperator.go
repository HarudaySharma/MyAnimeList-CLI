package utils

import (
	"fmt"
	"strings"
)

func FormatNumberStringWithSeparator(numStr string, sep string) string {
	nLen := len(numStr)

	// Handle negative numbers
	start := 0
	if numStr[0] == '-' {
		start = 1
	}

	// Insert the separator every three digits (thousands)
	var sb strings.Builder
	for i, digit := range numStr[start:] {
		if i > 0 && (nLen-i)%3 == 0 {
			sb.WriteString(sep)
		}
		sb.WriteRune(digit)
	}

	if start == 1 {
		return "-" + sb.String() // Re-add the negative sign for negative numbers
	}

	if sb.Len() == 0 {
		return "0"
	}
	return sb.String()
}

func FormatNumberInterfaceWithSeparator(ni interface{}, sep string) string {
	switch num := ni.(type) {
	case string:
		return FormatNumberStringWithSeparator(num, sep)
	case int:
		return FormatNumberWithSeparator(int64(num), sep)
	}

	return "-"
}

func FormatNumberWithSeparator(ne int64, sep string) string {
	numStr := fmt.Sprintf("%d", ne)
	nLen := len(numStr)

	// Handle negative numbers
	start := 0
	if numStr[0] == '-' {
		start = 1
	}

	// Insert the separator every three digits (thousands)
	var sb strings.Builder
	for i, digit := range numStr[start:] {
		if i > 0 && (nLen-i)%3 == 0 {
			sb.WriteString(sep)
		}
		sb.WriteRune(digit)
	}

	if start == 1 {
		return "-" + sb.String() // Re-add the negative sign for negative numbers
	}

	if sb.Len() == 0 {
		return "0"
	}
	return sb.String()
}
