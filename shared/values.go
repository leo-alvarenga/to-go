package shared

// The name of the YAML file used for storing configs -> TODO
const configFile string = ""

const (
	pending            string = "pending"
	inProgress         string = "doing"
	done               string = "done"
	low                string = "low"
	medium             string = "medium"
	high               string = "high"
	LenLongestStatus   int    = len(pending)
	LenLongestPriority int    = len(medium)
)

var ValidStatuses map[string]string = map[string]string{
	pending:    pending,
	inProgress: inProgress,
	done:       done,
}

var ValidPriorities map[string]string = map[string]string{
	low:    low,
	medium: medium,
	high:   high,
}

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
