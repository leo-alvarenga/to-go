package task

type Task struct {
	Id          string
	Title       string
	Description string
	Priority    string
	Status      string
}

/*
 - title (required)
 - createdIn (autoset)
 - description (default="...")
 - priority (default=low)
 - dueTo (default=unset)
*/
