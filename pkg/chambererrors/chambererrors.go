package chambererrors

import "fmt"

var (
	BadConfigError       = ChamberError{StatusCode: 1, Message: "Error Reading Config File. Check Syntax"}
	UnsupportedHostError = ChamberError{StatusCode: 2, Message: "Unsupported Template Repo Host"}
	DownloadError        = ChamberError{StatusCode: 3, Message: "Download Error"}
)

type ChamberError struct {
	StatusCode int
	Message    string
}

func (e *ChamberError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.StatusCode, e.Message)
}
