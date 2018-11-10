package models

type (
	PatientDetail struct {
		PatientName      string `json:"patientName"`
		PatientAge       int    `json:"patientAge"`
		PatientInsurance string `json:"patientInsurance"`
	}

	PatientInfo struct {
		PatientDetail PatientDetail `json:"patientDetail"`
	}
)
