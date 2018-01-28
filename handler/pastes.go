package handler

import (
	"net/http"

	"github.com/pastepaste/backend/service"
	"github.com/sirupsen/logrus"
)

// Pastes handles requests to the paste resource.
type Pastes struct {
	DB  service.DB
	Log logrus.FieldLogger
}

// Create creates a paste.
func (p *Pastes) Create(w http.ResponseWriter, r *http.Request) {
}

// Get gets a paste by its slug.
func (p *Pastes) Get(w http.ResponseWriter, r *http.Request) {
}
