package builder

import (
	"github.com/aynakeya/scene"
	"github.com/aynakeya/scene/registry"
	sgin "github.com/aynakeya/scene/scenes/gin"
	sgrpc "github.com/aynakeya/scene/scenes/grpc"
	sws "github.com/aynakeya/scene/scenes/websocket"
	"scene-template/echo"
	"scene-template/echo/delivery"
	"scene-template/echo/service"
)

type Builder struct {
	scene.Builder
}

func (b Builder) Init() scene.LensInit {
	return func() {
		// init function.
		registry.Register(service.NewEchoService())
	}
}

func (b Builder) Apps() []any {
	srv := registry.Use(echo.EchoService(nil))
	return []any{
		func() sws.WebsocketApplication {
			return delivery.NewWsApp(srv)
		},
		func() sgin.GinApplication {
			return delivery.NewGinApp(srv)
		},
		func() sgrpc.GrpcApplication {
			return delivery.NewGrpcApp(srv)
		},
	}
}
