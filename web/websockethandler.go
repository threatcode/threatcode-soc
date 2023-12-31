package web

import (
	"context"
	"errors"
	"net/http"

	"github.com/apex/log"
	"github.com/gorilla/websocket"
	"github.com/threatcode/threatcode-soc/json"
)

type WebSocketHandler struct {
	BaseHandler
}

func NewWebSocketHandler(host *Host) *WebSocketHandler {
	handler := &WebSocketHandler{}
	handler.Host = host
	handler.Impl = handler
	return handler
}

func (webSocketHandler *WebSocketHandler) HandleNow(ctx context.Context, writer http.ResponseWriter, request *http.Request) (int, interface{}, error) {
	upgrader := websocket.Upgrader{}
	connection, err := upgrader.Upgrade(writer, request, nil)
	ip := webSocketHandler.Host.GetSourceIp(request)
	if err != nil {
		log.WithError(err).WithFields(log.Fields{
			"remoteAddr": request.RemoteAddr,
			"sourceIp":   ip,
			"path":       request.URL.Path,
		}).Warn("Failed to upgrade websocket")
		return http.StatusBadRequest, nil, errors.New("Unable to upgrade request to websocket")
	}

	log.WithFields(log.Fields{
		"remoteAddr": request.RemoteAddr,
		"sourceIp":   ip,
		"path":       request.URL.Path,
	}).Info("WebSocket connected")
	conn := webSocketHandler.Host.AddConnection(connection, ip)

	defer connection.Close()
	for {
		messageType, messageBytes, err := connection.ReadMessage()
		if err != nil {
			break
		}
		log.WithFields(log.Fields{
			"remoteAddr": request.RemoteAddr,
			"sourceIp":   ip,
			"path":       request.URL.Path,
			"msg":        string(messageBytes),
			"type":       messageType,
		}).Info("WebSocket message received")

		msg := &WebSocketMessage{}
		json.LoadJson(messageBytes, msg)
		webSocketHandler.handleMessage(msg, conn)
	}
	log.WithFields(log.Fields{
		"remoteAddr": request.RemoteAddr,
		"sourceIp":   ip,
		"path":       request.URL.Path,
	}).Info("WebSocket disconnected")
	webSocketHandler.Host.RemoveConnection(connection)
	return http.StatusOK, nil, nil
}

func (webSocketHandler *WebSocketHandler) handleMessage(msg *WebSocketMessage, conn *Connection) {
	if msg.Kind == "Ping" {
		conn.UpdatePingTime()
	}
}
