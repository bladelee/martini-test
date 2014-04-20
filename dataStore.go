package main

import (
	"errors"
	"sync"
)

var (
	ErrAlreadyExists = errors.New("album already exists")
)

// The DB interface defines methods to manipulate the albums.
type DB interface {
	Get(id string) *TemplateResource
	GetAll() []*TemplateResource
	Add(ts *TemplateResource) (string, error)
	Update(ts *TemplateResource) error
	Delete(id string)
}

// Thread-safe in-memory map of albums.
type TemplatesDB struct {
	sync.RWMutex
	m   map[string]*TemplateResource
	seq int
}

// GetAll returns all albums from the database.
func (db *TemplatesDB) GetAll() []*TemplateResource {
	db.RLock()
	defer db.RUnlock()
	if len(db.m) == 0 {
		return nil
	}
	ar := make([]*TemplateResource, len(db.m))
	i := 0
	for _, v := range db.m {
		ar[i] = v
		i++
	}
	return ar
}

// Get returns the album identified by the id, or nil.
func (db *TemplatesDB) Get(id string) *TemplateResource {
	db.RLock()
	defer db.RUnlock()
	return db.m[id]
}

// Add creates a new album and returns its id, or an error.
func (db *TemplatesDB) Add(ts *TemplateResource) (string, error) {
	db.Lock()
	defer db.Unlock()
	// Return an error if band-title already exists
	if !db.isUnique(ts) {
		return "0", ErrAlreadyExists
	}
	// Get the unique ID
	db.seq++
	ts.Index = db.seq
	// Store
	db.m[ts.Id] = ts
	return ts.Id, nil
}

// Update changes the album identified by the id. It returns an error if the
// updated album is a duplicate.
func (db *TemplatesDB) Update(ts *TemplateResource) error {
	db.Lock()
	defer db.Unlock()
	if !db.isUnique(ts) {
		return ErrAlreadyExists
	}
	db.m[ts.Id] = ts
	return nil
}

// Delete removes the album identified by the id from the database. It is a no-op
// if the id does not exist.
func (db *TemplatesDB) Delete(id string) {
	db.Lock()
	defer db.Unlock()
	delete(db.m, id)
}

// Checks if the album already exists in the database, based on the Band and Title
// fields.
func (db *TemplatesDB) isUnique(ts *TemplateResource) bool {
	for _, v := range db.m {
		if v.Name == ts.Name && v.Id != ts.Id {
			return false
		}
	}
	return true
}

// The one and only database instance.
var db DB

func init() {
	db = &TemplatesDB{
		m:   make(map[string]*TemplateResource),
		seq: 4}
	// Fill the databaseni
	db.Add(&TemplateResource{Id: "1", Name: "Slayer", Content: "Reign In Blood", TemplateObj: &Template{}, Index: 1})
	db.Add(&TemplateResource{Id: "2", Name: "Slayer", Content: "Seasons In The Abyss", TemplateObj: &Template{}, Index: 2})
	db.Add(&TemplateResource{Id: "3", Name: "Bruce Springsteen", Content: "Born To Run", TemplateObj: &Template{}, Index: 3})
}
