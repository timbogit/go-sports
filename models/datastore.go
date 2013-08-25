package models

import (
	"appengine"
//	"appengine/datastore"
)

type DataStore struct {
	Context appengine.Context
}

func (ds *DataStore) GetAllCountries() (countries []*Country, err error){
//	datastore.NewQuery("Country")
	countries = make([]*Country, len(Store))
	i := 0
	for _, value := range Store {
		countries[i] = value
		i++
	}
	return countries, nil
}