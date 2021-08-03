package validation

import (
	"regexp"
	"strconv"

	"salon-booking-guru/store"
)

var s store.Store

type Validation struct {
	Errors []string `json:"errors"`
}

type ValidationFunc func(interface{}) Validation

func Init(st store.Store) {
	s = st
}

func (v *Validation) IsValid() bool {
	if len(v.Errors) == 0 {
		return true
	} else {
		return false
	}
}

func (v *Validation) IsEmail(field string, email string) {
	re := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !re.MatchString(email) || len(email) < 5 {
		v.Errors = append(v.Errors, field+" - Must be a valid email address")
	}
}

func (v *Validation) isIpAddress(field string, ipAddress string) {
	re := regexp.MustCompile(`^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$`)
	if !re.MatchString(ipAddress) {
		v.Errors = append(v.Errors, field+" - Must be a valid IP address")
	}
}

func (v *Validation) isSelected(field string, id int) {
	if id == 0 {
		v.Errors = append(v.Errors, field+" - Must be selected")
	}
}

func (v *Validation) maxLength(field string, text string, limit int) {
	if len(text) > limit {
		v.Errors = append(v.Errors, field+" - Maximum length is "+strconv.Itoa(limit))
	}
}

func (v *Validation) minLength(field string, text string, limit int) {
	if len(text) < limit {
		v.Errors = append(v.Errors, field+" - Minimum length is "+strconv.Itoa(limit))
	}
}

func (v *Validation) withinRange(field string, number int, min int, max int) {
	if number < min || number > max {
		v.Errors = append(v.Errors, field+" - Must be between "+strconv.Itoa(min)+" - "+strconv.Itoa(max))
	}
}

func (v *Validation) AddError(text string) {
	v.Errors = append(v.Errors, text)
}

func (v *Validation) passwordsMatch(password string, confirmPassword string) {
	if password != confirmPassword {
		v.Errors = append(v.Errors, "Passwords do not match")
	}
}

func (v *Validation) IsValidPassword(password string) {
	v.minLength("Password", password, 8)
}
