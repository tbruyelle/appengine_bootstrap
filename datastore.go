package appengine_bootstrap

import (
	"appengine"
	"appengine/datastore"
)

type Registration struct {
	ID   string
	Date int64
}

func registrationKey(c appengine.Context) *datastore.Key {
	return datastore.NewKey(c, "Registration", "default_registration", 0, nil)
}

func (r *Registration) Save(c appengine.Context) error {
	key := datastore.NewIncompleteKey(c, "Registration", registrationKey(c))
	_, err := datastore.Put(c, key, r)
	return err
}

func FindRegistration(c appengine.Context) *datastore.Query {
	return datastore.NewQuery("Registration").Ancestor(registrationKey(c))
}
