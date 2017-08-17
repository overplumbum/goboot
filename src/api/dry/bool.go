package dry

func BoolToIntStr(val bool) string {
	if val {
		return "1"
	} else {
		return "0"
	}
}
