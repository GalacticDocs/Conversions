package builders

import (
	"fmt"
	"strings"

	"github.com/GalacticDocs/Conversions/structs"
	"github.com/GalacticDocs/Conversions/functions"
	CLRS "github.com/iVitaliya/colors-go"
)

// Appends the string to the end of the original string.
func append(s string, str string) string {
	return s + str
}

// Returns the string representation of the value.
func toString(s string) string {
	return s
}

// Formats the string with provided args.
func format(s string, args ...interface{}) string {
	return fmt.Sprintf(s, args...)
}

// Returns a slice of the string from [start:end].
func slice(s string, start int, end int) string {
	return s[start:end]
}

// Splits the string s into substrings at each UTF-8 sequence of bytes specified by sep.
func split(s string, sep string) []string {
	return strings.Split(s, sep)
}

// Replaces the first n occurrences of old with new in s.
func replace(s string, old string, new string) string {
	return strings.Replace(s, old, new, -1)
}

// Replaces all occurrences of old with new in s.
func replaceAll(s string, old string, new string) string {
	return strings.ReplaceAll(s, old, new)
}

// Removes leading and trailing white spaces from s.
func trimSpace(s string) string {
	return strings.TrimSpace(s)
}

func StringBuilder() *structs.StringBuilder {
	return &structs.StringBuilder{
		New: func(s string) *structs.StrBuilder {
			builder := &structs.StrBuilder{
				Append: func(str string) string {
					return append(str)
				},
				ToString: func() string {
					return toString(s)
				},
				Format: func(args ...interface{}) string {
					return format(s, args...)
				},
				Slice: func(start int, end int) string {
					return slice(s, start, end)
				},
				Split: func(sep string) []string {
					return split(s, sep)
				},
				Replace: func(str string, old string, new string) string {
					return replace(s, old, new)
				},
				ReplaceAll: func(str string, old string, new string) string {
					return replaceAll(s, old, new)
				},
				TrimSpace: func(str string) string {
					return trimSpace(str)
				},
				TrimLeft: func(str string) string {
					
				},
				TrimRight: func(str string) string {
					
				},
				Trim: func(str string) string {
					
				},
				ToLower: func(str string) string {
					
				},
				ToUpper: func(str string) string {
					
				},
				Join: func(str []string, sep string) string {
					
				},
				Contains: func(str string, substr string) bool {

				},
				HasPrefix: func(str string, prefix string) bool {
				
				},
				HasSuffix: func(str string, suffix string) bool {

				},
				EqualFold: func(str string, s string) bool {

				},
				Equal: func(str string, s string) bool {

				},
				Compare: func(str string, s string) int {

				},
				CompareFold: func(str string, s string) int {

				},
				CompareFoldRunes: func(str string, s string) int {

				},
				CompareRunes: func(str string, s string) int {

				},
				CompareRunesFold: func(str string, s string) int {

				},
				CompareRunesFoldRunes: func(str string, s string) int {

				},
				CompareRunesReverse: func(str string, s string) int {

				},
				CompareRunesReverseFold: func(str string, s string) int {

				},
				CompareRunesReverseFoldRunes: func(str string, s string) int {

				},
	}
}
