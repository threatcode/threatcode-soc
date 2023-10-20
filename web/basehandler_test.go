/*
 */
package web

import (
	"errors"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestHandler struct {
	BaseHandler
}

func NewTestHandler() *TestHandler {
	handler := &TestHandler{}
	return handler
}

func TestGetPathParameter(tester *testing.T) {
	handler := NewTestHandler()
	var testTable = []struct {
		path     string
		index    int
		expected string
	}{
		{"", -1, ""},
		{"", 0, ""},
		{"", 1, ""},
		{"/", -1, ""},
		{"/", 0, ""},
		{"/", 1, ""},
		{"/123", -1, ""},
		{"/123", 0, "123"},
		{"/123", 1, ""},
		{"/123/", 0, "123"},
		{"/123/", 1, ""},
		{"/123/456", 0, "123"},
		{"/123/456", 1, "456"},
	}

	for _, test := range testTable {
		tester.Run("path="+test.path+", index="+strconv.Itoa(test.index), func(t *testing.T) {
			actual := handler.GetPathParameter(test.path, test.index)
			assert.Equal(tester, test.expected, actual)
		})
	}
}

func TestConvertErrorToSafeString(tester *testing.T) {
	handler := NewTestHandler()

	assert.Equal(tester, "ERROR_FOO", handler.convertErrorToSafeString(errors.New("ERROR_FOO")))
	assert.Equal(tester, GENERIC_ERROR_MESSAGE, handler.convertErrorToSafeString(errors.New("ERROR2_FOO")))
}
