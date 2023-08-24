package delivery

import (
	"github.com/aynakeya/scene"
	"github.com/gin-gonic/gin"
	"net/http"

	sgin "github.com/aynakeya/scene/scenes/gin"
)

type ginApp struct {
}

func (g ginApp) Name() scene.AppName {
	return "echo.app.gin"
}

func (g ginApp) Status() scene.AppStatus {
	return scene.AppStatusRunning
}

func (g ginApp) Error() error {
	return nil
}

func (g ginApp) Prefix() string {
	return "echo"
}

func (g ginApp) Create(engine *gin.Engine, router gin.IRouter) error {
	router.GET("/:msg", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"msg": c.Param("msg")})
	})
	return nil
}

func (g ginApp) Destroy() error {
	return nil
}

func NewGinApp() sgin.GinApplication {
	return &ginApp{}
}
