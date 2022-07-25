package converter

func ConvertDay(day int) string {
	var str string
	
	if day == 1 || day == 21 || day == 31 {
		str = IntToString(day) + "st"
	} else if day == 2 || day == 22 {
		str = IntToString(day) + "nd"
	} else if day == 3 || day == 23 {
		str = IntToString(day) + "rd"
	} else {
		str = IntToString(day) + "th"
	}

	return str
}

func ConvertDaysCap(month int) {
		
}