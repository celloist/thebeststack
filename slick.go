package slick

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/a-h/templ"
	"github.com/julienschmidt/httprouter"
)

type Handler func(c *Context) error

type Plug func(Handler) Handler

type Slick struct {
	ErrorHandler ErrorHandler
	router       *httprouter.Router
	middleware   []Plug
}

type ErrorHandler func(error, *Context) error

type Context struct {
	response http.ResponseWriter
	request  *http.Request
	ctx      context.Context
}

func defaultErrorHandler(err error, c *Context) error {
	slog.Error("error", "err", err)
	return nil
}

func (c *Context) Set(key string, value any) {
	c.ctx = context.WithValue(c.ctx, key, value)
}
func (c *Context) Get(key string) any {
	return c.ctx.Value(key)
}

func (c *Context) Render(comp templ.Component) error {
	return comp.Render(c.ctx, c.response)
}

func (s *Slick) Plug(plugs ...Plug) {
	s.middleware = append(s.middleware, plugs...)
}

func (s *Slick) Start(port string) error {
	return http.ListenAndServe(port, s.router)
}

func New() *Slick {
	return &Slick{
		router:       httprouter.New(),
		ErrorHandler: defaultErrorHandler,
	}
}

func (s *Slick) Get(path string, h Handler, plug ...Handler) {
	s.router.GET(path, s.makeHTTPRouterHandler(h))
}

func (s *Slick) makeHTTPRouterHandler(h Handler) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
		ctx := &Context{
			response: w,
			request:  r,
			ctx:      context.Background(),
		}

		for i := len(s.middleware) - 1; i >= 0; i-- {
			h = s.middleware[i](h)
		}
		if err := h(ctx); err != nil {
			s.ErrorHandler(err, ctx)
		}
	}
}
