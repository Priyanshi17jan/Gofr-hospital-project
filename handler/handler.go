package handler

import (
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"priyanshi_gofr/datastore"
	"priyanshi_gofr/model"
	"strconv"
)

type handler struct {
	store datastore.Patient
}

func New(s datastore.Patient) handler {
	return handler{store: s}
}
func (h handler) GetByID(ctx *gofr.Context) (interface{}, error) {
	// ctx.PathParam() returns the path parameter from HTTP request.
	id := ctx.PathParam("id")
	if id == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}

	if _, err := validateID(id); err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}

	resp, err := h.store.GetByID(ctx, id)
	if err != nil {
		return nil, errors.EntityNotFound{
			Entity: "entry",
			ID:     id,
		}
	}
	return resp, nil
}
func (h handler) Create(ctx *gofr.Context) (interface{}, error) {
	var student model.Patient

	// ctx.Bind() binds the incoming data from the HTTP request to a provided interface (i).
	if err := ctx.Bind(&student); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	resp, err := h.store.Create(ctx, &student)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (h handler) Update(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	id, err := validateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}
	var patient model.Patient
	if err = ctx.Bind(&patient); err != nil {
		ctx.Logger.Errorf("error in binding: %v", err)
		return nil, errors.InvalidParam{Param: []string{"body"}}
	}
	patient.PatientId = id
	resp, err := h.store.Update(ctx, &patient)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
func (h handler) Delete(ctx *gofr.Context) (interface{}, error) {
	i := ctx.PathParam("id")
	if i == "" {
		return nil, errors.MissingParam{Param: []string{"id"}}
	}
	id, err := validateID(i)
	if err != nil {
		return nil, errors.InvalidParam{Param: []string{"id"}}
	}
	if err := h.store.Delete(ctx, id); err != nil {
		return nil, err
	}
	return "Deleted successfully", nil
}
func validateID(id string) (int, error) {
	res, err := strconv.Atoi(id)
	if err != nil {
		return 0, err
	}
	return res, err
}
