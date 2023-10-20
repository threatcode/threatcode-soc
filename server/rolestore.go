package server

import (
	"context"

	"github.com/threatcode/threatcode-soc/model"
)

type Rolestore interface {
	GetAssignments(ctx context.Context) (map[string][]string, error)
	PopulateUserRoles(ctx context.Context, user *model.User) error
}
