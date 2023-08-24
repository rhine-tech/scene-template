package builder

import (
	"github.com/aynakeya/scene"
	"scene-template/echo/delivery"
)

type Builder struct {
	scene.Builder
}

func (b Builder) Init() scene.LensInit {
	return func() {
		// init function.
		// do nothing here.
	}
}

func (b Builder) Apps() []any {
	return []any{
		delivery.NewWsApp,
		delivery.NewGinApp,
	}
}
