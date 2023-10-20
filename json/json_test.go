

package json

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJson(tester *testing.T) {
	testFile := "/tmp/threatsensor_test.json"
	defer os.Remove(testFile)
	obj := make(map[string]string)
	obj["MyKey"] = "MyValue"
	err := WriteJsonFile(testFile, obj)
	assert.Nil(tester, err)
	obj = make(map[string]string)
	err = LoadJsonFile(testFile, &obj)
	if assert.Nil(tester, err) {
		assert.Equal(tester, "MyValue", obj["MyKey"])
	}
}
