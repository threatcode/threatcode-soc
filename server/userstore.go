

package server

import (
  "context"
  "github.com/threatcode/threatcode-soc/model"
)

type Userstore interface {
  GetUsers(ctx context.Context) ([]*model.User, error)
  DeleteUser(ctx context.Context, id string) error
  GetUser(ctx context.Context, id string) (*model.User, error)
  UpdateUser(ctx context.Context, id string, user *model.User) error
}
