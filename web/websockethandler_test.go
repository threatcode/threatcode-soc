

package web

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestHandlePingMessage(tester *testing.T) {
	webSocketHandler := NewWebSocketHandler(nil)
	conn := NewConnection(nil, "")
	oldPingTime := conn.lastPingTime
	time.Sleep(3 * time.Millisecond)
	msg := &WebSocketMessage{Kind: "Ping"}
	webSocketHandler.handleMessage(msg, conn)
	newPingTime := conn.lastPingTime
	assert.GreaterOrEqual(tester, newPingTime.Sub(oldPingTime).Milliseconds(), int64(3))
}
