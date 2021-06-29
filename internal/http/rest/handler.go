package rest

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/lorgioedtech/go-rest-dev/internal/adding"
	"github.com/lorgioedtech/go-rest-dev/internal/entity"
	"github.com/lorgioedtech/go-rest-dev/internal/listing"
)

func Handler(adding adding.Service, listing listing.Service) http.Handler {
	router := httprouter.New()

	router.GET("/institutions", getInstitutions(listing))
	router.GET("/institutions/:id", getInstitution(listing))
	router.POST("/institutions", AddInstitution(adding))
	return router
}

// AddInstitution returns a handler for POST /institutions requests
func AddInstitution(s adding.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		decoder := json.NewDecoder(r.Body)

		var newInstitution entity.Institution
		err := decoder.Decode(&newInstitution)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		s.AddInstitution(newInstitution)
		// error handling omitted for simplicity

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode("New Institution added.")
	}
}

// getInstitutions returns a handler for GET /institutions requests
func getInstitutions(s listing.Service) func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		list := s.GetInstitutions()
		json.NewEncoder(w).Encode(list)
	}
}

// getInstitution returns a handler for GET /institutions/:id requests
func getInstitution(s listing.Service) func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		institution, err := s.GetInstitution(p.ByName("id"))
		if err == listing.ErrNotFound {
			http.Error(w, "The institution you requested does not exist.", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(institution)
	}
}
