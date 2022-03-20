package public

import (
	"errors"
	"fmt"
)

var Register = register{

	PasswordMin: 8,
	PasswordMax: 128,

	InvalidLength: errors.New(fmt.Sprintf("Invalid password format: min length %d and max %d", 8, 128)),
	InvalidFormat: errors.New(fmt.Sprintf("Invalid password format: At least one number, lower/upper letter with min length %d and max %d", 8, 128)),
	KeyNotFound:   errors.New("key not found"),
}

type register struct {
	PasswordMin int
	PasswordMax int

	InvalidLength error
	InvalidFormat error
	KeyNotFound   error
}
