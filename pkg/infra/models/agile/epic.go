package agile

type EpicScheme struct {
	ID      int              `json:"id,omitempty"`
	Key     string           `json:"key,omitempty"`
	Self    string           `json:"self,omitempty"`
	Name    string           `json:"name,omitempty"`
	Summary string           `json:"summary,omitempty"`
	Color   *EpicColorScheme `json:"color,omitempty"`
	Done    bool             `json:"done,omitempty"`
}
type EpicColorScheme struct {
	Key string `json:"key,omitempty"`
}
