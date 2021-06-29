package json

import (
	"encoding/json"
	"log"
	"path"
	"runtime"
	"time"

	"github.com/lorgioedtech/go-rest-dev/internal/entity"
	"github.com/lorgioedtech/go-rest-dev/internal/listing"
	"github.com/lorgioedtech/go-rest-dev/internal/storage"
	scribble "github.com/nanobox-io/golang-scribble"
)

const (
	// dir defines the name of the directory where the files are stored
	dir = "/data/"

	// CollectionInstitution identifier for the JSON collection of beers
	CollectionInstitution = "institutions"
)

type Institution struct {
	entity.Institution
}

// Storage stores institution data in JSON files
type Storage struct {
	db *scribble.Driver
}

// NewStorage returns a new JSON  storage
func NewStorage() (*Storage, error) {
	var err error

	s := new(Storage)

	_, filename, _, _ := runtime.Caller(0)
	p := path.Dir(filename)

	s.db, err = scribble.New(p+dir, nil)
	if err != nil {
		return nil, err
	}

	return s, nil
}

// AddInstitution saves the given institution to the repository
func (s *Storage) AddInstitution(newInstitution entity.Institution) error {
	id, err := storage.GetID("institution")
	if err != nil {
		log.Fatal(err)
	}

	institutionModel := entity.Institution{
		ID:      id,
		Name:    newInstitution.Name,
		Created: time.Now(),
	}

	if err := s.db.Write(CollectionInstitution, institutionModel.ID, institutionModel); err != nil {
		return err
	}
	return nil
}

// Get returns a institution with the specified ID
func (s *Storage) GetInstitution(id string) (listing.Institution, error) {
	var institutionPersistence Institution
	var institution listing.Institution

	if err := s.db.Read(CollectionInstitution, id, &institutionPersistence); err != nil {
		// err handling omitted for simplicity
		return institution, listing.ErrNotFound
	}

	institution.ID = institutionPersistence.Institution.ID
	institution.Name = institutionPersistence.Institution.Name
	institution.Created = institutionPersistence.Institution.Created

	return institution, nil
}

// GetAll returns all beers
func (s *Storage) GetAllInstitutions() []listing.Institution {
	list := []listing.Institution{}

	records, err := s.db.ReadAll(CollectionInstitution)
	if err != nil {
		// err handling omitted for simplicity
		return list
	}

	for _, r := range records {
		var i Institution
		var institution listing.Institution

		if err := json.Unmarshal([]byte(r), &i); err != nil {
			// err handling omitted for simplicity
			return list
		}

		institution.ID = i.Institution.ID
		institution.Name = i.Institution.Name
		institution.Created = i.Institution.Created

		list = append(list, institution)
	}

	return list
}
