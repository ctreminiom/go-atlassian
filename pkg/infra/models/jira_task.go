package models

// TaskScheme represents a task in Jira.
type TaskScheme struct {
	Self           string `json:"self"`           // The URL of the task.
	ID             string `json:"id"`             // The ID of the task.
	Description    string `json:"description"`    // The description of the task.
	Status         string `json:"status"`         // The status of the task.
	Result         string `json:"result"`         // The result of the task.
	SubmittedBy    int    `json:"submittedBy"`    // The ID of the user who submitted the task.
	Progress       int    `json:"progress"`       // The progress of the task.
	ElapsedRuntime int    `json:"elapsedRuntime"` // The elapsed runtime of the task.
	Submitted      int64  `json:"submitted"`      // The timestamp when the task was submitted.
	Started        int64  `json:"started"`        // The timestamp when the task started.
	Finished       int64  `json:"finished"`       // The timestamp when the task finished.
	LastUpdate     int64  `json:"lastUpdate"`     // The timestamp of the last update to the task.
}
