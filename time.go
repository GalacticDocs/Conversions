package conversions

import (
	"time"
)

func ConvertTime(hour int, minute int, second int) string {
	var (
		str string
		hr  = Stringify().IntToString(hour)
		min = Stringify().IntToString(minute)
		sec = Stringify().IntToString(second)
	)

	if hour > 11 {
		str = hr + ":" + min + ":" + sec + " PM"
	} else if hour < 12 {
		str = hr + ":" + min + ":" + sec + " AM"
	}

	return str
}

func ConvertDay(day int) string {
	var str string

	if day == 1 || day == 21 || day == 31 {
		str = Stringify().IntToString(day) + "st"
	} else if day == 2 || day == 22 {
		str = Stringify().IntToString(day) + "nd"
	} else if day == 3 || day == 23 {
		str = Stringify().IntToString(day) + "rd"
	} else {
		str = Stringify().IntToString(day) + "th"
	}

	return str
}

func CheckLeapYear(year int) bool {
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		return true
	}

	return false
}

func DaysCapacity(month int, year int) int {
	var cap int

	switch month {
	case 1:
		cap = 31
	case 2:
		if CheckLeapYear(year) {
			cap = 29
			break
		}
		cap = 28
	case 3:
		cap = 31
	case 4:
		cap = 30
	case 5:
		cap = 31
	case 6:
		cap = 30
	case 7:
		cap = 31
	case 8:
		cap = 31
	case 9:
		cap = 30
	case 10:
		cap = 31
	case 11:
		cap = 30
	case 12:
		cap = 31
	default:
		cap = 30
	}

	return cap
}

func ConvertMonth(month int) string {
	var str string

	switch month {
	case 1:
		str = "January"
	case 2:
		str = "February"
	case 3:
		str = "March"
	case 4:
		str = "April"
	case 5:
		str = "May"
	case 6:
		str = "June"
	case 7:
		str = "July"
	case 8:
		str = "August"
	case 9:
		str = "September"
	case 10:
		str = "October"
	case 11:
		str = "November"
	case 12:
		str = "December"
	default:
		str = "January"
	}

	return str
}

type IDate struct {
	Current string
	// Format: day, month, year, hour, minute, second
	Custom func(int, int, int, int, int, int) string
}

func Date() *IDate {
	var (
		standard_builder = StringBuilder("")
		custom_builder   = StringBuilder("")
		now              = time.Now().Local()
		timeStr          = ConvertTime(now.Hour(), now.Minute(), now.Second())
		dayStr           = ConvertDay(now.Day())
		monthStr         = ConvertMonth(int(now.Month()))
	)

	standard_builder.Append(dayStr)
	standard_builder.Append(" of ")
	standard_builder.Append(monthStr)
	standard_builder.Append(" ")
	standard_builder.Append(Stringify().IntToString(now.Year()))
	standard_builder.Append(" | ")
	standard_builder.Append(timeStr)

	return &IDate{
		Current: standard_builder.Build(),
		Custom: func(day int, month int, year int, hour int, minute int, second int) string {
			custom_builder.Append(ConvertDay(day))
			custom_builder.Append(" of ")
			custom_builder.Append(ConvertMonth(month))
			custom_builder.Append(" ")
			custom_builder.Append(Stringify().IntToString(year))
			custom_builder.Append(" | ")
			custom_builder.Append(ConvertTime(hour, minute, second))

			return custom_builder.Build()
		},
	}
}
