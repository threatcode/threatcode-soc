

package server

import (
	"github.com/threatcode/threatcode-soc/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewServer(tester *testing.T) {
	cfg := &config.ServerConfig{}
	srv := NewServer(cfg, "")
	assert.NotNil(tester, srv.Host)
	assert.NotNil(tester, srv.stoppedChan)
}

func TestDeveloperAuthorization(tester *testing.T) {
	cfg := &config.ServerConfig{}
	srv := NewServer(cfg, "")
	cfg.DeveloperEnabled = true

	authErr := srv.CheckAuthorized(nil, "read", "users")
	assert.NoError(tester, authErr)
}

func TestMissingAuthorization(tester *testing.T) {
	cfg := &config.ServerConfig{}
	srv := NewServer(cfg, "")

	authErr := srv.CheckAuthorized(nil, "read", "users")
	assert.Error(tester, authErr)
	assert.Equal(tester, "Missing Authorizer module", authErr.Error())
}

func TestFailedAuthorization(tester *testing.T) {
	srv := NewFakeUnauthorizedServer()

	authErr := srv.CheckAuthorized(nil, "read", "users")
	assert.Error(tester, authErr)
	assert.Contains(tester, authErr.Error(), "not authorized to perform operation 'read' on target 'users'")
}
