package task

func IsThisAStatus(s string) bool {
	for _, status := range Statuses {
		if status == s {
			return true
		}
	}

	return false
}

func IsThisAPriority(s string) bool {
	for _, p := range Priorities {
		if p == s {
			return true
		}
	}

	return false
}
