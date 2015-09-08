package appengine_bootstrap

import (
	"appengine"
	"appengine/datastore"
)

type Registration struct {
	ID   int64 `datastore:"-"`
	Name string
	Date int64
}

func registrationKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Registration", "default_registration", 0, nil)
}

func (r *Registration) Key(c appengine.Context) *datastore.Key {
	if r.ID == 0 {
		return datastore.NewIncompleteKey(c, "Registration", registrationKey(c))
	}
	return datastore.NewKey(c, "Registration", "default_registration", r.ID, nil)
}

func (r *Registration) Save(c appengine.Context) error {
	_, err := datastore.Put(c, r.Key(c), r)
	return err
}

func FindRegistration(c appengine.Context) *datastore.Query {
	return datastore.NewQuery("Registration").Ancestor(registrationKey(c))
}
