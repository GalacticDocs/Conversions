package converter

func StringToBool(str string) bool {
  if str == "true" {
    return true
  } else if str == "false" {
    return false
  }
}
