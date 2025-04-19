package domain

import "time"

// статусы записи
const (
	SCHEDULED = 1
	COMPLETED = 2
	CANCELED  = 3
)

type Appointment struct {
	Id        int       `json:"appointmentID"`
	DoctorID  int       `json:"doctorID"`
	PatientID int       `json:"patientID"`
	Date      time.Time `json:"date"`
	Time      time.Time `json:"time"`
	Status    string    `json:"status"`
	Rating    float64   `json:"rating"`
}

type AppointmentDTO struct {
	Id        int     `json:"appointmentID"`
	DoctorID  int     `json:"doctorID"`
	PatientID int     `json:"patientID"`
	Date      string  `json:"date"`
	Time      string  `json:"time"`
	Status    string  `json:"status"`
	Rating    float64 `json:"rating"`
}
