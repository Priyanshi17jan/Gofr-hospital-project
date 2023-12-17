package datastore

import (
	"gofr.dev/pkg/gofr"
	"priyanshi_gofr/model"
)

type Patient interface {
	// GetByID retrieves a book record based on its ID.
	GetByID(ctx *gofr.Context, id string) (*model.Patient, error)
	// Create inserts a new book record into the database.
	Create(ctx *gofr.Context, model *model.Patient) (*model.Patient, error)
	// Update updates an existing book with the provided information.
	Update(ctx *gofr.Context, model *model.Patient) (*model.Patient, error)
	// Delete removes a entry record from the database based on its ID.
	Delete(ctx *gofr.Context, id int) error
}
