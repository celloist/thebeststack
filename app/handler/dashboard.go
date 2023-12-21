package handler

import (
	"athmare/thebeststack/slick"
	"athmare/thebeststack/slick/app/view/dashboard"
)

func HandleDashboard(c *slick.Context) error {

	return c.Render(dashboard.Index())
}
