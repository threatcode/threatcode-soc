

package module

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeetsPrerequisites(tester *testing.T) {
	mgr := NewModuleManager()
	mcm := make(ModuleConfigMap)

	prereqs := make([]string, 0)
	prereqs = append(prereqs, "foo")
	prereqs = append(prereqs, "bar")

	actual := mgr.meetsPrerequisites(prereqs, mcm)
	assert.False(tester, actual)

	mcm["foo"] = make(ModuleConfig)
	actual = mgr.meetsPrerequisites(prereqs, mcm)
	assert.False(tester, actual)

	mcm["bar"] = make(ModuleConfig)
	actual = mgr.meetsPrerequisites(prereqs, mcm)
	assert.True(tester, actual)
}
