

package thehive

import (
	"testing"

	"github.com/threatcode/threatcode-soc/module"
	"github.com/threatcode/threatcode-soc/server"
	"github.com/stretchr/testify/assert"
)

func TestTheHiveInit(tester *testing.T) {
	thehive := NewTheHive(server.NewFakeUnauthorizedServer())
	cfg := make(module.ModuleConfig)
	err := thehive.Init(cfg)
	assert.Nil(tester, err)

	// Fail if casestore already initialized
	err = thehive.Init(cfg)
	assert.Error(tester, err)
}
