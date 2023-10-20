

package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerifyServer(tester *testing.T) {
	cfg := &ServerConfig{}
	err := cfg.Verify()
	if assert.Error(tester, err) {
		assert.Equal(tester, DEFAULT_MAX_PACKET_COUNT, cfg.MaxPacketCount)
		assert.Equal(tester, DEFAULT_IDLE_CONNECTION_TIMEOUT_MS, cfg.IdleConnectionTimeoutMs)
		assert.Equal(tester, DEFAULT_MAX_UPLOAD_SIZE_BYTES, cfg.MaxUploadSizeBytes)
		assert.False(tester, cfg.DeveloperEnabled)
	}

	cfg.BindAddress = "http://some.where"
	cfg.MaxPacketCount = 123
	err = cfg.Verify()
	if assert.Nil(tester, err) {
		assert.Equal(tester, 123, cfg.MaxPacketCount)
		assert.Equal(tester, "/opt/threatsensor/scripts/timezones.sh", cfg.TimezoneScript)
		assert.False(tester, cfg.DeveloperEnabled)
	}
}
