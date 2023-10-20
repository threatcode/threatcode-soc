

package config

import (
  "errors"
  "github.com/threatcode/threatcode-soc/module"
  "os"
)

const DEFAULT_POLL_INTERVAL_MS = 1000

type AgentConfig struct {
  NodeId                string                 `json:"nodeId"`
  Role                  string                 `json:"role"`
  Description           string                 `json:"description"`
  Address               string                 `json:"address"`
  Model                 string                 `json:"model"`
  ServerUrl             string                 `json:"serverUrl"`
  VerifyCert            bool                   `json:"verifyCert"`
  PollIntervalMs        int                    `json:"pollIntervalMs"`
  Modules               module.ModuleConfigMap `json:"modules"`
  ModuleFailuresIgnored bool                   `json:"moduleFailuresIgnored"`
}

func (config *AgentConfig) Verify() error {
  var err error
  if err == nil && config.PollIntervalMs <= 0 {
    config.PollIntervalMs = DEFAULT_POLL_INTERVAL_MS
  }
  if err == nil && config.NodeId == "" {
    config.NodeId, err = os.Hostname()
  }
  if err == nil && config.ServerUrl == "" {
    err = errors.New("Agent.ServerUrl configuration value is required")
  }
  return err
}
