

package web

import (
  "context"
  "github.com/stretchr/testify/assert"
  "net/http"
  "testing"
)

func TestPreprocessPriority(tester *testing.T) {
  handler := NewBasePreprocessor()
  assert.Zero(tester, handler.PreprocessPriority())
}

func TestPreprocess(tester *testing.T) {
  handler := NewBasePreprocessor()
  request, _ := http.NewRequest("GET", "", nil)
  ctx, statusCode, err := handler.Preprocess(context.Background(), request)
  assert.NoError(tester, err)
  assert.Zero(tester, statusCode)

  actualId := ctx.Value(ContextKeyRequestId).(string)
  assert.Len(tester, actualId, 36, "Expected valid UUID")
}
