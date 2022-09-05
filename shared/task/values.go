package task

const (
	pending            string = "pending"
	doing              string = "doing"
	done               string = "done"
	low                string = "low"
	medium             string = "medium"
	high               string = "high"
	LenLongestStatus   int    = len(pending)
	LenLongestPriority int    = len(medium)
)

var Statuses map[string]string = map[string]string{
	pending: pending,
	doing:   doing,
	done:    done,
}

var Priorities map[string]string = map[string]string{
	low:    low,
	medium: medium,
	high:   high,
}
