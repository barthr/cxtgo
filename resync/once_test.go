package resync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnceReset(t *testing.T) {
	assert := assert.New(t)
	var calls int
	var c Once
	c.Do(func() {
		calls++
	})
	c.Do(func() {
		calls++
	})
	c.Do(func() {
		calls++
	})
	assert.Equal(calls, 1)
	c.Reset()
	c.Do(func() {
		calls++
	})
	c.Do(func() {
		calls++
	})
	c.Do(func() {
		calls++
	})
	assert.Equal(calls, 2)
}
