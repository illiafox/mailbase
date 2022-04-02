package public

import "errors"

var Admin = struct {
	NoRights    error
	NotSuper    error
	WrongFormat error

	NotAdmin     error
	AdminAlready error
	EditSuper    error

	ReportsNotFound error

	IDNotFound error
}{
	NoRights: errors.New("you are not admin<br>ask <a href='https://t.me/ebashu_gerych'>owners</a> to become one"),

	NotSuper: errors.New("only super admin can use this"),

	WrongFormat: errors.New("wrong field format"),

	ReportsNotFound: errors.New("reports not found"),

	AdminAlready: errors.New("user is already admin"),
	NotAdmin:     errors.New("user is not admin"),
	EditSuper:    errors.New("managing super accounts is forbidden"),
	IDNotFound:   errors.New("id not found"),
}
