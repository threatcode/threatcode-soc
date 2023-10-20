

package web

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestIsAuthorized(tester *testing.T) {
	conn := NewConnection(nil, "")
	assert.True(tester, conn.IsAuthorized("test"))
}

func TestUpdatePingTime(tester *testing.T) {
	conn := NewConnection(nil, "")
	oldPingTime := conn.lastPingTime
	time.Sleep(3 * time.Millisecond)
	conn.UpdatePingTime()
	newPingTime := conn.lastPingTime

	assert.True(tester, newPingTime.Sub(oldPingTime).Milliseconds() >= 3, "expected 3s increase in lastPingTime")
}
