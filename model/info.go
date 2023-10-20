

package model

import (
  "github.com/threatcode/threatcode-soc/config"
)

type Info struct {
  Version        string                   `json:"version"`
  License        string                   `json:"license"`
  Parameters     *config.ClientParameters `json:"parameters"`
  ElasticVersion string                   `json:"elasticVersion"`
  WazuhVersion   string                   `json:"wazuhVersion"`
  UserId         string                   `json:"userId"`
  Timezones      []string                 `json:"timezones"`
}
