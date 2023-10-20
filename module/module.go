

package module

type ModuleConfig map[string]interface{}

type ModuleConfigMap map[string]ModuleConfig

type Module interface {
  PrerequisiteModules() []string
  Init(config ModuleConfig) error
  Start() error
  Stop() error
  IsRunning() bool
}
