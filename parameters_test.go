package cxtgo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParams_GetParameters(t *testing.T) {
	assert := assert.New(t)

	params := Params{
		"testInt":     1,
		"testInt32":   int32(1),
		"testInt64":   int64(1),
		"testString":  "test",
		"testFloat32": float32(0.1),
		"testFloat64": float64(0.1),
	}

	val, ok := params.GetInt("testInt")
	assert.True(ok)
	assert.Equal(1, val)

	val32, ok := params.GetInt32("testInt32")
	assert.True(ok)
	assert.Equal(int32(1), val32)

	val64, ok := params.GetInt64("testInt64")
	assert.True(ok)
	assert.Equal(int64(1), val64)

	sValue, ok := params.GetString("testString")
	assert.True(ok)
	assert.Equal("test", sValue)

	fval32, ok := params.GetFloat32("testFloat32")
	assert.True(ok)
	assert.Equal(float32(0.1), fval32)

	fval64, ok := params.GetFloat64("testFloat64")
	assert.True(ok)
	assert.Equal(float64(0.1), fval64)

	// invalid
	zeroValue, ok := params.GetFloat64("non-existing")
	assert.False(ok)
	assert.Equal(float64(0), zeroValue)

	// invalid type
	invalid, ok := params.GetString("testInt64")
	assert.False(ok)
	assert.Equal("", invalid)

}
