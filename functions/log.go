package functions

import (
	"time"

	"github.com/GalacticDocs/Conversions/structs"
	"github.com/GalacticDocs/Conversions/builders"
	"github.com/GalacticDocs/Conversions/converter"
	CLRS "github.com/iVitaliya/colors-go"
)

func bolder() *structs.IBolder {
	return &structs.IBolder{
		Green: func(msg string) string {
			return CLRS.Bold(CLRS.Green(msg))
		},
		Red: func(msg string) string {
			return CLRS.Bold(CLRS.Red(msg))
		},
		Yellow: func(msg string) string {
			return CLRS.Bold(CLRS.BrightYellow(msg))
		},
		Gray: func(msg string) string {
			return CLRS.Bold(CLRS.Gray(msg))
		},
		White: func(msg string) string {
			return CLRS.Bold(CLRS.White(msg))
		},
	}
}

func format(format_type string, str string) string {
	var (
		timer string = time.Now().Format("DD MMMM YYYY HH:mm:ss A")
		_str  string = ""
	)

	switch format_type {
	case "time":
		_str = bolder().Gray("[") + bolder().Yellow(timer) + bolder().Gray("]")
	case "exception":
		_str = bolder().Gray("(") + bolder().Green(str) + bolder().Gray(")")
	case "message":
		_str = bolder().Gray(":") + bolder().White(str)
	}

	return _str
}

func LogException(exception string, msg string) {
	now := time.Now()
	built := builders.StringBuilder().New("")

	built
}
