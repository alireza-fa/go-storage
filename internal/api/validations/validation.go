package validations

type Validator interface {
	Validate(input any) error
}

var Validators map[string]Validator = map[string]Validator{
	"email":    EmailValidator{},
	"username": UsernameValidator{},
	"password": PasswordValidator{},
	"code":     CodeValidator{},
}
