

package kratos

import (
	"context"
	"github.com/threatcode/threatcode-soc/web"
	"net/http"
)

type KratosPreprocessor struct {
	userstore *KratosUserstore
}

func NewKratosPreprocessor(impl *KratosUserstore) *KratosPreprocessor {
	return &KratosPreprocessor{
		userstore: impl,
	}
}

func (proc *KratosPreprocessor) PreprocessPriority() int {
	return 110
}

func (proc *KratosPreprocessor) Preprocess(ctx context.Context, request *http.Request) (context.Context, int, error) {
	var statusCode int
	var err error

	userId := request.Header.Get("x-user-id")
	if userId != "" {
		ctx = context.WithValue(ctx, web.ContextKeyRequestorId, userId)
		user, err := proc.userstore.GetUser(ctx, userId)
		if err == nil {
			ctx = context.WithValue(ctx, web.ContextKeyRequestor, user)
		}
	}

	return ctx, statusCode, err
}
