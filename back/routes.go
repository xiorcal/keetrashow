package back

import (
	"net/http"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/xiorcal/keetrashow/back/jwt"
)

func CreateBackRoutes() {
	http.HandleFunc("/signin", jwt.Signin)
	http.HandleFunc("/welcome", jwt.Welcome)
	http.HandleFunc("/refresh", jwt.Refresh)

	http.Handle("/", &app.Handler{
		Name:        "Hello",
		Description: "An Hello World! example",
		Styles: []string{
			"/web/keetrashow.css", // Loads keetrashow.css file.
		},
	})

}
