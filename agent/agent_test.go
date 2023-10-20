

package agent

import (
	"testing"

	"github.com/threatcode/threatcode-soc/config"
	"github.com/stretchr/testify/assert"
)

func TestNewAgent(tester *testing.T) {
	cfg := &config.AgentConfig{}
	cfg.ServerUrl = "http://some.where"
	agent := NewAgent(cfg, "")
	assert.NotNil(tester, agent.Client)
	assert.NotNil(tester, agent.JobMgr)
	assert.NotNil(tester, agent.stoppedChan)
}
