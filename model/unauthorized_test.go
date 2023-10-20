

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnauthorized(tester *testing.T) {
	event := NewUnauthorized("mysubject", "myop", "mytarget")
	assert.NotZero(tester, event.CreateTime)
	assert.Equal(tester, "mysubject", event.Subject)
	assert.Equal(tester, "myop", event.Operation)
	assert.Equal(tester, "mytarget", event.Target)
}
