package task

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
