package listing

import (
	"errors"
)

// ErrNotFound is used when a beer could not be found.
var ErrNotFound = errors.New("institution not found")

// Repository provides access to the beer and review storage.
type Repository interface {
	// GetInstitution returns the beer with given ID.
	GetInstitution(string) (Institution, error)
	// GetAllInstitutions returns all beers saved in storage.
	GetAllInstitutions() []Institution
}

// Service provides institution listing operations.
type Service interface {
	GetInstitution(string) (Institution, error)
	GetInstitutions() []Institution
}

type service struct {
	repository Repository
}

// NewService creates a listing service with the necessary dependencies
func NewService(repository Repository) Service {
	return &service{repository}
}

// GetInstitutions returns all institutions
func (s *service) GetInstitutions() []Institution {
	return s.repository.GetAllInstitutions()
}

// GetInstitution returns a institution
func (s *service) GetInstitution(id string) (Institution, error) {
	return s.repository.GetInstitution(id)
}
