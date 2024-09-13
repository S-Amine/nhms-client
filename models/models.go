package models

type Patient struct {
	NIN                  string `json:"nin"`
	FirstName            string `json:"firstName"`
	LastName             string `json:"lastName"`
	DateOfBirth          string `json:"dateOfBirth"`
	Sex                  string `json:"sex"`
	MotherNIN            string `json:"motherNin"`
	FatherNIN            string `json:"fatherNin"`
	FamilyMedicalHistory string `json:"familyMedicalHistory"`
	Allergy              string `json:"allergy"`
	ChronicIllnesses     string `json:"chronicIllnesses"`
	AmendedFrom          string `json:"amendedFrom"`
}
