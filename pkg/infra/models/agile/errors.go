package agile

import "errors"

var (
	ErrNoBoardIDError  = errors.New("agile: no board id set")
	ErrNoFilterIDError = errors.New("agile: no filter id set")
	ErrNoEpicIDError   = errors.New("agile: no epic id set")
	ErrNoSprintIDError = errors.New("agile: no sprint id set")
)
