

package statickeyauth

import (
	"testing"

	"github.com/threatcode/threatcode-soc/config"
	"github.com/threatcode/threatcode-soc/module"
	"github.com/threatcode/threatcode-soc/server"
	"github.com/stretchr/testify/assert"
)

func TestAuthInit(tester *testing.T) {
	scfg := &config.ServerConfig{}
	srv := server.NewServer(scfg, "")
	auth := NewStaticKeyAuth(srv)
	cfg := make(module.ModuleConfig)

	authInit(tester, auth, cfg, true, "")

	cfg["apiKey"] = "abc"
	authInit(tester, auth, cfg, true, "")

	expectedCidr := "172.17.0.0/24"
	cfg["anonymousCidr"] = expectedCidr
	authInit(tester, auth, cfg, false, expectedCidr)
}

func authInit(tester *testing.T, auth *StaticKeyAuth, cfg module.ModuleConfig, failure bool, expectedCidr string) {
	assert.Len(tester, auth.server.Host.Preprocessors(), 1)
	err := auth.Init(cfg)
	if failure {
		assert.Error(tester, err, "Expected Init error")
	} else {
		if assert.Nil(tester, err) {
			assert.Equal(tester, expectedCidr, auth.impl.anonymousNetwork.String())
			assert.Len(tester, auth.server.Host.Preprocessors(), 2)
		}
	}
}
