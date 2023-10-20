

package server

import (
  "context"
  "github.com/threatcode/threatcode-soc/config"
  "github.com/threatcode/threatcode-soc/model"
)

type FakeAuthorizer struct {
  authorized bool
}

func (fake FakeAuthorizer) CheckContextOperationAuthorized(ctx context.Context, operation string, target string) error {
  if fake.authorized {
    return nil
  }
  return model.NewUnauthorized("fake-subject", operation, target)
}

type FakeRolestore struct {
  roleMap map[string][]string
}

func (impl *FakeRolestore) GetAssignments(ctx context.Context) (map[string][]string, error) {
  return impl.roleMap, nil
}

func (impl *FakeRolestore) PopulateUserRoles(ctx context.Context, user *model.User) error {
  user.Roles = impl.roleMap[user.Email]
  return nil
}

func NewFakeServer(authorized bool, roleMap map[string][]string) *Server {
  cfg := &config.ServerConfig{}
  srv := NewServer(cfg, "")
  srv.Authorizer = &FakeAuthorizer{
    authorized: authorized,
  }
  srv.Rolestore = &FakeRolestore{
    roleMap: roleMap,
  }
  return srv
}

func NewFakeAuthorizedServer(roleMap map[string][]string) *Server {
  return NewFakeServer(true, roleMap)
}

func NewFakeUnauthorizedServer() *Server {
  return NewFakeServer(false, make(map[string][]string))
}
