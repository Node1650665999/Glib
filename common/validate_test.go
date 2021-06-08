package common_test

import (
	. "Glib/common"
	validation "github.com/go-ozzo/ozzo-validation/v3"
	"regexp"
	"testing"
)

func TestValidWithMap(t *testing.T) {
	name := "tclxxx"
	age  := 30
	valids := []ValidRule{
		{
			name,
			[]validation.Rule{validation.Required,validation.Length(5,10)},
		},
		{
			age,
			[]validation.Rule{validation.Required, validation.Min(100)},
		},
	}

	err := ValidWithMap(valids)
	t.Log(err)
}

type Address struct {
	Street string
	City   string
	State  string
	Zip    string
}

func ValidateWithStruct(a Address) error {
	return validation.ValidateStruct(&a,
		// Street cannot be empty, and the length must between 5 and 50
		validation.Field(&a.Street, validation.Required, validation.Length(5, 50)),
		// City cannot be empty, and the length must between 5 and 50
		validation.Field(&a.City, validation.Required, validation.Length(5, 50), validation.Skip),
		// State cannot be empty, and must be a string consisting of two letters in upper case
		validation.Field(&a.State, validation.Required, validation.Match(regexp.MustCompile("^[A-Z]{2}$"))),
		// State cannot be empty, and must be a string consisting of five digits
		validation.Field(&a.Zip, validation.Required, validation.Match(regexp.MustCompile("^[0-9]{5}$"))),
	)
}

func TestValidWithStruct(t *testing.T) {
	a := Address{
		Street: "123",
		City:   "Unknown",
		State:  "Virginia",
		Zip:    "12345",
	}

	err := ValidateWithStruct(a)
	t.Log(err)
}

