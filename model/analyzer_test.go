

package model

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestGetModule(tester *testing.T) {
  analyzer := NewAnalyzer("id", "path")
  assert.Equal(tester, "id.id", analyzer.GetModule())
  assert.Equal(tester, "path/site-packages", analyzer.GetSitePackagesPath())
  assert.Equal(tester, "path/source-packages", analyzer.GetSourcePackagesPath())
  assert.Equal(tester, "path/requirements.txt", analyzer.GetRequirementsPath())
}
