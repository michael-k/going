package validate_test

import (
	"testing"

	. "github.com/markbates/validate-go"
	"github.com/stretchr/testify/assert"
)

type v1 struct{}

func (v *v1) IsValid(errors *ValidationErrors) {
	errors.Add("v1", "there's an error with v1")
}

type v2 struct{}

func (v *v2) IsValid(errors *ValidationErrors) {
	errors.Add("v2", "there's an error with v2")
}

func TestValidate(t *testing.T) {
	assert := assert.New(t)
	errors := Validate(&v1{}, &v2{})
	assert.Equal(errors.Count(), 2)
	assert.Equal(errors.HasAny(), true)
	assert.Equal(errors.Errors["v1"], []string{"there's an error with v1"})
	assert.Equal(errors.Errors["v2"], []string{"there's an error with v2"})
}