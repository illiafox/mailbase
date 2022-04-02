package public

import (
	"errors"
	"fmt"
)

var Register = struct {
	PasswordMin int
	PasswordMax int

	InvalidLength error
	InvalidFormat error
	KeyNotFound   error
}{
	PasswordMin: 8,
	PasswordMax: 128,

	InvalidLength: fmt.Errorf("invalid password format: min length %d and max %d", 8, 128),
	InvalidFormat: fmt.Errorf(
		"invalid password format: At least one number, lower/upper letter with min length %d and max %d", 8, 128,
	),
	KeyNotFound: errors.New("key not found or expired"),
}
