package api

import (
//	"appengine"
//	"appengine/datastore"
	"models"
	"github.com/ant0ine/go-json-rest"
	"net/http"
)

func GetCountry(w *rest.ResponseWriter, r *rest.Request) {
	code := r.PathParam("code")
	country := models.Store[code]
	if country == nil {
		rest.NotFound(w, r)
		return
	}
	w.WriteJson(&country)
}
func GetAllCountries(w *rest.ResponseWriter, r *rest.Request) {
	countries := make([]*models.Country, len(models.Store))
	i := 0
	for _, country := range models.Store {
		countries[i] = country
		i++
	}
	w.WriteJson(&countries)
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
	models.Store[country.Code] = &country
	w.WriteJson(&country)
}

func DeleteCountry(w *rest.ResponseWriter, r *rest.Request) {
	code := r.PathParam("code")
	delete(models.Store, code)
}
