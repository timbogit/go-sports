package models

import (
	"appengine"
//	"appengine/datastore"
)

type DataStore struct {
	Context appengine.Context
}

func (ds *DataStore) UpdateRace(race *Race) (err error) {
	//_, err := datastore.Put(ds.Context, datastore.NewKey(ds.Context, "Race", race.Key, 0, nil), race)
	RaceStore[race.Key] = race
	return nil

}

func (ds *DataStore) GetRace(key string) (race *Race, err error) {
	//err := datastore.Get(ds.Context, datastore.NewKey(ds.Context, "Race", key, 0, nil), country)
	race = RaceStore[key]
	return race, nil
}

func (ds *DataStore) DeleteRace(key string) (err error) {
	//err := datastore.Delete(ds.Context, datastore.NewKey(ds.Context, "Race", key, 0, nil))
	delete(RaceStore, key)
	return nil
}

func (ds *DataStore) GetAllRaces() (races []*Race, err error){
	// see for details: https://developers.google.com/appengine/docs/go/datastore/reference
//	datastore.NewQuery("Race")
	races = make([]*Race, len(RaceStore))
	i := 0
	for _, value := range RaceStore {
		races[i] = value
		i++
	}
	return races, nil
}
