package cmandroid

import (
	"net/http"

	_ "github.com/JoelPagliuca/cmandroid-server/go/statik" // necessary
	"github.com/gorilla/mux"
	"github.com/rakyll/statik/fs"
)

// Route all the data we need for a router entry
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes list of routes
type Routes []Route

// NewRouter generate the router from the handlers below
func NewRouter(app App) *mux.Router {

	var routes = Routes{
		Route{
			"Index",
			"GET",
			"/",
			app.Index,
		},

		Route{
			"GetDevicelist",
			"GET",
			"/devicelist",
			app.GetDevicelist,
		},

		Route{
			"PostTap",
			"POST",
			"/tap",
			app.PostTap,
		},

		Route{
			"PostPackage",
			"POST",
			"/startpackage",
			app.PostPackage,
		},

		Route{
			"PostKeyboard",
			"POST",
			"/keyboard",
			app.PostKeyboard,
		},
	}

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		handler = Logger(handler, route.Name)
		handler = BasicHeaders(handler)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}

	// statik swaggerui
	statikFS, err := fs.New()
	if err != nil {
		panic(err)
	}

	statikServer := http.FileServer(statikFS)
	swaggerHandler := http.StripPrefix("/swaggerui/", statikServer)
	swaggerHandler = Logger(swaggerHandler, "SwaggerUI")
	router.PathPrefix("/swaggerui/").Handler(swaggerHandler)

	return router
}
