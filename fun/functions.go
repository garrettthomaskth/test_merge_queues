package fun

import "strings"

func ConcatinateStringsWithSpace(strs ...string) string {
	return strings.Join(strs, " ")
}