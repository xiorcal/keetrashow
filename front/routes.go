package front

import (
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

func CreateFrontRoutes() {
	app.Route("/", &Login{})

}
