package methodmethod

func GetMessage(s string) string {
	if len(s) > 5 {
		return p1() + s + (", what a long name.")
	} else {
		return p1() + s + (", what a short name.")
	}
}

func p1() string {
	return "Hello "
}