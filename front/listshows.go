package front

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/xiorcal/keetrashow/data"
)

type ListShows struct {
	app.Compo
	Shows []data.Show
}

func (l *ListShows) Render() app.UI {
	return app.Div().Class("list-page").Body(
		app.H1().
			Class("title").
			Text("Please log in"),
		app.Form().Body(
			app.Div().Body(
				app.Label().Text("username"),
				app.Input().Type("text")),
			app.Div().Body(
				app.Label().Text("password"),
				app.Input().Type("password")),
		),
	)
}
