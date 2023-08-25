package delivery

import (
	"github.com/aynakeya/scene"
	"github.com/gorilla/websocket"
	"scene-template/echo"

	sws "github.com/aynakeya/scene/scenes/websocket"
)

type websocketApp struct {
	srv echo.EchoService
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
			return conn.WriteMessage(msgType, []byte(e.srv.Echo(string(msg))))
		}
	})
	return nil
}

func (e *websocketApp) Name() scene.AppName {
	return "echo.app.websocket"
}

func NewWsApp(srv echo.EchoService) sws.WebsocketApplication {
	return &websocketApp{srv: srv}
}
