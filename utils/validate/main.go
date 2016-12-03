package validate

import "gopkg.in/bluesuncorp/validator.v9"

var Instance *validator.Validate

func New() *validator.Validate{
	Instance = validator.New()
	return Instance
}
