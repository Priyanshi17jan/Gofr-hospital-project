package datastore

import (
	"database/sql"
	"gofr.dev/pkg/errors"
	"gofr.dev/pkg/gofr"
	"priyanshi_gofr/model"
)

type patient struct {
	PatientId   int    `json:"patient_id"`
	PatientName string `json:"patient_name"`
	PhoneNo     int    `json:"phone_number"`
	BillDue     int    `json:"bill_due"`
	WardNo      int    `json:"ward_number"`
}

func New() *patient {
	return &patient{}
}
func (s *patient) GetByID(ctx *gofr.Context, id string) (*model.Patient, error) {
	var resp model.Patient

	err := ctx.DB().QueryRowContext(ctx, " SELECT patient_id,patient_name,phone_number,bill_due,ward_number FROM patient where patient_id= ?", id).
		Scan(&resp.PatientId, &resp.PatientName, &resp.PhoneNo, &resp.BillDue, &resp.WardNo)
	switch err {
	case sql.ErrNoRows:
		return &model.Patient{}, errors.EntityNotFound{Entity: "entries", ID: id}
	case nil:
		return &resp, nil
	default:
		return &model.Patient{}, err
	}
}
func (s *patient) Create(ctx *gofr.Context, patient *model.Patient) (*model.Patient, error) {
	var resp model.Patient
	result, err := ctx.DB().ExecContext(ctx, "INSERT INTO patient (patient_id, patient_name, phone_number, bill_due, ward_number) VALUES (?,?,?,?,?)",
		patient.PatientId, patient.PatientName, patient.PhoneNo, patient.BillDue, patient.WardNo)
	if err != nil {
		return &model.Patient{}, errors.DB{Err: err}
	}
	// Remove the line since lastInsertID is not being used
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return &model.Patient{}, errors.DB{Err: err}
	}
	// Set the ID in the response using entries.BookId
	resp.PatientId = int(lastInsertID)
	resp.PatientName = patient.PatientName
	resp.PhoneNo = patient.PhoneNo
	resp.BillDue = patient.BillDue
	resp.WardNo = patient.WardNo
	return &resp, nil
}
func (s *patient) Update(ctx *gofr.Context, patient *model.Patient) (*model.Patient, error) {
	_, err := ctx.DB().ExecContext(ctx, "UPDATE patient SET  patient_name= ? ,phone_number= ? , bill_due = ? , ward_number = ?   WHERE patient_id= ?",
		patient.PatientName, patient.PhoneNo, patient.BillDue, patient.WardNo, patient.PatientId)
	if err != nil {
		return &model.Patient{}, errors.DB{Err: err}
	}
	return patient, nil
}
func (s *patient) Delete(ctx *gofr.Context, id int) error {
	_, err := ctx.DB().ExecContext(ctx, "DELETE FROM patient where patient_id= ?", id)
	if err != nil {
		return errors.DB{Err: err}
	}
	return nil
}
