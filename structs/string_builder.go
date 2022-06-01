package structs

type builder struct {
	Append                func(string) string
	ToString              func() string
	Format                func(string, ...interface{}) string
	Slice                 func(int, int) string
	Split                 func(string, int) []string
	Replace               func(string, string, string) string
	ReplaceAll            func(string, string, string) string
	TrimSpace             func(string) string
	TrimLeft              func(string) string
	TrimRight             func(string) string
	Trim                  func(string) string
	ToLower               func(string) string
	ToUpper               func(string) string
	Join                  func([]string, string) string
	Contains              func(string, string) bool
	HasPrefix             func(string, string) bool
	HasSuffix             func(string, string) bool
	EqualFold             func(string, string) bool
	Equal                 func(string, string) bool
	Compare               func(string, string) int
	CompareFold           func(string, string) int
	CompareFoldRunes      func(string, string) int
	CompareRunes          func(string, string) int
	CompareRunesFold      func(string, string) int
	CompareRunesFoldRunes func(string, string) int
}

type StringBuilder struct {
	New func(string) *builder
}
