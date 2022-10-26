package utils

func Contains[E comparable](list []E, item E) bool {
	for _, el := range list {
		if el == item {
			return true
		}
	}

	return false
}
