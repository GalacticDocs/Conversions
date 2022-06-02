package structs

type IBolder struct {
	Green  func(string) string
	Red    func(string) string
	Yellow func(string) string
	Gray   func(string) string
	White  func(string) string
}
