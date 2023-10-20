

package server

import (
  "context"
  "errors"
  "github.com/threatcode/threatcode-soc/model"
  "github.com/threatcode/threatcode-soc/web"
  "net/http"
)

type NodeHandler struct {
  web.BaseHandler
  server *Server
}

func NewNodeHandler(srv *Server) *NodeHandler {
  handler := &NodeHandler{}
  handler.Host = srv.Host
  handler.server = srv
  handler.Impl = handler
  return handler
}

func (nodeHandler *NodeHandler) HandleNow(ctx context.Context, writer http.ResponseWriter, request *http.Request) (int, interface{}, error) {
  switch request.Method {
  case http.MethodPost:
    return nodeHandler.post(ctx, writer, request)
  }
  return http.StatusMethodNotAllowed, nil, errors.New("Method not supported")
}

func (nodeHandler *NodeHandler) post(ctx context.Context, writer http.ResponseWriter, request *http.Request) (int, interface{}, error) {
  var job *model.Job
  node := model.NewNode("")
  err := nodeHandler.ReadJson(request, node)
  if err == nil {
    node, err = nodeHandler.server.Datastore.UpdateNode(ctx, node)
    if err == nil {
      nodeHandler.server.Metrics.UpdateNodeMetrics(ctx, node)
      nodeHandler.Host.Broadcast("node", node)
      job = nodeHandler.server.Datastore.GetNextJob(ctx, node.Id)
    }
  }
  return http.StatusOK, job, err
}
