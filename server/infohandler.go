

package server

import (
  "context"
  "errors"
  "github.com/threatcode/threatcode-soc/model"
  "github.com/threatcode/threatcode-soc/web"
  "net/http"
  "os"
)

type InfoHandler struct {
  web.BaseHandler
  server    *Server
  timezones []string
}

func NewInfoHandler(srv *Server) *InfoHandler {
  handler := &InfoHandler{}
  handler.Host = srv.Host
  handler.server = srv
  handler.Impl = handler
  handler.timezones = srv.GetTimezones()
  return handler
}

func (infoHandler *InfoHandler) HandleNow(ctx context.Context, writer http.ResponseWriter, request *http.Request) (int, interface{}, error) {
  switch request.Method {
  case http.MethodGet:
    return infoHandler.get(ctx, writer, request)
  }
  return http.StatusMethodNotAllowed, nil, errors.New("Method not supported")
}

func (infoHandler *InfoHandler) get(ctx context.Context, writer http.ResponseWriter, request *http.Request) (int, interface{}, error) {
  var err error
  var info *model.Info
  if user, ok := request.Context().Value(web.ContextKeyRequestor).(*model.User); ok {
    info = &model.Info{
      Version:        infoHandler.Host.Version,
      License:        "GPL v2",
      Parameters:     &infoHandler.server.Config.ClientParams,
      ElasticVersion: os.Getenv("ELASTIC_VERSION"),
      WazuhVersion:   os.Getenv("WAZUH_VERSION"),
      UserId:         user.Id,
      Timezones:      infoHandler.timezones,
    }
  } else {
    err = errors.New("Unable to determine logged in user from context")
  }
  return http.StatusOK, info, err
}
