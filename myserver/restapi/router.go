package restapi

import (
	"net/http"

	"myserver/restapi/handlers"
	"myserver/infrastructure"

	"github.com/gorilla/mux"
)

var controller = &handlers.Controller{Repository: infrastructure.Repository{}}

// Route defines a route
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes defines the list of routes of our API
type Routes []Route

var routes = Routes{
	Route{
		"GetInvoices",
		"GET",
		"/",
		controller.GetInvoices,
	},
	Route{
		"GetParnerInvoices",
		"GET",
		"/partnerinvoices",
		controller.GetPartnerInvoices,
	}
	Route{
		"AddInvoice",
		"POST",
		"/",
		controller.AddInvoice,
	},
	Route{
		"UpdateInvoice",
		"PUT",
		"/",
		controller.UpdateInvoice,
	},
	Route{
		"DeleteInvoice",
		"DELETE",
		"/",
		controller.DeleteInvoice,
	},
}

//NewRouter configures a new router to the API
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return router
}
