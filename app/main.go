package main

import (
	"athmare/thebeststack/slick"
	"athmare/thebeststack/slick/app/handler"

	"github.com/google/uuid"
)

func main() {
	app := slick.New()

	app.Plug(withRequestID, WithAuth)
	app.Get("/profile", handler.HandleUserProfile)
	app.Get("/dashboard", handler.HandleDashboard)

	app.Start(":3000")
}

func WithAuth(handler slick.Handler) slick.Handler {
	return func(c *slick.Context) error {
		c.Set("email", "a@a.a")
		return handler(c)
	}
}

func withRequestID(h slick.Handler) slick.Handler {
	return func(c *slick.Context) error {
		c.Set("requestID", uuid.New())
		return h(c)
	}
}
