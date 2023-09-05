package main

import (
	"github.com/rhine-tech/scene"
	"github.com/rhine-tech/scene/engines"
	authentication "github.com/rhine-tech/scene/lens/middlewares/authentication/builder"
	permission "github.com/rhine-tech/scene/lens/middlewares/permission/builder"
	"github.com/rhine-tech/scene/registry"
	sgrpc "github.com/rhine-tech/scene/scenes/grpc"
	sws "github.com/rhine-tech/scene/scenes/websocket"
	echo "scene-template/echo/builder"

	config "github.com/rhine-tech/scene/lens/infrastructure/config/builder"

	asynctask "github.com/rhine-tech/scene/lens/infrastructure/asynctask/builder"
	datasource "github.com/rhine-tech/scene/lens/infrastructure/datasource/builder"
	ingestion "github.com/rhine-tech/scene/lens/infrastructure/ingestion/builder"
	logger "github.com/rhine-tech/scene/lens/infrastructure/logger/builder"

	sgin "github.com/rhine-tech/scene/scenes/gin"
)

const configFile = ".env.dev"

func main() {
	config.Init(configFile)
	builders := scene.BuildableArray{
		logger.Builder{},
		asynctask.Builder{},
		datasource.Builder{},
		ingestion.DummyBuilder{},
		echo.Builder{},
		permission.Builder{},
		authentication.Builder{},
	}
	scene.BuildInitArray(builders).Inits()
	registry.Logger.Infof("using config file: %s", configFile)
	engine := engines.NewEngine(registry.Logger,
		sgin.NewAppContainer(
			registry.Config.GetString("scene.app.gin.addr"),
			scene.BuildApps[sgin.GinApplication](builders)...),
		sws.NewContainer(
			registry.Config.GetString("scene.app.websocket.addr"),
			scene.BuildApps[sws.WebsocketApplication](builders)...),
		sgrpc.NewContainer(
			registry.Config.GetString("scene.app.grpc.addr"),
			scene.BuildApps[sgrpc.GrpcApplication](builders)...))
	if err := engine.Run(); err != nil {
		registry.Logger.Errorf("engine error: %s", err)
	}
}
