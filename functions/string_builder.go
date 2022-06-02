package functions

import (
	"fmt"
	"strings"
)

// Appends the string to the end of the original string.
func Append(s string, str string) string {
	return s + str
}

// Returns the string representation of the value.
func ToString(s string) string {
	return s
}

// Formats the string with provided args.
func Format(s string, args ...interface{}) string {
	return fmt.Sprintf(s, args...)
}

// Returns a slice of the string from [start:end].
func Slice(s string, start int, end int) string {
	return s[start:end]
}

// Splits the string s into substrings at each UTF-8 sequence of bytes specified by sep.
func Split(s string, sep string) []string {
	return strings.Split(s, sep)
}

// Replaces the first n occurrences of old with new in s.
func Replace(s string, old string, new string) string {
	return strings.Replace(s, old, new, -1)
}

// Replaces all occurrences of old with new in s.
func ReplaceAll(s string, old string, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// Removes leading and trailing white spaces from s.
func TrimSpace(s string) string {
	return strings.TrimSpace(s)
}

// Returns a slice of the string s with all leading Unicode code points contained in cutset removed.
//
// To remove a prefix from a string, use TrimPrefix instead.
func TrimLeft(s string, sep string) string {
	return strings.TrimLeft(s, sep)
}

// Returns a slice of the string s, with all trailing Unicode code points contained in cutset removed.
//
// To remove a suffix from a string, use TrimSuffix instead.
func TrimRight(s string, sep string) string {
	return strings.TrimRight(s, sep)
}

// Returns s without the provided leading prefix string.
func TrimPrefix(s string, sep string) string {
	return strings.TrimPrefix(s, sep)
}

// Returns s without the provided trailing suffix string.
func TrimSuffix(s string, sep string) string {
	return strings.TrimSuffix(s, sep)
}

// Returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.
func Trim(s string, sep string) string {
	return strings.Trim(s, sep)
}

// Returns s with all Unicode letters mapped to their lower case.
func ToLower(s string) string {
	return strings.ToLower(s)
}

// Returns s with all Unicode letters mapped to their upper case.
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// Join concatenates the elements of its first argument to create a single string. The separator string sep is placed between elements in the resulting string.
func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}