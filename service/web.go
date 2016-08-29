package service

import (
	"fmt"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
)

// webController is the private interface all controllers
// must adhere to
type webController interface {
	Route(r *echo.Echo)
}

type web struct {
	Host     string
	Port     string
	instance *echo.Echo
}

// Route registers routes in the server
func (w *web) Route(controllers ...webController) *web {
	w.ready()

	for _, c := range controllers {
		c.Route(w.instance)
	}

	return w
}

// Run starts a webserver
func (w *web) Run() {
	w.ready()

	addr := fmt.Sprintf("%s:%s", w.Host, w.Port)
	Logger.Info("Starting web server at: %s", addr)
	w.instance.Run(standard.New(addr))
}

// ready ensures the web server is prepared to be run
func (s *web) ready() {
	if s.instance == nil {
		s.instance = echo.New()
	}
}
