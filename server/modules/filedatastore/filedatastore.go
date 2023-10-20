

package filedatastore

import (
  "github.com/threatcode/threatcode-soc/module"
  "github.com/threatcode/threatcode-soc/server"
)

type FileDatastore struct {
  config module.ModuleConfig
  server *server.Server
  impl   *FileDatastoreImpl
}

func NewFileDatastore(srv *server.Server) *FileDatastore {
  return &FileDatastore{
    server: srv,
    impl:   NewFileDatastoreImpl(srv),
  }
}

func (fdmodule *FileDatastore) PrerequisiteModules() []string {
  return nil
}

func (fdmodule *FileDatastore) Init(cfg module.ModuleConfig) error {
  fdmodule.config = cfg
  err := fdmodule.impl.Init(cfg)
  if err == nil {
    fdmodule.server.Datastore = fdmodule.impl
  }
  return err
}

func (fdmodule *FileDatastore) Start() error {
  return nil
}

func (fdmodule *FileDatastore) Stop() error {
  return nil
}

func (fdmodule *FileDatastore) IsRunning() bool {
  return false
}
