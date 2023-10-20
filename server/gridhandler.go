

package server

import (
  "context"
  "errors"
  "github.com/threatcode/threatcode-soc/web"
  "net/http"
)

type GridHandler struct {
  web.BaseHandler
  server *Server
}

func NewGridHandler(srv *Server) *GridHandler {
  handler := &GridHandler{}
  handler.Host = srv.Host
  handler.server = srv
  handler.Impl = handler
  return handler
}

func (gridHandler *GridHandler) HandleNow(ctx context.Context, writer http.ResponseWriter, request *http.Request) (int, interface{}, error) {
  switch request.Method {
  case http.MethodGet:
    return gridHandler.get(ctx, writer, request)
  }
  return http.StatusMethodNotAllowed, nil, errors.New("Method not supported")
}

func (gridHandler *GridHandler) get(ctx context.Context, writer http.ResponseWriter, request *http.Request) (int, interface{}, error) {
  return http.StatusOK, gridHandler.server.Datastore.GetNodes(ctx), nil
}
