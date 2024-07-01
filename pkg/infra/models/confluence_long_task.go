package models

// LongTaskPageScheme represents a page of long tasks in Confluence.
type LongTaskPageScheme struct {
	Results []*LongTaskScheme `json:"results,omitempty"` // The long tasks in the page.
	Start   int               `json:"start,omitempty"`   // The start index of the long tasks in the page.
	Limit   int               `json:"limit,omitempty"`   // The limit of the long tasks in the page.
	Size    int               `json:"size,omitempty"`    // The size of the long tasks in the page.
}

// LongTaskScheme represents a long task in Confluence.
type LongTaskScheme struct {
	ID                 string                   `json:"id,omitempty"`                 // The ID of the long task.
	Name               *LongTaskNameScheme      `json:"name,omitempty"`               // The name of the long task.
	ElapsedTime        int                      `json:"elapsedTime,omitempty"`        // The elapsed time of the long task.
	PercentageComplete int                      `json:"percentageComplete,omitempty"` // The percentage completion of the long task.
	Successful         bool                     `json:"successful,omitempty"`         // Indicates if the long task is successful.
	Finished           bool                     `json:"finished,omitempty"`           // Indicates if the long task is finished.
	Messages           []*LongTaskMessageScheme `json:"messages,omitempty"`           // The messages of the long task.
	Status             string                   `json:"status,omitempty"`             // The status of the long task.
	Errors             []*LongTaskMessageScheme `json:"errors,omitempty"`             // The errors of the long task.
	AdditionalDetails  *LongTaskDetailsScheme   `json:"additionalDetails,omitempty"`  // The additional details of the long task.
}

// LongTaskNameScheme represents the name of a long task in Confluence.
type LongTaskNameScheme struct {
	Key string `json:"key,omitempty"` // The key of the name.
}

// LongTaskMessageScheme represents a message of a long task in Confluence.
type LongTaskMessageScheme struct {
	Translation string   `json:"translation,omitempty"` // The translation of the message.
	Args        []string `json:"args,omitempty"`        // The arguments of the message.
}

// LongTaskDetailsScheme represents the additional details of a long task in Confluence.
type LongTaskDetailsScheme struct {
	DestinationID        string `json:"destinationId,omitempty"`        // The ID of the destination of the long task.
	DestinationURL       string `json:"destinationUrl,omitempty"`       // The URL of the destination of the long task.
	TotalPageNeedToCopy  int    `json:"totalPageNeedToCopy,omitempty"`  // The total number of pages needed to copy for the long task.
	AdditionalProperties string `json:"additionalProperties,omitempty"` // The additional properties of the long task.
}
