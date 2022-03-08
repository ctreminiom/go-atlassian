package models

type LongTaskPageScheme struct {
	Results []*LongTaskScheme `json:"results,omitempty"`
	Start   int               `json:"start,omitempty"`
	Limit   int               `json:"limit,omitempty"`
	Size    int               `json:"size,omitempty"`
}

type LongTaskScheme struct {
	ID                 string                   `json:"id,omitempty"`
	Name               *LongTaskNameScheme      `json:"name,omitempty"`
	ElapsedTime        int                      `json:"elapsedTime,omitempty"`
	PercentageComplete int                      `json:"percentageComplete,omitempty"`
	Successful         bool                     `json:"successful,omitempty"`
	Finished           bool                     `json:"finished,omitempty"`
	Messages           []*LongTaskMessageScheme `json:"messages,omitempty"`
	Status             string                   `json:"status,omitempty"`
	Errors             []*LongTaskMessageScheme `json:"errors,omitempty"`
	AdditionalDetails  *LongTaskDetailsScheme   `json:"additionalDetails,omitempty"`
}

type LongTaskNameScheme struct {
	Key string `json:"key,omitempty"`
}

type LongTaskMessageScheme struct {
	Translation string   `json:"translation,omitempty"`
	Args        []string `json:"args,omitempty"`
}

type LongTaskDetailsScheme struct {
	DestinationID        string `json:"destinationId,omitempty"`
	DestinationURL       string `json:"destinationUrl,omitempty"`
	TotalPageNeedToCopy  int    `json:"totalPageNeedToCopy,omitempty"`
	AdditionalProperties string `json:"additionalProperties,omitempty"`
}
