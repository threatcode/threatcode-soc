

package web

import (
  "github.com/stretchr/testify/assert"
  "testing"
)

func TestFormatUrl(tester *testing.T) {
  client := NewClient("http://some.where/path", true)
  var testTable = []struct {
    url      string
    path     string
    expected string
  }{
    {"http://far.out", "path", "http://far.out/path"},
    {"http://far.out", "/path", "http://far.out/path"},
    {"http://far.out/", "path", "http://far.out/path"},
    {"http://far.out/", "/path", "http://far.out/path"},
    {"http://far.out/", "/path/end", "http://far.out/path/end"},
  }

  for _, test := range testTable {
    tester.Run("url="+test.url+", path="+test.path, func(t *testing.T) {
      actual := client.FormatUrl(test.url, test.path)
      assert.Equal(tester, test.expected, actual)
    })
  }
}

type TestObject struct {
  Foo string
}

func TestMock(tester *testing.T) {
  client := NewClient("http://some.where/path", true)
  respObj := &TestObject{}
  respBody := `{"foo": "bar"}`
  client.MockStringResponse(respBody, 200, nil)
  client.SendObject("GET", "subpath", nil, respObj, false)
  assert.Equal(tester, "bar", respObj.Foo)
}
