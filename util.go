package main

import (
	"fmt"
	"os"
	"strings"
)

func checkHasSequancalWildcard(base string) bool {
	if strings.Contains(base, "*.*") {
		return true
	}
	return false
}

func checkArraySelector(base string) bool {
	if strings.Contains(base, "[") {
		return true
	}
	return false
}

func checkHasWildcard(base string) bool {
	if strings.Contains(base, "*") {
		return true
	}
	return false
}

func isWildCardMatches(base, value string) bool {
	// This means change all values and it's causes problems
	if base == "*" {
		return false
	}

	if !checkHasWildcard(base) {
		return false
	}

	if checkHasSequancalWildcard(base) {
		fmt.Printf("Cannot use consecutive wildcard in absolute config key: %s\n", base)
		os.Exit(1)
	}

	if checkArraySelector(base) {
		fmt.Printf("Cannot use array selector with wildcards in absolute config key: %s\n", base)
		os.Exit(1)
	}

	baseArr := strings.Split(base, ".")
	valueArr := strings.Split(value, ".")

	// If valueArr contains wildcard, it's not ok.
	for i := range valueArr {
		if valueArr[i] == "*" {
			return false
		}
	}
	isWildCard := false
	for i := range baseArr {

		if baseArr[i] == "*" {
			isWildCard = true
		}
	}

	if !isWildCard {
		return false
	}

	baseIndex := 0
	valueIndex := 0

	for {

		if len(baseArr) == baseIndex && len(valueArr) == valueIndex {
			return true
		}

		if len(baseArr) == baseIndex {
			return false
		}

		if len(valueArr) == valueIndex {
			return false
		}

		if valueArr[valueIndex] == baseArr[baseIndex] {
			baseIndex++
			valueIndex++
			continue
		}

		if len(baseArr) == baseIndex+1 {
			return false
		}

		if valueArr[valueIndex] == baseArr[baseIndex+1] {
			baseIndex++
			continue
		}

		if baseArr[baseIndex] == "*" {
			valueIndex++
			continue
		}

		return false

	}

	return true
}

func wildcardCount(values []string) int {
	count := 0
	for _, value := range values {
		if value == "*" {
			count++
		}
	}
	return count
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
