package model

type Patient struct {
	PatientId   int    `json:"patient_id"`
	PatientName string `json:"patient_name"`
	PhoneNo     int    `json:"phone_number"`
	BillDue     int    `json:"bill_due"`
	WardNo      int    `json:"ward_number"`
}
