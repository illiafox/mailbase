package public

import "errors"

var Session = session{
	SessionTimoutDays: 7, // типо сессия активна 7 дне

	// Errors
	NoSession:  errors.New("Session not found!\nPlease <a href=\"/login\">login</a> or <a href=\"/register\">register</a>"),
	OldSession: errors.New("Your session is old enough!\nPlease <a href=\"/login\">login</a> again"),
}

type session struct {
	SessionTimoutDays int // Is used in old sessions checks

	// // Errors
	NoSession  error
	OldSession error
}
