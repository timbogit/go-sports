/* Demonstrate simple POST GET and DELETE operations

The Curl Demo:

        curl -i -d '{"Code":"FR","Name":"France"}' http://127.0.0.1:8080/api/countries
        curl -i -d '{"Code":"US","Name":"United States"}' http://127.0.0.1:8080/api/countries
        curl -i http://127.0.0.1:8080/api/countries/FR
        curl -i http://127.0.0.1:8080/api/countries/US
        curl -i http://127.0.0.1:8080/api/countries
        curl -i -X DELETE http://127.0.0.1:8080/api/countries/FR
        curl -i http://127.0.0.1:8080/api/countries
        curl -i -X DELETE http://127.0.0.1:8080/api/countries/US
        curl -i http://127.0.0.1:8080/api/countries

*/
package api

import (
	"appengine"
	"models"
	"github.com/ant0ine/go-json-rest"
	"net/http"
)

func GetCountry(w *rest.ResponseWriter, r *rest.Request) {
	code := r.PathParam("code")
	ds := &models.DataStore{ appengine.NewContext(r.Request) }
	country, err := ds.GetCountry(code)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if country == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&country)
}
func GetAllCountries(w *rest.ResponseWriter, r *rest.Request) {
	ds := &models.DataStore{ appengine.NewContext(r.Request) }
	cs, err := ds.GetAllCountries()
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteJson(cs)
}

func PostCountry(w *rest.ResponseWriter, r *rest.Request) {
	country := models.Country{}
	err := r.DecodeJsonPayload(&country)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if country.Code == "" {
		rest.Error(w, "country code required", 400)
		return
	}
	if country.Name == "" {
		rest.Error(w, "country name required", 400)
		return
	}
	ds := &models.DataStore{ appengine.NewContext(r.Request) }
	err = ds.UpdateCountry(&country)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteJson(&country)
}

func DeleteCountry(w *rest.ResponseWriter, r *rest.Request) {
	code := r.PathParam("code")
	ds := &models.DataStore{ appengine.NewContext(r.Request) }
	err := ds.DeleteCountry(code)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
