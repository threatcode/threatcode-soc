

package statickeyauth

import (
	"github.com/threatcode/threatcode-soc/module"
	"github.com/threatcode/threatcode-soc/server"
)

type StaticKeyAuth struct {
	config module.ModuleConfig
	server *server.Server
	impl   *StaticKeyAuthImpl
}

func NewStaticKeyAuth(srv *server.Server) *StaticKeyAuth {
	return &StaticKeyAuth{
		server: srv,
		impl:   NewStaticKeyAuthImpl(srv),
	}
}

func (skmodule *StaticKeyAuth) PrerequisiteModules() []string {
	return nil
}

func (skmodule *StaticKeyAuth) Init(cfg module.ModuleConfig) error {
	skmodule.config = cfg
	key, err := module.GetString(cfg, "apiKey")
	if err == nil {
		var anonymousCidr string
		anonymousCidr, err = module.GetString(cfg, "anonymousCidr")
		if err == nil {
			err = skmodule.impl.Init(key, anonymousCidr)
			if err == nil {
				err = skmodule.server.Host.AddPreprocessor(skmodule.impl)
			}
		}
	}
	return err
}

func (skmodule *StaticKeyAuth) Start() error {
	return nil
}

func (skmodule *StaticKeyAuth) Stop() error {
	return nil
}

func (skmodule *StaticKeyAuth) IsRunning() bool {
	return false
}
