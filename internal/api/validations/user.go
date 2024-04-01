package validations

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

type EmailValidator struct{}

func (v EmailValidator) Validate(input any) error {
	email, ok := input.(string)
	if !ok {
		return errors.New("email is not a string")
	}

	if !isValidaEmail(email) {
		return fmt.Errorf("%s is not a valid email", email)
	}

	return nil
}

func isValidaEmail(email string) bool {
	emailPattern := "^[\\w.!#$%&'*+\\/=?^_`{|}~-]+@[\\w-]+(?:\\.[\\w-]+)+$"
	re := regexp.MustCompile(emailPattern)

	return re.MatchString(email)
}

type UsernameValidator struct{}

func (v UsernameValidator) Validate(input any) error {
	username, ok := input.(string)
	if !ok {
		return errors.New("username is not a string")
	}

	if !isValidaUsername(username) {
		return fmt.Errorf("%s is not a valid username", username)
	}

	return nil
}

func isValidaUsername(username string) bool {
	usernamePattern := `^[a-zA-Z0-9]{5,32}$`
	re := regexp.MustCompile(usernamePattern)

	return re.MatchString(username)
}

type PasswordValidator struct{}

func (v PasswordValidator) Validate(input any) error {
	password, ok := input.(string)
	if !ok {
		return errors.New("password is not string")
	}

	if !isValidPassword(password) {
		return fmt.Errorf("%s is not a valid password", password)
	}

	return nil
}

func isValidPassword(password string) bool {
	if len(password) < 8 || len(password) > 64 {
		return false
	}
	return true
}

type CodeValidator struct{}

func (v CodeValidator) Validate(input any) error {
	code, ok := input.(string)
	if !ok {
		return errors.New("code is not string")
	}

	if !isValidCode(code) {
		return errors.New("invalid code")
	}

	return nil
}

func isValidCode(code string) bool {
	if len(code) != 6 {
		return false
	}
	_, err := strconv.Atoi(code)
	if err != nil {
		return false
	}

	return true
}
