package main

import (
	"fmt"
	"strings"
)

func isWildCardMatches(base, value string) bool {
	// This means change all values and it's causes problems
	if base == "*" {
		return false
	}

	return true
}

func isMatchesForArray(base string, value string) bool {
	if base == value {
		return true
	}

	baseArr := strings.Split(base, ".")
	valueArr := strings.Split(value, ".")

	for baseIndex, baseValue := range baseArr {
		for valueIndex, valueValue := range valueArr {
			if baseIndex == valueIndex && baseValue == valueValue {
				continue
			}

			if baseIndex == valueIndex && isSame(base, value) {
				continue
			}

			if baseIndex != valueIndex {
				continue
			}

			return false
		}
	}

	return true
}

func isSame(base, value string) bool {
	var aStack Stack
	var bStack Stack

	for _, v := range base {
		aStack.Push(fmt.Sprintf("%v", v))
	}

	for _, v := range value {
		bStack.Push(fmt.Sprintf("%v", v))
	}

	for !bStack.IsEmpty() {
		aValue, aHasValue := aStack.Pop()
		bValue, bHasValue := bStack.Pop()

		if aHasValue && bHasValue && aValue == bValue {
			continue
		}

		if aHasValue && bHasValue && aValue != bValue {
			bNextValue, bHasNextValue := bStack.Pop()
			if bHasNextValue && bNextValue == aValue {
				continue
			}

		}
		return false
	}

	return true
}
