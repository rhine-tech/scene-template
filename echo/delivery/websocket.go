package delivery

import (
	"github.com/aynakeya/scene"
	"github.com/gorilla/websocket"

	sws "github.com/aynakeya/scene/scenes/websocket"
)

type websocketApp struct {
}

func (e *websocketApp) Status() scene.AppStatus {
	return scene.AppStatusRunning
}

func (e *websocketApp) Error() error {
	return nil
}

func (e *websocketApp) Prefix() string {
	return "echo"
}

func (e *websocketApp) Create(mux sws.WebsocketMux) error {
	mux.HandleFunc("", func(conn *websocket.Conn, clos func()) sws.WebsocketMessageHandler {
		return func(msgType int, msg []byte, err error) error {
			if err != nil {
				return err
			}
			return conn.WriteMessage(msgType, msg)
		}
	})
	return nil
}

func (e *websocketApp) Name() scene.AppName {
	return "echo.app.websocket"
}

func NewWsApp() sws.WebsocketApplication {
	return &websocketApp{}
}
