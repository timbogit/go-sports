/* Demonstrate simple POST GET and DELETE operations

The Curl Demo:

        curl -i -d '{"Key":"Miami703","Name":"Miami Half-Triathlon", "Date":"2013-11-21 13:00:00 UTC"}' http://127.0.0.1:8080/api/races
        curl -i http://127.0.0.1:8080/api/races/Miami703
        curl -i http://127.0.0.1:8080/api/races
        curl -i -X DELETE http://127.0.0.1:8080/api/races/FR
        curl -i http://127.0.0.1:8080/api/races
        curl -i -X DELETE http://127.0.0.1:8080/api/races/US
        curl -i http://127.0.0.1:8080/api/races

*/
package api

import (
//   "appengine"
//   "models"
   "github.com/ant0ine/go-json-rest"
//   "net/http"
)

func GetRace(w *rest.ResponseWriter, r *rest.Request) {
  // code := r.PathParam("code")
  // ds := &models.DataStore{ appengine.NewContext(r.Request) }
  // Race, err := ds.GetRace(code)
  // if err != nil {
  //   rest.Error(w, err.Error(), http.StatusInternalServerError)
  //   return
  // }
  // if Race == nil {
  //   rest.NotFound(w, r)
  //   return
  // }
  // w.WriteJson(&Race)
}
func GetAllraces(w *rest.ResponseWriter, r *rest.Request) {
  // ds := &models.DataStore{ appengine.NewContext(r.Request) }
  // cs, err := ds.GetAllraces()
  // if err != nil {
  //   rest.Error(w, err.Error(), http.StatusInternalServerError)
  //   return
  // }
  // w.WriteJson(cs)
}

func PostRace(w *rest.ResponseWriter, r *rest.Request) {
  // Race := models.Race{}
  // err := r.DecodeJsonPayload(&Race)
  // if err != nil {
  //   rest.Error(w, err.Error(), http.StatusInternalServerError)
  //   return
  // }
  // if Race.Code == "" {
  //   rest.Error(w, "Race code required", 400)
  //   return
  // }
  // if Race.Name == "" {
  //   rest.Error(w, "Race name required", 400)
  //   return
  // }
  // ds := &models.DataStore{ appengine.NewContext(r.Request) }
  // err = ds.UpdateRace(&Race)
  // if err != nil {
  //   rest.Error(w, err.Error(), http.StatusInternalServerError)
  //   return
  // }

  // w.WriteJson(&Race)
}

func DeleteRace(w *rest.ResponseWriter, r *rest.Request) {
  // code := r.PathParam("code")
  // ds := &models.DataStore{ appengine.NewContext(r.Request) }
  // err := ds.DeleteRace(code)
  // if err != nil {
  //   rest.Error(w, err.Error(), http.StatusInternalServerError)
  //   return
  // }
}
