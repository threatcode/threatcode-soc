package web

type WebSocketMessage struct {
	Kind   string
	Object interface{}
}
