package chefondate

func ChefOnDate(x, y int) (result string) {
	if x <= 0 || y <= 0 {
		return ""
	}

	if x >= y {
		return "YES"
	}

	return "NO"
}
