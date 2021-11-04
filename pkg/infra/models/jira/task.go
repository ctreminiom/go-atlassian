package jira

type TaskScheme struct {
	Self           string `json:"self"`
	ID             string `json:"id"`
	Description    string `json:"description"`
	Status         string `json:"status"`
	Result         string `json:"result"`
	SubmittedBy    int    `json:"submittedBy"`
	Progress       int    `json:"progress"`
	ElapsedRuntime int    `json:"elapsedRuntime"`
	Submitted      int64  `json:"submitted"`
	Started        int64  `json:"started"`
	Finished       int64  `json:"finished"`
	LastUpdate     int64  `json:"lastUpdate"`
}
