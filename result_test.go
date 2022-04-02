package result

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Example struct {
	Name string
}

var ErrFailed = errors.New("I have failed you")

var Fixture = Example{"Óðinn"}

func Fails() (*Example, error) {
	return nil, ErrFailed
}

func Succeeds() (*Example, error) {
	return &Fixture, nil
}

func TestResultWrapError(t *testing.T) {
	r := Wrap(Fails())

	assert.False(t, r.Valid())
	assert.Error(t, r.Error())
	assert.Nil(t, r.Value())
	assert.Equal(t, ErrFailed, r.Error())
}

func TestResultWrapSuccess(t *testing.T) {
	r := Wrap(Succeeds())

	assert.True(t, r.Valid())
	assert.NoError(t, r.Error())
	assert.NotNil(t, r.Value())
	assert.Equal(t, &Fixture, r.Value())
}

func TestResultUnwrapError(t *testing.T) {
	r, err := Unwrap(Wrap(Fails()))

	assert.Error(t, err)
	assert.Nil(t, r)
	assert.Equal(t, ErrFailed, err)
}

func TestResultUnwrapSuccess(t *testing.T) {
	r, err := Unwrap(Wrap(Succeeds()))

	assert.NoError(t, err)
	assert.NotNil(t, r)
	assert.Equal(t, &Fixture, r)
}
