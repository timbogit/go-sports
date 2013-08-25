package hello

import (
	"appengine"
	"appengine/datastore"
	"appengine/user"
	"fmt"
	"html/template"
	"math/rand"
	"github.com/ant0ine/go-json-rest"
	"net/http"
	"time"
	"api"
)

type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}

func init() {
// 	 http.HandleFunc("/", root)
//       http.HandleFunc("/sign", sign)
	handler := rest.ResourceHandler{
                EnableRelaxedContentType: true,
        }
	handler.SetRoutes(
		rest.Route{"GET", "/api/countries", api.GetAllCountries},
		rest.Route{"POST", "/api/countries", api.PostCountry},
		rest.Route{"GET", "/api/countries/:code", api.GetCountry},
		rest.Route{"DELETE", "/api/countries/:code", api.DeleteCountry},
	)

	http.Handle("/", &handler)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Greeting").Order("-Date").Limit(10)
	greetings := make([]Greeting, 0, 10)
	if _, err := q.GetAll(c, &greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := guestbookTemplate.Execute(w, greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	rc := make(chan string)
	rand.Seed(time.Now().Unix())

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		rc <- "random 1"
	}()

	go func() {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		rc <- "random 2"
	}()

	str := "empty"
	select {
	case str = <-rc:
	}
	fmt.Fprint(w, str)
}

var guestbookTemplate = template.Must(template.New("book").Parse(guestbookTemplateHTML))

const guestbookTemplateHTML = `
<html>
  <body>
    {{range .}}
      {{with .Author}}
        <p><b>{{.}}</b> wrote:</p>
      {{else}}
        <p>An anonymous person wrote:</p>
      {{end}}
      <pre>{{.Content}}</pre>
    {{end}}
    <form action="/sign" method="post">
      <div><textarea name="content" rows="3" cols="60"></textarea></div>
      <div><input type="submit" value="Sign Guestbook"></div>
    </form>
  </body>
</html>
`

func sign(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	g := Greeting{
		Content: r.FormValue("content"),
		Date:    time.Now(),
	}
	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}
	_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Greeting", nil), &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/", http.StatusFound)
}
