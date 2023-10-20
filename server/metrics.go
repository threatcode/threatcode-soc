

package server

import (
	"context"
	"github.com/threatcode/threatcode-soc/model"
)

type Metrics interface {
	GetGridEps(ctx context.Context) int
	UpdateNodeMetrics(ctx context.Context, node *model.Node) bool
}
