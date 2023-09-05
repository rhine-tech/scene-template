package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/rhine-tech/scene"
	"net/http"
	"scene-template/echo"

	sgin "github.com/rhine-tech/scene/scenes/gin"
)

type ginApp struct {
	srv echo.EchoService
}

func (g *ginApp) Name() scene.AppName {
	return "echo.app.gin"
}

func (g *ginApp) Status() scene.AppStatus {
	return scene.AppStatusRunning
}

func (g *ginApp) Error() error {
	return nil
}

func (g *ginApp) Prefix() string {
	return "echo"
}

func (g *ginApp) Create(engine *gin.Engine, router gin.IRouter) error {
	router.GET("/:msg", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": g.srv.Echo(c.Param("msg"))})
	})
	return nil
}

func (g *ginApp) Destroy() error {
	return nil
}

func NewGinApp(srv echo.EchoService) sgin.GinApplication {
	return &ginApp{srv: srv}
}
