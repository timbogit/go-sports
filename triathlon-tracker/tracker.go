package tracker

import (
	"github.com/ant0ine/go-json-rest"
	"net/http"
	"api"
)

func init() {
	handler := rest.ResourceHandler{
                EnableRelaxedContentType: true,
        }
	handler.SetRoutes(
		rest.Route{"GET", "/api/countries", api.GetAllCountries},
		rest.Route{"POST", "/api/countries", api.PostCountry},
		rest.Route{"GET", "/api/countries/:code", api.GetCountry},
		rest.Route{"DELETE", "/api/countries/:code", api.DeleteCountry},
		// races routes
		rest.Route{"GET", "/api/races", api.GetAllRaces},
		rest.Route{"POST", "/api/races", api.PostRace},
		rest.Route{"GET", "/api/races/:key", api.GetRace},
		rest.Route{"DELETE", "/api/races/:key", api.DeleteRace},
	)

	http.Handle("/", &handler)
}
