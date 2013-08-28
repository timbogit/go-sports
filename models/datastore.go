package models

import (
	"appengine"
//	"appengine/datastore"
)

type DataStore struct {
	Context appengine.Context
}

func (ds *DataStore) GetAllCountries() (countries []*Country, err error){
	// see for details: https://developers.google.com/appengine/docs/go/datastore/reference
//	datastore.NewQuery("Country")
	countries = make([]*Country, len(Store))
	i := 0
	for _, value := range Store {
		countries[i] = value
		i++
	}
	return countries, nil
}

func (ds *DataStore) UpdateCountry(country *Country) (err error) {
	//_, err := datastore.Put(ds.Context, datastore.NewKey(ds.Context, "Country", country.Code, 0, nil), country)
	Store[country.Code] = country
	return nil

}

func (ds *DataStore) GetCountry(code string) (country *Country, err error) {
	//err := datastore.Get(ds.Context, datastore.NewKey(ds.Context, "Country", code, 0, nil), country)
	country = Store[code]
	return country, nil
}


func (ds *DataStore) DeleteCountry(code string) (err error) {
	//err := datastore.Delete(ds.Context, datastore.NewKey(ds.Context, "Country", code, 0, nil))
	delete(Store, code)
	return nil
}