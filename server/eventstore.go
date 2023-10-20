

package server

import (
	"context"
	"github.com/threatcode/threatcode-soc/model"
)

type Eventstore interface {
	Search(context context.Context, criteria *model.EventSearchCriteria) (*model.EventSearchResults, error)
	Index(ctx context.Context, index string, document map[string]interface{}, id string) (*model.EventIndexResults, error)
	Update(context context.Context, criteria *model.EventUpdateCriteria) (*model.EventUpdateResults, error)
	Delete(context context.Context, index string, id string) error
	Acknowledge(context context.Context, criteria *model.EventAckCriteria) (*model.EventUpdateResults, error)
}
