package result

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTernary(t *testing.T) {
	r1 := Ternary(true, 69, 420)
	assert.Equal(t, 69, r1)

	r2 := Ternary(false, 69, 420)
	assert.Equal(t, 420, r2)
}

func TestTernaryFn(t *testing.T) {
	r1 := TernaryFn(true,
		func() int { return 69 },
		func() int { return 420 },
	)
	assert.Equal(t, 69, r1)

	r2 := TernaryFn(false,
		func() int { return 69 },
		func() int { return 420 },
	)
	assert.Equal(t, 420, r2)
}

func TestTernaryResult(t *testing.T) {
	r1 := TernaryResult(true,
		Wrap(Fails()),
		Wrap(Succeeds()),
	)
	assert.False(t, r1.Valid())
	assert.Error(t, r1.Error())
	assert.Nil(t, r1.Value())
	assert.Equal(t, ErrFailed, r1.Error())

	r2 := TernaryResult(false,
		Wrap(Fails()),
		Wrap(Succeeds()),
	)
	assert.True(t, r2.Valid())
	assert.NoError(t, r2.Error())
	assert.NotNil(t, r2.Value())
	assert.Equal(t, &Fixture, r2.Value())
}
