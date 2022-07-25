package reader

import (
	"regexp"

	"github.com/iVitaliya/logger-go"
)

func ReadOnlyNumbers(str string) string {
	var (
		re         = regexp.MustCompile(`^(?:(?:0|[1-9]\d*)(?:\.\d*)?|\.\d+)$`)
		searchable = re.FindString(str)
	)

	if Contains(searchable).String("") {
		logger.Error("The string which needed to be searched for numbers couldn't find anything, are you sure your string contains any numbers?")
	}

	return searchable
}
