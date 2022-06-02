package structs

type StrBuilder struct {
	// Appends the string to the end of the original string.
	Append                		 func(string) string
	// Returns the string representation of the value.
	ToString              		 func() string
	// Formats the string with provided args.
	Format                		 func(string, ...interface{}) string
	// Returns a slice of the string from [start:end].
	Slice                 		 func(int, int) string
	// Splits the string s into substrings at each UTF-8 sequence of bytes specified by sep.
	Split                 		 func(string) []string
	// Replaces the first occurrences of old with new.
	Replace               		 func(string, string, string) string
	// Replaces all occurrences of old with new.
	ReplaceAll            		 func(string, string, string) string
	// Removes leading and trailing white spaces from str.
	TrimSpace             		 func(string) string
	// Returns a slice of the string s with all leading Unicode code points contained in cutset removed.
	//
	// To remove a prefix from a string, use TrimPrefix instead.
	TrimLeft              		 func(string) string
	// Returns a slice of the string s, with all trailing Unicode code points contained in cutset removed.
	//
	// To remove a suffix from a string, use TrimSuffix instead.
	TrimRight             		 func(string) string
	// Returns s without the provided leading prefix string.
	TrimPrefix					 func(string) string
	// Returns s without the provided trailing suffix string.
	TrimSuffix 		  		 	 func(string) string
	// Returns a slice of the string s with all leading and trailing Unicode code points contained in cutset removed.
	Trim                  		 func(string) string
	// Returns s with all Unicode letters mapped to their lower case.
	ToLower               		 func(string) string
	// Returns s with all Unicode letters mapped to their upper case.
	ToUpper              		 func(string) string
	// Join concatenates the elements of its first argument to create a single string. The separator string sep is placed between elements in the resulting string.
	Join                  		 func([]string, string) string
	
	Contains              		 func(string, string) bool

	HasPrefix             		 func(string, string) bool

	HasSuffix             		 func(string, string) bool

	EqualFold             		 func(string, string) bool

	Equal                 		 func(string, string) bool

	Compare               		 func(string, string) int

	CompareFold           		 func(string, string) int
	
	CompareFoldRunes      		 func(string, string) int
	
	CompareRunes          		 func(string, string) int
	
	CompareRunesFold      		 func(string, string) int
	
	CompareRunesFoldRunes 	 	 func(string, string) int

	CompareRunesReverse    		 func(string, string) int

	CompareRunesReverseFold 	 func(string, string) int

	CompareRunesReverseFoldRunes func(string, string) int
}

type StringBuilder struct {
	New func(string) *StrBuilder
}
