package hecho

import (
	"github.com/Kamva/hexa"
	"github.com/labstack/echo/v4"
)

// SetContextLogger set the hexa logger on each context.
func SetContextLogger(cfg hexa.Config, logger hexa.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(ctx echo.Context) error {
			// Set context logger
			hexaContext := ctx.Get(ContextKeyHexaCtx).(hexa.Context)
			ctx.SetLogger(HexaToEchoLogger(cfg, logger.With(hexaContext)))
			return next(ctx)
		}
	}
}
