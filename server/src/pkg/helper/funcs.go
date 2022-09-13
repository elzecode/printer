package helper

func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func PercentageChange(old, new int) (delta int64) {
	diff := float64(new - old)
	delta = int64(100 + ((diff / float64(old)) * 100))
	return
}
