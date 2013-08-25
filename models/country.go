package models

//import (
//	"appengine"
//	"appengine/datastore"
//)

type Country struct {
	Code string
	Name string
}

var Store = map[string]*Country{ 
	"FR": &Country{"FR", "France" },
	"DE": &Country{"DE", "Germany" },
	"US": &Country{"US", "United States of America" },
}

