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

var Statuses map[string]string = map[string]string{
	pending:    pending,
	inProgress: inProgress,
	done:       done,
}

var Priorities map[string]string = map[string]string{
	low:    low,
	medium: medium,
	high:   high,
}
