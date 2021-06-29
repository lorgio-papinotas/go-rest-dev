package adding

import (
	"errors"

	"github.com/lorgioedtech/go-rest-dev/internal/entity"
	"github.com/lorgioedtech/go-rest-dev/internal/listing"
)

// ErrDuplicate is used when a institution already exists.
var ErrDuplicate = errors.New("institution already exists")

// Service provides beer adding operations.
type Service interface {
	AddInstitution(entity.Institution) error
}

// Repository provides access to institution repository.
type Repository interface {
	// AddInstitution saves a given institution to the repository.
	AddInstitution(entity.Institution) error
	// GetAllInstitutions returns all institutions saved in storage.
	GetAllInstitutions() []listing.Institution
}

type service struct {
	r Repository
}

// NewService creates an adding service with the necessary dependencies
func NewService(r Repository) Service {
	return &service{r}
}

// AddInstitution persists the given institution to storage
func (s *service) AddInstitution(institution entity.Institution) error {
	// make sure we don't add any duplicates
	existingInstitutions := s.r.GetAllInstitutions()
	for _, existing := range existingInstitutions {
		if institution.Name == existing.Name {
			return ErrDuplicate
		}
	}

	// any other validation can be done here
	_ = s.r.AddInstitution(institution)

	return nil
}
