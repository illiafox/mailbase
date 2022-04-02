package public

import "errors"

var Session = struct {
	SessionTimoutDays int // Is used in old sessions checks

	// // Errors
	NoSession  error
	OldSession error
}{

	SessionTimoutDays: 7, // active for 7 days

	NoSession: errors.New(
		"Session not found!\nPlease <a href=\"/login\">login</a> or <a href=\"/register\">register</a>",
	),
	OldSession: errors.New("Your session is old enough!\nPlease <a href=\"/login\">login</a> again"),
}
