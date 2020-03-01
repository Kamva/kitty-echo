package kecho

import (
	"fmt"
	"github.com/Kamva/kitty"
	"github.com/labstack/echo/v4"
)

type (
	GetResource interface {
		Get(ctx echo.Context) error
	}

	CreateResource interface {
		Create(ctx echo.Context) error
	}

	UpdateResource interface {
		Update(ctx echo.Context) error
	}

	PatchResource interface {
		Patch(ctx echo.Context) error
	}

	DeleteResource interface {
		Delete(ctx echo.Context) error
	}

	Resource struct {
	}
)

// Ctx method extract the kitty context from the echo context.
func (r Resource) Ctx(c echo.Context) kitty.Context {
	return c.Get(ContextKeyKittyCtx).(kitty.Context)
}

// Resource define each method that exists in provided resource.
func ResourceAPI(group *echo.Group, resource interface{}, prefix string, m ...echo.MiddlewareFunc) {
	if r, ok := resource.(GetResource); ok {
		group.GET("/:id", r.Get, m...).Name = routeName(prefix, "get")
	}

	if r, ok := resource.(CreateResource); ok {
		group.POST("/:id", r.Create, m...).Name = routeName(prefix, "post")
	}

	if r, ok := resource.(UpdateResource); ok {
		group.PUT("/:id", r.Update, m...).Name = routeName(prefix, "put")
	}

	if r, ok := resource.(PatchResource); ok {
		group.PATCH("/:id", r.Patch, m...).Name = routeName(prefix, "patch")
	}

	if r, ok := resource.(DeleteResource); ok {
		group.DELETE("/:id", r.Delete, m...).Name = routeName(prefix, "delete")
	}
}

func routeName(prefix, name string) string {
	return fmt.Sprintf("%s::%s", prefix, name)
}
