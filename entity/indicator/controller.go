package indicator

import (
	"net/http"

	"github.com/labstack/echo"
)

type Controller struct{}

func (c Controller) Route(r *echo.Echo) {
	r.GET("/indicator", c.Index)
}

func (c Controller) Index(ctx echo.Context) error {
	indicators, _ := Repository{}.Fetch()
	return ctx.JSON(http.StatusOK, indicators)
}
