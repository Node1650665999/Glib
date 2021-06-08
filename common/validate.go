package common

import (
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v3"
)

type ValidRule struct {
	Data  interface{}
	Rules []validation.Rule
}

func ValidWithMap(sv []ValidRule) error {
	for _, valid := range sv {
		err := validation.Validate(valid.Data, valid.Rules...)
		if err != nil {
			return fmt.Errorf("%v %v", valid.Data,err)
		}
	}
	return nil
}


