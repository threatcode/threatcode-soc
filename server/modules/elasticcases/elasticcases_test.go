

package elasticcases

import (
	"testing"

	"github.com/threatcode/threatcode-soc/module"
	"github.com/threatcode/threatcode-soc/server"
	"github.com/stretchr/testify/assert"
)

func TestElasticCasesInit(tester *testing.T) {
	somodule := NewElasticCases(server.NewFakeUnauthorizedServer())
	cfg := make(module.ModuleConfig)
	err := somodule.Init(cfg)
	assert.Nil(tester, err)

	// Fail if casestore already initialized
	err = somodule.Init(cfg)
	assert.Error(tester, err)
}
