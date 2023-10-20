

package web

import (
	"context"
	"github.com/google/uuid"
	"net/http"
)

type ContextKey string

const ContextKeyRequestId = ContextKey("ContextKeyRequestId")
const ContextKeyRequestorId = ContextKey("ContextKeyRequestorId")
const ContextKeyRequestor = ContextKey("ContextKeyRequestor")

type BasePreprocessor struct {
}

func NewBasePreprocessor() *BasePreprocessor {
	return &BasePreprocessor{}
}

func (Processor *BasePreprocessor) PreprocessPriority() int {
	return 0
}

func (processor *BasePreprocessor) Preprocess(ctx context.Context, req *http.Request) (context.Context, int, error) {
	uuid := uuid.New().String()
	ctx = context.WithValue(ctx, ContextKeyRequestId, uuid)
	return ctx, 0, nil
}
