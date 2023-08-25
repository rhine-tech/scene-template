package delivery

import (
	context "context"
	"github.com/aynakeya/scene"
	grpc "google.golang.org/grpc"
	"scene-template/echo"

	sgrpc "github.com/aynakeya/scene/scenes/grpc"
)

type grpcApp struct {
	srv echo.EchoService
	echo.UnimplementedEchoServer
}

func (g *grpcApp) Name() scene.AppName {
	return "echo.app.grpc"
}

func (g *grpcApp) Status() scene.AppStatus {
	return scene.AppStatusRunning
}

func (g *grpcApp) Error() error {
	return nil
}

func (g *grpcApp) Create(server *grpc.Server) error {
	echo.RegisterEchoServer(server, g)
	return nil
}

func (g *grpcApp) Echo(ctx context.Context, request *echo.EchoRequest) (*echo.EchoReply, error) {
	return &echo.EchoReply{Message: g.srv.Echo(request.Message)}, nil
}

func NewGrpcApp(srv echo.EchoService) sgrpc.GrpcApplication {
	return &grpcApp{srv: srv}
}
