

package modules

import (
	"testing"

	"github.com/threatcode/threatcode-soc/module"
	"github.com/stretchr/testify/assert"
)

func TestBuildModuleMap(tester *testing.T) {
	mm := BuildModuleMap(nil)
	findModule(tester, mm, "elastic")
	findModule(tester, mm, "elasticcases")
	findModule(tester, mm, "filedatastore")
	findModule(tester, mm, "httpcase")
	findModule(tester, mm, "kratos")
	findModule(tester, mm, "influxdb")
	findModule(tester, mm, "sostatus")
	findModule(tester, mm, "statickeyauth")
	findModule(tester, mm, "thehive")
}

func findModule(tester *testing.T, mm map[string]module.Module, module string) {
	_, ok := mm[module]
	assert.True(tester, ok)
}
