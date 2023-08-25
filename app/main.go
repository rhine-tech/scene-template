package main

import (
	"github.com/aynakeya/scene"
	"github.com/aynakeya/scene/engines"
	"github.com/aynakeya/scene/registry"
	sgrpc "github.com/aynakeya/scene/scenes/grpc"
	sws "github.com/aynakeya/scene/scenes/websocket"
	echo "scene-template/echo/builder"

	config "github.com/aynakeya/scene/lens/infrastructure/config/builder"

	asynctask "github.com/aynakeya/scene/lens/infrastructure/asynctask/builder"
	datasource "github.com/aynakeya/scene/lens/infrastructure/datasource/builder"
	ingestion "github.com/aynakeya/scene/lens/infrastructure/ingestion/builder"
	logger "github.com/aynakeya/scene/lens/infrastructure/logger/builder"

	sgin "github.com/aynakeya/scene/scenes/gin"
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
