package front

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Login struct {
	app.Compo
}

func (l *Login) Render() app.UI {
	return app.Div().Class("login-page").Body(
		app.H1().
			Class("title").
			Text("Please log in"),
		app.Form().Class("login-form").Body(
			app.Div().Class("form-row").Body(
				app.Label().Text("username"),
				app.Input().Type("text")),
			app.Div().Class("form-row").Body(
				app.Label().Text("password"),
				app.Input().Type("password")),
			app.Div().Class("form-row").Body(
				app.Input().Type("button").Value("Login").OnClick(l.onSubmit),
			),
		),
	)
}

func (c *Login) onSubmit(ctx app.Context, e app.Event) {
	v := e.Type()
	app.Log(v)
}
