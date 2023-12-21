package handler

import (
	"athmare/thebeststack/slick"
	"athmare/thebeststack/slick/app/view/profile"
)

func HandleUserProfile(c *slick.Context) error {
	user := profile.User{
		FirstName: "first",
		LastName:  "last",
		Email:     "a@a",
	}
	return c.Render(profile.Index(user))
}
