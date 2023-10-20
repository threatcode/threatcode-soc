

package server

import (
  "context"
  "github.com/threatcode/threatcode-soc/model"
  "io"
)

type Datastore interface {
  CreateNode(ctx context.Context, id string) *model.Node
  GetNodes(ctx context.Context) []*model.Node
  AddNode(ctx context.Context, node *model.Node) error
  UpdateNode(ctx context.Context, newNode *model.Node) (*model.Node, error)
  GetNextJob(ctx context.Context, nodeId string) *model.Job
  CreateJob(ctx context.Context) *model.Job
  GetJob(ctx context.Context, jobId int) *model.Job
  GetJobs(ctx context.Context, kind string, parameters map[string]interface{}) []*model.Job
  AddJob(ctx context.Context, job *model.Job) error
  AddPivotJob(ctx context.Context, job *model.Job) error
  UpdateJob(ctx context.Context, job *model.Job) error
  DeleteJob(ctx context.Context, jobId int) (*model.Job, error)
  GetPackets(ctx context.Context, jobId int, offset int, count int, unwrap bool) ([]*model.Packet, error)
  SavePacketStream(ctx context.Context, jobId int, reader io.ReadCloser) error
  GetPacketStream(ctx context.Context, jobId int, unwrap bool) (io.ReadCloser, string, int64, error)
}
