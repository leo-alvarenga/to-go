package task

func IsThisAStatus(s string) bool {
	for _, status := range ValidStatuses {
		if status == s {
			return true
		}
	}

	return false
}

func IsThisAPriority(s string) bool {
	for _, p := range ValidPriorities {
		if p == s {
			return true
		}
	}

	return false
}
