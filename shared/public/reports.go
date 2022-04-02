package public

import "errors"

var Reports = struct {
	StillProcessing error

	IsChecked error
	NotFound  error
}{
	StillProcessing: errors.New("your previous report is still processed"),
	IsChecked:       errors.New("report is already checked"),
	NotFound:        errors.New("report not found"),
}
